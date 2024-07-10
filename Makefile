.PHONY: build build-UpdateTokenServicoFunction

build-UpdateTokenServicoFunction:
	GOOS=linux GOARCH=amd64 go build -o bootstrap main.go

build:
	sam build
