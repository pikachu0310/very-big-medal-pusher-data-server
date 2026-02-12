import { Link } from 'react-router';

function PrivacyPolicyPage() {
  return (
    <div className="policy-page">
      <article className="policy-card">
        <Link to="/" className="policy-back-link">
          ← ホームに戻る
        </Link>
        <h1 className="policy-title">プライバシーポリシー</h1>

        <div className="policy-sections">
          <section>
            <h2>1. 収集する情報</h2>
            <p>
              本サービスは、VRChatワールド「クソでっけぇプッシャーゲーム」のプレイデータを収集・保存します。
              収集されるデータには以下が含まれます：
            </p>
            <ul>
              <li>VRChatユーザー名</li>
              <li>ゲームプレイに関する統計情報（メダル数、プレイ時間、スコアなど）</li>
              <li>実績データ</li>
              <li>セーブデータの更新日時</li>
            </ul>
          </section>

          <section>
            <h2>2. 情報の利用目的</h2>
            <p>収集した情報は以下の目的で利用されます：</p>
            <ul>
              <li>プレイヤーのゲームデータの保存と復元</li>
              <li>グローバルランキングの生成と表示</li>
              <li>ゲームの統計情報の提供</li>
              <li>サービスの改善と機能追加</li>
            </ul>
          </section>

          <section>
            <h2>3. 情報の共有と開示</h2>
            <p>本サービスでは、以下の情報が公開されます：</p>
            <ul>
              <li>グローバルランキング（ユーザー名と統計値）</li>
              <li>実績取得率などの集計情報</li>
            </ul>
            <p>
              個人を特定できる情報を第三者に販売または貸与することはありません。
            </p>
          </section>

          <section>
            <h2>4. データの保存期間</h2>
            <p>
              収集したデータは、サービスの提供に必要な期間保存されます。
              ユーザーがデータの削除を希望する場合は、運営者にお問い合わせください。
            </p>
          </section>

          <section>
            <h2>5. セキュリティ</h2>
            <p>
              本サービスは、収集したデータを保護するために適切なセキュリティ対策を実施しています。
              データの送受信にはHMAC-SHA256署名による認証を使用しています。
            </p>
          </section>

          <section>
            <h2>6. プライバシーポリシーの変更</h2>
            <p>
              本プライバシーポリシーは、法令の変更やサービスの改善に伴い、予告なく変更される場合があります。
              変更後のプライバシーポリシーは、本ページに掲載された時点で効力を生じるものとします。
            </p>
          </section>

          <section>
            <h2>7. お問い合わせ</h2>
            <p>
              本プライバシーポリシーに関するご質問やご意見、データの削除依頼などは、
              サービス運営者までお問い合わせください。
            </p>
          </section>
        </div>

        <p className="policy-updated-at">最終更新日：2025年10月11日</p>
      </article>
    </div>
  );
}

export default PrivacyPolicyPage;
