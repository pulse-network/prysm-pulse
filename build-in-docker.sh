#!/bin/bash

#################################################################
# Building working docker images currently required Ubuntu-18.04.
# This will build beacon-chain and validator docker images within
# docker (Bazel-in-Docker), and import the built images into the
# host docker.
#################################################################

# exit on failures
set -e

# build the prysm-builder image
docker build --platform linux/amd64 -f Dockerfile.builder -t prysm-builder .

# run the build
docker run --platform linux/amd64 -it -v $(pwd):/app -v $(pwd)/.bazel-in-docker-cache:/root/.cache prysm-builder

# source the outputs file to build the correct symlink directories bazel-*
source .bazel-in-docker-outputs
rm .bazel-in-docker-outputs

# load new images into docker
docker load -i bazel-bin/cmd/beacon-chain/image_bundle.tar
docker load -i bazel-bin/cmd/validator/image_bundle.tar
docker load -i bazel-bin/cmd/prysmctl/image_bundle.tar

# finish up
echo ""
echo "Build Complete! Bazel outputs are available in the following directories:"
ls | grep bazel- | tr '\n' '\t'
echo ""

# print new docker images
echo ""
echo "Docker images imported:"
docker image ls | head
