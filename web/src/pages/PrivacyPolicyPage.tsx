import { Link } from 'react-router';

function PrivacyPolicyPage() {
  return (
    <div className="policy-page">
      <article className="policy-card">
        <Link to="/" className="policy-back-link">
          ← エイプリルフール会場に戻る
        </Link>
        <h1 className="policy-title">プライバシーポリシー (4/1 特別版)</h1>

        <div className="policy-sections">
          <section>
            <h2>1. 収集する情報</h2>
            <p>
              本サービスは、VRChatワールド「クソでっけぇプッシャーゲーム」のプレイデータを収集・保存します。
              収集対象は主にゲームの統計値、セーブデータ更新時刻、ユーザー識別に必要な最低限の情報です。
            </p>
            <ul>
              <li>VRChatユーザー名</li>
              <li>ゲーム統計（メダル数、プレイ時間、各種スコア）</li>
              <li>実績データ</li>
              <li>セーブデータの更新日時</li>
            </ul>
          </section>

          <section>
            <h2>2. 利用目的</h2>
            <p>取得したデータは以下のために使用します。</p>
            <ul>
              <li>プレイデータの保存・復元</li>
              <li>ランキングや統計情報の生成・表示</li>
              <li>ゲームバランス調整と不具合調査</li>
              <li>サービス改善</li>
            </ul>
            <p>なお、4月1日だけテンションは高いですが、データの扱いは通常どおりです。</p>
          </section>

          <section>
            <h2>3. 公開される情報</h2>
            <p>ランキングや統計ページでは、ユーザー名と統計値が表示される場合があります。</p>
            <p>個人を特定可能な情報を第三者へ販売・貸与することはありません。</p>
          </section>

          <section>
            <h2>4. 保存期間</h2>
            <p>
              データはサービス提供に必要な期間保存します。
              削除希望がある場合は運営へ連絡してください。確認後、対応可能な範囲で削除します。
            </p>
          </section>

          <section>
            <h2>5. セキュリティ</h2>
            <p>
              データ保護のため、送受信時の認証やアクセス制御などの対策を実施しています。
              API 認証には署名検証（HMAC-SHA256）を利用しています。
            </p>
          </section>

          <section>
            <h2>6. 改定</h2>
            <p>
              本ポリシーは法令やサービス内容の変更に応じて改定される場合があります。
              改定後はこのページに掲載した時点で効力を持ちます。
            </p>
          </section>

          <section>
            <h2>7. お問い合わせ</h2>
            <p>
              データ取り扱いに関する質問、削除依頼、誤表示の報告は運営者までお問い合わせください。
              4月1日でも対応は真面目です。
            </p>
          </section>
        </div>

        <p className="policy-updated-at">最終更新日: 2026年4月1日</p>
      </article>
    </div>
  );
}

export default PrivacyPolicyPage;
