#!/bin/sh

# verify if exists go.mod and go.sum files, if not exists, create it
if [ ! -f go.mod ] || [ ! -f go.sum ]; then
  go mod init github.com/gt2rz/status-bar/api
fi

# install air for hot reload
go install github.com/cosmtrek/air@latest

# install dependencies
go mod tidy

# download dependencies
go mod download

# run air
air -c .air.toml
