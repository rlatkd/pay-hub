package handlers

import (
	"auth/internal/models"
	"auth/internal/repository"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

 const AccessTokenDuration = 24 * time.Hour

var (
	jwtSecret = []byte("JWT_SECRET_KEY") // TODO
	tmpl      = template.Must(template.ParseFiles("templates/login.html"))
)

type AuthHandler struct {
	DB *sql.DB
}

func (h *AuthHandler) getUserFromCookie(r *http.Request) (*models.User, error) {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("유효하지 않은 토큰")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("클레임 파싱 실패")
	}

	userID := claims["sub"].(string)
	return repository.GetUserByID(h.DB, userID)
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
		Expires:  time.Now().Add(AccessTokenDuration),
	})
	
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token_id",
		Value:    refreshTokenID,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(7 * AccessTokenDuration),
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "login_at",
		Value:    time.Now().Format(time.RFC3339),
		Path:     "/",
		HttpOnly: false,
		Expires:  time.Now().Add(AccessTokenDuration),
	})

	w.Header().Set("HX-Redirect", "http://localhost:3000")
	w.WriteHeader(http.StatusOK)
}

func generateAccessToken(u *models.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":  u.UserID,
		"name": u.Name,
		"role": u.RoleID,
		"exp":  time.Now().Add(AccessTokenDuration).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func (h *AuthHandler) RefreshProcess(w http.ResponseWriter, r *http.Request) {
	user, err := h.getUserFromCookie(r) 
	if err != nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte("<div class='error'>세션 정보가 유효하지 않습니다.</div>"))
		return
	}

	newAccessToken, err := generateAccessToken(user)
	if err != nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte("<div class='error'>토큰 생성 오류 발생</div>"))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    newAccessToken,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(AccessTokenDuration),
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "login_at",
		Value:    time.Now().Format(time.RFC3339),
		Path:     "/",
		HttpOnly: false,
		Expires:  time.Now().Add(AccessTokenDuration),
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"success": true}`))
}
