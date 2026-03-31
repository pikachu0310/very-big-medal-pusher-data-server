import { lazy, Suspense, useCallback, useEffect, useRef, useState, type CSSProperties } from 'react';
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
  IconExternalLink,
  IconLock,
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

type CoinBurst = {
  id: number;
  particles: CoinParticle[];
};

type FakeAlert = {
  id: number;
  text: string;
  tone: 'info' | 'warn' | 'critical';
};

type AchievementKey = 'redEngineer' | 'slotWinner' | 'bossSlayer' | 'overdrive' | 'godMode' | 'mascotFriend';

type AchievementState = Record<AchievementKey, boolean>;

const rouletteFaces = ['7', '7', '🍒', '🍒', '🪙', '👑'];
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
  "Happy April Fool's Day!! 今年は4/13ごろまでゆるっとイベント運用中です。",
  '赤いボタンは押した分だけ盛り上がるタイプです。遠慮なくどうぞ。',
  '連打してる姿、けっこうカッコいいです。たぶん。',
  'いっぱい遊ぶほど演出がはしゃぎます。仕様です。',
  '今日の豆知識: 連打は哲学、放置は美学。'
];
const fakeAdMessages = [
  '期間限定: メダル1枚でメダル1枚が当たる激アツ抽選会!',
  'あなたへの特別提案: 1日3回の深呼吸でジャックポット体質へ。',
  'この広告を閉じると、閉じた事実だけが残ります。',
  '新機能: ボタンを見つめるとボタンも見つめ返します。'
];
const fakeAlertPool: { text: string; tone: FakeAlert['tone'] }[] = [
  { text: '通知: 大量連打を検知、演出負荷を解禁しました。', tone: 'warn' },
  { text: '業務連絡: 開発者は現在カフェインで稼働しています。', tone: 'info' },
  { text: '緊急: 押しすぎ検知。指に休暇を与えてください。', tone: 'critical' },
  { text: '朗報: 今日のバグは全部エンタメ扱いです。', tone: 'info' },
  { text: '注意: その演出、意味はありませんが勢いはあります。', tone: 'warn' }
];
const updateResults = [
  '更新完了: 演出がさらに元気になりました。',
  '更新完了: 何も変わってないようで、なんかいい感じです。',
  '更新完了: バージョン表記だけちょっと伸びました。',
  '更新完了: 0.7秒だけドラマチック成分を増量しました。'
];
const konamiCode = ['arrowup', 'arrowup', 'arrowdown', 'arrowdown', 'arrowleft', 'arrowright', 'arrowleft', 'arrowright', 'b', 'a'];

function buildCoinBurst(count: number) {
  return Array.from({ length: count }, (_, index) => ({
    id: index,
    x: Math.random() * 100,
    drift: (Math.random() - 0.5) * 46,
    duration: 2.3 + Math.random() * 2,
    delay: Math.random() * 0.7,
    emoji: coinEmojis[Math.floor(Math.random() * coinEmojis.length)],
    size: 1.1 + Math.random() * 1.5,
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
  const twitterHashUrl = 'https://x.com/search?q=%23%E3%81%A7%E3%81%8B%E3%83%97%20OR%20%23VR%E3%81%A7%E3%81%8B%E3%83%97&f=live';
  const [showDeferredSections, setShowDeferredSections] = useState(false);
  const [chaosLevel, setChaosLevel] = useState(8);
  const [chaosOverdrive, setChaosOverdrive] = useState(false);
  const [todayFortune, setTodayFortune] = useState('「運勢を受け取る」を押すと、今日の流れをゆるく占えます。');
  const [slotResult, setSlotResult] = useState(rollSlotFaces().join(''));
  const [headline, setHeadline] = useState(makeHeadline());
  const [tickerMessage, setTickerMessage] = useState(randomFrom(fakeTickerMessages));
  const [tickerKey, setTickerKey] = useState(0);
  const [coinBursts, setCoinBursts] = useState<CoinBurst[]>([]);
  const [flipMode, setFlipMode] = useState(false);
  const [glitchMode, setGlitchMode] = useState(false);
  const [godMode, setGodMode] = useState(false);
  const [maintenanceVisible, setMaintenanceVisible] = useState(false);
  const [maintenanceMessage, setMaintenanceMessage] = useState('メンテナンス処理をゆるっと開始しました...');
  const [jackpotCount, setJackpotCount] = useState(18427001);
  const [mascotClickCount, setMascotClickCount] = useState(0);
  const [fakeAlerts, setFakeAlerts] = useState<FakeAlert[]>([]);
  const [bossHp, setBossHp] = useState(100);
  const [bossLog, setBossLog] = useState('ボス「巨大メダル山」がやる気満々で待機中');
  const [clickStreak, setClickStreak] = useState(0);
  const [redButtonCount, setRedButtonCount] = useState(0);
  const [adPopupVisible, setAdPopupVisible] = useState(false);
  const [adMessage, setAdMessage] = useState(randomFrom(fakeAdMessages));
  const [isUpdateChecking, setIsUpdateChecking] = useState(false);
  const [updateProgress, setUpdateProgress] = useState(0);
  const [updateResult, setUpdateResult] = useState('');
  const [jackpotFever, setJackpotFever] = useState(false);
  const [achievements, setAchievements] = useState<AchievementState>({
    redEngineer: false,
    slotWinner: false,
    bossSlayer: false,
    overdrive: false,
    godMode: false,
    mascotFriend: false
  });

  const deferredSectionAnchorRef = useRef<HTMLDivElement | null>(null);
  const maintenanceTimersRef = useRef<number[]>([]);
  const chaosActionAtRef = useRef(Date.now());
  const alertIdRef = useRef(0);
  const burstIdRef = useRef(0);
  const burstTimersRef = useRef<number[]>([]);
  const konamiProgressRef = useRef(0);

  const setTicker = useCallback((message: string) => {
    setTickerMessage(message);
    setTickerKey((value) => value + 1);
  }, []);

  const registerChaosAction = useCallback((delta: number) => {
    chaosActionAtRef.current = Date.now();
    setChaosLevel((value) => Math.min(100, value + delta));
  }, []);

  const updateAchievement = useCallback((key: AchievementKey) => {
    setAchievements((current) => {
      if (current[key]) {
        return current;
      }
      return { ...current, [key]: true };
    });
  }, []);

  const triggerCoinRain = useCallback((count = 70, chaosGain = 9, durationMs = 5200) => {
    const burstId = ++burstIdRef.current;
    const burst: CoinBurst = {
      id: burstId,
      particles: buildCoinBurst(count)
    };

    setCoinBursts((current) => [...current, burst]);

    if (chaosGain > 0) {
      registerChaosAction(chaosGain);
    }

    const timerId = window.setTimeout(() => {
      setCoinBursts((current) => current.filter((item) => item.id !== burstId));
    }, durationMs);

    burstTimersRef.current.push(timerId);
  }, [registerChaosAction]);

  const pushFakeAlert = useCallback(() => {
    const source = randomFrom(fakeAlertPool);
    const nextAlert: FakeAlert = {
      id: ++alertIdRef.current,
      text: source.text,
      tone: source.tone
    };

    setFakeAlerts((current) => [nextAlert, ...current].slice(0, 6));
  }, []);

  const completedAchievements = Object.values(achievements).filter(Boolean).length;
  const chaosGauge = Math.round(Math.min(100, chaosLevel));
  const chaosIntensity = Math.max(0.45, Math.min(2.2, chaosLevel / 46));

  const streakComment =
    clickStreak >= 35
      ? '連打神: 指先がサーバーより先に温まってます。'
      : clickStreak >= 20
        ? '連打職人: いいリズム! その調子です。'
        : clickStreak >= 8
          ? '連打見習い: だいぶノってきました。'
          : 'ウォームアップ中: まずは気楽に押してみましょう。';

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
      setJackpotCount((current) => current + Math.floor(Math.random() * 6000) + 1400 + chaosLevel * 2 + (chaosOverdrive ? 2200 : 0));
    }, 1700);

    return () => window.clearInterval(timerId);
  }, [chaosLevel, chaosOverdrive]);

  useEffect(() => {
    const tickerId = window.setInterval(() => {
      setTicker(randomFrom(fakeTickerMessages));
    }, 5000);

    return () => window.clearInterval(tickerId);
  }, [setTicker]);

  useEffect(() => {
    const adId = window.setInterval(() => {
      if (Math.random() > 0.48) {
        setAdMessage(randomFrom(fakeAdMessages));
        setAdPopupVisible(true);
      }
    }, 24000);

    return () => window.clearInterval(adId);
  }, []);

  useEffect(() => {
    const decayId = window.setInterval(() => {
      const idleMs = Date.now() - chaosActionAtRef.current;
      if (idleMs < 5000 || chaosOverdrive) {
        return;
      }

      setChaosLevel((value) => {
        if (value <= 5) {
          return 5;
        }
        return Math.max(5, value * 0.9);
      });
    }, 1000);

    return () => window.clearInterval(decayId);
  }, [chaosOverdrive]);

  useEffect(() => {
    if (chaosLevel < 100 || chaosOverdrive) {
      return;
    }

    setChaosOverdrive(true);
    setGlitchMode(true);
    updateAchievement('overdrive');
    setTicker('OVERDRIVE発動! 演出が本気モードに入りました。');
    pushFakeAlert();

    triggerCoinRain(160, 0, 3800);
    const stormId = window.setInterval(() => {
      triggerCoinRain(55, 0, 2800);
    }, 900);

    const endId = window.setTimeout(() => {
      window.clearInterval(stormId);
      setChaosOverdrive(false);
      setChaosLevel(82);
    }, 7500);

    return () => {
      window.clearInterval(stormId);
      window.clearTimeout(endId);
    };
  }, [chaosLevel, chaosOverdrive, pushFakeAlert, setTicker, triggerCoinRain, updateAchievement]);

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
          updateAchievement('godMode');
          registerChaosAction(20);
          setTodayFortune('隠しコマンド成功! GOD MODEになりました。今日は演出が勝ち確です。');
          setTicker('隠しコマンドを検知! 運営が静かにガッツポーズしています。');
        }
      } else {
        konamiProgressRef.current = key === konamiCode[0] ? 1 : 0;
      }
    };

    window.addEventListener('keydown', handleKeyDown);
    return () => window.removeEventListener('keydown', handleKeyDown);
  }, [registerChaosAction, setTicker, updateAchievement]);

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
          registerChaosAction(6);
        }
        return next;
      });
    }, 300);

    return () => window.clearInterval(progressId);
  }, [isUpdateChecking, registerChaosAction]);

  useEffect(() => {
    return () => {
      burstTimersRef.current.forEach((timerId) => window.clearTimeout(timerId));
      maintenanceTimersRef.current.forEach((timerId) => window.clearTimeout(timerId));
    };
  }, []);

  const triggerEmergencyMaintenance = () => {
    if (maintenanceVisible) {
      return;
    }

    setMaintenanceVisible(true);
    setMaintenanceMessage('メンテナンス処理をゆるっと開始しました...');
    registerChaosAction(12);

    const firstTimer = window.setTimeout(() => {
      setMaintenanceMessage('メンテナンス完了! お待たせしました。');
    }, 1600);
    const secondTimer = window.setTimeout(() => {
      setMaintenanceVisible(false);
    }, 4200);

    maintenanceTimersRef.current.push(firstTimer, secondTimer);
  };

  const rollFortune = () => {
    setTodayFortune(randomFrom(fortunePool));
    registerChaosAction(4);
  };

  const runSlot = () => {
    const faces = rollSlotFaces();
    const result = faces.join('');
    setSlotResult(result);
    registerChaosAction(6);

    if (faces.every((face) => face === faces[0])) {
      setTodayFortune('777演出発生: 今日はなにを押しても正解です。');
      setJackpotCount((value) => value + 777777);
      updateAchievement('slotWinner');
      setJackpotFever(true);
      registerChaosAction(18);
      triggerCoinRain(160, 0, 3600);
      const feverTimer = window.setTimeout(() => setJackpotFever(false), 3600);
      burstTimersRef.current.push(feverTimer);
      pushFakeAlert();
    }
  };

  const jamRedButton = () => {
    setClickStreak((value) => value + 1);
    setRedButtonCount((value) => {
      const next = value + 1;
      if (next >= 10) {
        updateAchievement('redEngineer');
      }
      return next;
    });
    registerChaosAction(Math.floor(Math.random() * 11) + 8);
    if (Math.random() > 0.62) {
      setHeadline(makeHeadline());
    }
    if (Math.random() > 0.74) {
      pushFakeAlert();
    }
  };

  const mascotClick = () => {
    setMascotClickCount((value) => {
      const next = value + 1;
      if (next >= 7) {
        updateAchievement('mascotFriend');
      }
      return next;
    });

    if (mascotClickCount >= 6) {
      setGlitchMode(true);
      setTodayFortune('隠し条件クリア! マスコットがページを楽しんでいます。');
      setAdPopupVisible(true);
      setAdMessage('マスコットより: つついてくれてありがとう! もう3回いける?');
    }
  };

  const attackBoss = () => {
    if (bossHp <= 0) {
      return;
    }

    const damage = Math.floor(Math.random() * 22) + 9;
    const nextHp = Math.max(0, bossHp - damage);
    setBossHp(nextHp);
    setBossLog(`あなたの連打が ${damage}% のダメージを与えました!`);
    registerChaosAction(5);

    if (nextHp === 0) {
      setBossLog('撃破成功! 巨大メダル山は「また明日」と言って崩れました。');
      setTodayFortune('ボス討伐ボーナス! 今日は演出面でほぼ最強です。');
      updateAchievement('bossSlayer');
      triggerCoinRain(120, 0, 3200);
      pushFakeAlert();
    }
  };

  const resetBoss = () => {
    setBossHp(100);
    setBossLog('ボス「巨大メダル山」を再召喚しました!');
  };

  return (
    <div
      className={`home-stack april-home ${glitchMode ? 'april-glitch' : ''} ${godMode ? 'april-god-mode' : ''} ${chaosOverdrive ? 'april-chaos-overdrive' : ''}`}
      style={{ '--chaos-intensity': chaosIntensity.toFixed(2) } as CSSProperties}
    >
      <section className="april-hero" aria-labelledby="april-title">
        <p className="april-ribbon">Happy April Fool's Day!! (4/1〜4/13ごろまでイベント中)</p>
        <div className="april-hero-headline-row">
          <h1 id="april-title" className="home-title april-title">
            クソでっけぇプッシャーゲーム
            <span>イベント会場</span>
          </h1>
          <div className="april-hero-side">
            <button className="april-mascot-button" onClick={mascotClick} type="button" aria-label="マスコットをつつく">
              🐟
            </button>
            <img src="/MMP_logo_480.webp" alt="Massive Medal Pusher ロゴ" className="april-mini-logo" />
          </div>
        </div>
        <p className="home-subtitle april-subtitle">
          今日はちょっとお祭りテンションのページです。落ち着いた通常版は <a href="/classic" className="april-inline-link">/classic</a> からどうぞ。
        </p>
        <p className="april-breaking">📰 {headline}</p>

        <div className="april-live-ticker" aria-live="polite">
          <span key={tickerKey}>{tickerMessage}</span>
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

      <section className="april-gimmick-grid" aria-label="イベントギミック">
        <article className="april-card april-card-loud">
          <h2><IconConfetti size={18} /> メダル豪雨ボタン</h2>
          <p>押すたびに新しい豪雨を追加発動。重なるほどド派手になります。</p>
          <button type="button" className="april-action-button" onClick={() => triggerCoinRain()}>
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
          <p>画面全体をひっくり返します。酔いやすい方はゆっくりどうぞ。</p>
          <button type="button" className="april-action-button" onClick={() => setFlipMode((value) => !value)}>
            {flipMode ? '元に戻す' : '世界をひっくり返す'}
          </button>
        </article>

        <article className="april-card">
          <h2><IconBolt size={18} /> 緊急メンテ演出</h2>
          <p>運営画面の緊張感を、3秒でちょっとだけ体験できます。</p>
          <button type="button" className="april-action-button" onClick={triggerEmergencyMaintenance}>
            メンテ開始
          </button>
        </article>

        <article className="april-card">
          <h2><IconSparkles size={18} /> 覚醒モード</h2>
          <p>ページ全体がぶるぶるします。カオス指数が高いほど揺れ幅もアップ。</p>
          <button type="button" className="april-action-button" onClick={() => setGlitchMode((value) => !value)}>
            {glitchMode ? '覚醒を解除' : '覚醒する'}
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
          <h2><IconRefresh size={18} /> 更新チェック</h2>
          <p>押すたびに、なんか進んだ気分になれます。</p>
          <div className="april-update-bar" aria-hidden="true">
            <span style={{ width: `${updateProgress}%` }} />
          </div>
          <p className="april-update-text">進捗: {updateProgress}%</p>
          {updateResult && <p className="april-update-result">{updateResult}</p>}
          <button type="button" className="april-action-button" onClick={() => setIsUpdateChecking(true)} disabled={isUpdateChecking}>
            {isUpdateChecking ? '確認中...' : '更新を確認'}
          </button>
        </article>
      </section>

      <section className="april-alert-feed" aria-live="polite" aria-label="通知フィード">
        <h2><IconAlertTriangle size={18} /> 通知フィード</h2>
        {fakeAlerts.length === 0 ? (
          <p>まだ通知はありません。遊ぶほどにぎやかになります。</p>
        ) : (
          <ul>
            {fakeAlerts.map((alert) => (
              <li key={alert.id} className={`tone-${alert.tone}`}>{alert.text}</li>
            ))}
          </ul>
        )}
      </section>

      <section className="april-mission-board" aria-labelledby="achievement-title">
        <div>
          <h2 id="achievement-title">イベント実績</h2>
          <p>条件を満たすと自動で達成。気楽に集めてください。</p>
        </div>
        <ul>
          <li className={achievements.redEngineer ? 'done' : ''}>赤ボタンを10回押す ({Math.min(redButtonCount, 10)}/10)</li>
          <li className={achievements.slotWinner ? 'done' : ''}>3秒スロットで当たりを引く</li>
          <li className={achievements.bossSlayer ? 'done' : ''}>レイドボスを撃破する</li>
          <li className={achievements.overdrive ? 'done' : ''}>カオス指数100%でOVERDRIVEを発動</li>
          <li className={achievements.godMode ? 'done' : ''}>隠しコマンドでGOD MODEを解禁</li>
          <li className={achievements.mascotFriend ? 'done' : ''}>マスコットを7回つつく</li>
        </ul>
        <p className="april-mission-progress">達成率: {completedAchievements}/6</p>
      </section>

      <section className="link-card april-links-card">
        <h2 className="section-title">でかプ公式リンク集 / MMP Quick Links</h2>
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
            <MmpPrimaryButton
              href="https://vrchat.com/home/launch?worldId=wrld_1af53798-92a3-4c3f-99ae-a7c42ec6084d"
              target="_blank"
              icon={<IconWorld size={22} />}
              size="lg"
              className="mmp-link-hero-button"
            >
              VRChatワールドリンク
            </MmpPrimaryButton>
            <MmpPrimaryButton
              href={twitterHashUrl}
              target="_blank"
              icon={<IconExternalLink size={22} />}
              size="lg"
              className="mmp-link-hero-button"
            >
              #でかプ / #VRでかプ (X投稿)
            </MmpPrimaryButton>
          </div>
        </div>
      </section>

      <section className="link-card april-links-card">
        <h2 className="section-title">開発者向けリンク集 / Links for Developers</h2>
        <div className="link-grid link-grid-three">
          <MmpOutlineButton
            href="/swagger/index.html"
            target="_blank"
            icon={<IconExternalLink size={18} />}
            size="lg"
            heightMultiplier={1.1}
          >
            SwaggerUI (API一覧)
          </MmpOutlineButton>
          <MmpOutlineButton
            href="https://push.trap.show/?server=mariadb.ns-system.svc.cluster.local&username=nsapp_c27d6f571f88ffff360fe2&db=nsapp_c27d6f571f88ffff360fe2"
            target="_blank"
            icon={<IconLock size={18} />}
            size="lg"
            heightMultiplier={1.1}
          >
            データベース
          </MmpOutlineButton>
          <MmpOutlineButton
            href="https://github.com/pikachu0310/very-big-medal-pusher-data-server"
            target="_blank"
            icon={<IconBrandGithub size={18} />}
            size="lg"
            heightMultiplier={1.1}
          >
            Data Server GitHub
          </MmpOutlineButton>
        </div>
      </section>

      <section className="april-jackpot-strip" aria-live="polite">
        <p>本日のなんとなく増える世界メダル総量</p>
        <strong>{jackpotCount.toLocaleString()} 枚</strong>
      </section>

      <section className="april-patchnote-card" aria-label="イベントパッチノート">
        <h2>イベントパッチノート</h2>
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
        <p>この先は通常機能です。イベント演出に疲れたらここへ避難してください。</p>
      </section>

      {showDeferredSections ? (
        <Suspense fallback={<DeferredPlaceholders />}>
          <DeferredSections />
        </Suspense>
      ) : (
        <DeferredPlaceholders />
      )}

      {coinBursts.length > 0 && (
        <div className="april-coin-rain" aria-hidden="true">
          {coinBursts.map((burst) => (
            <div key={burst.id} className="april-coin-layer">
              {burst.particles.map((coin) => (
                <span
                  key={`${burst.id}-${coin.id}`}
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
          ))}
        </div>
      )}

      {maintenanceVisible && (
        <div className="april-maintenance-overlay" role="alertdialog" aria-modal="true" aria-live="assertive">
          <div className="april-maintenance-card">
            <p>⚠ メンテナンス通知 ⚠</p>
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

      {chaosOverdrive && (
        <div className="april-overdrive-overlay" aria-hidden="true">
          <p>OVERDRIVE</p>
        </div>
      )}

      {jackpotFever && (
        <div className="april-jackpot-fever" aria-hidden="true">
          <p>JACKPOT FEVER</p>
        </div>
      )}
    </div>
  );
}

export default HomePage;
