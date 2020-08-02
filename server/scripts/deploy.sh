#!/usr/bin/env bash
set -exo

# 使い方
# 1. 引数なしで実行すると、全ての関数をデプロイする。
# 2. `./deploy.sh Start` のように第一引数には、デプロイしたい関数名を入力する。

cd $(dirname $0)

projectID="$(gcloud config get-value project)"
if [ -z "$projectID" ]; then
    echo "please set project id to the Gcloud SDK."
    exit 1
fi

if [ "$1" ]; then
    functions=(
        $1
    )
else
    functions=(
        "Start"
        "End"
        "GetLastRecord"
        "ListRecord"
    )
fi

cd ../src

for func in "${functions[@]}"; do
    gcloud functions deploy $func \
        --runtime go113 --trigger-http --set-env-vars GOOGLE_CLOUD_PROJECT=$projectID --allow-unauthenticated
done

# ref: https://cloud.google.com/functions/docs/env-var?hl=ja
