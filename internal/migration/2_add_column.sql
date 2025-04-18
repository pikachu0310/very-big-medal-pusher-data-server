-- +goose Up
DROP TABLE IF EXISTS game_data;

CREATE TABLE game_data
(
    id            INT AUTO_INCREMENT PRIMARY KEY,
    user_id       VARCHAR(255) NOT NULL,
    version       VARCHAR(16)  NOT NULL,
    in_medal      INT          NOT NULL,
    out_medal     INT          NOT NULL,
    slot_hit      INT          NOT NULL,
    get_shirbe    INT          NOT NULL,
    start_slot    INT          NOT NULL,
    shirbe_buy300 INT          NOT NULL,
    medal_1       INT          NOT NULL,
    medal_2       INT          NOT NULL,
    medal_3       INT          NOT NULL,
    medal_4       INT          NOT NULL,
    medal_5       INT          NOT NULL,
    R_medal       INT          NOT NULL,
    second        INT          NOT NULL,
    minute        INT          NOT NULL,
    hour          INT          NOT NULL,
    have_medal    INT          NOT NULL DEFAULT 0,
    fever         INT          NOT NULL DEFAULT 0,
    created_at    TIMESTAMP             DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS game_data;
