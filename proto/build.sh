#!/bin/sh

protoc --proto_path=. user/*.proto --go_out=plugins=grpc:.
