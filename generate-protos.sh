#!/bin/bash

rm -rf generated_protos/*
protoc -I=. --go_out=plugins=grpc:. protos/game.proto
cp -r protos/*.go generated-protos/
