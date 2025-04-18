-- +goose Up
CREATE TABLE IF NOT EXISTS game_data
(
    id            VARCHAR(36)  NOT NULL,
    user_id       VARCHAR(255) NOT NULL,
    version       VARCHAR(16)  NOT NULL,
    in_medal      INTEGER      NOT NULL,
    out_medal     INTEGER      NOT NULL,
    slot_hit      INTEGER      NOT NULL,
    get_shirbe    INTEGER      NOT NULL,
    start_slot    INTEGER      NOT NULL,
    shirbe_buy300 INTEGER      NOT NULL,
    medal_1       INTEGER      NOT NULL,
    medal_2       INTEGER      NOT NULL,
    medal_3       INTEGER      NOT NULL,
    medal_4       INTEGER      NOT NULL,
    medal_5       INTEGER      NOT NULL,
    R_medal       INTEGER      NOT NULL,
    second        FLOAT        NOT NULL,
    minute        INTEGER      NOT NULL,
    hour          INTEGER      NOT NULL,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);
