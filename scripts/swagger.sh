#!/usr/bin/env bash

cd ../internal/app

swag init --parseDependency -g ../../cmd/gin_app/main.go
