#!/bin/sh

rm bin/lsfr*
env GOOS=linux GOARCH=amd64 go build -o bin/lfsr
env GOOS=windows GOARCH=amd64 go build -o bin/lfsr.exe

