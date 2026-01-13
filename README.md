# Massive Medal Pusher Data Server
![Coverage](https://img.shields.io/badge/coverage-20.5%25-brightgreen)
## https://push.trap.games/
VRChat ワールド「クソでっけぇプッシャーゲーム」向けのクラウドセーブ/API サーバーです。  
クラウドセーブ送信・最新セーブ取得・ランキング/統計を提供し、`oapi-codegen` で型安全に実装されています。クライアント実装は `/mnt/h/github/VRCWorld-MassiveMedalPusher` の Udon# です。

## エンドポイントとドキュメント
- ベース URL: `http://localhost:8080/api` / `https://push.trap.games/api` / `https://push-test.trap.games/api`
- Swagger UI: `GET /swagger/index.html` (リダイレクト含む)
- OpenAPI: `GET /api/openapi.yaml`
- API バージョン: `v1/v2/v3` は互換維持のみ（HTTP 410 応答）、`v4` が現行。`/ping` は `general` タグに集約。

## クイックスタート（ローカル）
```bash
docker compose watch   # API + MariaDB + Adminer
# ブラウザ: http://localhost:8080/api , http://localhost:8081 (Adminer)
```

フロントエンド（開発用ダッシュボード等）
```bash
cd web
npm install   # 初回
npm run dev   # http://localhost:3000
```

## 環境変数（主要）
```bash
SECRET_KEY=A                 # HMAC メイン
SAVE=B                       # セーブ署名キー
LOAD=C                       # ロード署名キー
APP_ADDR=:8080
DB_HOST=localhost DB_PORT=3306 DB_USER=root DB_PASSWORD=pass DB_NAME=app
# NeoShowcase 環境では NS_MARIADB_* 系を自動検出
```

## 開発フロー（OpenAPI-first）
1) `openapi/openapi.yaml` を編集  
2) `make oapi` で `openapi/server.gen.go`, `openapi/models/models.gen.go` を再生成  
3) ハンドラ実装：`internal/handler/`  
4) DB 変更が必要なら `internal/migration/*.sql` を追加（Goose）  
5) ドメイン/リポジトリ更新：`internal/domain/`, `internal/repository/`  
6) `go test ./...`

### よく使うコマンド
```bash
make oapi           # コード生成
make test           # 全テスト
make build          # バイナリ作成
make lint           # golangci-lint --fix
```

## データベース
- MariaDB 11 系。起動時に Goose で `internal/migration/*.sql` を Up 適用。
- スキーマの正とするドキュメントは `DATABASE.md`（migration を踏まえた最新版）。その他のスキーマファイルは削除済み。
- 主要テーブル：
  - `v2_save_data` + サブテーブル（実績/メダル/ボール/パレット/トーテム/パーク）
  - `v3_user_latest_save_data` (+ `_achievements`) … 最新セーブのランキング用集約
  - `v1_game_data` … 旧版互換

### 新規カラム追加手順（サーバー）
1. `openapi/openapi.yaml` に項目を追加（`x-oapi-codegen-extra-tags` で db カラム名を付与）。  
2. `internal/migration/NN_description.sql` を作成し、既存 migration を確認して型/制約を揃える。  
3. `make oapi` でコード再生成。  
4. `internal/domain/data_v2.go`（パース/モデル変換）、`internal/repository/*.go`（Insert/Select）を更新。  
5. 必要なら `DATABASE.md` のスキーマ表を更新。  
6. `go test ./...` と動作確認。

## アーキテクチャ概要
- `main.go` : Echo 起動、Swagger 配信、Goose migration。  
- `internal/handler/` : v1–v4 API（署名検証、Base64/URL decode、最新セーブ返却）。  
- `internal/domain/` : SaveData パース/変換（Base64 または URL エンコード対応、long 対応カラム）。  
- `internal/repository/` : sqlx で CRUD・ランキング取得。  
- `internal/migration/` : Goose SQL（Up/Down）。  
- `openapi/` : 生成コード。  
- `tools/` : oapi-codegen 設定。  
- `web/` : ダッシュボード/デバッグ UI。

## 運用メモ
- 本番/ステージングの Swagger からも spec を参照可能（UI で prod/stg/local を切替）。  
- セーブ送信は Base64URL、ロード応答は標準 Base64 + HMAC-SHA256 署名（LOAD シークレット）。  
- v1–v3 は互換スタブ（HTTP 410）。新規機能は v4 で実装。

## 関連リポジトリ
- クライアント (VRChat/Udon#): `/mnt/h/github/VRCWorld-MassiveMedalPusher`
