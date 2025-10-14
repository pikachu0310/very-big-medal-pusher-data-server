# 現状のDB定義と接続方法
## データベース定義の調査結果

### 接続方法

#### Docker Compose経由での接続
```bash
# コンテナに直接接続
docker exec myapp_db mariadb -u root -ppass

# 特定のデータベースに接続
docker exec myapp_db mariadb -u root -ppass app

# コマンドを直接実行
docker exec myapp_db mariadb -u root -ppass -e "SHOW DATABASES;"
```

#### 外部からの接続
```bash
# ホストから接続（ポート3306が公開されている場合）
mariadb -h localhost -P 3306 -u root -ppass

# 特定のデータベースに接続
mariadb -h localhost -P 3306 -u root -ppass app
```

#### Adminer（Web管理ツール）
- **URL**: http://localhost:8081
- **サーバー**: `db`
- **ユーザー名**: `root`
- **パスワード**: `pass`
- **データベース**: `app`

### Docker Compose設定
- **コンテナ名**: `myapp_db`
- **イメージ**: `mariadb:latest`
- **バージョン**: `11.7.2-MariaDB-ubu2404`
- **ポート**: `3306:3306`
- **ルートパスワード**: `pass`
- **デフォルトデータベース**: `app`
- **文字セット**: `utf8mb4`
- **照合順序**: `utf8mb4_unicode_ci`

### 存在するデータベース
1. **app** (メインアプリケーションデータベース)
2. **information_schema** (システム情報)
3. **mysql** (システムデータベース)
4. **performance_schema** (パフォーマンス情報)
5. **sys** (システムビュー)

### appデータベースのテーブル定義

#### 1. `v1_game_data` テーブル (v1 API用・非推奨)
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
  `created_at` timestamp NULL DEFAULT current_timestamp(),
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci
```

#### 2. `v2_save_data` テーブル (メインセーブデータ)
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
  `sqr_get` int(11) NOT NULL DEFAULT 0,
  `sqr_step` bigint(20) NOT NULL DEFAULT 0,
  `jack_get` int(11) NOT NULL DEFAULT 0,
  `jack_startmax` bigint(20) NOT NULL DEFAULT 0,
  `jack_totalmax` int(11) NOT NULL DEFAULT 0,
  `ult_get` int(11) NOT NULL DEFAULT 0,
  `ult_combomax` int(11) NOT NULL DEFAULT 0,
  `ult_totalmax` int(11) NOT NULL DEFAULT 0,
  `rmshbi_get` int(11) NOT NULL DEFAULT 0,
  `bstp_step` int(11) NOT NULL DEFAULT 0,
  `bstp_rwd` int(11) NOT NULL DEFAULT 0,
  `buy_total` int(11) NOT NULL DEFAULT 0,
  `sp_use` int(11) NOT NULL DEFAULT 0,
  `hide_record` int(11) NOT NULL DEFAULT 0,
  `cpm_max` double NOT NULL DEFAULT 0,
  `palball_get` int(11) NOT NULL DEFAULT 0,
  `pallot_lot_t0` int(11) NOT NULL DEFAULT 0,
  `pallot_lot_t1` int(11) NOT NULL DEFAULT 0,
  `pallot_lot_t2` int(11) NOT NULL DEFAULT 0,
  `pallot_lot_t3` int(11) NOT NULL DEFAULT 0,
  `task_cnt` int(11) NOT NULL DEFAULT 0,
  `totem_altars` int(11) NOT NULL DEFAULT 0,              -- Save Version 12
  `totem_altars_credit` bigint(20) NOT NULL DEFAULT 0,     -- Save Version 12
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci
```

#### 3. `v2_save_data_achievements` テーブル (実績データ)
```sql
CREATE TABLE `v2_save_data_achievements` (
  `save_id` int(11) NOT NULL,
  `achievement_id` varchar(255) NOT NULL,
  PRIMARY KEY (`save_id`,`achievement_id`),
  KEY `idx_save_data_v2_achieve_save` (`save_id`),
  KEY `idx_save_data_v2_achievements_achievement_id` (`achievement_id`),
  KEY `idx_save_data_v2_achievements_user_achievement` (`save_id`,`achievement_id`),
  CONSTRAINT `v2_save_data_achievements_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci
```

#### 4. `v2_save_data_ball_chain` テーブル (ボールチェインデータ)
```sql
CREATE TABLE `v2_save_data_ball_chain` (
  `save_id` int(11) NOT NULL,
  `ball_id` varchar(255) NOT NULL,
  `chain_count` int(11) NOT NULL,
  PRIMARY KEY (`save_id`,`ball_id`),
  KEY `idx_save_data_v2_ball_chain_save` (`save_id`),
  KEY `idx_ball_chain_ballid_saveid_count` (`ball_id`,`save_id`,`chain_count`),
  CONSTRAINT `v2_save_data_ball_chain_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci
```

#### 5. `v2_save_data_ball_get` テーブル (ボール取得データ)
```sql
CREATE TABLE `v2_save_data_ball_get` (
  `save_id` int(11) NOT NULL,
  `ball_id` varchar(255) NOT NULL,
  `count` int(11) NOT NULL,
  PRIMARY KEY (`save_id`,`ball_id`),
  KEY `idx_save_data_v2_ball_get_save` (`save_id`),
  CONSTRAINT `v2_save_data_ball_get_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci
```

#### 6. `v2_save_data_medal_get` テーブル (メダル取得データ)
```sql
CREATE TABLE `v2_save_data_medal_get` (
  `save_id` int(11) NOT NULL,
  `medal_id` varchar(255) NOT NULL,
  `count` int(11) NOT NULL,
  PRIMARY KEY (`save_id`,`medal_id`),
  KEY `idx_save_data_v2_medal_get_save` (`save_id`),
  CONSTRAINT `v2_save_data_medal_get_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci
```

#### 7. `v2_save_data_palball_get` テーブル (パレッタボール取得データ)
```sql
CREATE TABLE `v2_save_data_palball_get` (
  `save_id` int(11) NOT NULL,
  `ball_id` varchar(255) NOT NULL,
  `count` int(11) NOT NULL,
  PRIMARY KEY (`save_id`,`ball_id`),
  KEY `idx_save_data_v2_palball_get_save` (`save_id`),
  CONSTRAINT `v2_save_data_palball_get_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci
```

#### 8. `v2_save_data_palball_jp` テーブル (パレッタボールジャックポットデータ)
```sql
CREATE TABLE `v2_save_data_palball_jp` (
  `save_id` int(11) NOT NULL,
  `ball_id` varchar(255) NOT NULL,
  `count` int(11) NOT NULL,
  PRIMARY KEY (`save_id`,`ball_id`),
  KEY `idx_save_data_v2_palball_jp_save` (`save_id`),
  CONSTRAINT `v2_save_data_palball_jp_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci
```

#### 9. `v2_save_data_perks` テーブル (パークレベルデータ)
```sql
CREATE TABLE `v2_save_data_perks` (
  `save_id` int(11) NOT NULL,
  `perk_id` int(11) NOT NULL,
  `level` int(11) NOT NULL,
  PRIMARY KEY (`save_id`,`perk_id`),
  KEY `idx_save_data_v2_perks_save` (`save_id`),
  CONSTRAINT `v2_save_data_perks_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci
```

#### 10. `v2_save_data_perks_credit` テーブル (パーク消費クレジットデータ)
```sql
CREATE TABLE `v2_save_data_perks_credit` (
  `save_id` int(11) NOT NULL,
  `perk_id` int(11) NOT NULL,
  `credits` bigint(20) NOT NULL DEFAULT 0,
  PRIMARY KEY (`save_id`,`perk_id`),
  KEY `idx_save_data_v2_perks_credit_save` (`save_id`),
  CONSTRAINT `v2_save_data_perks_credit_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci
```

#### 11. `v2_save_data_totems` テーブル (トーテムレベルデータ・Save Version 12)
```sql
CREATE TABLE `v2_save_data_totems` (
  `save_id` int(11) NOT NULL,
  `totem_id` int(11) NOT NULL,
  `level` int(11) NOT NULL,
  PRIMARY KEY (`save_id`,`totem_id`),
  KEY `idx_v2_save_data_totems_save` (`save_id`),
  CONSTRAINT `v2_save_data_totems_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci
```

#### 12. `v2_save_data_totems_credit` テーブル (トーテム消費クレジットデータ・Save Version 12)
```sql
CREATE TABLE `v2_save_data_totems_credit` (
  `save_id` int(11) NOT NULL,
  `totem_id` int(11) NOT NULL,
  `credits` bigint(20) NOT NULL,
  PRIMARY KEY (`save_id`,`totem_id`),
  KEY `idx_v2_save_data_totems_credit_save` (`save_id`),
  CONSTRAINT `v2_save_data_totems_credit_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci
```

#### 13. `v2_save_data_totems_placement` テーブル (トーテム配置データ・Save Version 12)
```sql
CREATE TABLE `v2_save_data_totems_placement` (
  `save_id` int(11) NOT NULL,
  `placement_idx` int(11) NOT NULL,
  `totem_id` int(11) NOT NULL,
  PRIMARY KEY (`save_id`,`placement_idx`),
  KEY `idx_v2_save_data_totems_placement_save` (`save_id`),
  CONSTRAINT `v2_save_data_totems_placement_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci
```

#### 14. `v3_user_latest_save_data` テーブル (v3/v4 API用最新セーブデータキャッシュ)
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
  `sp_use` int(11) NOT NULL DEFAULT 0,
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
  KEY `idx_v3_user_latest_save_data_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
```

#### 15. `v3_user_latest_save_data_achievements` テーブル (v3/v4 API用最新実績データキャッシュ)
```sql
CREATE TABLE `v3_user_latest_save_data_achievements` (
  `user_id` varchar(255) NOT NULL,
  `achievement_id` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`user_id`,`achievement_id`),
  KEY `idx_v3_user_latest_achievements_user_id` (`user_id`),
  KEY `idx_v3_achievements_achievement_user` (`achievement_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
```

#### 16. `goose_db_version` テーブル（マイグレーション管理）
```sql
CREATE TABLE `goose_db_version` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `version_id` bigint(20) NOT NULL,
  `is_applied` tinyint(1) NOT NULL,
  `tstamp` timestamp NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci
```

### テーブル統計情報 (2025-10-14現在)

| テーブル名 | 行数 | データサイズ | インデックスサイズ | 照合順序 |
|------------|------|--------------|-------------------|----------|
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

### 外部キー制約

| 制約名 | テーブル | カラム | 参照先テーブル | 参照先カラム |
|--------|----------|--------|----------------|--------------|
| v2_save_data_achievements_ibfk_1 | v2_save_data_achievements | save_id | v2_save_data | id |
| v2_save_data_ball_chain_ibfk_1 | v2_save_data_ball_chain | save_id | v2_save_data | id |
| v2_save_data_ball_get_ibfk_1 | v2_save_data_ball_get | save_id | v2_save_data | id |
| v2_save_data_medal_get_ibfk_1 | v2_save_data_medal_get | save_id | v2_save_data | id |
| v2_save_data_palball_get_ibfk_1 | v2_save_data_palball_get | save_id | v2_save_data | id |
| v2_save_data_palball_jp_ibfk_1 | v2_save_data_palball_jp | save_id | v2_save_data | id |
| v2_save_data_perks_ibfk_1 | v2_save_data_perks | save_id | v2_save_data | id |
| v2_save_data_perks_credit_ibfk_1 | v2_save_data_perks_credit | save_id | v2_save_data | id |
| v2_save_data_totems_ibfk_1 | v2_save_data_totems | save_id | v2_save_data | id |
| v2_save_data_totems_credit_ibfk_1 | v2_save_data_totems_credit | save_id | v2_save_data | id |
| v2_save_data_totems_placement_ibfk_1 | v2_save_data_totems_placement | save_id | v2_save_data | id |

### マイグレーション履歴

現在適用済みのマイグレーション（goose_db_version テーブルより）：
- バージョン 0-22: すべて適用済み
- 最新マイグレーション: バージョン 22（2025-10-14 11:58:22 適用）

### データベースの特徴
- **メイン**: `v1_game_data`（v1 API用・非推奨）、`v2_save_data`（v2/v3/v4 API用）
- **関連**: `v2_save_data_*` は `v2_save_data` に外部キー
- **キャッシュ**: `v3_user_latest_save_data*` は v3/v4 API用の高速化テーブル
- **管理**: `goose_db_version` でマイグレーション管理
- **エンジン**: InnoDB
- **文字セット**: utf8mb4
- **照合順序**: utf8mb4_unicode_ci / utf8mb4_uca1400_ai_ci
- **外部キー**: すべて ON DELETE CASCADE 設定

### Save Version 12 の追加要素（2025-10-14）
- `v2_save_data.totem_altars` - 解放済みの台座の数
- `v2_save_data.totem_altars_credit` - 台座解放に消費したクレジット
- `v2_save_data_totems` - トーテムレベルデータ
- `v2_save_data_totems_credit` - トーテムごとの消費クレジット
- `v2_save_data_totems_placement` - セット中のトーテムID

### よく使用するクエリ例

#### データベース情報確認
```sql
-- 全データベース一覧
SHOW DATABASES;

-- テーブル一覧
SHOW TABLES;

-- テーブル構造確認
DESCRIBE table_name;
SHOW CREATE TABLE table_name;

-- インデックス確認
SHOW INDEX FROM table_name;
```

#### データ確認
```sql
-- レコード数確認
SELECT COUNT(*) FROM table_name;

-- 最新のセーブデータ確認
SELECT * FROM v2_save_data ORDER BY created_at DESC LIMIT 10;

-- ユーザー別のゲームデータ確認
SELECT user_id, COUNT(*) as record_count FROM v2_save_data GROUP BY user_id;

-- v3最新セーブデータ確認
SELECT * FROM v3_user_latest_save_data ORDER BY created_at DESC LIMIT 10;
```

#### マイグレーション確認
```sql
-- マイグレーション履歴確認
SELECT * FROM goose_db_version ORDER BY version_id;

-- 最新のマイグレーション確認
SELECT MAX(version_id) as latest_version FROM goose_db_version;
```

#### トーテムデータ確認（Save Version 12）
```sql
-- トーテムレベル確認
SELECT s.user_id, t.totem_id, t.level
FROM v2_save_data s
JOIN v2_save_data_totems t ON s.id = t.save_id
WHERE s.user_id = 'your_user_id'
ORDER BY s.created_at DESC LIMIT 1;

-- トーテム配置確認
SELECT s.user_id, tp.placement_idx, tp.totem_id
FROM v2_save_data s
JOIN v2_save_data_totems_placement tp ON s.id = tp.save_id
WHERE s.user_id = 'your_user_id'
ORDER BY s.created_at DESC, tp.placement_idx;
```
