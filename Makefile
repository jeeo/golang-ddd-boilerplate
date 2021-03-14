build: download wire-gen
	@go build -o app github.com/Jeeo/golang-ddd-boilerplate/cmd/some-app

wire-gen:
	@cd ./cmd/some-app && wire

download:
	@go mod download

build-docker:
	@docker build . -t jeeo/go_api 

test:
	@go test ./...

coverage:
	@go test -coverprofile=coverage.out
