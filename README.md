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

mise という tool を利用しています

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
