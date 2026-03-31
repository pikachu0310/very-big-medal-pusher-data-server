import { lazy, Suspense, useEffect, useMemo, useRef, useState, type CSSProperties } from 'react';
import {
  IconAlertTriangle,
  IconBook2,
  IconBolt,
  IconBomb,
  IconBrandDiscord,
  IconBrandGithub,
  IconCalendarSmile,
  IconConfetti,
  IconDeviceGamepad2,
  IconRefresh,
  IconRotateClockwise2,
  IconSparkles,
  IconWorld
} from '@tabler/icons-react';
import { MmpOutlineButton, MmpPrimaryButton } from '../components/MmpButton';

const DeferredSections = lazy(() => import('../components/DeferredSections'));

type CoinParticle = {
  id: number;
  x: number;
  drift: number;
  duration: number;
  delay: number;
  emoji: string;
  size: number;
  rotate: number;
};

type FakeAlert = {
  id: number;
  text: string;
  tone: 'info' | 'warn' | 'critical';
};

const rouletteFaces = ['7', '🍒', '🎰', '💥', '🪙', '🐟', '🍣', '👑'];
const coinEmojis = ['🪙', '🥇', '🎖️', '💿', '🍩'];
const fortunePool = [
  '大吉: 今日のあなたは1クリックで777メダル分の顔をします。',
  '中吉: 2度寝してもメダルは逃げません。たぶん。',
  '小吉: レバーを引く前に深呼吸。深呼吸だけで終わる可能性あり。',
  '末吉: 画面の端にいるカニを信じてください。',
  '凶: そのガチャ、演出だけ豪華です。',
  '超大吉: 今日だけUIのバグが仕様です。'
];
const fakeHeadlines = [
  '運営声明: メダルが重すぎてサーバーが床を抜いた件について',
  '速報: 伝説のプッシャー職人、ボタンを押さずに優勝',
  '検証: 連打力は筋トレになるのか、編集部が24時間押し続けた結果',
  '悲報: 公式マスコット、メンテ画面を勝手に量産',
  '特報: でかプ史上初「押すほど軽くなるUI」実装へ'
];
const fakeTickerMessages = [
  '速報: 今日だけメダル税 0%。明日は未定。',
  '注意: 赤ボタンは押した分だけ赤くなります。',
  '運営からのお知らせ: 仕様書は今日だけ俳句形式です。',
  '臨時告知: クリックしすぎると画面に友情が芽生えます。',
  '豆知識: 連打は哲学、放置は美学。'
];
const fakeAdMessages = [
  '期間限定: メダル1枚でメダル1枚が当たる激アツ抽選会!',
  'あなたへの特別提案: 1日3回の深呼吸でジャックポット体質へ。',
  'この広告を閉じると、閉じた事実だけが残ります。',
  '新機能: ボタンを見つめるとボタンも見つめ返します。'
];
const fakeAlertPool: { text: string; tone: FakeAlert['tone'] }[] = [
  { text: '偽速報: メダルが自我を持ったため会議中です。', tone: 'warn' },
  { text: '業務連絡: 開発者は現在カフェインで稼働しています。', tone: 'info' },
  { text: '緊急: 押しすぎ検知。指に休暇を与えてください。', tone: 'critical' },
  { text: '朗報: 今日のバグは全部エンタメ扱いです。', tone: 'info' },
  { text: '注意: その演出、意味はありませんが勢いはあります。', tone: 'warn' }
];
const updateResults = [
  '更新完了: バグを2件増やして臨場感を強化しました。',
  '更新完了: 何も変わっていません。気のせいです。',
  '更新完了: 表示だけ1.0.0.0.1になりました。',
  '更新完了: 演出が0.7秒だけドラマチックになりました。'
];
const konamiCode = ['arrowup', 'arrowup', 'arrowdown', 'arrowdown', 'arrowleft', 'arrowright', 'arrowleft', 'arrowright', 'b', 'a'];

function buildCoinBurst(count: number) {
  return Array.from({ length: count }, (_, index) => ({
    id: index,
    x: Math.random() * 100,
    drift: (Math.random() - 0.5) * 36,
    duration: 2.5 + Math.random() * 2,
    delay: Math.random() * 0.8,
    emoji: coinEmojis[Math.floor(Math.random() * coinEmojis.length)],
    size: 1 + Math.random() * 1.15,
    rotate: Math.random() * 360
  } satisfies CoinParticle));
}

function randomFrom<T>(values: T[]) {
  return values[Math.floor(Math.random() * values.length)];
}

function rollSlotFaces() {
  return [randomFrom(rouletteFaces), randomFrom(rouletteFaces), randomFrom(rouletteFaces)];
}

function makeHeadline() {
  return randomFrom(fakeHeadlines);
}

function DeferredPlaceholders() {
  return (
    <>
      <section className="deferred-placeholder deferred-placeholder-lg" aria-live="polite" role="status">
        <span className="spinner" aria-hidden="true" />
      </section>
      <section className="deferred-placeholder" aria-live="polite" role="status">
        <span className="spinner spinner-sm" aria-hidden="true" />
      </section>
    </>
  );
}

function HomePage() {
  const [showDeferredSections, setShowDeferredSections] = useState(false);
  const [chaosLevel, setChaosLevel] = useState(8);
  const [todayFortune, setTodayFortune] = useState('「運勢を受け取る」を押すと未来が雑に確定します。');
  const [slotResult, setSlotResult] = useState(rollSlotFaces().join(''));
  const [headline, setHeadline] = useState(makeHeadline());
  const [tickerMessage, setTickerMessage] = useState(randomFrom(fakeTickerMessages));
  const [coinBurst, setCoinBurst] = useState<CoinParticle[]>([]);
  const [flipMode, setFlipMode] = useState(false);
  const [glitchMode, setGlitchMode] = useState(false);
  const [partyMode, setPartyMode] = useState(false);
  const [godMode, setGodMode] = useState(false);
  const [maintenanceVisible, setMaintenanceVisible] = useState(false);
  const [maintenanceMessage, setMaintenanceMessage] = useState('サーバー再起動のフリをしています...');
  const [jackpotCount, setJackpotCount] = useState(18427001);
  const [mascotClickCount, setMascotClickCount] = useState(0);
  const [fakeAlerts, setFakeAlerts] = useState<FakeAlert[]>([]);
  const [bossHp, setBossHp] = useState(100);
  const [bossLog, setBossLog] = useState('ボス「巨大メダル山」が待機中');
  const [clickStreak, setClickStreak] = useState(0);
  const [adPopupVisible, setAdPopupVisible] = useState(false);
  const [adMessage, setAdMessage] = useState(randomFrom(fakeAdMessages));
  const [isUpdateChecking, setIsUpdateChecking] = useState(false);
  const [updateProgress, setUpdateProgress] = useState(0);
  const [updateResult, setUpdateResult] = useState('');

  const deferredSectionAnchorRef = useRef<HTMLDivElement | null>(null);
  const coinClearTimerRef = useRef<number | null>(null);
  const maintenanceTimersRef = useRef<number[]>([]);
  const alertIdRef = useRef(0);
  const konamiProgressRef = useRef(0);

  const missionState = useMemo(
    () => ({
      rain: coinBurst.length > 0,
      wild: chaosLevel >= 45,
      flip: flipMode,
      mascot: mascotClickCount >= 7,
      boss: bossHp === 0,
      god: godMode
    }),
    [coinBurst.length, chaosLevel, flipMode, mascotClickCount, bossHp, godMode]
  );

  const completedMissions = Object.values(missionState).filter(Boolean).length;
  const chaosGauge = Math.min(100, chaosLevel);

  const streakComment =
    clickStreak >= 35
      ? '連打神: 指先がサーバーより先に温まっています。'
      : clickStreak >= 20
        ? '連打職人: いいリズムです。近所には配慮してください。'
        : clickStreak >= 8
          ? '連打見習い: だいぶ押せてます。いい汗です。'
          : 'ウォームアップ中: まずは押して慣らしましょう。';

  useEffect(() => {
    if (showDeferredSections) {
      return;
    }

    let fallbackTimeoutId: number | undefined;
    const revealDeferredSections = () => setShowDeferredSections(true);
    const anchorElement = deferredSectionAnchorRef.current;

    if (anchorElement && typeof window.IntersectionObserver === 'function') {
      const observer = new window.IntersectionObserver(
        (entries) => {
          if (entries.some((entry) => entry.isIntersecting)) {
            revealDeferredSections();
            observer.disconnect();
          }
        },
        { rootMargin: '320px 0px' }
      );
      observer.observe(anchorElement);
      return () => observer.disconnect();
    }

    fallbackTimeoutId = window.setTimeout(revealDeferredSections, 1200);
    return () => {
      if (typeof fallbackTimeoutId === 'number') {
        window.clearTimeout(fallbackTimeoutId);
      }
    };
  }, [showDeferredSections]);

  useEffect(() => {
    document.body.classList.toggle('april-flip-mode', flipMode);
    return () => document.body.classList.remove('april-flip-mode');
  }, [flipMode]);

  useEffect(() => {
    const timerId = window.setInterval(() => {
      setJackpotCount((current) => current + Math.floor(Math.random() * 6000) + 1400 + chaosLevel * 2 + (partyMode ? 1300 : 0));
    }, 1700);

    return () => window.clearInterval(timerId);
  }, [chaosLevel, partyMode]);

  useEffect(() => {
    const tickerId = window.setInterval(() => {
      setTickerMessage(randomFrom(fakeTickerMessages));
    }, 5000);

    return () => window.clearInterval(tickerId);
  }, []);

  useEffect(() => {
    const adId = window.setInterval(() => {
      if (Math.random() > 0.45) {
        setAdMessage(randomFrom(fakeAdMessages));
        setAdPopupVisible(true);
      }
    }, 23000);

    return () => window.clearInterval(adId);
  }, []);

  useEffect(() => {
    const handleKeyDown = (event: KeyboardEvent) => {
      const target = event.target as HTMLElement | null;
      if (target?.tagName === 'INPUT' || target?.tagName === 'TEXTAREA') {
        return;
      }

      const key = event.key.toLowerCase();
      const expected = konamiCode[konamiProgressRef.current];

      if (key === expected) {
        konamiProgressRef.current += 1;

        if (konamiProgressRef.current === konamiCode.length) {
          konamiProgressRef.current = 0;
          setGodMode(true);
          setGlitchMode(true);
          setChaosLevel((value) => Math.min(100, value + 20));
          setTodayFortune('隠しコマンド成功: GOD MODE 解禁。今日だけ何を押しても演出が勝ちます。');
          setTickerMessage('隠しコマンド検知: 運営が静かに拍手しています。');
        }
      } else {
        konamiProgressRef.current = key === konamiCode[0] ? 1 : 0;
      }
    };

    window.addEventListener('keydown', handleKeyDown);
    return () => window.removeEventListener('keydown', handleKeyDown);
  }, []);

  useEffect(() => {
    if (!partyMode) {
      return;
    }

    const partyTimerId = window.setInterval(() => {
      setCoinBurst(buildCoinBurst(24));
      setChaosLevel((value) => Math.min(100, value + 2));

      if (coinClearTimerRef.current !== null) {
        window.clearTimeout(coinClearTimerRef.current);
      }

      coinClearTimerRef.current = window.setTimeout(() => {
        setCoinBurst([]);
        coinClearTimerRef.current = null;
      }, 2100);
    }, 4600);

    return () => window.clearInterval(partyTimerId);
  }, [partyMode]);

  useEffect(() => {
    if (!isUpdateChecking) {
      return;
    }

    setUpdateProgress(0);
    setUpdateResult('');

    const progressId = window.setInterval(() => {
      setUpdateProgress((current) => {
        const next = Math.min(100, current + Math.floor(Math.random() * 18) + 7);
        if (next >= 100) {
          window.clearInterval(progressId);
          setIsUpdateChecking(false);
          setUpdateResult(randomFrom(updateResults));
          setChaosLevel((value) => Math.min(100, value + 6));
        }
        return next;
      });
    }, 300);

    return () => window.clearInterval(progressId);
  }, [isUpdateChecking]);

  useEffect(() => {
    return () => {
      if (coinClearTimerRef.current !== null) {
        window.clearTimeout(coinClearTimerRef.current);
      }
      maintenanceTimersRef.current.forEach((timerId) => window.clearTimeout(timerId));
    };
  }, []);

  const triggerCoinRain = () => {
    setCoinBurst(buildCoinBurst(70));
    setChaosLevel((value) => Math.min(100, value + 9));

    if (coinClearTimerRef.current !== null) {
      window.clearTimeout(coinClearTimerRef.current);
    }

    coinClearTimerRef.current = window.setTimeout(() => {
      setCoinBurst([]);
      coinClearTimerRef.current = null;
    }, 4700);
  };

  const pushFakeAlert = () => {
    const source = randomFrom(fakeAlertPool);
    const nextAlert: FakeAlert = {
      id: ++alertIdRef.current,
      text: source.text,
      tone: source.tone
    };

    setFakeAlerts((current) => [nextAlert, ...current].slice(0, 6));
    setHeadline(makeHeadline());
    setChaosLevel((value) => Math.min(100, value + 3));
  };

  const triggerEmergencyMaintenance = () => {
    if (maintenanceVisible) {
      return;
    }

    setMaintenanceVisible(true);
    setMaintenanceMessage('サーバー再起動のフリをしています...');
    setChaosLevel((value) => Math.min(100, value + 12));

    const firstTimer = window.setTimeout(() => {
      setMaintenanceMessage('うそです。通常営業です。驚いた?');
    }, 1600);
    const secondTimer = window.setTimeout(() => {
      setMaintenanceVisible(false);
    }, 4200);

    maintenanceTimersRef.current.push(firstTimer, secondTimer);
  };

  const rollFortune = () => {
    setTodayFortune(randomFrom(fortunePool));
    setChaosLevel((value) => Math.min(100, value + 4));
  };

  const runSlot = () => {
    const faces = rollSlotFaces();
    const result = faces.join('');
    setSlotResult(result);
    setChaosLevel((value) => Math.min(100, value + 5));

    if (faces.every((face) => face === faces[0])) {
      setTodayFortune('777演出発生: 今日はなにを押しても正解です。');
      setJackpotCount((value) => value + 777777);
      triggerCoinRain();
      pushFakeAlert();
    }
  };

  const jamRedButton = () => {
    setClickStreak((value) => value + 1);
    setChaosLevel((value) => Math.min(100, value + Math.floor(Math.random() * 11) + 8));
    if (Math.random() > 0.62) {
      setHeadline(makeHeadline());
    }
    if (Math.random() > 0.7) {
      pushFakeAlert();
    }
  };

  const mascotClick = () => {
    setMascotClickCount((value) => value + 1);
    if (mascotClickCount >= 6) {
      setGlitchMode(true);
      setTodayFortune('隠し条件達成: マスコットがページを乗っ取りました。');
      setAdPopupVisible(true);
      setAdMessage('マスコット広告: つついてくれてありがとう。もう3回どうぞ。');
    }
  };

  const attackBoss = () => {
    if (bossHp <= 0) {
      return;
    }

    const damage = Math.floor(Math.random() * 22) + 9;
    const nextHp = Math.max(0, bossHp - damage);
    setBossHp(nextHp);
    setBossLog(`あなたの連打が ${damage}% のダメージを与えた!`);
    setChaosLevel((value) => Math.min(100, value + 4));

    if (nextHp === 0) {
      setBossLog('撃破成功: 巨大メダル山は「また明日」と言い残して崩れ落ちた。');
      setTodayFortune('ボス討伐ボーナス: 今日のあなたは演出面で最強です。');
      triggerCoinRain();
      pushFakeAlert();
    }
  };

  const resetBoss = () => {
    setBossHp(100);
    setBossLog('ボス「巨大メダル山」が再召喚された!');
  };

  return (
    <div className={`home-stack april-home ${glitchMode ? 'april-glitch' : ''} ${godMode ? 'april-god-mode' : ''}`}>
      <section className="april-hero" aria-labelledby="april-title">
        <p className="april-ribbon">2026 APRIL FOOLS SPECIAL MODE</p>
        <div className="april-hero-headline-row">
          <h1 id="april-title" className="home-title april-title">
            クソでっけぇプッシャーゲーム
            <span>嘘アプデ祭り会場</span>
          </h1>
          <button className="april-mascot-button" onClick={mascotClick} type="button" aria-label="マスコットをつつく">
            🐟
          </button>
        </div>
        <p className="home-subtitle april-subtitle">
          文言、見た目、テンション、すべて4月1日仕様です。まともな顔は明日戻る予定です。
        </p>
        <p className="april-breaking">📰 {headline}</p>

        <div className="april-live-ticker" aria-live="polite">
          <span>{tickerMessage}</span>
        </div>

        <div className="april-chaos-panel" role="status" aria-live="polite">
          <div>
            <p className="april-chaos-label">カオス指数</p>
            <p className="april-chaos-value">{chaosGauge}%</p>
          </div>
          <div className="april-chaos-bar" aria-hidden="true">
            <span style={{ width: `${chaosGauge}%` }} />
          </div>
          <button type="button" className="april-danger-button" onClick={jamRedButton}>
            <IconBomb size={18} aria-hidden="true" />
            とりあえず赤いボタンを押す
          </button>
          <p className="april-streak-note">連打カウント: {clickStreak} / {streakComment}</p>
        </div>
      </section>

      <section className="april-gimmick-grid" aria-label="エイプリルフールギミック">
        <article className="april-card april-card-loud">
          <h2><IconConfetti size={18} /> メダル豪雨ボタン</h2>
          <p>押すとだいたい景気が良くなります。CPU使用率も少し景気良くなります。</p>
          <button type="button" className="april-action-button" onClick={triggerCoinRain}>
            メダルを降らせる
          </button>
        </article>

        <article className="april-card">
          <h2><IconCalendarSmile size={18} /> 本日の運勢</h2>
          <p className="april-fortune">{todayFortune}</p>
          <button type="button" className="april-action-button" onClick={rollFortune}>
            運勢を受け取る
          </button>
        </article>

        <article className="april-card">
          <h2><IconDeviceGamepad2 size={18} /> 3秒スロット</h2>
          <p className="april-slot-result" aria-live="polite">{slotResult}</p>
          <button type="button" className="april-action-button" onClick={runSlot}>
            レバーを引く
          </button>
        </article>

        <article className="april-card">
          <h2><IconRotateClockwise2 size={18} /> 逆さまモード</h2>
          <p>画面全体をひっくり返します。酔いやすい方はご注意ください。</p>
          <button type="button" className="april-action-button" onClick={() => setFlipMode((value) => !value)}>
            {flipMode ? '元に戻す' : '世界をひっくり返す'}
          </button>
        </article>

        <article className="april-card">
          <h2><IconBolt size={18} /> 緊急メンテ演出</h2>
          <p>押すと一瞬だけ運営の気持ちになれます。</p>
          <button type="button" className="april-action-button" onClick={triggerEmergencyMaintenance}>
            メンテ開始(うそ)
          </button>
        </article>

        <article className="april-card">
          <h2><IconSparkles size={18} /> 覚醒モード</h2>
          <p>文字が暴れます。責任は4月1日が持ちます。</p>
          <button type="button" className="april-action-button" onClick={() => setGlitchMode((value) => !value)}>
            {glitchMode ? '正気を取り戻す' : '覚醒する'}
          </button>
        </article>

        <article className="april-card">
          <h2><IconAlertTriangle size={18} /> 偽ニュース生成機</h2>
          <p>最新のそれっぽい速報を自動生成して、みんなを軽く混乱させます。</p>
          <button type="button" className="april-action-button" onClick={pushFakeAlert}>
            偽速報を投下
          </button>
        </article>

        <article className="april-card">
          <h2><IconDeviceGamepad2 size={18} /> レイドボス戦</h2>
          <p className="april-boss-log">{bossLog}</p>
          <div className="april-boss-hp" aria-hidden="true">
            <span style={{ width: `${bossHp}%` }} />
          </div>
          <p className="april-boss-percent">ボスHP: {bossHp}%</p>
          {bossHp > 0 ? (
            <button type="button" className="april-action-button" onClick={attackBoss}>
              連打で攻撃
            </button>
          ) : (
            <button type="button" className="april-action-button" onClick={resetBoss}>
              ボス再召喚
            </button>
          )}
        </article>

        <article className="april-card">
          <h2><IconRefresh size={18} /> 嘘アップデート確認</h2>
          <p>チェックするたびに、ちょっとだけ意味のない達成感を得られます。</p>
          <div className="april-update-bar" aria-hidden="true">
            <span style={{ width: `${updateProgress}%` }} />
          </div>
          <p className="april-update-text">進捗: {updateProgress}%</p>
          {updateResult && <p className="april-update-result">{updateResult}</p>}
          <button type="button" className="april-action-button" onClick={() => setIsUpdateChecking(true)} disabled={isUpdateChecking}>
            {isUpdateChecking ? '確認中...' : '更新を確認'}
          </button>
        </article>

        <article className="april-card">
          <h2><IconConfetti size={18} /> パーティーモード</h2>
          <p>ONにすると定期的にメダルが降ります。理性は降りません。</p>
          <button type="button" className="april-action-button" onClick={() => setPartyMode((value) => !value)}>
            {partyMode ? 'パーティー終了' : 'パーティー開始'}
          </button>
        </article>
      </section>

      <section className="april-alert-feed" aria-live="polite" aria-label="偽速報フィード">
        <h2>偽速報フィード</h2>
        {fakeAlerts.length === 0 ? (
          <p>まだ速報はありません。ボタンを押して世間をざわつかせてください。</p>
        ) : (
          <ul>
            {fakeAlerts.map((alert) => (
              <li key={alert.id} className={`tone-${alert.tone}`}>{alert.text}</li>
            ))}
          </ul>
        )}
      </section>

      <section className="april-mission-board" aria-labelledby="mission-title">
        <div>
          <h2 id="mission-title">本日のミッション</h2>
          <p>クリアしても称号はもらえません。達成感だけ配布します。</p>
        </div>
        <ul>
          <li className={missionState.rain ? 'done' : ''}>メダル豪雨を1回発動</li>
          <li className={missionState.wild ? 'done' : ''}>カオス指数45%以上</li>
          <li className={missionState.flip ? 'done' : ''}>逆さまモードを体験</li>
          <li className={missionState.mascot ? 'done' : ''}>マスコットを7回つつく</li>
          <li className={missionState.boss ? 'done' : ''}>レイドボスを撃破</li>
          <li className={missionState.god ? 'done' : ''}>隠しコマンドでGOD MODE発動</li>
        </ul>
        <p className="april-mission-progress">達成率: {completedMissions}/6</p>
      </section>

      <section className="link-card april-links-card">
        <h2 className="section-title">メイン導線(ここだけ真面目)</h2>
        <div className="link-grid link-grid-two">
          <div className="link-column">
            <MmpPrimaryButton
              href="https://discord.com/invite/CgnYyXecKm"
              target="_blank"
              icon={<IconBrandDiscord size={32} />}
              size="xl"
              className="mmp-link-hero-button"
              heightMultiplier={1.35}
            >
              公式Discord
            </MmpPrimaryButton>
            <MmpPrimaryButton
              href="https://wikiwiki.jp/vr_bigpusher/"
              target="_blank"
              icon={<IconBook2 size={22} />}
              size="lg"
              className="mmp-link-hero-button"
            >
              公式Wiki
            </MmpPrimaryButton>
          </div>
          <div className="link-column">
            <MmpPrimaryButton
              href="https://vrchat.com/home/group/grp_5900a25d-0bb9-48d4-bab1-f3bd5c9a5e73"
              target="_blank"
              icon={<IconWorld size={22} />}
              size="lg"
              className="mmp-link-hero-button"
            >
              公式グループ(VRChat)
            </MmpPrimaryButton>
            <MmpOutlineButton
              href="https://github.com/pikachu0310/very-big-medal-pusher-data-server"
              target="_blank"
              icon={<IconBrandGithub size={20} />}
              size="lg"
            >
              Data Server GitHub
            </MmpOutlineButton>
          </div>
        </div>
      </section>

      <section className="april-jackpot-strip" aria-live="polite">
        <p>本日のなんとなく増える世界メダル総量</p>
        <strong>{jackpotCount.toLocaleString()} 枚</strong>
      </section>

      <section className="april-patchnote-card" aria-label="嘘パッチノート">
        <h2>4/1 嘘パッチノート</h2>
        <ul>
          <li>メダルの角を丸くしてやさしい握り心地に改善。</li>
          <li>押すと鳴るボタンの効果音を「ﾄﾞﾔｧ」に変更。</li>
          <li>サーバー負荷が高いとき、代わりに開発者が汗をかきます。</li>
          <li>ランキング1位の称号を「今日の圧」に変更。</li>
          <li>長押し判定の代わりに「心の強さ判定」を実験導入。</li>
          <li>広告ポップアップを閉じる速度に応じて精神力を表示。</li>
          <li>仕様書の脚注にだけ本当の変更点が紛れ込むよう改善。</li>
        </ul>
      </section>

      <div ref={deferredSectionAnchorRef} className="deferred-anchor" aria-hidden="true" />

      <section className="april-serious-zone" aria-label="真面目ゾーン">
        <h2>真面目ゾーン (統計・開発者向け)</h2>
        <p>この先は通常機能です。エイプリルフールに疲れたらここへ避難してください。</p>
      </section>

      {showDeferredSections ? (
        <Suspense fallback={<DeferredPlaceholders />}>
          <DeferredSections />
        </Suspense>
      ) : (
        <DeferredPlaceholders />
      )}

      {coinBurst.length > 0 && (
        <div className="april-coin-rain" aria-hidden="true">
          {coinBurst.map((coin) => (
            <span
              key={coin.id}
              className="april-coin"
              style={{
                left: `${coin.x}%`,
                '--coin-drift': `${coin.drift}px`,
                '--coin-duration': `${coin.duration}s`,
                '--coin-delay': `${coin.delay}s`,
                '--coin-size': `${coin.size}rem`,
                '--coin-rotate': `${coin.rotate}deg`
              } as CSSProperties}
            >
              {coin.emoji}
            </span>
          ))}
        </div>
      )}

      {maintenanceVisible && (
        <div className="april-maintenance-overlay" role="alertdialog" aria-modal="true" aria-live="assertive">
          <div className="april-maintenance-card">
            <p>⚠ 緊急メンテナンス速報 ⚠</p>
            <strong>{maintenanceMessage}</strong>
            <span>この演出は 4.2 秒で自動終了します。</span>
          </div>
        </div>
      )}

      {adPopupVisible && (
        <div className="april-ad-popup" role="dialog" aria-modal="true" aria-live="assertive">
          <div className="april-ad-card">
            <p className="april-ad-label">提供: でかプ広告委員会(架空)</p>
            <strong>{adMessage}</strong>
            <button type="button" onClick={() => setAdPopupVisible(false)}>
              閉じる(勇気)
            </button>
          </div>
        </div>
      )}
    </div>
  );
}

export default HomePage;
