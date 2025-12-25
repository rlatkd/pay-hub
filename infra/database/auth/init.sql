-- 권한
CREATE TABLE auth_role (
    role_id VARCHAR(50) PRIMARY KEY,
    description VARCHAR(255)
);
INSERT INTO auth_role (role_id, description) VALUES ('ROLE_READ', '읽기 전용');
INSERT INTO auth_role (role_id, description) VALUES ('ROLE_WRITE', '쓰기 전용');

-- 사용자
CREATE TABLE auth_user (
    user_id VARCHAR(255) PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    status VARCHAR(50) DEFAULT 'ACTIVE',
    role_id VARCHAR(50) NOT NULL, 
    last_login_at DATETIME(6),
    created_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6),
    updated_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
    CONSTRAINT fk_user_role FOREIGN KEY (role_id) REFERENCES auth_role(role_id)
);

-- 리프레시 토큰
CREATE TABLE auth_refresh_token (
    token_id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    rotation_counter INT DEFAULT 0,
    issued_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6),
    expires_at DATETIME(6) NOT NULL,
    device_info VARCHAR(255),
    CONSTRAINT fk_token_user FOREIGN KEY (user_id) REFERENCES auth_user(user_id) ON DELETE CASCADE
);

CREATE INDEX idx_auth_user_email ON auth_user(email);
