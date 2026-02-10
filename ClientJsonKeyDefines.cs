/// <summary>
/// JSON キー定義用クラス
/// </summary>
public static class JsonKeyDefines
{
    #region JSON キー定義 - セーブ値

    /// <summary>
    /// [JSON キー] レガシーフラグ (古いデータから移行したものか)
    /// </summary>
    /// <remarks>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string Legacy = "legacy";

    /// <summary>
    /// [JSON キー] セーブデータバージョン
    /// </summary>
    /// <remarks>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string Version = "version";

    /// <summary>
    /// [JSON キー] プレイ時間総計
    /// </summary>
    /// <remarks>
    /// データ型 - JSON: [Double], ゲーム内: [Long]
    /// </remarks>
    public const string PlayTime = "playtime";

    /// <summary>
    /// [JSON キー] 所持クレジット
    /// </summary>
    /// <remarks>
    /// データ型 - JSON: [String], ゲーム内: [Long]
    /// </remarks>
    public const string Credit = "credit";

    /// <summary>
    /// [JSON キー] 総獲得クレジット
    /// </summary>
    /// <remarks>
    /// データ型 - JSON: [String], ゲーム内: [Long]
    /// </remarks>
    public const string CreditTotal = "credit_all";

    /// <summary>
    /// [JSON キー] 投入したメダルの数
    /// </summary>
    /// <remarks>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string MedalInsert = "medal_in";

    /// <summary>
    /// [JSON キー] 獲得したメダルの数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 5<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Long]
    /// </remarks>
    public const string MedalGet = "medal_get";

    /// <summary>
    /// [JSON キー] シャルベボール獲得回数
    /// </summary>
    /// <remarks>
    /// データ型 - JSON: [Double], ゲーム内: [Long]
    /// </remarks>
    public const string BallGet = "ball_get";

    /// <summary>
    /// [JSON キー] シャルベボール最大チェイン数
    /// </summary>
    /// <remarks>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string BallChainMax = "ball_chain";

    /// <summary>
    /// [JSON キー] ルーレット始動回数
    /// </summary>
    /// <remarks>
    /// データ型 - JSON: [Double], ゲーム内: [Long]
    /// </remarks>
    public const string SlotStart = "slot_start";

    /// <summary>
    /// [JSON キー] フィーバー付きルーレット始動回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 5<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Long]
    /// </remarks>
    public const string SlotStartFever = "slot_startfev";

    /// <summary>
    /// [JSON キー] ルーレット勝利回数
    /// </summary>
    /// <remarks>
    /// データ型 - JSON: [Double], ゲーム内: [Long]
    /// </remarks>
    public const string SlotHit = "slot_hit";

    /// <summary>
    /// [JSON キー] ルーレットフィーバー開始数
    /// </summary>
    /// <remarks>
    /// データ型 - JSON: [Double], ゲーム内: [Long]
    /// </remarks>
    public const string SlotFever = "slot_getfev";

    /// <summary>
    /// [JSON キー] すごろくを進行させた回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 5<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Long]
    /// </remarks>
    public const string SugorokuGet = "sqr_get";

    /// <summary>
    /// [JSON キー] すごろくで進んだマス数
    /// </summary>
    /// <remarks>
    /// データ型 - JSON: [Double], ゲーム内: [Long]
    /// </remarks>
    public const string SugorokuStep = "sqr_step";

    /// <summary>
    /// [JSON キー] ジャックポット獲得回数
    /// </summary>
    /// <remarks>
    /// データ型 - JSON: [Double], ゲーム内: [Long]
    /// </remarks>
    public const string JackpotGet = "jack_get";

    /// <summary>
    /// [JSON キー] ジャックポット最大スタート値
    /// </summary>
    /// <remarks>
    /// データ型 - JSON: [Double], ゲーム内: [Long]
    /// </remarks>
    public const string JackpotStartMax = "jack_startmax";

    /// <summary>
    /// [JSON キー] ジャックポット最大値 (v1)
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 5<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string JackpotTotalMax = "jack_totalmax";

    /// <summary>
    /// [JSON キー] ULTIMATE MODE 発生回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 5<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string UltimateGet = "ult_get";

    /// <summary>
    /// [JSON キー] ULTIMATE MODE 最大コンボ回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 5<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string UltimateComboMax = "ult_combomax";

    /// <summary>
    /// [JSON キー] ULTIMATE MODE 最終結果最大値 (v1)
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 5<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string UltimateTotalMax = "ult_totalmax";

    /// <summary>
    /// [JSON キー] お部屋シャルベ獲得回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 5<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string RoomkeeperGet = "rmshbi_get";

    /// <summary>
    /// [JSON キー] ボーナスステップ進行回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 5<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Long]
    /// </remarks>
    public const string BonusStepStep = "bstp_step";

    /// <summary>
    /// [JSON キー] ボーナスステップリワード回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 5<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Long]
    /// </remarks>
    public const string BonusStepReward = "bstp_rwd";

    /// <summary>
    /// [JSON キー] ショップの購入ボタンを押した回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 7<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string BuyTotalCount = "buy_total";

    /// <summary>
    /// [JSON キー] ショップでシャルベを購入した回数
    /// </summary>
    /// <remarks>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string BuySherbi = "buy_shbi";

    /// <summary>
    /// [JSON キー] スキルポイント消費数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 6<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Long]
    /// </remarks>
    public const string SkillPointUsed = "sp_use";

    /// <summary>
    /// [JSON キー] セーブデータの作成日時 (Long で持ちたいので文字列)<br/>
    /// `DateTimeOffset.Now.ToUnixTimeSeconds().ToString()` で保存
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 5<br/>
    /// データ型 - JSON: [String], ゲーム内: [Long]
    /// </remarks>
    public const string StrDateFirstBoot = "firstboot";

    /// <summary>
    /// [JSON キー] 最終セーブ日時 (Long で持ちたいので文字列)<br/>
    /// `DateTimeOffset.Now.ToUnixTimeSeconds().ToString()` で保存
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 5<br/>
    /// データ型 - JSON: [String], ゲーム内: [Long]
    /// </remarks>
    public const string StrDateLastSave = "lastsave";

    /// <summary>
    /// [JSON キー] [DataDictionary] 各メダル ID ごとの獲得数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 5<br/>
    /// データ型 - JSON: [Dictionary], ゲーム内: [Dictionary]
    /// </remarks>
    public const string DictMedalGet = "dc_medal_get";

    /// <summary>
    /// [JSON キー] [DataDictionary] 各シャルベボール ID ごとの獲得数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 5<br/>
    /// データ型 - JSON: [Dictionary], ゲーム内: [Dictionary]
    /// </remarks>
    public const string DictBallGet = "dc_ball_get";

    /// <summary>
    /// [JSON キー] [DataDictionary] 各シャルベボール ID ごとの最大チェイン数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 5<br/>
    /// データ型 - JSON: [Dictionary], ゲーム内: [Dictionary]
    /// </remarks>
    public const string DictBallChain = "dc_ball_chain";

    /// <summary>
    /// [JSON キー] [DataList] 解除したアチーブメント ID のリスト
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 5<br/>
    /// データ型 - JSON: [List], ゲーム内: [List]
    /// </remarks>
    public const string ListAchieveUnlock = "l_achieve";

    #endregion

    #region JSON キー定義 - セーブ v9

    /// <summary>
    /// [JSON キー] [int] レコード非公開フラグ
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 9<br/>
    /// データ型 - JSON: [Int], ゲーム内: [Int]
    /// </remarks>
    public const string HideRecord = "hide_record";

    /// <summary>
    /// [JSON キー] [double] 記録された最大 CPM (毎分シャルベクレジット)
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 9<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Double]
    /// </remarks>
    public const string CpMMax = "cpm_max";

    #endregion

    #region JSON キー定義 - セーブ v10

    /// <summary>
    /// [JSON キー] ジャックポット最大値 (v2)
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 10<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Long]
    /// </remarks>
    public const string JackpotTotalMaxV2 = "jack_totalmax_v2";

    /// <summary>
    /// [JSON キー] ULTIMATE MODE 最終結果最大値 (v2)
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 10<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Long]
    /// </remarks>
    public const string UltimateTotalMaxV2 = "ult_totalmax_v2";

    /// <summary>
    /// [JSON キー] パレッタボール獲得回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 10<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string PalettaBallGet = "palball_get";

    /// <summary>
    /// [JSON キー] パレッタ抽選機 Tier 0 挑戦回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 10<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string PalettaLotteryAttemptTier0 = "pallot_lot_t0";

    /// <summary>
    /// [JSON キー] パレッタ抽選機 Tier 1 挑戦回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 10<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string PalettaLotteryAttemptTier1 = "pallot_lot_t1";

    /// <summary>
    /// [JSON キー] パレッタ抽選機 Tier 2 挑戦回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 10<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string PalettaLotteryAttemptTier2 = "pallot_lot_t2";

    /// <summary>
    /// [JSON キー] パレッタ抽選機 Tier 3 挑戦回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 10<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string PalettaLotteryAttemptTier3 = "pallot_lot_t3";

    /// <summary>
    /// [JSON キー] スーパージャックポット (All Tier) 獲得回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 10<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string JackpotSuperGetTotal = "jacksp_get_all";

    /// <summary>
    /// [JSON キー] スーパージャックポット (Tier 0) 獲得回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 10<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string JackpotSuperGetTier0 = "jacksp_get_t0";

    /// <summary>
    /// [JSON キー] スーパージャックポット (Tier 1) 獲得回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 10<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string JackpotSuperGetTier1 = "jacksp_get_t1";

    /// <summary>
    /// [JSON キー] スーパージャックポット (Tier 2) 獲得回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 10<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string JackpotSuperGetTier2 = "jacksp_get_t2";

    /// <summary>
    /// [JSON キー] スーパージャックポット (Tier 3) 獲得回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 10<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string JackpotSuperGetTier3 = "jacksp_get_t3";

    /// <summary>
    /// [JSON キー] スーパージャックポット最大スタート値
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 10<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Long]
    /// </remarks>
    public const string JackpotSuperStartMax = "jacksp_startmax";

    /// <summary>
    /// [JSON キー] スーパージャックポット最大値
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 10<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Long]
    /// </remarks>
    public const string JackpotSuperTotalMax = "jacksp_totalmax";

    /// <summary>
    /// [JSON キー] タスク完了回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 10<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string TaskCompleteCount = "task_cnt";

    /// <summary>
    /// [JSON キー] [DataDictionary] 各パレッタボール ID ごとの獲得数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 10<br/>
    /// データ型 - JSON: [Dictionary], ゲーム内: [Dictionary] (Int)
    /// </remarks>
    public const string DictPalettaBallGet = "dc_palball_get";

    /// <summary>
    /// [JSON キー] [DataDictionary] 各パレッタボール ID ごとのJACKPOT獲得数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 10<br/>
    /// データ型 - JSON: [Dictionary], ゲーム内: [Dictionary] (Int)
    /// </remarks>
    public const string DictPalettaBallJackpot = "dc_palball_jp";

    /// <summary>
    /// [JSON キー] [DataList] パークレベル
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 10<br/>
    /// データ型 - JSON: [List], ゲーム内: [List] (Int)
    /// </remarks>
    public const string ListPerkLevels = "l_perks";

    /// <summary>
    /// [JSON キー] [DataList] パークごとの消費したクレジット数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 10<br/>
    /// データ型 - JSON: [List], ゲーム内: [List] (Long)
    /// </remarks>
    public const string ListPerkUsedCredits = "l_perks_credit";

    #endregion

    #region JSON キー定義 - セーブ v12

    /// <summary>
    /// [JSON キー] [DataList] 解放済みの台座の数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 12<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string TotemAltarUnlockCount = "totem_altars";

    /// <summary>
    /// [JSON キー] [DataList] 台座解放に消費したクレジット (ロード時の整合性確認用)
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 12<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Long]
    /// </remarks>
    public const string TotemAltarUnlockUsedCredits = "totem_altars_credit";

    /// <summary>
    /// [JSON キー] [DataList] トーテムレベル
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 12<br/>
    /// データ型 - JSON: [List], ゲーム内: [List] (Int)
    /// </remarks>
    public const string ListTotemLevels = "l_totems";

    /// <summary>
    /// [JSON キー] [DataList] トーテムごとの消費したクレジット数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 12<br/>
    /// データ型 - JSON: [List], ゲーム内: [List] (Long)
    /// </remarks>
    public const string ListTotemUsedCredits = "l_totems_credit";

    /// <summary>
    /// [JSON キー] [DataList] セット中のトーテムIDのリスト
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 12<br/>
    /// データ型 - JSON: [List], ゲーム内: [List] (int)
    /// </remarks>
    public const string ListTotemPlacements = "l_totems_set";

    #endregion

    #region JSON キー定義 - セーブ v14

    /// <summary>
    /// [JSON キー] スキルポイント
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 14<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Long]
    /// </remarks>
    public const string SkillPoint = "sp";

    /// <summary>
    /// [Temp] [JSON キー] BlackBox（仮名）所持数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 14<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string BlackBox = "bbox";

    /// <summary>
    /// [Temp] [JSON キー] BlackBox（仮名）総獲得数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 14<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string BlackBoxTotal = "bbox_all";

    #endregion

    #region JSON キー定義 - セーブ v16

    /// <summary>
    /// [JSON キー] パレッタ抽選機 Tier 4 挑戦回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 10<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string PalettaLotteryAttemptTier4 = "pallot_lot_t4";

    /// <summary>
    /// [JSON キー] スーパージャックポット (Tier 4) 獲得回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 16<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string JackpotSuperGetTier4 = "jacksp_get_t4";

    #endregion

    #region JSON キー定義 - セーブ v19

    /// <summary>
    /// [JSON キー] フェレッタボール獲得回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 19<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string FerrettaBallGet = "ferball_get";

    /// <summary>
    /// [JSON キー] フェレッタ抽選機挑戦回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 19<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string FerrettaLotteryAttempt = "ferlot_lot";

    /// <summary>
    /// [JSON キー] フェレッタジャックポット獲得数 (任意) 獲得回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 19<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string JackpotFerrettaGetTotal = "jackfr_get_all";

    /// <summary>
    /// [JSON キー] フェレッタジャックポット獲得数 (シングル) 獲得回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 19<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string JackpotFerrettaGetTier0 = "jackfr_get_t0";

    /// <summary>
    /// [JSON キー] フェレッタジャックポット獲得数 (ダブル) 獲得回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 19<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string JackpotFerrettaGetTier1 = "jackfr_get_t1";

    /// <summary>
    /// [JSON キー] フェレッタジャックポット獲得数 (マッシブ) 獲得回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 19<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string JackpotFerrettaGetTier2 = "jackfr_get_t2";

    /// <summary>
    /// [JSON キー] フェレッタジャックポット獲得数 (ヘブン) 獲得回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 19<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string JackpotFerrettaGetTier3 = "jackfr_get_t3";

    /// <summary>
    /// [JSON キー] フェレッタジャックポット獲得数 (ギャラクシー) 獲得回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 19<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string JackpotFerrettaGetTier4 = "jackfr_get_t4";

    /// <summary>
    /// [JSON キー] フェレッタジャックポット最大スタート値
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 19<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Long]
    /// </remarks>
    public const string JackpotFerrettaStartMax = "jackfr_startmax";

    /// <summary>
    /// [JSON キー] フェレッタジャックポット最終結果最大値
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 19<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Long]
    /// </remarks>
    public const string JackpotFerrettaTotalMax = "jackfr_totalmax";

    /// <summary>
    /// [JSON キー] フェレッタチャンス結果が HIT になった回数 (アイテム効果は含まない)
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 19<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string FerrettaLotteryHit = "ferlot_hit";

    /// <summary>
    /// [JSON キー] フェレッタチャンス結果が LOSE になった回数 (アイテム効果は含まない)
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 19<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string FerrettaLotteryLose = "ferlot_lose";

    /// <summary>
    /// [JSON キー] フェレッタチャンス結果が CHANCE になった回数 (アイテム効果は含まない)
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 19<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string FerrettaLotteryChance = "ferlot_chance";

    /// <summary>
    /// [JSON キー] フェレッタチャンスで獲得したマス数の合計 (アイテム含む、保証埋めしたマスは含まない)
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 19<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string FerrettaLotteryActives = "ferlot_act";

    /// <summary>
    /// [JSON キー] フェレッタチャンスで獲得したライン数の合計
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 19<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string FerrettaLotteryLines = "ferlot_lines";

    /// <summary>
    /// [JSON キー] 黒箱ショップ合計利用回数
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 19<br/>
    /// データ型 - JSON: [Double], ゲーム内: [Int]
    /// </remarks>
    public const string BlackBoxShopUsed = "bbox_shop";

    /// <summary>
    /// [JSON キー] 黒箱ショップ利用回数の Dictionary (キーはアイテム ID)
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 19<br/>
    /// データ型 - JSON: [Dictionary], ゲーム内: [Dictionary] (string key, double value)
    /// </remarks>
    public const string DictBlackBoxShopUsed = "dc_bbox_shop";

    /// <summary>
    /// [JSON キー] フェレッタチャンス抽選で獲得したアイテムの Dictionary (キーはアイテム ID)
    /// </summary>
    /// <remarks>
    /// 追加した Save Version: 19<br/>
    /// データ型 - JSON: [Dictionary], ゲーム内: [Dictionary] (string key, double value)
    /// </remarks>
    public const string DictFerrettaLotteryItem = "dc_ferlot_item";

    #endregion

    #region JSON キー定義 - 一時セーブ

    /// <summary>
    /// [Temp] [JSON キー] スキルポイント
    /// </summary>
    [System.Obsolete("Moved to JsonKeyDefines.SkillPoint")]
    public const string TempSkillPoint = "sp";

    /// <summary>
    /// [Temp] [JSON キー] スキルメダル出現カウンター
    /// </summary>
    public const string TempChargeMedalSkill = "chg_mdl_sp";

    /// <summary>
    /// [Temp] [JSON キー] ボール出現カウンター
    /// </summary>
    public const string TempChargeItemBall = "chg_ball";

    /// <summary>
    /// [Temp] [JSON キー] Ultimate 発生カウンター
    /// </summary>
    public const string TempChargeUltimate = "chg_ult";

    /// <summary>
    /// [Temp] [JSON キー] ルーレットのボーナスステップ現在値
    /// </summary>
    public const string TempRouletteBonusStep = "slot_bstp";

    /// <summary>
    /// [Temp] [JSON キー] すごろくの現在周回数
    /// </summary>
    public const string TempSugorokuLoop = "sqr_loop";

    /// <summary>
    /// [Temp] [JSON キー] すごろくのクリームストック数
    /// </summary>
    public const string TempSugorokuCream = "sqr_cream";

    /// <summary>
    /// [Temp] [JSON キー] パレッタ抽選機メダル投入カウンター値
    /// </summary>
    public const string TempSpChusenMedalInsert = "pallot_mdl";

    /// <summary>
    /// [Temp] [JSON キー] パレッタ抽選機ゴールデンボール率ボーナス値
    /// </summary>
    public const string TempSpChusenGoldenProbBonus = "pallot_gpb";

    /// <summary>
    /// [Temp] [JSON キー] パレッタ抽選機プログレッシブカウンター Tier 0 上昇値
    /// </summary>
    public const string TempSpChusenProgCounterTier0 = "pallot_pc0";

    /// <summary>
    /// [Temp] [JSON キー] パレッタ抽選機プログレッシブカウンター Tier 1 上昇値
    /// </summary>
    public const string TempSpChusenProgCounterTier1 = "pallot_pc1";

    /// <summary>
    /// [Temp] [JSON キー] パレッタ抽選機プログレッシブカウンター Tier 2 上昇値
    /// </summary>
    public const string TempSpChusenProgCounterTier2 = "pallot_pc2";

    /// <summary>
    /// [Temp] [JSON キー] パレッタ抽選機プログレッシブカウンター Tier 3 上昇値
    /// </summary>
    public const string TempSpChusenProgCounterTier3 = "pallot_pc3";

    /// <summary>
    /// [Temp] [JSON キー] パレッタ抽選機プログレッシブカウンター Tier 4 上昇値
    /// </summary>
    public const string TempSpChusenProgCounterTier4 = "pallot_pc4";

    /// <summary>
    /// [Temp] [JSON キー] BlackBox（仮名）所有数
    /// </summary>
    [System.Obsolete("Moved to JsonKeyDefines.BlackBox")]
    public const string TempBlackBoxCredits = "blackbox_credits";

    /// <summary>
    /// [Temp] [JSON キー] BlackBox（仮名）獲得数
    /// </summary>
    [System.Obsolete("Moved to JsonKeyDefines.BlackBox")]
    public const string TempGetBlackBox = "blackbox_credits";

    /// <summary>
    /// [Temp] [JSON キー] デイリータスクの日付情報
    /// </summary>
    public const string TempTaskDay = "task_day";

    /// <summary>
    /// [Temp] [JSON キー] デイリータスクのリスト
    /// </summary>
    public const string TempTaskList = "task_list";

    /// <summary>
    /// [Temp] フェレッタチャンスのビンゴ状態
    /// </summary>
    public const string TempFerrettaBingoState = "ferlot_bingo";

    /// <summary>
    /// [Temp] フェレッタチャンスのプログレッシブカウンター上昇値
    /// </summary>
    public const string TempFerrettaProgCounter = "ferlot_pc";

    /// <summary>
    /// [Temp] フェレッタボール出現カウンター
    /// </summary>
    public const string TempFerrettaSummonCount = "ferlot_nx";

    /// <summary>
    /// [Temp] フェレッタアイテムが抽選経由で獲得できているかのフラグ
    /// </summary>
    public const string TempIsItemFromLottery = "ferlot_item_ch";

    /// <summary>
    /// [Temp] フェレッタアイテム状態
    /// </summary>
    public const string TempFerrettaBingoItemState = "ferlot_item";

    #endregion
}
