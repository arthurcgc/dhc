.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/server cmd/server/main.go

.PHONY: build-arm
build-arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o build/server cmd/server/main.go

.PHONY: run
run: build
	./build/server
	
.PHONY: docker-push
docker-push: build
	docker build -t arthurcgc/dhc:latest .
	docker push arthurcgc/dhc

.PHONY: docker-push-arm64
docker-push-arm64: build-arm
	docker build -t arthurcgc/dhc:arm64 .
	docker push arthurcgc/dhc:arm64
