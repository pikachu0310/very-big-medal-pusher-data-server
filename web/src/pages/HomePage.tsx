import { lazy, Suspense, useEffect, useMemo, useRef, useState, type CSSProperties } from 'react';
import {
  IconBrandDiscord,
  IconBrandGithub,
  IconBook2,
  IconBolt,
  IconBomb,
  IconCalendarSmile,
  IconConfetti,
  IconDeviceGamepad2,
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

function buildCoinBurst(count: number) {
  return Array.from({ length: count }, (_, index) => ({
    id: index,
    x: Math.random() * 100,
    drift: (Math.random() - 0.5) * 36,
    duration: 2.6 + Math.random() * 2.1,
    delay: Math.random() * 0.9,
    emoji: coinEmojis[Math.floor(Math.random() * coinEmojis.length)],
    size: 1 + Math.random() * 1.1,
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
  const [coinBurst, setCoinBurst] = useState<CoinParticle[]>([]);
  const [flipMode, setFlipMode] = useState(false);
  const [glitchMode, setGlitchMode] = useState(false);
  const [maintenanceVisible, setMaintenanceVisible] = useState(false);
  const [maintenanceMessage, setMaintenanceMessage] = useState('サーバー再起動のフリをしています...');
  const [jackpotCount, setJackpotCount] = useState(18427001);
  const [mascotClickCount, setMascotClickCount] = useState(0);
  const deferredSectionAnchorRef = useRef<HTMLDivElement | null>(null);
  const coinClearTimerRef = useRef<number | null>(null);
  const maintenanceTimersRef = useRef<number[]>([]);

  const missionState = useMemo(
    () => ({
      rain: coinBurst.length > 0,
      wild: chaosLevel >= 45,
      flip: flipMode,
      mascot: mascotClickCount >= 7
    }),
    [coinBurst.length, chaosLevel, flipMode, mascotClickCount]
  );

  const completedMissions = Object.values(missionState).filter(Boolean).length;

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
      setJackpotCount((current) => current + Math.floor(Math.random() * 6000) + 1400 + chaosLevel * 2);
    }, 1700);

    return () => window.clearInterval(timerId);
  }, [chaosLevel]);

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
    }
  };

  const jamRedButton = () => {
    setChaosLevel((value) => Math.min(100, value + Math.floor(Math.random() * 11) + 8));
    if (Math.random() > 0.62) {
      setHeadline(makeHeadline());
    }
  };

  const mascotClick = () => {
    setMascotClickCount((value) => value + 1);
    if (mascotClickCount >= 6) {
      setGlitchMode(true);
      setTodayFortune('隠し条件達成: マスコットがページを乗っ取りました。');
    }
  };

  const chaosGauge = Math.min(100, chaosLevel);

  return (
    <div className={`home-stack april-home ${glitchMode ? 'april-glitch' : ''}`}>
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
        </ul>
        <p className="april-mission-progress">達成率: {completedMissions}/4</p>
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
          <li>プライバシーポリシーを少しだけ正直にしました。</li>
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
    </div>
  );
}

export default HomePage;
