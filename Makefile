GOPATH:=$(shell go env GOPATH)
PREFIX=$(shell pwd)

.DEFAULT: all

all:build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o osec_log main.go

	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o mac_osec_log main.go
	
