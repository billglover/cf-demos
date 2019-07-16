#!/bin/sh

set -xe

sed -E "s/(- name: .*)/\1-$(cat ./app-info/next-app.txt)/" cf-demos/blue-green/manifest.yml > cf-demos/blue-green/manifest.yml.tmp
sed -E "s/(.*)unknown(.*)/\1$(cat ./app-info/next-app.txt)\2/" cf-demos/blue-green/manifest.yml.tmp > artefacts/manifest.yml
cat artefacts/manifest.yml
