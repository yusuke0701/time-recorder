# Time Recorder

[![CircleCI](https://circleci.com/gh/yusuke0701/time-recorder.svg?style=svg)](https://circleci.com/gh/yusuke0701/time-recorder)

## 概要

時間を記録するツールです。

## 主な構成

- クライアント
  - Chrome 拡張ツール
  - Vue.js(https://github.com/Kocal/vue-web-extension)
- サーバー
  - Google Cloud Functions
  - Golang
- データベース
  - Google Cloud Datastore

## 動作確認方法

```
./server/scripts/deploy.sh <your_project_id>
cd client
npm install
npm run build:dev
```

上記手順でサーバー側のデプロイと client/dist フォルダができる。

ブラウザで `chrome://extentions` を開き、dist フォルダを読み込ませる。

詳細な手順は以下のサイトが参考になる。

https://toranoana-lab.hatenablog.com/entry/2020/04/23/174421

## やることリスト

- ポップアップを非表示にすると ID が消えるので、保存する場所を変更する
- 日付が分かりにくい
- API の URL を env にする
- 認証認可
  - https://cloud.google.com/functions/docs/securing/authenticating?hl=ja#google_sign-in
