
all: deps build test lint

deps:
	@echo "Installing dependencies"
	@glide install

test:
	@echo "Running tests.."
	@go test `glide novendor`

lint:
	@echo "Performing linting"
	@gometalinter --config=metalinter.json

build:
	@echo "Building project"
	@go build
