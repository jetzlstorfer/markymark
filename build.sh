#!/bin/bash

# Mac
env GOOS=darwin GOARCH=amd64 go build -o bin/markymark-osx

# Linux 
env GOOS=linux GOARCH=amd64 go build -o bin/markymark-linux

# Windows
env GOOS=windows GOARCH=amd64 go build -o bin/markymark-win.exe