#!/bin/bash
set -e

# Variablen anpassen:
IMAGE_NAME="healthcheck"
QUAY_NAMESPACE="seidner"  # z.B. meinaccount
QUAY_REPO="healthcheck"
TAG="latest"

# Volle Image-URL
IMAGE_FULL_NAME="quay.io/${QUAY_NAMESPACE}/${QUAY_REPO}:${TAG}"

echo "Building image ${IMAGE_NAME}..."
podman build -t ${IMAGE_NAME} .

echo "Tagging image as ${IMAGE_FULL_NAME}..."
podman tag ${IMAGE_NAME} ${IMAGE_FULL_NAME}

echo "Logging in to Quay.io..."
podman login quay.io

echo "Pushing image to Quay.io..."
podman push ${IMAGE_FULL_NAME}

echo "Fertig! Image gepusht: ${IMAGE_FULL_NAME}"
