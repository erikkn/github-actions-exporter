#!/usr/bin/env bash

REGISTRY_URL=docker.tw.ee
COMMIT=${CIRCLE_SHA1:0:8}
TAG=v0.2.0-${COMMIT}

echo "Building images"

docker build --pull -t "${REGISTRY_URL}/actions-api-exporter:${TAG}" .

echo "${DEPLOY_REGISTRY_PASSWORD}" | docker login ${REGISTRY_URL} --username "${DEPLOY_REGISTRY_USERNAME}" --password-stdin

if [ "$CIRCLE_BRANCH" == "main" ]
then
  echo "Pushing image.."
  docker push "${REGISTRY_URL}/actions-api-exporter:${TAG}"
fi

echo "done."
