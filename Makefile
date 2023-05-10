.PHONY: build
build:
	go build -o build/server cmd/server/main.go

.PHONY: run
run: build
	./build/server
	
