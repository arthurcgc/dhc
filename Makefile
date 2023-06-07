.PHONY: build
build:
	@if [ ! $(arch) ]; then \
					echo "architecture parameter is required... use: make build arch=<adm64/arm64>"; \
					exit 1; \
	fi
	CGO_ENABLED=0 GOOS=linux GOARCH=$(arch) go build -o build/server cmd/server/main.go

.PHONY: run
run: build
	./build/server
	
.PHONY: docker-push
docker-push: build
	docker build -t arthurcgc/dhc:$(arch) .
	docker push arthurcgc/dhc:$(arch)
