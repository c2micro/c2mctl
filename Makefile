BIN_DIR=$(PWD)/bin
C2M_DIR=$(PWD)/cmd/c2mctl
CC=gcc
CXX=g++

.PHONY: c2mctl
c2mctl:
	@mkdir -p ${BIN_DIR}
	@echo "Building management cli..."
	CGO_ENABLED=0 CC=${CC} CXX=${CXX} go build -o ${BIN_DIR}/c2mctl ${C2M_DIR}
	@strip bin/c2mctl

.PHONY: go-sync
go-sync:
	@go mod tidy && go mod vendor

.PHONY: dep-shared
dep-shared:
	@echo "Update shared components..."
	@export GOPRIVATE="github.com/c2micro" && go get -u github.com/c2micro/c2mshr/ && go mod tidy && go mod vendor

