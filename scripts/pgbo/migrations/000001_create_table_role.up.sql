CREATE TABLE IF NOT EXISTS role (
    role_id BIGINT NOT NULL AUTO_INCREMENT,

    role_name VARCHAR(255) NOT NULL,

    status VARCHAR(255) NOT NULL DEFAULT 'active',
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(255) NOT NULL,
    updated_at TIMESTAMP NULL DEFAULT NULL,
    updated_by VARCHAR(255) NULL,
    
    PRIMARY KEY (role_id)
);

CREATE INDEX idx_role ON role (role_id, status);