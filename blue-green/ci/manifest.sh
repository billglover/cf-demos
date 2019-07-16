#!/bin/sh

set -xe

sed -E "s/(- name: .*)/\1-$(cat ./app-info/next-app.txt)/" cf-demos/blue-green/manifest.yml > artefacts/manifest.yml
cat artefacts/manifest.yml
