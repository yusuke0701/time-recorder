#! /bin/sh -eux

cd `dirname $0`

cd ../src

projectID="`gcloud config get-value project`"

gcloud functions deploy HelloGet \
--runtime go113 --trigger-http --set-env-vars GOOGLE_CLOUD_PROJECT=$projectID --allow-unauthenticated

# ref: https://cloud.google.com/functions/docs/env-var?hl=ja