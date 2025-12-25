package handlers

import (
	"auth/internal/models"
	"auth/internal/repository"
	"database/sql"
	"html/template"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	jwtSecret = []byte("JWT_SECRET_KEY") // TODO
	tmpl = template.Must(template.ParseFiles("templates/login.html"))
)

type AuthHandler struct {
	DB *sql.DB
}

func (h *AuthHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}

func (h *AuthHandler) LoginProcess(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userID := r.FormValue("user_id")
	password := r.FormValue("password")

	user, err := repository.GetUserByID(h.DB, userID)
	if err != nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte("<div class='error'>아이디가 존재하지 않습니다.</div>"))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte("<div class='error'>비밀번호가 일치하지 않습니다.</div>"))
		return
	}

	accessToken, _ := generateAccessToken(user)
	
	userAgent := r.UserAgent()
	
	refreshTokenID, err := repository.SaveRefreshToken(h.DB, user.UserID, userAgent)
	if err != nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte("<div class='error'>서버 오류 발생</div>"))
		return
	}
	go repository.UpdateLastLogin(h.DB, user.UserID)

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(15 * time.Minute),
	})
	
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token_id",
		Value:    refreshTokenID,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
	})

	w.Header().Set("HX-Redirect", "http://localhost:3000")
	w.WriteHeader(http.StatusOK)
}

func generateAccessToken(u *models.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":  u.UserID,
		"name": u.Name,
		"role": u.RoleID,
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
