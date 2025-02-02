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

# 環境構築
mise という tool を利用しています
```
$ brew install mise

$ mise install node@22.13.1
$ mise use node@22.13.1

$ mise install pnpm@[mise.toml pnpm version]
$ mise use pnpm@[mise.toml pnpm version]

$ mise install go@[mise.toml go verion]
$ mise use go@[mise.toml go verion]
```