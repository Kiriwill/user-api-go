#!/usr/bin/env bash
#
# Usage
# ./ensure.sh
# to ensure that all the CLI tools required to develop and run the project are installed
set -e


if ! command -v migrate &> /dev/null
then
    go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.14.1
fi

if ! command -v air &> /dev/null
then
    curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
fi

if ! command -v docker &> /dev/null
then
    echo "MISSING DEPENDENCY: docker is not installed"
    exit 1
fi

# IF THE PROJECT IS USING DOCKER-COMPOSE
# if ! command -v docker-compose &> /dev/null
# then
#     echo "MISSING DEPENDENCY: docker-compose is not installed"
#     exit 1
# fi

# you should also run the docker compose up to start the dependencies here
# docker compose up -d
