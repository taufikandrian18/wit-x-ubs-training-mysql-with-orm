CREATE TABLE IF NOT EXISTS profile (
    profile_id BIGINT NOT NULL AUTO_INCREMENT,

    username VARCHAR(255) NOT NULL,
    nama_lengkap VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    telephone VARCHAR(255) NOT NULL,
    user_password VARCHAR(255),
    is_verified BOOLEAN DEFAULT false,
    is_banned BOOLEAN DEFAULT false,
    token VARCHAR(255),

    role_id BIGINT NOT NULL, -- New column for foreign key

    status VARCHAR(255) NOT NULL DEFAULT 'active',
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255) NOT NULL,
    updated_at TIMESTAMP NULL DEFAULT NULL,
    updated_by VARCHAR(255) NULL,
    
    PRIMARY KEY (profile_id),
    FOREIGN KEY (role_id) REFERENCES role(role_id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE INDEX idx_profile ON profile (profile_id, status);