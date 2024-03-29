#!/bin/sh

set -xe

cf api $PWS_API
cf login -u $PWS_USER -p $PWS_PWD -o "$PWS_ORG" -s "$PWS_SPACE"

set +e
cf apps | grep "$PWS_APP_SUFFIX.$PWS_APP_DOMAIN" | grep "$PWS_APP_SUFFIX-green"
if [ $? -eq 0 ]
then
  echo "green" > ./app-info/current-app.txt
  echo "blue" > ./app-info/next-app.txt
else
  echo "blue" > ./app-info/current-app.txt
  echo "green" > ./app-info/next-app.txt
fi
set -xe

echo "Current main app routes to app instance $(cat ./app-info/current-app.txt)"
echo "New version of app to be deployed to instance $(cat ./app-info/next-app.txt)"
