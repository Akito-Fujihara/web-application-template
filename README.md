# web-application-template

## フロントエンド

- TypeScript
  - フレームワーク: NextJS
  - UI コンポーネント: Mantine

## バックエンド

- Go
  - Web フレームワーク: echo
  - ORM: gorm-gen

で実装された Web アプリケーションの雛形です

## API スキーマ
- OpenAPI
  - TypeSpec

# 環境構築

## ツールのインストール
mise を利用しています

```
$ brew install mise

$ mise install direnv@[mise.toml direnv version]
$ mise use direnv@[mise.toml direnv version]

$ mise install node@[mise.toml node version]
$ mise use node@[mise.toml node version]

$ mise install pnpm@[mise.toml pnpm version]
$ mise use pnpm@[mise.toml pnpm version]

$ mise install go@[mise.toml go verion]
$ mise use go@[mise.toml go verion]
```

環境変数の追加で .envrc を設定するために
```
direnv allow .
```

pnpm の install
```
$ pnpm install
```

go の install
```
$ go mod tidy
```

## コンテナの起動
mysql などコンテナを設定
```
docker compose up --build
```

# 開発ツール

## migrate
```
# migrate up
$ make migrate.up

# migrate down
$ make migrate.down

# migrate new (migrate file の生成)
$ make migrate.new name=[table名]
```
