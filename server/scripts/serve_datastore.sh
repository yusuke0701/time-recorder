#!/usr/bin/env bash
set -exo

cd $(dirname $0)

gcloud beta emulators datastore start
