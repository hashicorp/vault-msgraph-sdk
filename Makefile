APIS ?= Applications.yml

default: generate

.PHONY: default generate test

%.yml:
	curl -q https://raw.githubusercontent.com/microsoftgraph/msgraph-sdk-powershell/dev/openApiDocs/v1.0/$@ > $@
	docker run -v ./$(shell echo $(subst .yml,,$@) | tr '[:upper:]' '[:lower:]'):/app/output \
		-v ./$@:/app/openapi.yaml \
		mcr.microsoft.com/openapi/kiota generate \
		--language go \
		-c ApiClient \
		-n github.com/hashicorp/vault-msgraph-sdk/$(shell echo $(subst .yml,,$@) | tr '[:upper:]' '[:lower:]') \
		--exclude-backward-compatible \
		--additional-data=False
	go mod tidy
	copywrite headers

generate: $(APIS)

test:
	go build .
	go test ./...