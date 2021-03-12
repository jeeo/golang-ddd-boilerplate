build: wire-gen
	go build -o bin/app github.com/Jeeo/golang-ddd-boilerplate/cmd/some-app

wire-gen:
	cd ./cmd/some-app && wire