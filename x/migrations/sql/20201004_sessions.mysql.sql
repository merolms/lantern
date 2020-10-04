CREATE TABLE IF NOT EXISTS sessions
(
    `id`          CHAR(36) CHARACTER SET ascii NOT NULL,
    PRIMARY KEY (`id`),
    `entity_id`   CHAR(36) CHARACTER SET ascii NOT NULL,
    `entity_type` INT                          NOT NULL,
    `expires_at`  DATETIME                     NOT NULL,
    `created_at`  DATETIME                     NOT NULL,
    `updated_at`  DATETIME                     NOT NULL
)
CREATE INDEX idx_sessions_expires_at ON sessions(expires_at);