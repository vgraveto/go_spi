#!/usr/bin/env bash
# Build windows version
# env GOOS=windows GOARCH=amd64 go build -o bin/spi_mcp3009.exe main.go
# Build linux version
env GOOS=linux GOARCH=amd64 go build -o bin/spi_mcp3009 main.go
# Build mac version
# env GOOS=darwin GOARCH=amd64 go build -o bin/spi_mcp3009_mac main.go
# Build RPi version
#env GOOS=linux GOARCH=arm GOARM=6 go build -o bin/spi_mcp3009_rpi main.go