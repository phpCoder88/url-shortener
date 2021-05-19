APP = server
BUILD_DIR = build
REPO = $(shell go list -m)
BUILD_DATE = $(shell date +%FT%T%Z)
BUILD_COMMIT = $(shell git rev-parse HEAD)
VERSION = $(if $(TAG),$(TAG),$(if $(BRANCH_NAME),$(BRANCH_NAME),$(shell git describe --tags --exact-match || git symbolic-ref -q --short HEAD)))

GO_BUILD_ARGS = \
  -ldflags " \
    -X '$(REPO)/internal/version.Version=$(VERSION)' \
    -X '$(REPO)/internal/version.BuildCommit=$(BUILD_COMMIT)' \
    -X '$(REPO)/internal/version.BuildDate=$(BUILD_DATE)' \
  " \

.PHONY: build
build:
	@echo "+ $@"
	@mkdir -p $(BUILD_DIR)
	go build -race $(GO_BUILD_ARGS) -o $(BUILD_DIR) ./cmd/server

.PHONY: test
test:
	@echo "+ $@"
	go test -cover ./...

.PHONY: test-cover
test-cover:
	@echo "+ $@"
	go test -coverprofile=profile.out ./...
	go tool cover -html=profile.out
	rm profile.out

.PHONY: check
check:
	golangci-lint run

.PHONY: run
run: clean build
	@echo "+ $@"
	./${BUILD_DIR}/${APP}

.PHONY: clean
clean:
	@rm -rf $(BUILD_DIR)

#.PHONY: watch
#watch:
#	reflex -r "\.txt$" echo {}
#	#reflex -r '\.go$' -s -- sh -c "go run ./cmd/server/main.go"

.PHONY: watch
watch: go_prep_watch
	#reflex -s -r '\.go$$' make run
	reflex -r '\.go$$' -s -- sh -c "go run ./cmd/server/main.go"

go_prep_watch:
	@echo "\nPreparing environment...."
	go get github.com/cespare/reflex
