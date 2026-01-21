START TRANSACTION;

SET @from_id = 293762;
SET @to_id = 293766;
SET @base_id = 292728;

-- Sanity check: confirm same user_id for all three saves
SELECT id, user_id
FROM v2_save_data
WHERE id IN (@from_id, @to_id, @base_id);

-- v2_save_data: base + (to - from)
INSERT INTO v2_save_data (
  user_id,
  legacy,
  version,
  credit,
  credit_all,
  medal_in,
  medal_get,
  ball_get,
  ball_chain,
  slot_start,
  slot_startfev,
  slot_hit,
  slot_getfev,
  sqr_get,
  sqr_step,
  jack_get,
  jack_startmax,
  jack_totalmax,
  ult_get,
  ult_combomax,
  ult_totalmax,
  rmshbi_get,
  bstp_step,
  bstp_rwd,
  buy_total,
  skill_point,
  blackbox,
  blackbox_total,
  sp_use,
  hide_record,
  cpm_max,
  palball_get,
  pallot_lot_t0,
  pallot_lot_t1,
  pallot_lot_t2,
  pallot_lot_t3,
  pallot_lot_t4,
  task_cnt,
  totem_altars,
  totem_altars_credit,
  buy_shbi,
  firstboot,
  lastsave,
  playtime,
  created_at,
  updated_at,
  jack_totalmax_v2,
  ult_totalmax_v2,
  jacksp_get_all,
  jacksp_get_t0,
  jacksp_get_t1,
  jacksp_get_t2,
  jacksp_get_t3,
  jacksp_get_t4,
  jacksp_startmax,
  jacksp_totalmax
)
SELECT
  base.user_id,
  base.legacy + (t.legacy - f.legacy),
  base.version + (t.version - f.version),
  base.credit + (t.credit - f.credit),
  base.credit_all + (t.credit_all - f.credit_all),
  base.medal_in + (t.medal_in - f.medal_in),
  base.medal_get + (t.medal_get - f.medal_get),
  base.ball_get + (t.ball_get - f.ball_get),
  base.ball_chain + (t.ball_chain - f.ball_chain),
  base.slot_start + (t.slot_start - f.slot_start),
  base.slot_startfev + (t.slot_startfev - f.slot_startfev),
  base.slot_hit + (t.slot_hit - f.slot_hit),
  base.slot_getfev + (t.slot_getfev - f.slot_getfev),
  base.sqr_get + (t.sqr_get - f.sqr_get),
  base.sqr_step + (t.sqr_step - f.sqr_step),
  base.jack_get + (t.jack_get - f.jack_get),
  base.jack_startmax + (t.jack_startmax - f.jack_startmax),
  base.jack_totalmax + (t.jack_totalmax - f.jack_totalmax),
  base.ult_get + (t.ult_get - f.ult_get),
  base.ult_combomax + (t.ult_combomax - f.ult_combomax),
  base.ult_totalmax + (t.ult_totalmax - f.ult_totalmax),
  base.rmshbi_get + (t.rmshbi_get - f.rmshbi_get),
  base.bstp_step + (t.bstp_step - f.bstp_step),
  base.bstp_rwd + (t.bstp_rwd - f.bstp_rwd),
  base.buy_total + (t.buy_total - f.buy_total),
  base.skill_point + (t.skill_point - f.skill_point),
  base.blackbox + (t.blackbox - f.blackbox),
  base.blackbox_total + (t.blackbox_total - f.blackbox_total),
  base.sp_use + (t.sp_use - f.sp_use),
  base.hide_record + (t.hide_record - f.hide_record),
  base.cpm_max + (t.cpm_max - f.cpm_max),
  base.palball_get + (t.palball_get - f.palball_get),
  base.pallot_lot_t0 + (t.pallot_lot_t0 - f.pallot_lot_t0),
  base.pallot_lot_t1 + (t.pallot_lot_t1 - f.pallot_lot_t1),
  base.pallot_lot_t2 + (t.pallot_lot_t2 - f.pallot_lot_t2),
  base.pallot_lot_t3 + (t.pallot_lot_t3 - f.pallot_lot_t3),
  base.pallot_lot_t4 + (t.pallot_lot_t4 - f.pallot_lot_t4),
  base.task_cnt + (t.task_cnt - f.task_cnt),
  base.totem_altars + (t.totem_altars - f.totem_altars),
  base.totem_altars_credit + (t.totem_altars_credit - f.totem_altars_credit),
  base.buy_shbi + (t.buy_shbi - f.buy_shbi),
  base.firstboot + (t.firstboot - f.firstboot),
  base.lastsave + (t.lastsave - f.lastsave),
  base.playtime + (t.playtime - f.playtime),
  NOW(),
  NOW(),
  base.jack_totalmax_v2 + (t.jack_totalmax_v2 - f.jack_totalmax_v2),
  base.ult_totalmax_v2 + (t.ult_totalmax_v2 - f.ult_totalmax_v2),
  base.jacksp_get_all + (t.jacksp_get_all - f.jacksp_get_all),
  base.jacksp_get_t0 + (t.jacksp_get_t0 - f.jacksp_get_t0),
  base.jacksp_get_t1 + (t.jacksp_get_t1 - f.jacksp_get_t1),
  base.jacksp_get_t2 + (t.jacksp_get_t2 - f.jacksp_get_t2),
  base.jacksp_get_t3 + (t.jacksp_get_t3 - f.jacksp_get_t3),
  base.jacksp_get_t4 + (t.jacksp_get_t4 - f.jacksp_get_t4),
  base.jacksp_startmax + (t.jacksp_startmax - f.jacksp_startmax),
  base.jacksp_totalmax + (t.jacksp_totalmax - f.jacksp_totalmax)
FROM v2_save_data base
JOIN v2_save_data f ON f.id = @from_id
JOIN v2_save_data t ON t.id = @to_id
WHERE base.id = @base_id
  AND base.user_id = f.user_id
  AND base.user_id = t.user_id;

SET @new_save_id = LAST_INSERT_ID();

-- v2_save_data_achievements: base + (to - from)
INSERT INTO v2_save_data_achievements (save_id, achievement_id)
SELECT @new_save_id, achievement_id
FROM (
  SELECT achievement_id
  FROM v2_save_data_achievements
  WHERE save_id = @base_id
  UNION
  SELECT t.achievement_id
  FROM v2_save_data_achievements t
  LEFT JOIN v2_save_data_achievements f
    ON f.save_id = @from_id AND f.achievement_id = t.achievement_id
  WHERE t.save_id = @to_id AND f.achievement_id IS NULL
) s;

-- v2_save_data_ball_chain
WITH diff AS (
  SELECT k.ball_id,
         COALESCE(t.chain_count, 0) - COALESCE(f.chain_count, 0) AS diff_chain
  FROM (
    SELECT ball_id FROM v2_save_data_ball_chain WHERE save_id = @to_id
    UNION
    SELECT ball_id FROM v2_save_data_ball_chain WHERE save_id = @from_id
  ) k
  LEFT JOIN v2_save_data_ball_chain t ON t.save_id = @to_id AND t.ball_id = k.ball_id
  LEFT JOIN v2_save_data_ball_chain f ON f.save_id = @from_id AND f.ball_id = k.ball_id
),
base AS (
  SELECT ball_id, chain_count FROM v2_save_data_ball_chain WHERE save_id = @base_id
),
key_rows AS (
  SELECT ball_id FROM base
  UNION
  SELECT ball_id FROM diff
)
INSERT INTO v2_save_data_ball_chain (save_id, ball_id, chain_count)
SELECT @new_save_id, k.ball_id,
       COALESCE(b.chain_count, 0) + COALESCE(d.diff_chain, 0)
FROM key_rows k
LEFT JOIN base b ON b.ball_id = k.ball_id
LEFT JOIN diff d ON d.ball_id = k.ball_id
WHERE COALESCE(b.chain_count, 0) + COALESCE(d.diff_chain, 0) <> 0;

-- v2_save_data_ball_get
WITH diff AS (
  SELECT k.ball_id,
         COALESCE(t.`count`, 0) - COALESCE(f.`count`, 0) AS diff_count
  FROM (
    SELECT ball_id FROM v2_save_data_ball_get WHERE save_id = @to_id
    UNION
    SELECT ball_id FROM v2_save_data_ball_get WHERE save_id = @from_id
  ) k
  LEFT JOIN v2_save_data_ball_get t ON t.save_id = @to_id AND t.ball_id = k.ball_id
  LEFT JOIN v2_save_data_ball_get f ON f.save_id = @from_id AND f.ball_id = k.ball_id
),
base AS (
  SELECT ball_id, `count` FROM v2_save_data_ball_get WHERE save_id = @base_id
),
key_rows AS (
  SELECT ball_id FROM base
  UNION
  SELECT ball_id FROM diff
)
INSERT INTO v2_save_data_ball_get (save_id, ball_id, `count`)
SELECT @new_save_id, k.ball_id,
       COALESCE(b.`count`, 0) + COALESCE(d.diff_count, 0)
FROM key_rows k
LEFT JOIN base b ON b.ball_id = k.ball_id
LEFT JOIN diff d ON d.ball_id = k.ball_id
WHERE COALESCE(b.`count`, 0) + COALESCE(d.diff_count, 0) <> 0;

-- v2_save_data_medal_get
WITH diff AS (
  SELECT k.medal_id,
         COALESCE(t.`count`, 0) - COALESCE(f.`count`, 0) AS diff_count
  FROM (
    SELECT medal_id FROM v2_save_data_medal_get WHERE save_id = @to_id
    UNION
    SELECT medal_id FROM v2_save_data_medal_get WHERE save_id = @from_id
  ) k
  LEFT JOIN v2_save_data_medal_get t ON t.save_id = @to_id AND t.medal_id = k.medal_id
  LEFT JOIN v2_save_data_medal_get f ON f.save_id = @from_id AND f.medal_id = k.medal_id
),
base AS (
  SELECT medal_id, `count` FROM v2_save_data_medal_get WHERE save_id = @base_id
),
key_rows AS (
  SELECT medal_id FROM base
  UNION
  SELECT medal_id FROM diff
)
INSERT INTO v2_save_data_medal_get (save_id, medal_id, `count`)
SELECT @new_save_id, k.medal_id,
       COALESCE(b.`count`, 0) + COALESCE(d.diff_count, 0)
FROM key_rows k
LEFT JOIN base b ON b.medal_id = k.medal_id
LEFT JOIN diff d ON d.medal_id = k.medal_id
WHERE COALESCE(b.`count`, 0) + COALESCE(d.diff_count, 0) <> 0;

-- v2_save_data_palball_get
WITH diff AS (
  SELECT k.ball_id,
         COALESCE(t.`count`, 0) - COALESCE(f.`count`, 0) AS diff_count
  FROM (
    SELECT ball_id FROM v2_save_data_palball_get WHERE save_id = @to_id
    UNION
    SELECT ball_id FROM v2_save_data_palball_get WHERE save_id = @from_id
  ) k
  LEFT JOIN v2_save_data_palball_get t ON t.save_id = @to_id AND t.ball_id = k.ball_id
  LEFT JOIN v2_save_data_palball_get f ON f.save_id = @from_id AND f.ball_id = k.ball_id
),
base AS (
  SELECT ball_id, `count` FROM v2_save_data_palball_get WHERE save_id = @base_id
),
key_rows AS (
  SELECT ball_id FROM base
  UNION
  SELECT ball_id FROM diff
)
INSERT INTO v2_save_data_palball_get (save_id, ball_id, `count`)
SELECT @new_save_id, k.ball_id,
       COALESCE(b.`count`, 0) + COALESCE(d.diff_count, 0)
FROM key_rows k
LEFT JOIN base b ON b.ball_id = k.ball_id
LEFT JOIN diff d ON d.ball_id = k.ball_id
WHERE COALESCE(b.`count`, 0) + COALESCE(d.diff_count, 0) <> 0;

-- v2_save_data_palball_jp
WITH diff AS (
  SELECT k.ball_id,
         COALESCE(t.`count`, 0) - COALESCE(f.`count`, 0) AS diff_count
  FROM (
    SELECT ball_id FROM v2_save_data_palball_jp WHERE save_id = @to_id
    UNION
    SELECT ball_id FROM v2_save_data_palball_jp WHERE save_id = @from_id
  ) k
  LEFT JOIN v2_save_data_palball_jp t ON t.save_id = @to_id AND t.ball_id = k.ball_id
  LEFT JOIN v2_save_data_palball_jp f ON f.save_id = @from_id AND f.ball_id = k.ball_id
),
base AS (
  SELECT ball_id, `count` FROM v2_save_data_palball_jp WHERE save_id = @base_id
),
key_rows AS (
  SELECT ball_id FROM base
  UNION
  SELECT ball_id FROM diff
)
INSERT INTO v2_save_data_palball_jp (save_id, ball_id, `count`)
SELECT @new_save_id, k.ball_id,
       COALESCE(b.`count`, 0) + COALESCE(d.diff_count, 0)
FROM key_rows k
LEFT JOIN base b ON b.ball_id = k.ball_id
LEFT JOIN diff d ON d.ball_id = k.ball_id
WHERE COALESCE(b.`count`, 0) + COALESCE(d.diff_count, 0) <> 0;

-- v2_save_data_perks
WITH diff AS (
  SELECT k.perk_id,
         COALESCE(t.level, 0) - COALESCE(f.level, 0) AS diff_level
  FROM (
    SELECT perk_id FROM v2_save_data_perks WHERE save_id = @to_id
    UNION
    SELECT perk_id FROM v2_save_data_perks WHERE save_id = @from_id
  ) k
  LEFT JOIN v2_save_data_perks t ON t.save_id = @to_id AND t.perk_id = k.perk_id
  LEFT JOIN v2_save_data_perks f ON f.save_id = @from_id AND f.perk_id = k.perk_id
),
base AS (
  SELECT perk_id, level FROM v2_save_data_perks WHERE save_id = @base_id
),
key_rows AS (
  SELECT perk_id FROM base
  UNION
  SELECT perk_id FROM diff
)
INSERT INTO v2_save_data_perks (save_id, perk_id, level)
SELECT @new_save_id, k.perk_id,
       COALESCE(b.level, 0) + COALESCE(d.diff_level, 0)
FROM key_rows k
LEFT JOIN base b ON b.perk_id = k.perk_id
LEFT JOIN diff d ON d.perk_id = k.perk_id
WHERE COALESCE(b.level, 0) + COALESCE(d.diff_level, 0) <> 0;

-- v2_save_data_perks_credit
WITH diff AS (
  SELECT k.perk_id,
         COALESCE(t.credits, 0) - COALESCE(f.credits, 0) AS diff_credits
  FROM (
    SELECT perk_id FROM v2_save_data_perks_credit WHERE save_id = @to_id
    UNION
    SELECT perk_id FROM v2_save_data_perks_credit WHERE save_id = @from_id
  ) k
  LEFT JOIN v2_save_data_perks_credit t ON t.save_id = @to_id AND t.perk_id = k.perk_id
  LEFT JOIN v2_save_data_perks_credit f ON f.save_id = @from_id AND f.perk_id = k.perk_id
),
base AS (
  SELECT perk_id, credits FROM v2_save_data_perks_credit WHERE save_id = @base_id
),
key_rows AS (
  SELECT perk_id FROM base
  UNION
  SELECT perk_id FROM diff
)
INSERT INTO v2_save_data_perks_credit (save_id, perk_id, credits)
SELECT @new_save_id, k.perk_id,
       COALESCE(b.credits, 0) + COALESCE(d.diff_credits, 0)
FROM key_rows k
LEFT JOIN base b ON b.perk_id = k.perk_id
LEFT JOIN diff d ON d.perk_id = k.perk_id
WHERE COALESCE(b.credits, 0) + COALESCE(d.diff_credits, 0) <> 0;

-- v2_save_data_totems
WITH diff AS (
  SELECT k.totem_id,
         COALESCE(t.level, 0) - COALESCE(f.level, 0) AS diff_level
  FROM (
    SELECT totem_id FROM v2_save_data_totems WHERE save_id = @to_id
    UNION
    SELECT totem_id FROM v2_save_data_totems WHERE save_id = @from_id
  ) k
  LEFT JOIN v2_save_data_totems t ON t.save_id = @to_id AND t.totem_id = k.totem_id
  LEFT JOIN v2_save_data_totems f ON f.save_id = @from_id AND f.totem_id = k.totem_id
),
base AS (
  SELECT totem_id, level FROM v2_save_data_totems WHERE save_id = @base_id
),
key_rows AS (
  SELECT totem_id FROM base
  UNION
  SELECT totem_id FROM diff
)
INSERT INTO v2_save_data_totems (save_id, totem_id, level)
SELECT @new_save_id, k.totem_id,
       COALESCE(b.level, 0) + COALESCE(d.diff_level, 0)
FROM key_rows k
LEFT JOIN base b ON b.totem_id = k.totem_id
LEFT JOIN diff d ON d.totem_id = k.totem_id
WHERE COALESCE(b.level, 0) + COALESCE(d.diff_level, 0) <> 0;

-- v2_save_data_totems_credit
WITH diff AS (
  SELECT k.totem_id,
         COALESCE(t.credits, 0) - COALESCE(f.credits, 0) AS diff_credits
  FROM (
    SELECT totem_id FROM v2_save_data_totems_credit WHERE save_id = @to_id
    UNION
    SELECT totem_id FROM v2_save_data_totems_credit WHERE save_id = @from_id
  ) k
  LEFT JOIN v2_save_data_totems_credit t ON t.save_id = @to_id AND t.totem_id = k.totem_id
  LEFT JOIN v2_save_data_totems_credit f ON f.save_id = @from_id AND f.totem_id = k.totem_id
),
base AS (
  SELECT totem_id, credits FROM v2_save_data_totems_credit WHERE save_id = @base_id
),
key_rows AS (
  SELECT totem_id FROM base
  UNION
  SELECT totem_id FROM diff
)
INSERT INTO v2_save_data_totems_credit (save_id, totem_id, credits)
SELECT @new_save_id, k.totem_id,
       COALESCE(b.credits, 0) + COALESCE(d.diff_credits, 0)
FROM key_rows k
LEFT JOIN base b ON b.totem_id = k.totem_id
LEFT JOIN diff d ON d.totem_id = k.totem_id
WHERE COALESCE(b.credits, 0) + COALESCE(d.diff_credits, 0) <> 0;

-- v2_save_data_totems_placement: take changes from to_id
WITH key_rows AS (
  SELECT placement_idx FROM v2_save_data_totems_placement WHERE save_id = @base_id
  UNION
  SELECT placement_idx FROM v2_save_data_totems_placement WHERE save_id = @from_id
  UNION
  SELECT placement_idx FROM v2_save_data_totems_placement WHERE save_id = @to_id
),
base AS (
  SELECT placement_idx, totem_id FROM v2_save_data_totems_placement WHERE save_id = @base_id
),
f AS (
  SELECT placement_idx, totem_id FROM v2_save_data_totems_placement WHERE save_id = @from_id
),
t AS (
  SELECT placement_idx, totem_id FROM v2_save_data_totems_placement WHERE save_id = @to_id
)
INSERT INTO v2_save_data_totems_placement (save_id, placement_idx, totem_id)
SELECT @new_save_id, s.placement_idx, s.totem_id
FROM (
  SELECT k.placement_idx,
         CASE
           WHEN t.totem_id IS NOT NULL THEN t.totem_id
           WHEN f.totem_id IS NOT NULL AND t.totem_id IS NULL THEN NULL
           ELSE b.totem_id
         END AS totem_id
  FROM key_rows k
  LEFT JOIN base b ON b.placement_idx = k.placement_idx
  LEFT JOIN f ON f.placement_idx = k.placement_idx
  LEFT JOIN t ON t.placement_idx = k.placement_idx
) s
WHERE s.totem_id IS NOT NULL;

COMMIT;
