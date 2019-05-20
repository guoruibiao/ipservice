SERVICE ?= ipservice
BIN ?= ./${SERVICE}

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BIN}
