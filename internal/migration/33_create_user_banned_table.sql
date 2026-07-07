-- +goose Up
-- Add global ranking shadow ban table

CREATE TABLE IF NOT EXISTS user_banned (
    user_id    VARCHAR(255) NOT NULL,
    enabled    TINYINT(1)   NOT NULL DEFAULT 1,
    note       TEXT         NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id),
    KEY idx_user_banned_enabled (enabled)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +goose Down
DROP TABLE IF EXISTS user_banned;
