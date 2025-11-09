/*M!999999\- enable the sandbox mode */ 
-- MariaDB dump 10.19-11.7.2-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: app
-- ------------------------------------------------------
-- Server version	11.7.2-MariaDB-ubu2404

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*M!100616 SET @OLD_NOTE_VERBOSITY=@@NOTE_VERBOSITY, NOTE_VERBOSITY=0 */;

--
-- Table structure for table `goose_db_version`
--

DROP TABLE IF EXISTS `goose_db_version`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `goose_db_version` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `version_id` bigint(20) NOT NULL,
  `is_applied` tinyint(1) NOT NULL,
  `tstamp` timestamp NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `v1_game_data`
--

DROP TABLE IF EXISTS `v1_game_data`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
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
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `v2_save_data`
--

DROP TABLE IF EXISTS `v2_save_data`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
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
  `skill_point` int(11) NOT NULL DEFAULT 0,
  `blackbox` int(11) NOT NULL DEFAULT 0,
  `blackbox_total` bigint(20) NOT NULL DEFAULT 0,
  `sp_use` int(11) NOT NULL DEFAULT 0,
  `hide_record` int(11) NOT NULL DEFAULT 0,
  `cpm_max` double NOT NULL DEFAULT 0,
  `palball_get` int(11) NOT NULL DEFAULT 0,
  `pallot_lot_t0` int(11) NOT NULL DEFAULT 0,
  `pallot_lot_t1` int(11) NOT NULL DEFAULT 0,
  `pallot_lot_t2` int(11) NOT NULL DEFAULT 0,
  `pallot_lot_t3` int(11) NOT NULL DEFAULT 0,
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
/*!40101 SET character_set_client = @saved_cs_client */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_uca1400_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`%`*/ /*!50003 TRIGGER update_v3_user_latest_save_data_after_insert
AFTER INSERT ON v2_save_data
FOR EACH ROW
BEGIN
    DECLARE v_achievements_count INT(11) DEFAULT 0;
    DECLARE v_max_chain_rainbow INT(11) DEFAULT 0;
    DECLARE v_golden_palball_get INT(11) DEFAULT 0;

    -- 実績数を計算
    SELECT COUNT(*) INTO v_achievements_count
    FROM v2_save_data_achievements
    WHERE save_id = NEW.id;

    -- 最大レインボーチェイン数を計算
    SELECT COALESCE(MAX(chain_count), 0) INTO v_max_chain_rainbow
    FROM v2_save_data_ball_chain
    WHERE save_id = NEW.id AND ball_id = '3';

    -- golden_palball_getを計算（ball_id = 100の数）
    SELECT COALESCE(SUM(count), 0) INTO v_golden_palball_get
    FROM v2_save_data_palball_get
    WHERE save_id = NEW.id AND ball_id = '100';

    -- v3_user_latest_save_data を更新
    INSERT INTO v3_user_latest_save_data (
        user_id, save_id, version, credit_all, playtime, achievements_count, jacksp_startmax, golden_palball_get,
        cpm_max, max_chain_rainbow, jack_totalmax_v2, ult_combomax, ult_totalmax_v2, blackbox_total, sp_use
    ) VALUES (
        NEW.user_id, NEW.id, NEW.version, NEW.credit_all, NEW.playtime, v_achievements_count, NEW.jacksp_startmax, v_golden_palball_get,
        NEW.cpm_max, v_max_chain_rainbow, NEW.jack_totalmax_v2, NEW.ult_combomax, NEW.ult_totalmax_v2, NEW.blackbox_total, NEW.sp_use
    ) ON DUPLICATE KEY UPDATE
        version = NEW.version,
        credit_all = NEW.credit_all,
        playtime = NEW.playtime,
        save_id = NEW.id,
        achievements_count = v_achievements_count,
        jacksp_startmax = NEW.jacksp_startmax,
        golden_palball_get = v_golden_palball_get,
        cpm_max = NEW.cpm_max,
        max_chain_rainbow = v_max_chain_rainbow,
        jack_totalmax_v2 = NEW.jack_totalmax_v2,
        ult_combomax = NEW.ult_combomax,
        ult_totalmax_v2 = NEW.ult_totalmax_v2,
        blackbox_total = NEW.blackbox_total,
        sp_use = NEW.sp_use,
        updated_at = CURRENT_TIMESTAMP;
END */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;

--
-- Table structure for table `v2_save_data_achievements`
--

DROP TABLE IF EXISTS `v2_save_data_achievements`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `v2_save_data_achievements` (
  `save_id` int(11) NOT NULL,
  `achievement_id` varchar(255) NOT NULL,
  PRIMARY KEY (`save_id`,`achievement_id`),
  KEY `idx_save_data_v2_achieve_save` (`save_id`),
  KEY `idx_save_data_v2_achievements_achievement_id` (`achievement_id`),
  KEY `idx_save_data_v2_achievements_user_achievement` (`save_id`,`achievement_id`),
  CONSTRAINT `v2_save_data_achievements_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_uca1400_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`%`*/ /*!50003 TRIGGER update_v3_achievements_count_after_insert
AFTER INSERT ON v2_save_data_achievements
FOR EACH ROW
BEGIN
    DECLARE v_user_id VARCHAR(255);
    DECLARE v_achievements_count INT(11) DEFAULT 0;
    
    -- ユーザーIDを取得
    SELECT user_id INTO v_user_id FROM v2_save_data WHERE id = NEW.save_id;
    
    -- 実績数を再計算
    SELECT COUNT(*) INTO v_achievements_count
    FROM v2_save_data_achievements
    WHERE save_id = NEW.save_id;
    
    -- v3_user_latest_save_data を更新
    UPDATE v3_user_latest_save_data
    SET achievements_count = v_achievements_count, updated_at = CURRENT_TIMESTAMP
    WHERE user_id = v_user_id;
END */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;

--
-- Table structure for table `v2_save_data_ball_chain`
--

DROP TABLE IF EXISTS `v2_save_data_ball_chain`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
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
/*!40101 SET character_set_client = @saved_cs_client */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_uca1400_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`%`*/ /*!50003 TRIGGER update_v3_max_chain_rainbow_after_insert
AFTER INSERT ON v2_save_data_ball_chain
FOR EACH ROW
BEGIN
    DECLARE v_user_id VARCHAR(255);
    DECLARE v_max_chain_rainbow INT(11) DEFAULT 0;
    
    -- ユーザーIDを取得
    SELECT user_id INTO v_user_id FROM v2_save_data WHERE id = NEW.save_id;
    
    -- レインボーチェインの最大値を再計算
    SELECT COALESCE(MAX(chain_count), 0) INTO v_max_chain_rainbow
    FROM v2_save_data_ball_chain
    WHERE save_id = NEW.save_id AND ball_id = '3';
    
    -- v3_user_latest_save_data を更新
    UPDATE v3_user_latest_save_data
    SET max_chain_rainbow = v_max_chain_rainbow, updated_at = CURRENT_TIMESTAMP
    WHERE user_id = v_user_id;
END */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;

--
-- Table structure for table `v2_save_data_ball_get`
--

DROP TABLE IF EXISTS `v2_save_data_ball_get`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `v2_save_data_ball_get` (
  `save_id` int(11) NOT NULL,
  `ball_id` varchar(255) NOT NULL,
  `count` bigint(20) NOT NULL,
  PRIMARY KEY (`save_id`,`ball_id`),
  KEY `idx_save_data_v2_ball_get_save` (`save_id`),
  CONSTRAINT `v2_save_data_ball_get_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `v2_save_data_medal_get`
--

DROP TABLE IF EXISTS `v2_save_data_medal_get`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `v2_save_data_medal_get` (
  `save_id` int(11) NOT NULL,
  `medal_id` varchar(255) NOT NULL,
  `count` int(11) NOT NULL,
  PRIMARY KEY (`save_id`,`medal_id`),
  KEY `idx_save_data_v2_medal_get_save` (`save_id`),
  CONSTRAINT `v2_save_data_medal_get_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `v2_save_data_palball_get`
--

DROP TABLE IF EXISTS `v2_save_data_palball_get`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `v2_save_data_palball_get` (
  `save_id` int(11) NOT NULL,
  `ball_id` varchar(255) NOT NULL,
  `count` int(11) NOT NULL,
  PRIMARY KEY (`save_id`,`ball_id`),
  KEY `idx_save_data_v2_palball_get_save` (`save_id`),
  CONSTRAINT `v2_save_data_palball_get_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_uca1400_ai_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`%`*/ /*!50003 TRIGGER update_v3_golden_palball_get_after_insert
AFTER INSERT ON v2_save_data_palball_get
FOR EACH ROW
BEGIN
    DECLARE v_user_id VARCHAR(255);
    DECLARE v_golden_palball_get INT(11) DEFAULT 0;
    
    -- ユーザーIDを取得
    SELECT user_id INTO v_user_id FROM v2_save_data WHERE id = NEW.save_id;
    
    -- golden_palball_getを再計算（ball_id = 100の数）
    SELECT COALESCE(SUM(count), 0) INTO v_golden_palball_get
    FROM v2_save_data_palball_get
    WHERE save_id = NEW.save_id AND ball_id = '100';
    
    -- v3_user_latest_save_data を更新
    UPDATE v3_user_latest_save_data
    SET golden_palball_get = v_golden_palball_get, updated_at = CURRENT_TIMESTAMP
    WHERE user_id = v_user_id;
END */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;

--
-- Table structure for table `v2_save_data_palball_jp`
--

DROP TABLE IF EXISTS `v2_save_data_palball_jp`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `v2_save_data_palball_jp` (
  `save_id` int(11) NOT NULL,
  `ball_id` varchar(255) NOT NULL,
  `count` int(11) NOT NULL,
  PRIMARY KEY (`save_id`,`ball_id`),
  KEY `idx_save_data_v2_palball_jp_save` (`save_id`),
  CONSTRAINT `v2_save_data_palball_jp_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `v2_save_data_perks`
--

DROP TABLE IF EXISTS `v2_save_data_perks`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `v2_save_data_perks` (
  `save_id` int(11) NOT NULL,
  `perk_id` int(11) NOT NULL,
  `level` int(11) NOT NULL,
  PRIMARY KEY (`save_id`,`perk_id`),
  KEY `idx_save_data_v2_perks_save` (`save_id`),
  CONSTRAINT `v2_save_data_perks_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `v2_save_data_perks_credit`
--

DROP TABLE IF EXISTS `v2_save_data_perks_credit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `v2_save_data_perks_credit` (
  `save_id` int(11) NOT NULL,
  `perk_id` int(11) NOT NULL,
  `credits` bigint(20) NOT NULL DEFAULT 0,
  PRIMARY KEY (`save_id`,`perk_id`),
  KEY `idx_save_data_v2_perks_credit_save` (`save_id`),
  CONSTRAINT `v2_save_data_perks_credit_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `v2_save_data_totems`
--

DROP TABLE IF EXISTS `v2_save_data_totems`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `v2_save_data_totems` (
  `save_id` int(11) NOT NULL,
  `totem_id` int(11) NOT NULL,
  `level` int(11) NOT NULL,
  PRIMARY KEY (`save_id`,`totem_id`),
  KEY `idx_v2_save_data_totems_save` (`save_id`),
  CONSTRAINT `v2_save_data_totems_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `v2_save_data_totems_credit`
--

DROP TABLE IF EXISTS `v2_save_data_totems_credit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `v2_save_data_totems_credit` (
  `save_id` int(11) NOT NULL,
  `totem_id` int(11) NOT NULL,
  `credits` bigint(20) NOT NULL,
  PRIMARY KEY (`save_id`,`totem_id`),
  KEY `idx_v2_save_data_totems_credit_save` (`save_id`),
  CONSTRAINT `v2_save_data_totems_credit_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `v2_save_data_totems_placement`
--

DROP TABLE IF EXISTS `v2_save_data_totems_placement`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `v2_save_data_totems_placement` (
  `save_id` int(11) NOT NULL,
  `placement_idx` int(11) NOT NULL,
  `totem_id` int(11) NOT NULL,
  PRIMARY KEY (`save_id`,`placement_idx`),
  KEY `idx_v2_save_data_totems_placement_save` (`save_id`),
  CONSTRAINT `v2_save_data_totems_placement_ibfk_1` FOREIGN KEY (`save_id`) REFERENCES `v2_save_data` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `v3_user_latest_save_data`
--

DROP TABLE IF EXISTS `v3_user_latest_save_data`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
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
  `sp_use` int(11) NOT NULL DEFAULT 0,
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
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `v3_user_latest_save_data_achievements`
--

DROP TABLE IF EXISTS `v3_user_latest_save_data_achievements`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8mb4 */;
CREATE TABLE `v3_user_latest_save_data_achievements` (
  `user_id` varchar(255) NOT NULL,
  `achievement_id` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`user_id`,`achievement_id`),
  KEY `idx_v3_user_latest_achievements_user_id` (`user_id`),
  KEY `idx_v3_achievements_achievement_user` (`achievement_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*M!100616 SET NOTE_VERBOSITY=@OLD_NOTE_VERBOSITY */;

-- Dump completed on 2025-11-09  0:18:32
