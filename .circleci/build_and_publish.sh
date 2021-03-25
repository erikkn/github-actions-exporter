#!/usr/bin/env bash

REGISTRY_URL=docker.tw.ee
COMMIT=${GITHUB_SHA:0:8}
TAG=v0.1.0-${COMMIT}

echo "Building images"

docker build --pull -t "${REGISTRY_URL}/actions-api-exporter:${TAG}" .

echo "${ARTIFACTORY_PASSWORD}" | docker login ${REGISTRY_URL} --username "${ARTIFACTORY_USER}" --password-stdin

echo "branch name $GITHUB_REF"
if [ "$GITHUB_REF" == *"refs/heads/main"* ]
then
  echo "Pushing image.."
  docker push "${REGISTRY_URL}/actions-api-exporter:${TAG}"
fi

echo "done."
