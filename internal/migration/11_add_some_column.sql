-- +goose Up
ALTER TABLE save_data_v2
    /* ボーナスステップ進行回数 */
    ADD COLUMN bstp_step INT NOT NULL DEFAULT 0 AFTER rmshbi_get,
    /* ボーナスステップ報酬回数 */
    ADD COLUMN bstp_rwd  INT NOT NULL DEFAULT 0 AFTER bstp_step,
    /* ショップ購入ボタン押下総数 */
    ADD COLUMN buy_total INT NOT NULL DEFAULT 0 AFTER bstp_rwd,
    /* スキルポイント消費数 */
    ADD COLUMN sp_use    INT NOT NULL DEFAULT 0 AFTER buy_total;

-- +goose Down
ALTER TABLE save_data_v2
    DROP COLUMN sp_use,
    DROP COLUMN buy_total,
    DROP COLUMN bstp_rwd,
    DROP COLUMN bstp_step;
