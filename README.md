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

1. クライアント側のビルド

```
cd client
npm install
npm run build:dev
```

上記手順で client/dist フォルダができる。

2. ブラウザで `chrome://extentions` を開き、client/dist フォルダを読み込ませる。

詳細な手順は以下のサイトが参考になる。

https://toranoana-lab.hatenablog.com/entry/2020/04/23/174421
