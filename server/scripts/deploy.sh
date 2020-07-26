#! /bin/sh -eux

cd `dirname $0`

cd ../src

gcloud functions deploy HelloGet \
--runtime go113 --trigger-http --allow-unauthenticated