BINARY_NAME = demo-smtp
TEST_COMMAND = gotest


.PHONY: build
build:
	go build -v -o $(BINARY_NAME) ./cmd/$(BINARY_NAME)


# auto restart using air file main.go in cmd/$(BINARY_NAME)/main.go
.PHONY: dev
dev:
	air -c .air.toml



