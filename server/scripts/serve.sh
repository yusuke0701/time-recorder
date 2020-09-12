#!/usr/bin/env bash
set -exo

cd $(dirname $0)

export GOOGLE_CLOUD_PROJECT=local

$(gcloud beta emulators datastore env-init)

cd ../src

go run cmd/main.go
