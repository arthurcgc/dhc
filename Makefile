.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/server cmd/server/main.go

.PHONY: run
run: build
	./build/server
	
