-- +goose Up
-- 1) For both “orange” and “rainbow” chain queries, you’re filtering on bc.ball_id
--    and then joining on bc.save_id → sd.id.  So give MySQL a composite that
--    covers the WHERE and the join, plus the chain_count (the “value”) if you want
--    it to be fully covering:

CREATE INDEX idx_ball_chain_ballid_saveid_count
    ON save_data_v2_ball_chain (ball_id, save_id, chain_count);

-- 2) For your max_total_jackpot query, you GROUP BY user_id, take MAX(jack_totalmax),
--    then re-join on sd.jack_totalmax = that value and pick up sd.created_at.
--    A composite on (user_id, jack_totalmax, created_at) will let the engine
--    do the entire grouping + join via the index:

CREATE INDEX idx_save_data_v2_user_jacktotal_created
    ON save_data_v2 (user_id, jack_totalmax, created_at);

-- 3) You already group by user_id in several places (and in the total_medals you
--    do MAX(id) per user).  If you don’t already have a simple index on user_id
--    itself, add that too:

CREATE INDEX idx_save_data_v2_user
    ON save_data_v2 (user_id);

ALTER TABLE save_data_v2
    ADD INDEX idx_user_jt_created (user_id, jack_totalmax, created_at),
    ADD INDEX idx_user_id_id (user_id, id);
