BINARY_NAME = demo-smtp
TEST_COMMAND = gotest

# auto restart using air file main.go in cmd/$(BINARY_NAME)/main.go
.PHONY: dev
dev:
	air -c .air.toml



