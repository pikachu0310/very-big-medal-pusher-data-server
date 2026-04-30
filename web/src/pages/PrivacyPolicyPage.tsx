import { Link } from 'react-router';

function PrivacyPolicyPage() {
  return (
    <div className="policy-page">
      <article className="policy-card">
        <Link to="/" className="policy-back-link">
          ← トップに戻る
        </Link>
        <h1 className="policy-title">プライバシーポリシー</h1>

        <div className="policy-sections">
          <section>
            <h2>1. 収集する情報</h2>
            <p>
              本サービスは、VRChatワールド「Massive Medal Pusher」のクラウドセーブと統計機能を提供するため、
              ゲームクライアントから送信される次の情報を保存します。
            </p>
            <ul>
              <li>ユーザー識別子（user_id）</li>
              <li>セーブデータに含まれるゲーム進行状況、所持メダル、プレイ時間、各種統計値</li>
              <li>実績、パーク、トーテムなどの解除・利用状況</li>
              <li>セーブデータの作成日時・更新日時</li>
              <li>APIアクセスログ（リクエスト日時、URL、HTTPメソッド、ステータス、接続元IPアドレス）</li>
            </ul>
          </section>

          <section>
            <h2>2. 利用目的</h2>
            <p>取得した情報は、次の目的で利用します。</p>
            <ul>
              <li>プレイデータの保存・復元</li>
              <li>ランキング、実績取得率、メダル総量などの統計情報の生成・表示</li>
              <li>ゲームバランス調整と不具合調査</li>
              <li>APIの不正利用防止、署名検証、障害調査</li>
            </ul>
          </section>

          <section>
            <h2>3. 公開される情報</h2>
            <p>
              公開APIや統計ページでは、ランキング対象の user_id、統計値、更新日時が表示される場合があります。
              セーブデータで非表示設定が有効な場合、そのデータはランキングから除外されます。
            </p>
            <p>
              個別の最新セーブデータ、セーブ履歴、実績解除履歴の取得には署名認証が必要です。
            </p>
          </section>

          <section>
            <h2>4. 第三者提供</h2>
            <p>
              実装上、保存したセーブデータを外部サービスへ送信する処理はありません。
              法令に基づく場合を除き、取得した情報を第三者へ販売・貸与しません。
            </p>
          </section>

          <section>
            <h2>5. 保存期間と削除</h2>
            <p>
              セーブデータと統計用データは、サービス提供に必要な期間保存します。
              削除希望がある場合は運営者へ連絡してください。本人確認後、対応可能な範囲で削除します。
            </p>
          </section>

          <section>
            <h2>6. セキュリティ</h2>
            <p>
              セーブデータの送受信と個別データ取得では、HMAC-SHA256による署名検証を行います。
              署名検証用の秘密情報はサーバー側で管理し、通常のAPIレスポンスには含めません。
            </p>
          </section>

          <section>
            <h2>7. 改定とお問い合わせ</h2>
            <p>
              本ポリシーは、サービス内容や実装の変更に応じて改定される場合があります。
              データの取り扱い、削除依頼、誤表示の報告は運営者までお問い合わせください。
            </p>
          </section>
        </div>

        <p className="policy-updated-at">最終更新日: 2026年4月30日</p>
      </article>
    </div>
  );
}

export default PrivacyPolicyPage;
