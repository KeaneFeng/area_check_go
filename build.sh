#!/bin/sh
rm -rf ./release
mkdir  release
go build -o area
chmod +x ./area
cp area ./release/
#cp favicon.ico ./release/
#cp -arf ./asset ./release/
cp -arf ./view ./release/
nohup ./release/area >>./area.log 2>&1 &