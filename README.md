# Time Recorder

[![CircleCI](https://circleci.com/gh/yusuke0701/time-recorder.svg?style=svg)](https://circleci.com/gh/yusuke0701/time-recorder)

## 概要

時間を記録するツールです。

## 主な構成

- クライアント
  - Firebase Hosting
  - Vue 2
- サーバー
  - Google Cloud Functions
  - Golang
- データベース
  - Google Cloud Datastore

## 動作確認方法
下記URLにて動作確認可能
https://hoge-hoge-123456789.web.app

以下の方法で、ユーザー登録できる

1. https://hoge-hoge-123456789.web.app/signup へアクセス
1. 登録画面へ繋がるので、メールアドレスとパスワードを入力する

以下の方法で、ログインできる

1. https://hoge-hoge-123456789.web.app/signin へアクセス
1. ログイン画面へ繋がるので、メールアドレスとパスワードを入力する