#!/bin/bash
set -e

REMOTE_USER=lubos
REMOTE_HOST=pi
REMOTE_PATH=/home/lubos/containers/eden/eden-ark
BUILD_DIR=build
GO_MAIN=cmd/app/main.go

# ======================
# Local build
# ======================
echo "-- Building Go executable"
GOARCH=arm64 GOOS=linux go build -o $BUILD_DIR/eden-ark $GO_MAIN

echo "-- Building React frontend"
pnpm --dir ./web build

echo "-- Preparing build folder"
mkdir -p $BUILD_DIR/web
cp -r web/dist $BUILD_DIR/web

echo "-- Zipping build folder"
cd $BUILD_DIR
zip -r build.zip .
cd ..

echo "-- Sending build to Raspberry Pi"
scp $BUILD_DIR/build.zip $REMOTE_USER@$REMOTE_HOST:$REMOTE_PATH/

echo "-- Cleaning local build"
rm -rf $BUILD_DIR

# ======================
# Remote deployment
# ======================
ssh $REMOTE_USER@$REMOTE_HOST << EOF
cd $REMOTE_PATH

echo "-- Extracting build"
unzip -oq build.zip
rm build.zip

echo "-- Stopping previous containers"
podman-compose down

echo "-- Building and starting containers"
podman-compose up -d --build --force-recreate

echo "-- Removing dangling images"
podman image prune -f

echo "-- Deployment complete"
EOF
