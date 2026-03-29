.PHONY: gen proto.gen proto.lint proto.update proto.format proto.check
.PHONY: test tidy

BUF ?= buf
PROTO_TEMPLATE ?= buf.go.gen.yaml
GO_ENV ?= GOWORK=off

# Run these targets from each protocol module directory, e.g. server/tcp.
gen: proto.gen

proto.gen:
	$(BUF) generate --template $(PROTO_TEMPLATE)

proto.lint:
	$(BUF) lint

proto.update:
	$(BUF) dep update

proto.format:
	$(BUF) format -w

proto.check: proto.lint proto.gen

test:
	$(GO_ENV) go test ./...

tidy:
	$(GO_ENV) go mod tidy
