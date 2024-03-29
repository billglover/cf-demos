#!/bin/sh

set -xe

cf api $PWS_API
cf login -u $PWS_USER -p $PWS_PWD -o "$PWS_ORG" -s "$PWS_SPACE"

export PWS_DOMAIN_NAME=$PWS_APP_DOMAIN
export MAIN_ROUTE_HOSTNAME=$PWS_APP_SUFFIX

export NEXT_APP_COLOR=$(cat ./app-info/next-app.txt)
export NEXT_APP_HOSTNAME=$PWS_APP_SUFFIX-$NEXT_APP_COLOR

export CURRENT_APP_COLOR=$(cat ./app-info/current-app.txt)
export CURRENT_APP_HOSTNAME=$PWS_APP_SUFFIX-$CURRENT_APP_COLOR

echo "Mapping main app route to point to $NEXT_APP_HOSTNAME instance"
cf map-route $NEXT_APP_HOSTNAME $PWS_DOMAIN_NAME --hostname $MAIN_ROUTE_HOSTNAME

cf routes

echo "Removing previous main app route that pointed to $CURRENT_APP_HOSTNAME instance"

set +e
cf unmap-route $CURRENT_APP_HOSTNAME $PWS_DOMAIN_NAME --hostname $MAIN_ROUTE_HOSTNAME
set -e

echo "Routes updated"

cf routes
