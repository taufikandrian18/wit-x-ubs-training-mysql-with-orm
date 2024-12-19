CREATE TABLE IF NOT EXISTS authentication (
    authentication_id BIGINT NOT NULL AUTO_INCREMENT,
    profile_id BIGINT NOT NULL,
    secret VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    old_password VARCHAR(255) NULL,
    password VARCHAR(255) NOT NULL,
    password_confirmation VARCHAR(255) NULL,
    phone_number VARCHAR(20) NOT NULL,
    request_type VARCHAR(50) NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL,

    PRIMARY KEY (authentication_id),
    INDEX idx_profile_id (profile_id),
    INDEX idx_email (email)
);
