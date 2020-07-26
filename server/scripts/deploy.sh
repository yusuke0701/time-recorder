#!/usr/bin/env bash
set -exo

cd $(dirname $0)

projectID="$(gcloud config get-value project)"
if [ -z "$projectID" ]; then
    echo "please set project id to the Gcloud SDK."
    exit 1
fi

cd ../src

functions=(
    "Start"
    "End"
)
for func in "${functions[@]}"; do
    gcloud functions deploy $func \
        --runtime go113 --trigger-http --set-env-vars GOOGLE_CLOUD_PROJECT=$projectID --allow-unauthenticated
done

# ref: https://cloud.google.com/functions/docs/env-var?hl=ja
