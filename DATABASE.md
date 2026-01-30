# データベース定義（正規ドキュメント）

> 本ドキュメントがスキーマの唯一の参照源です。内容は `internal/migration/*.sql` を踏まえて手動更新しています（v2_save_data の long 対応含む）。

このドキュメントは、現在稼働中の MariaDB（`app` データベース）を手動で把握・更新するためのリファレンスです。以下のフォーマットを維持しておけば、今後 AI で自動更新する際も扱いやすいはずです。

---

## 1. 接続方法

### 1.1 Docker コンテナ経由（推奨）

```bash
# シェル付きで root 接続
docker exec -it myapp_db mariadb -u root -ppass

# app データベースを指定して接続
docker exec -it myapp_db mariadb -u root -ppass app

# 1 クエリだけ実行
docker exec myapp_db mariadb -u root -ppass -e "SHOW DATABASES;"
```

### 1.2 Adminer (Web UI)
- URL: http://localhost:8081
- サーバー: `db`
- ユーザー: `root`
- パスワード: `pass`
- データベース: `app`

---

## 2. 実行中コンテナと DB 設定

| 項目 | 値 |
| ---- | --- |
| コンテナ名 | `myapp_db` |
| イメージ | `mariadb:latest` |
| バージョン | 11.7.2-MariaDB-ubu2404 |
| 公開ポート | `3306:3306` |
| ルートパスワード | `pass` |
| デフォルト DB | `app` |
| 文字セット | `utf8mb4` |
| 照合順序 | `utf8mb4_unicode_ci`（一部 `utf8mb4_uca1400_ai_ci`） |

存在するデータベース:
1. `app` (メイン)
2. `information_schema`
3. `mysql`
4. `performance_schema`
5. `sys`

---

## 3. app データベース概要

### 3.1 テーブル一覧（役割サマリ）

| テーブル | 役割/用途 | 備考 |
|----------|-----------|------|
| `goose_db_version` | goose マイグレーション履歴 | 自動生成テーブル |
| `v1_game_data` | 旧 v1 API 用ゲーム統計 | 新規更新は想定せず |
| `v2_save_data` | ユーザーセーブデータ本体 | もっとも参照されるテーブル |
| `v2_save_data_*` | v2 セーブデータの詳細 (実績/各種詳細カウンタ等) | すべて `v2_save_data.id` に FK |
| `v3_user_latest_save_data` | 最新セーブのサマリ | v3/v4 API のランキング高速化 |
| `v3_user_latest_save_data_achievements` | 最新セーブの実績一覧 | v3 用キャッシュ |

### 3.2 テーブルサイズ（`SHOW TABLE STATUS` 抜粋）

| テーブル | 行数 | データサイズ | インデックスサイズ | Collation |
|----------|------|--------------|---------------------|-----------|
| v1_game_data | 42,009 | 6,672 KB | 5,152 KB | utf8mb4_uca1400_ai_ci |
| v2_save_data | 115,328 | 37,440 KB | 50,304 KB | utf8mb4_uca1400_ai_ci |
| v2_save_data_achievements | 15,599,961 | 533,504 KB | 891,744 KB | utf8mb4_uca1400_ai_ci |
| v2_save_data_ball_chain | 1,987,451 | 86,672 KB | 122,624 KB | utf8mb4_uca1400_ai_ci |
| v2_save_data_ball_get | 1,945,962 | 60,000 KB | 24,112 KB | utf8mb4_uca1400_ai_ci |
| v2_save_data_medal_get | 1,534,639 | 47,696 KB | 18,992 KB | utf8mb4_uca1400_ai_ci |
| v2_save_data_palball_get | 71,063 | 2,576 KB | 1,552 KB | utf8mb4_uca1400_ai_ci |
| v2_save_data_palball_jp | 58,767 | 2,576 KB | 1,552 KB | utf8mb4_uca1400_ai_ci |
| v2_save_data_perks | 430,272 | 13,840 KB | 6,672 KB | utf8mb4_uca1400_ai_ci |
| v2_save_data_perks_credit | 429,880 | 15,920 KB | 6,672 KB | utf8mb4_uca1400_ai_ci |
| v2_save_data_totems | 0 | 16 KB | 16 KB | utf8mb4_uca1400_ai_ci |
| v2_save_data_totems_credit | 0 | 16 KB | 16 KB | utf8mb4_uca1400_ai_ci |
| v2_save_data_totems_placement | 0 | 16 KB | 16 KB | utf8mb4_uca1400_ai_ci |
| v3_user_latest_save_data | 24,156 | 7,696 KB | 17,072 KB | utf8mb4_unicode_ci |
| v3_user_latest_save_data_achievements | 2,505,741 | 139,056 KB | 223,328 KB | utf8mb4_unicode_ci |
| goose_db_version | 20 | 16 KB | 16 KB | utf8mb4_uca1400_ai_ci |

---

## 4. 外部キー制約

| 制約名 | 子テーブル | カラム | 親テーブル | 親カラム | 挙動 |
|--------|------------|--------|------------|----------|------|
| v2_save_data_achievements_ibfk_1 | v2_save_data_achievements | save_id | v2_save_data | id | ON DELETE CASCADE |
| v2_save_data_ball_chain_ibfk_1 | v2_save_data_ball_chain | save_id | v2_save_data | id | ON DELETE CASCADE |
| v2_save_data_ball_get_ibfk_1 | v2_save_data_ball_get | save_id | v2_save_data | id | ON DELETE CASCADE |
| v2_save_data_medal_get_ibfk_1 | v2_save_data_medal_get | save_id | v2_save_data | id | ON DELETE CASCADE |
| v2_save_data_palball_get_ibfk_1 | v2_save_data_palball_get | save_id | v2_save_data | id | ON DELETE CASCADE |
| v2_save_data_palball_jp_ibfk_1 | v2_save_data_palball_jp | save_id | v2_save_data | id | ON DELETE CASCADE |
| v2_save_data_perks_ibfk_1 | v2_save_data_perks | save_id | v2_save_data | id | ON DELETE CASCADE |
| v2_save_data_perks_credit_ibfk_1 | v2_save_data_perks_credit | save_id | v2_save_data | id | ON DELETE CASCADE |
| v2_save_data_totems_ibfk_1 | v2_save_data_totems | save_id | v2_save_data | id | ON DELETE CASCADE |
| v2_save_data_totems_credit_ibfk_1 | v2_save_data_totems_credit | save_id | v2_save_data | id | ON DELETE CASCADE |
| v2_save_data_totems_placement_ibfk_1 | v2_save_data_totems_placement | save_id | v2_save_data | id | ON DELETE CASCADE |

---

## 5. テーブル DDL 一覧

> 取得コマンド例: `docker exec myapp_db mariadb -u root -ppass app -e 'SHOW CREATE TABLE \`v2_save_data\`\G'`

### 5.1 goose_db_version

```sql
CREATE TABLE `goose_db_version` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `version_id` bigint(20) NOT NULL,
  `is_applied` tinyint(1) NOT NULL,
  `tstamp` timestamp NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
```

### 5.2 v1_game_data

```sql
CREATE TABLE `v1_game_data` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) NOT NULL,
  `version` int(11) NOT NULL,
  `have_medal` int(11) NOT NULL DEFAULT 0,
  `in_medal` int(11) NOT NULL,
  `out_medal` int(11) NOT NULL,
  `slot_hit` int(11) NOT NULL,
  `get_shirbe` int(11) NOT NULL,
  `start_slot` int(11) NOT NULL,
  `shirbe_buy300` int(11) NOT NULL,
  `medal_1` int(11) NOT NULL,
  `medal_2` int(11) NOT NULL,
  `medal_3` int(11) NOT NULL,
  `medal_4` int(11) NOT NULL,
  `medal_5` int(11) NOT NULL,
  `R_medal` int(11) NOT NULL,
  `total_play_time` int(11) NOT NULL,
  `fever` int(11) NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `max_chain_item` int(11) NOT NULL DEFAULT 0,
  `max_chain_orange` int(11) NOT NULL DEFAULT 0,
  `max_chain_rainbow` int(11) NOT NULL DEFAULT 0,
  `sugoroku_steps` int(11) NOT NULL DEFAULT 0,
  `jackpots` int(11) NOT NULL DEFAULT 0,
  `max_jackpot_win` int(11) NOT NULL DEFAULT 0,
  `max_total_jackpot` int(11) NOT NULL DEFAULT 0,
  `max_total_ultimate` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `idx_game_data_user_created_at` (`user_id`,`created_at`),
  KEY `idx_game_data_user_total_play_time` (`user_id`,`total_play_time`)
) ENGINE=InnoDB AUTO_INCREMENT=44001 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
```

### 5.3 v2_save_data

```sql
CREATE TABLE `v2_save_data` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) NOT NULL,
  `legacy` tinyint(4) NOT NULL,
  `version` int(11) NOT NULL,
  `credit` bigint(20) NOT NULL DEFAULT 0,
  `credit_all` bigint(20) NOT NULL DEFAULT 0,
  `medal_in` int(11) NOT NULL DEFAULT 0,
  `medal_get` bigint(20) NOT NULL DEFAULT 0,
  `ball_get` bigint(20) NOT NULL DEFAULT 0,
  `ball_chain` int(11) NOT NULL DEFAULT 0,
  `slot_start` bigint(20) NOT NULL DEFAULT 0,
  `slot_startfev` bigint(20) NOT NULL DEFAULT 0,
  `slot_hit` bigint(20) NOT NULL DEFAULT 0,
  `slot_getfev` bigint(20) NOT NULL DEFAULT 0,
  `sqr_get` bigint(20) NOT NULL DEFAULT 0,
  `sqr_step` bigint(20) NOT NULL DEFAULT 0,
  `jack_get` bigint(20) NOT NULL DEFAULT 0,
  `jack_startmax` bigint(20) NOT NULL DEFAULT 0,
  `jack_totalmax` int(11) NOT NULL DEFAULT 0,
  `ult_get` int(11) NOT NULL DEFAULT 0,
  `ult_combomax` int(11) NOT NULL DEFAULT 0,
  `ult_totalmax` int(11) NOT NULL DEFAULT 0,
  `rmshbi_get` int(11) NOT NULL DEFAULT 0,
  `bstp_step` bigint(20) NOT NULL DEFAULT 0,
  `bstp_rwd` bigint(20) NOT NULL DEFAULT 0,
  `buy_total` int(11) NOT NULL DEFAULT 0,
  `skill_point` bigint(20) NOT NULL DEFAULT 0,
  `blackbox` int(11) NOT NULL DEFAULT 0,
  `blackbox_total` bigint(20) NOT NULL DEFAULT 0,
  `sp_use` bigint(20) NOT NULL DEFAULT 0,
  `hide_record` int(11) NOT NULL DEFAULT 0,
  `cpm_max` double NOT NULL DEFAULT 0,
  `palball_get` int(11) NOT NULL DEFAULT 0,
  `pallot_lot_t0` int(11) NOT NULL DEFAULT 0,
  `pallot_lot_t1` int(11) NOT NULL DEFAULT 0,
  `pallot_lot_t2` int(11) NOT NULL DEFAULT 0,
  `pallot_lot_t3` int(11) NOT NULL DEFAULT 0,
  `pallot_lot_t4` int(11) NOT NULL DEFAULT 0,
  `task_cnt` int(11) NOT NULL DEFAULT 0,
  `totem_altars` int(11) NOT NULL DEFAULT 0,
  `totem_altars_credit` bigint(20) NOT NULL DEFAULT 0,
  `buy_shbi` int(11) NOT NULL DEFAULT 0,
  `firstboot` bigint(20) NOT NULL,
  `lastsave` bigint(20) NOT NULL,
  `playtime` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `jack_totalmax_v2` bigint(20) NOT NULL DEFAULT 0,
  `ult_totalmax_v2` bigint(20) NOT NULL DEFAULT 0,
  `jacksp_get_all` int(11) NOT NULL DEFAULT 0,
  `jacksp_get_t0` int(11) NOT NULL DEFAULT 0,
  `jacksp_get_t1` int(11) NOT NULL DEFAULT 0,
  `jacksp_get_t2` int(11) NOT NULL DEFAULT 0,
  `jacksp_get_t3` int(11) NOT NULL DEFAULT 0,
  `jacksp_get_t4` int(11) NOT NULL DEFAULT 0,
  `jacksp_startmax` bigint(20) NOT NULL DEFAULT 0,
  `jacksp_totalmax` bigint(20) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `idx_save_data_v2_user_created_at` (`user_id`,`created_at`),
  KEY `idx_save_data_v2_user_playtime` (`user_id`,`playtime`),
  KEY `idx_sd_user_jptotal_created` (`user_id`,`jack_totalmax`,`created_at`),
  KEY `idx_sd_user` (`user_id`),
  KEY `idx_save_data_v2_user_jacktotal_created` (`user_id`,`jack_totalmax`,`created_at`),
  KEY `idx_save_data_v2_user` (`user_id`),
  KEY `idx_user_jt_created` (`user_id`,`jack_totalmax`,`created_at`),
  KEY `idx_user_id_id` (`user_id`,`id`)
) ENGINE=InnoDB AUTO_INCREMENT=122532 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
```

### 5.4 v2_save_data_achievements

```sql
CREATE TABLE `v2_save_data_achievements` (
  `save_id` int(11) NOT NULL,
  `achievement_id` varchar(255) NOT NULL,
  PRIMARY KEY (`save_id`,`achievement_id`),
  KEY `idx_save_data_v2_achieve_save` (`save_id`),
  KEY `idx_save_data_v2_achievements_achievement_id` (`achievement_id`),
  KEY `idx_save_data_v2_achievements_user_achievement` (`save_id`,`achievement_id`),
  CONSTRAINT `v2_save_data_achievements_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
```

### 5.5 v2_save_data_ball_chain

```sql
CREATE TABLE `v2_save_data_ball_chain` (
  `save_id` int(11) NOT NULL,
  `ball_id` varchar(255) NOT NULL,
  `chain_count` int(11) NOT NULL,
  PRIMARY KEY (`save_id`,`ball_id`),
  KEY `idx_save_data_v2_ball_chain_save` (`save_id`),
  KEY `idx_bc_ballid_chaincnt_saveid` (`ball_id`,`chain_count`,`save_id`),
  KEY `idx_ball_chain_ballid_saveid_count` (`ball_id`,`save_id`,`chain_count`),
  CONSTRAINT `v2_save_data_ball_chain_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
```

### 5.6 v2_save_data_ball_get

```sql
CREATE TABLE `v2_save_data_ball_get` (
  `save_id` int(11) NOT NULL,
  `ball_id` varchar(255) NOT NULL,
  `count` bigint(20) NOT NULL,
  PRIMARY KEY (`save_id`,`ball_id`),
  KEY `idx_save_data_v2_ball_get_save` (`save_id`),
  CONSTRAINT `v2_save_data_ball_get_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
```

### 5.7 v2_save_data_medal_get

```sql
CREATE TABLE `v2_save_data_medal_get` (
  `save_id` int(11) NOT NULL,
  `medal_id` varchar(255) NOT NULL,
  `count` int(11) NOT NULL,
  PRIMARY KEY (`save_id`,`medal_id`),
  KEY `idx_save_data_v2_medal_get_save` (`save_id`),
  CONSTRAINT `v2_save_data_medal_get_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
```

### 5.8 v2_save_data_palball_get

```sql
CREATE TABLE `v2_save_data_palball_get` (
  `save_id` int(11) NOT NULL,
  `ball_id` varchar(255) NOT NULL,
  `count` int(11) NOT NULL,
  PRIMARY KEY (`save_id`,`ball_id`),
  KEY `idx_save_data_v2_palball_get_save` (`save_id`),
  CONSTRAINT `v2_save_data_palball_get_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
```

### 5.9 v2_save_data_palball_jp

```sql
CREATE TABLE `v2_save_data_palball_jp` (
  `save_id` int(11) NOT NULL,
  `ball_id` varchar(255) NOT NULL,
  `count` int(11) NOT NULL,
  PRIMARY KEY (`save_id`,`ball_id`),
  KEY `idx_save_data_v2_palball_jp_save` (`save_id`),
  CONSTRAINT `v2_save_data_palball_jp_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
```

### 5.10 v2_save_data_perks

```sql
CREATE TABLE `v2_save_data_perks` (
  `save_id` int(11) NOT NULL,
  `perk_id` int(11) NOT NULL,
  `level` int(11) NOT NULL,
  PRIMARY KEY (`save_id`,`perk_id`),
  KEY `idx_save_data_v2_perks_save` (`save_id`),
  CONSTRAINT `v2_save_data_perks_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
```

### 5.11 v2_save_data_perks_credit

```sql
CREATE TABLE `v2_save_data_perks_credit` (
  `save_id` int(11) NOT NULL,
  `perk_id` int(11) NOT NULL,
  `credits` bigint(20) NOT NULL DEFAULT 0,
  PRIMARY KEY (`save_id`,`perk_id`),
  KEY `idx_save_data_v2_perks_credit_save` (`save_id`),
  CONSTRAINT `v2_save_data_perks_credit_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
```

### 5.12 v2_save_data_totems

```sql
CREATE TABLE `v2_save_data_totems` (
  `save_id` int(11) NOT NULL,
  `totem_id` int(11) NOT NULL,
  `level` int(11) NOT NULL,
  PRIMARY KEY (`save_id`,`totem_id`),
  KEY `idx_v2_save_data_totems_save` (`save_id`),
  CONSTRAINT `v2_save_data_totems_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
```

### 5.13 v2_save_data_totems_credit

```sql
CREATE TABLE `v2_save_data_totems_credit` (
  `save_id` int(11) NOT NULL,
  `totem_id` int(11) NOT NULL,
  `credits` bigint(20) NOT NULL,
  PRIMARY KEY (`save_id`,`totem_id`),
  KEY `idx_v2_save_data_totems_credit_save` (`save_id`),
  CONSTRAINT `v2_save_data_totems_credit_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
```

### 5.14 v2_save_data_totems_placement

```sql
CREATE TABLE `v2_save_data_totems_placement` (
  `save_id` int(11) NOT NULL,
  `placement_idx` int(11) NOT NULL,
  `totem_id` int(11) NOT NULL,
  PRIMARY KEY (`save_id`,`placement_idx`),
  KEY `idx_v2_save_data_totems_placement_save` (`save_id`),
  CONSTRAINT `v2_save_data_totems_placement_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
```

### 5.15 v3_user_latest_save_data

```sql
CREATE TABLE `v3_user_latest_save_data` (
  `user_id` varchar(255) NOT NULL,
  `save_id` int(11) NOT NULL,
  `version` int(11) NOT NULL DEFAULT 0,
  `credit_all` bigint(20) NOT NULL DEFAULT 0,
  `playtime` bigint(20) NOT NULL DEFAULT 0,
  `achievements_count` int(11) NOT NULL DEFAULT 0,
  `jacksp_startmax` bigint(20) NOT NULL DEFAULT 0,
  `golden_palball_get` int(11) NOT NULL DEFAULT 0,
  `cpm_max` double NOT NULL DEFAULT 0,
  `max_chain_rainbow` int(11) NOT NULL DEFAULT 0,
  `jack_totalmax_v2` int(11) NOT NULL DEFAULT 0,
  `ult_combomax` int(11) NOT NULL DEFAULT 0,
  `ult_totalmax_v2` int(11) NOT NULL DEFAULT 0,
  `blackbox_total` bigint(20) NOT NULL DEFAULT 0,
  `sp_use` bigint(20) NOT NULL DEFAULT 0,
  `hide_record` int(11) NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`user_id`),
  KEY `idx_v3_user_latest_save_data_credit_all` (`credit_all` DESC),
  KEY `idx_v3_user_latest_save_data_achievements_count` (`achievements_count` DESC),
  KEY `idx_v3_user_latest_save_data_jacksp_startmax` (`jacksp_startmax` DESC),
  KEY `idx_v3_user_latest_save_data_golden_palball_get` (`golden_palball_get` DESC),
  KEY `idx_v3_user_latest_save_data_cpm_max` (`cpm_max` DESC),
  KEY `idx_v3_user_latest_save_data_max_chain_rainbow` (`max_chain_rainbow` DESC),
  KEY `idx_v3_user_latest_save_data_jack_totalmax_v2` (`jack_totalmax_v2` DESC),
  KEY `idx_v3_user_latest_save_data_ult_combomax` (`ult_combomax` DESC),
  KEY `idx_v3_user_latest_save_data_ult_totalmax_v2` (`ult_totalmax_v2` DESC),
  KEY `idx_v3_user_latest_save_data_sp_use` (`sp_use` DESC),
  KEY `idx_v3_user_latest_save_data_created_at` (`created_at`),
  KEY `idx_v3_latest_hide_record` (`hide_record`),
  KEY `idx_v3_latest_achievements_hide` (`hide_record`,`achievements_count` DESC,`created_at`),
  KEY `idx_v3_latest_jacksp_hide` (`hide_record`,`jacksp_startmax` DESC,`created_at`),
  KEY `idx_v3_latest_golden_hide` (`hide_record`,`golden_palball_get` DESC,`created_at`),
  KEY `idx_v3_latest_cpm_hide` (`hide_record`,`cpm_max` DESC,`created_at`),
  KEY `idx_v3_latest_chain_hide` (`hide_record`,`max_chain_rainbow` DESC,`created_at`),
  KEY `idx_v3_latest_jack_hide` (`hide_record`,`jack_totalmax_v2` DESC,`created_at`),
  KEY `idx_v3_latest_ult_combo_hide` (`hide_record`,`ult_combomax` DESC,`created_at`),
  KEY `idx_v3_latest_ult_total_hide` (`hide_record`,`ult_totalmax_v2` DESC,`created_at`),
  KEY `idx_v3_latest_sp_hide` (`hide_record`,`sp_use` DESC,`created_at`),
  KEY `idx_v3_user_latest_save_data_blackbox_total` (`blackbox_total` DESC),
  KEY `idx_v3_latest_blackbox_total_hide` (`hide_record`,`blackbox_total` DESC,`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

### 5.16 v3_user_latest_save_data_achievements

```sql
CREATE TABLE `v3_user_latest_save_data_achievements` (
  `user_id` varchar(255) NOT NULL,
  `achievement_id` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`user_id`,`achievement_id`),
  KEY `idx_v3_user_latest_achievements_user_id` (`user_id`),
  KEY `idx_v3_achievements_achievement_user` (`achievement_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
```

---

## 6. よく使うクエリ

```sql
-- DB/テーブル情報
SHOW DATABASES;
SHOW TABLES;
SHOW CREATE TABLE table_name;
SHOW INDEX FROM table_name;
DESCRIBE table_name;

-- データ確認
SELECT COUNT(*) FROM table_name;
SELECT * FROM v2_save_data ORDER BY created_at DESC LIMIT 10;
SELECT user_id, COUNT(*) AS record_count FROM v2_save_data GROUP BY user_id;
SELECT * FROM v3_user_latest_save_data ORDER BY created_at DESC LIMIT 10;

-- マイグレーション確認
SELECT * FROM goose_db_version ORDER BY version_id;
SELECT MAX(version_id) AS latest_version FROM goose_db_version;
```

---

## 7. DDL の更新手順メモ

1. 最新のテーブル一覧を取得  
   `docker exec myapp_db mariadb -u root -ppass -N -B app -e 'SHOW TABLES;'`
2. `SHOW CREATE TABLE` を一括でファイルに保存（例: `table_definitions.txt`）  
   ```bash
   tables=$(docker exec myapp_db mariadb -u root -ppass -N -B app -e 'SHOW TABLES;')
   rm -f table_definitions.txt
   for table in $tables; do
     {
       echo "===== $table ====="
       query=$(printf 'SHOW CREATE TABLE `%s`\\G' "$table")
       docker exec myapp_db mariadb -u root -ppass app --execute="$query"
       echo
     } >> table_definitions.txt
   done
   ```
3. 必要であれば `table_definitions_clean.sql` を生成し、このドキュメントのセクション 5 を更新。

---

## 8. メモ
- メイン API は v2 以降を参照する想定で、`v1_game_data` は互換維持のみ。
- `v3_user_latest_*` は集計結果のキャッシュであり、バックフィルが必要な場合はマイグレーション 16〜21 を参照。
- すべて InnoDB かつ utf8mb4 系文字セットで統一。新規テーブルも同方針で作成する。
