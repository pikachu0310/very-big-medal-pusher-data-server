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

#### 1. `game_data` テーブル
```sql
CREATE TABLE `game_data` (
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
```

#### 2. `save_data_v2` テーブル
```sql
CREATE TABLE `save_data_v2` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(255) NOT NULL,
  `legacy` tinyint(4) NOT NULL,
  `version` int(11) NOT NULL,
  `credit` bigint(20) NOT NULL DEFAULT 0,
  `credit_all` bigint(20) NOT NULL DEFAULT 0,
  `medal_in` int(11) NOT NULL DEFAULT 0,
  `medal_get` int(11) NOT NULL DEFAULT 0,
  `ball_get` int(11) NOT NULL DEFAULT 0,
  `ball_chain` int(11) NOT NULL DEFAULT 0,
  `slot_start` int(11) NOT NULL DEFAULT 0,
  `slot_startfev` int(11) NOT NULL DEFAULT 0,
  `slot_hit` int(11) NOT NULL DEFAULT 0,
  `slot_getfev` int(11) NOT NULL DEFAULT 0,
  `sqr_get` int(11) NOT NULL DEFAULT 0,
  `sqr_step` int(11) NOT NULL DEFAULT 0,
  `jack_get` int(11) NOT NULL DEFAULT 0,
  `jack_startmax` int(11) NOT NULL DEFAULT 0,
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
  `buy_shbi` int(11) NOT NULL DEFAULT 0,
  `firstboot` bigint(20) NOT NULL,
  `lastsave` bigint(20) NOT NULL,
  `playtime` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `jack_totalmax_v2` int(11) NOT NULL DEFAULT 0,
  `ult_totalmax_v2` int(11) NOT NULL DEFAULT 0,
  `jacksp_get_all` int(11) NOT NULL DEFAULT 0,
  `jacksp_get_t0` int(11) NOT NULL DEFAULT 0,
  `jacksp_get_t1` int(11) NOT NULL DEFAULT 0,
  `jacksp_get_t2` int(11) NOT NULL DEFAULT 0,
  `jacksp_get_t3` int(11) NOT NULL DEFAULT 0,
  `jacksp_startmax` bigint(20) NOT NULL DEFAULT 0,
  `jacksp_totalmax` bigint(20) NOT NULL DEFAULT 0,
  `palball_get` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `idx_save_data_v2_user_created_at` (`user_id`,`created_at`),
  KEY `idx_save_data_v2_user_playtime` (`user_id`,`playtime`),
  KEY `idx_save_data_v2_user_jacktotal_created` (`user_id`,`jack_totalmax`,`created_at`),
  KEY `idx_save_data_v2_user` (`user_id`),
  KEY `idx_user_jt_created` (`user_id`,`jack_totalmax`,`created_at`),
  KEY `idx_user_id_id` (`user_id`,`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci
```

#### 3. `save_data_v2_achievements` テーブル
```sql
CREATE TABLE `save_data_v2_achievements` (
  `save_id` int(11) NOT NULL,
  `achievement_id` varchar(255) NOT NULL,
  PRIMARY KEY (`save_id`,`achievement_id`),
  KEY `idx_save_data_v2_achieve_save` (`save_id`),
  KEY `idx_save_data_v2_achievements_achievement_id` (`achievement_id`),
  KEY `idx_save_data_v2_achievements_user_achievement` (`save_id`,`achievement_id`),
  CONSTRAINT `save_data_v2_achievements_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `save_data_v2` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci
```

#### 4. `save_data_v2_ball_chain` テーブル
```sql
CREATE TABLE `save_data_v2_ball_chain` (
  `save_id` int(11) NOT NULL,
  `ball_id` varchar(255) NOT NULL,
  `chain_count` int(11) NOT NULL,
  PRIMARY KEY (`save_id`,`ball_id`),
  KEY `idx_save_data_v2_ball_chain_save` (`save_id`),
  KEY `idx_ball_chain_ballid_saveid_count` (`ball_id`,`save_id`,`chain_count`),
  CONSTRAINT `save_data_v2_ball_chain_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `save_data_v2` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci
```

#### 5. `save_data_v2_ball_get` テーブル
```sql
CREATE TABLE `save_data_v2_ball_get` (
  `save_id` int(11) NOT NULL,
  `ball_id` varchar(255) NOT NULL,
  `count` int(11) NOT NULL,
  PRIMARY KEY (`save_id`,`ball_id`),
  KEY `idx_save_data_v2_ball_get_save` (`save_id`),
  CONSTRAINT `save_data_v2_ball_get_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `save_data_v2` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci
```

#### 6. `save_data_v2_medal_get` テーブル
```sql
CREATE TABLE `save_data_v2_medal_get` (
  `save_id` int(11) NOT NULL,
  `medal_id` varchar(255) NOT NULL,
  `count` int(11) NOT NULL,
  PRIMARY KEY (`save_id`,`medal_id`),
  KEY `idx_save_data_v2_medal_get_save` (`save_id`),
  CONSTRAINT `save_data_v2_medal_get_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `save_data_v2` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci
```

#### 7. `save_data_v2_palball_get` テーブル
```sql
CREATE TABLE `save_data_v2_palball_get` (
  `save_id` int(11) NOT NULL,
  `ball_id` varchar(255) NOT NULL,
  `count` int(11) NOT NULL,
  PRIMARY KEY (`save_id`,`ball_id`),
  KEY `idx_save_data_v2_palball_get_save` (`save_id`),
  CONSTRAINT `save_data_v2_palball_get_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `save_data_v2` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci
```

#### 8. `save_data_v2_palball_jp` テーブル
```sql
CREATE TABLE `save_data_v2_palball_jp` (
  `save_id` int(11) NOT NULL,
  `ball_id` varchar(255) NOT NULL,
  `count` int(11) NOT NULL,
  PRIMARY KEY (`save_id`,`ball_id`),
  KEY `idx_save_data_v2_palball_jp_save` (`save_id`),
  CONSTRAINT `save_data_v2_palball_jp_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `save_data_v2` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci
```

#### 9. `save_data_v2_perks` テーブル
```sql
CREATE TABLE `save_data_v2_perks` (
  `save_id` int(11) NOT NULL,
  `perk_id` int(11) NOT NULL,
  `level` int(11) NOT NULL,
  PRIMARY KEY (`save_id`,`perk_id`),
  KEY `idx_save_data_v2_perks_save` (`save_id`),
  CONSTRAINT `save_data_v2_perks_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `save_data_v2` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci
```

#### 10. `save_data_v2_perks_credit` テーブル
```sql
CREATE TABLE `save_data_v2_perks_credit` (
  `save_id` int(11) NOT NULL,
  `perk_id` int(11) NOT NULL,
  `credits` bigint(20) NOT NULL DEFAULT 0,
  PRIMARY KEY (`save_id`,`perk_id`),
  KEY `idx_save_data_v2_perks_credit_save` (`save_id`),
  CONSTRAINT `save_data_v2_perks_credit_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `save_data_v2` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci
```

#### 11. `goose_db_version` テーブル（マイグレーション管理）
```sql
CREATE TABLE `goose_db_version` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `version_id` bigint(20) NOT NULL,
  `is_applied` tinyint(1) NOT NULL,
  `tstamp` timestamp NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
```

### テーブル統計情報

| テーブル名 | 行数 | データサイズ | インデックスサイズ | 照合順序 |
|------------|------|--------------|-------------------|----------|
| game_data | 0 | 16KB | 32KB | utf8mb4_unicode_ci |
| save_data_v2 | 0 | 16KB | 96KB | utf8mb4_uca1400_ai_ci |
| save_data_v2_achievements | 15 | 16KB | 48KB | utf8mb4_uca1400_ai_ci |
| save_data_v2_ball_chain | 7 | 16KB | 32KB | utf8mb4_uca1400_ai_ci |
| save_data_v2_ball_get | 5 | 16KB | 16KB | utf8mb4_uca1400_ai_ci |
| save_data_v2_medal_get | 8 | 16KB | 16KB | utf8mb4_uca1400_ai_ci |
| save_data_v2_palball_get | 0 | 16KB | 16KB | utf8mb4_uca1400_ai_ci |
| save_data_v2_palball_jp | 0 | 16KB | 16KB | utf8mb4_uca1400_ai_ci |
| save_data_v2_perks | 0 | 16KB | 16KB | utf8mb4_uca1400_ai_ci |
| save_data_v2_perks_credit | 0 | 16KB | 16KB | utf8mb4_uca1400_ai_ci |
| goose_db_version | 14 | 16KB | 16KB | utf8mb4_unicode_ci |

### 外部キー制約

| 制約名 | テーブル | カラム | 参照先テーブル | 参照先カラム |
|--------|----------|--------|----------------|--------------|
| save_data_v2_achievements_ibfk_1 | save_data_v2_achievements | save_id | save_data_v2 | id |
| save_data_v2_ball_chain_ibfk_1 | save_data_v2_ball_chain | save_id | save_data_v2 | id |
| save_data_v2_ball_get_ibfk_1 | save_data_v2_ball_get | save_id | save_data_v2 | id |
| save_data_v2_medal_get_ibfk_1 | save_data_v2_medal_get | save_id | save_data_v2 | id |
| save_data_v2_palball_get_ibfk_1 | save_data_v2_palball_get | save_id | save_data_v2 | id |
| save_data_v2_palball_jp_ibfk_1 | save_data_v2_palball_jp | save_id | save_data_v2 | id |
| save_data_v2_perks_ibfk_1 | save_data_v2_perks | save_id | save_data_v2 | id |
| save_data_v2_perks_credit_ibfk_1 | save_data_v2_perks_credit | save_id | save_data_v2 | id |

### インデックス情報

#### game_data テーブル
- `PRIMARY` (id) - 主キー
- `idx_game_data_user_created_at` (user_id, created_at) - 複合インデックス
- `idx_game_data_user_total_play_time` (user_id, total_play_time) - 複合インデックス

#### save_data_v2 テーブル
- `PRIMARY` (id) - 主キー
- `idx_save_data_v2_user` (user_id) - 単一インデックス
- `idx_save_data_v2_user_created_at` (user_id, created_at) - 複合インデックス
- `idx_save_data_v2_user_jacktotal_created` (user_id, jack_totalmax, created_at) - 複合インデックス
- `idx_save_data_v2_user_playtime` (user_id, playtime) - 複合インデックス
- `idx_user_id_id` (user_id, id) - 複合インデックス
- `idx_user_jt_created` (user_id, jack_totalmax, created_at) - 複合インデックス

#### 関連テーブル（save_data_v2_*）
各テーブルには以下のインデックスが設定されています：
- `PRIMARY` (save_id, [その他の主キーカラム]) - 複合主キー
- `idx_[テーブル名]_save` (save_id) - 外部キー用インデックス
- その他、テーブル固有のインデックス

### マイグレーション履歴

現在適用済みのマイグレーション（goose_db_version テーブルより）：
- バージョン 0-15: すべて適用済み
- 最新マイグレーション: バージョン 15（2025-08-29 12:03:21 適用）

### データベースの特徴
- **メイン**: `game_data`（ゲームデータ）、`save_data_v2`（セーブデータ）
- **関連**: `save_data_v2_*` は `save_data_v2` に外部キー
- **管理**: `goose_db_version` でマイグレーション管理
- **エンジン**: InnoDB
- **文字セット**: utf8mb4
- **照合順序**: utf8mb4_unicode_ci / utf8mb4_uca1400_ai_ci
- **外部キー**: すべて ON DELETE CASCADE 設定

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
SELECT * FROM save_data_v2 ORDER BY created_at DESC LIMIT 10;

-- ユーザー別のゲームデータ確認
SELECT user_id, COUNT(*) as record_count FROM game_data GROUP BY user_id;
```

#### マイグレーション確認
```sql
-- マイグレーション履歴確認
SELECT * FROM goose_db_version ORDER BY version_id;

-- 最新のマイグレーション確認
SELECT MAX(version_id) as latest_version FROM goose_db_version;
```