# 自動保存するマークダウンエディタを作る予定のリポジトリ(Qiitaみたいなやつ)

# 使用技術

- Vite
- Typescript(Vue3)
- GO
- Docker

# 環境構築

- リポジトリをクローンする
`https://github.com/yuto-tomita/auto-save-markdown-editor.git`

- コンテナの起動
`docker compose up`

- DBコンテナ内に入り、draftテーブルを作成する

```
// dbコンテナ内に入る
$ docker compose exec db bash

// mysqlサーバーに入る
$ mysql -u root -p rootpassword

$ use test-db;

$ create table draft (id int auto_increment, markdown_text text, user_id int, index(id));
```
