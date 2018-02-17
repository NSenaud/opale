SERVER_OUT := "bin/opale-server"
CLIENT_OUT := "bin/opale"
API_OUT := "api/api.pb.go"
PKG := "github.com/NSenaud/opale"
SERVER_PKG_BUILD := "${PKG}/opale-server"
CLIENT_PKG_BUILD := "${PKG}/opale"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)

.PHONY: all api server cli

all: server cli

api: api/api.proto
	@protoc -I api/ \
		-I${GOPATH}/src \
		--go_out=plugins=grpc:api \
		api/api.proto

dep: ## Get the dependencies
	@go get -v -d ./...

server: dep api ## Build the binary file for server
	@go build -i -v -o $(SERVER_OUT) $(SERVER_PKG_BUILD)

cli: dep api ## Build the binary file for client
	@go build -i -v -o $(CLIENT_OUT) $(CLIENT_PKG_BUILD)

clean: ## Remove previous builds
	@rm $(SERVER_OUT) $(CLIENT_OUT) $(API_OUT)

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

