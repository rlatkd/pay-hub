package repository

import (
	"auth/internal/models"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

func GetUserByID(db *sql.DB, id string) (*models.User, error) {
	u := &models.User{}
	query := `
		SELECT user_id, email, password_hash, name, role_id, status 
		FROM auth_user 
		WHERE user_id = ? 
	`
	err := db.QueryRow(query, id).Scan(&u.UserID, &u.Email, &u.PasswordHash, &u.Name, &u.RoleID, &u.Status)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("존재하지 않는 아이디입니다.")
		}
		return nil, err
	}
	return u, nil
}

func UpdateLastLogin(db *sql.DB, userID string) {
	_, _ = db.Exec("UPDATE auth_user SET last_login_at = NOW() WHERE user_id = ?", userID)
}

func SaveRefreshToken(db *sql.DB, userID string, deviceInfo string) (string, error) {
	tokenID := uuid.New().String()
	expiresAt := time.Now().Add(time.Hour * 24 * 7) // 7일

	query := `
		INSERT INTO auth_refresh_token (token_id, user_id, rotation_counter, expires_at, device_info)
		VALUES (?, ?, 0, ?, ?)
	`
	_, err := db.Exec(query, tokenID, userID, expiresAt, deviceInfo)
	if err != nil {
		return "", err
	}
	return tokenID, nil
}

func CreateTestUser(db *sql.DB, passwordHash string) {
	var count int
	db.QueryRow("SELECT COUNT(*) FROM auth_user WHERE user_id='rlatkd'").Scan(&count)
	
	if count > 0 { return }

	_, err := db.Exec(`
		INSERT INTO auth_user (user_id, email, password_hash, name, role_id, status)
		VALUES ('rlatkd', 'rlatkdgns042@naver.com', ?, '관리자', 'ROLE_WRITE', 'ACTIVE')
	`, passwordHash)

	if err != nil { panic(err) }
}
