#!/bin/bash

GOARCH=arm64
GOOS=linux
GO_BUILD_OUTPUT=build/eden-ark
GO_MAIN=cmd/app/main.go

BUILD_DIR=build
DOCKERFILE=Dockerfile
ENV_FILE=.env
CERTS_DIR=certs

REMOTE_USER=lubos
REMOTE_HOST=pi
REMOTE_PATH=/home/lubos/containers/eden/eden-ark

IMAGE_NAME=eden-ark-image
CONTAINER_NAME=eden-ark
HOST_PORT=9091
CONTAINER_PORT=9091
NETWORK=lubos-pi

# ======================
# Local build
# ======================
echo "-- Building Go executable"
GOARCH=$GOARCH GOOS=$GOOS go build -o $GO_BUILD_OUTPUT $GO_MAIN

echo "-- Building react client"
pnpm --dir ./web build

echo "-- Copying files into build folder"
mkdir -p $BUILD_DIR
mkdir $BUILD_DIR/web
cp $DOCKERFILE $BUILD_DIR/
cp $ENV_FILE $BUILD_DIR/
cp -r web/dist $BUILD_DIR/web

echo "-- Zipping the build folder"
cd $BUILD_DIR
zip -r build.zip .
cd ..

echo "-- Sending build to the server"
scp $BUILD_DIR/build.zip $REMOTE_USER@$REMOTE_HOST:$REMOTE_PATH/

echo "-- Clearing local build folder"
rm -rf $BUILD_DIR

# ======================
# Remote deployment
# ======================
ssh $REMOTE_USER@$REMOTE_HOST << EOF
cd $REMOTE_PATH

unzip -oq build.zip
rm build.zip

cp -r ../certs .

echo "-- Stopping and removing old container if exists"
if podman ps -a --format '{{.Names}}' | grep -q "^$CONTAINER_NAME\$"; then
    podman stop $CONTAINER_NAME
    podman rm $CONTAINER_NAME
fi

echo "-- Removing old image if exists"
if podman images --format '{{.Repository}}:{{.Tag}}' | grep -q "^$IMAGE_NAME:latest\$"; then
    podman rmi -f $IMAGE_NAME
fi

echo "-- Building new image"
podman build -t $IMAGE_NAME .

echo "-- Running new container"
podman run -d --name $CONTAINER_NAME --network $NETWORK -p $HOST_PORT:$CONTAINER_PORT --env-file $ENV_FILE $IMAGE_NAME

echo "-- Removing dangling images"
podman image prune -f

rm -rf certs

echo "-- Deploy successful"
EOF
