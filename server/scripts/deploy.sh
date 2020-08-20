#!/usr/bin/env bash
set -exo

cd $(dirname $0)

projectID="$(gcloud config get-value project)"
if [ -z "$projectID" ]; then
    echo "please set project id to the Gcloud SDK."
    exit 1
fi

gcloud datastore indexes create index.yaml --project=$projectID --quiet

cd ../src

functions=(
    "GetGoogleID"
)

for func in "${functions[@]}"; do
    gcloud functions deploy $func \
        --runtime go113 --trigger-http --quiet --set-env-vars GOOGLE_CLOUD_PROJECT=$projectID

    gcloud functions add-iam-policy-binding $func \
            --member="serviceAccount:$projectID@appspot.gserviceaccount.com" \
            --role="roles/cloudfunctions.invoker"
done

allowUnauthenticatedFunctions=(
    "Records"
    "GetLastRecord"
    "ListRecord"
)

for func in "${allowUnauthenticatedFunctions[@]}"; do
    gcloud functions deploy $func \
        --runtime go113 --trigger-http --quiet --set-env-vars GOOGLE_CLOUD_PROJECT=$projectID --allow-unauthenticated
done

# ref: https://cloud.google.com/functions/docs/env-var?hl=ja
