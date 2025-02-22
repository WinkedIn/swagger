# force makefile to be run with bash instead of sh
SHELL := /bin/bash

default: tidy generate

.PHONY: clean
clean:
	rm -fr onboarding

.PHONY: generate
generate:
	go install -mod=readonly github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.1.0
	for api in onboarding ; do \
  		for version in v1; do \
			[[ ! -f api/$${version}/$${api}.yml ]] && continue ;\
			rm -fr $${api}/$${version}; \
			mkdir -p $${api}/$${version}; \
			oapi-codegen --config api/$${version}/config/model.yaml api/$${version}/$${api}.yml > $${api}/$${version}/models.gen.go ;\
			oapi-codegen --config api/$${version}/config/client.yaml api/$${version}/$${api}.yml > $${api}/$${version}/client.gen.go; \
			oapi-codegen --config api/$${version}/config/server.yaml api/$${version}/$${api}.yml > $${api}/$${version}/server.gen.go; \
  		done; \
  	done

.PHONY: lint
lint:
	npx @redocly/cli lint api/v1/onboarding.yml

.PHONY: tidy
tidy:
	go mod tidy -v
