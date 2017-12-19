#!/bin/bash

GOOS=darwin GOARCH=amd64 go build -o om-lockdown-darwin
GOOS=linux GOARCH=amd64 go build -o om-lockdown-linux
GOOS=windows GOARCH=amd64 go build -o om-lockdown-windows
