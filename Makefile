ifeq (,$(TAG))
	tag = latest
else
	tag = $(TAG)
endif

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
	docker buildx build --platform linux/$(arch) -t  arthurcgc/dhc:$(tag) .
	docker push arthurcgc/dhc:$(tag)

.PHONY: deploy
deploy-push: build
	@if [ ! $(PROJECT_ID) ]; then \
					echo "please set project ID; i.e: export PROJECT_ID=<projectid>"; \
					exit 1; \
	fi
	@if [ ! $(REGISTRY_URL) ]; then \
					echo "please set the registry url; i.e: export REGISTRY_URL=registry.io"; \
					exit 1; \
	fi
	docker buildx build --platform linux/$(arch) -t  $(REGISTRY_URL)/$(PROJECT_ID)/arthurcgc/dhc:$(tag) .
	docker push $(REGISTRY_URL)/$(PROJECT_ID)/arthurcgc/dhc:$(arch)