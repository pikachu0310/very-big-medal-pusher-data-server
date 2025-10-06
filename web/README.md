# Very Big Medal Pusher - Frontend

このディレクトリには、Very Big Medal Pusher のフロントエンドアプリケーションが含まれています。

## 技術スタック

- **React 19** + **TypeScript**
- **Vite** - ビルドツール
- **Mantine** - UIコンポーネントライブラリ
- **Tailwind CSS** - スタイリング
- **React Router** - ルーティング
- **SWR** - データフェッチング
- **Jotai** - 状態管理

## セットアップ

### 必要な環境

- Node.js 20以上
- npm または yarn

### インストール

```bash
# webディレクトリに移動
cd web

# 依存関係をインストール
npm install
```

### 開発サーバーの起動

```bash
# 開発サーバーを起動（http://localhost:3000）
npm run dev
```

開発サーバーは自動的にAPIサーバー（http://localhost:8080）にプロキシします。

### ビルド

```bash
# プロダクション用にビルド
npm run build

# ビルドしたアプリをプレビュー
npm run preview
```

## ディレクトリ構造

```
web/
├── src/
│   ├── pages/          # ページコンポーネント
│   ├── App.tsx         # メインアプリコンポーネント
│   ├── main.tsx        # エントリーポイント
│   └── index.css       # グローバルスタイル
├── index.html          # HTMLテンプレート
├── package.json        # 依存関係
├── vite.config.ts      # Vite設定
├── tsconfig.json       # TypeScript設定
├── tailwind.config.js  # Tailwind CSS設定
└── postcss.config.js   # PostCSS設定
```

## 開発時の注意事項

- APIサーバーが `http://localhost:8080` で起動している必要があります
- フロントエンドは `http://localhost:3000` で起動します
- `/api/*` へのリクエストは自動的にバックエンドにプロキシされます

