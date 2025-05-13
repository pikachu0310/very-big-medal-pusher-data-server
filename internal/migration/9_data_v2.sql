-- +goose Up
CREATE TABLE save_data_v2
(
    id            INT AUTO_INCREMENT PRIMARY KEY,
    user_id       VARCHAR(255) NOT NULL,
    legacy        TINYINT      NOT NULL,
    version       INT          NOT NULL,
    credit        BIGINT       NOT NULL DEFAULT 0,
    credit_all    BIGINT       NOT NULL DEFAULT 0,
    medal_in      INT          NOT NULL DEFAULT 0,
    medal_get     INT          NOT NULL DEFAULT 0,
    ball_get      INT          NOT NULL DEFAULT 0,
    ball_chain    INT          NOT NULL DEFAULT 0,
    slot_start    INT          NOT NULL DEFAULT 0,
    slot_startfev INT          NOT NULL DEFAULT 0,
    slot_hit      INT          NOT NULL DEFAULT 0,
    slot_getfev   INT          NOT NULL DEFAULT 0,
    sqr_get       INT          NOT NULL DEFAULT 0,
    sqr_step      INT          NOT NULL DEFAULT 0,
    jack_get      INT          NOT NULL DEFAULT 0,
    jack_startmax INT          NOT NULL DEFAULT 0,
    jack_totalmax INT          NOT NULL DEFAULT 0,
    ult_get       INT          NOT NULL DEFAULT 0,
    ult_combomax  INT          NOT NULL DEFAULT 0,
    ult_totalmax  INT          NOT NULL DEFAULT 0,
    rmshbi_get    INT          NOT NULL DEFAULT 0,
    buy_shbi      INT          NOT NULL DEFAULT 0,
    firstboot     BIGINT       NOT NULL,
    lastsave      BIGINT       NOT NULL,
    playtime      INT          NOT NULL,
    created_at    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE INDEX idx_save_data_v2_user_created_at ON save_data_v2 (user_id, created_at);
CREATE INDEX idx_save_data_v2_user_playtime ON save_data_v2 (user_id, playtime);

CREATE TABLE save_data_v2_medal_get
(
    save_id  INT          NOT NULL,
    medal_id VARCHAR(255) NOT NULL,
    count    INT          NOT NULL,
    PRIMARY KEY (save_id, medal_id),
    INDEX idx_save_data_v2_medal_get_save (save_id),
    FOREIGN KEY (save_id) REFERENCES save_data_v2 (id) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE save_data_v2_ball_get
(
    save_id INT          NOT NULL,
    ball_id VARCHAR(255) NOT NULL,
    count   INT          NOT NULL,
    PRIMARY KEY (save_id, ball_id),
    INDEX idx_save_data_v2_ball_get_save (save_id),
    FOREIGN KEY (save_id) REFERENCES save_data_v2 (id) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE save_data_v2_ball_chain
(
    save_id     INT          NOT NULL,
    ball_id     VARCHAR(255) NOT NULL,
    chain_count INT          NOT NULL,
    PRIMARY KEY (save_id, ball_id),
    INDEX idx_save_data_v2_ball_chain_save (save_id),
    FOREIGN KEY (save_id) REFERENCES save_data_v2 (id) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE save_data_v2_achievements
(
    save_id        INT          NOT NULL,
    achievement_id VARCHAR(255) NOT NULL,
    PRIMARY KEY (save_id, achievement_id),
    INDEX idx_save_data_v2_achieve_save (save_id),
    FOREIGN KEY (save_id) REFERENCES save_data_v2 (id) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

-- +goose Down
DROP TABLE IF EXISTS save_data_v2_achievements;
DROP TABLE IF EXISTS save_data_v2_ball_chain;
DROP TABLE IF EXISTS save_data_v2_ball_get;
DROP TABLE IF EXISTS save_data_v2_medal_get;

DROP INDEX IF EXISTS idx_save_data_v2_user_playtime ON save_data_v2;
DROP INDEX IF EXISTS idx_save_data_v2_user_created_at ON save_data_v2;
DROP TABLE IF EXISTS save_data_v2;
