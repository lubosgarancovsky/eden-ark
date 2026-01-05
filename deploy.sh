#!/bin/bash
set -e

REMOTE_USER=lubos
REMOTE_HOST=pi
REMOTE_PATH=/home/lubos/containers/eden/eden-ark
BUILD_DIR=build
GO_MAIN=cmd/app/main.go
SERVICE_NAME='eden-ark'

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

echo "-- Stop and disable service"
systemctl --user stop container-$SERVICE_NAME.service
systemctl --user disable container-$SERVICE_NAME.service

echo "-- Building the service image"
podman-compose up -d --build --no-recreate

echo "-- Generating systemd unit"
podman generate systemd --name $SERVICE_NAME --files --new
mkdir -p ~/.config/systemd/user/
mv container-$SERVICE_NAME.service ~/.config/systemd/user/
podman stop $SERVICE_NAME

echo "-- Staring and enabling service"
systemctl --user daemon-reload
systemctl --user enable container-$SERVICE_NAME.service
systemctl --user start container-$SERVICE_NAME.service

podman image prune -f

echo "-- Deployment complete"
EOF
