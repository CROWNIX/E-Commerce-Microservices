.PHONY: clean

clean:
	rm -rf dist/* generated build vendor
	find . -name "*.mock.gen.go" -type f -delete
	find . -name "*.out" -type f -delete
	find . -name "wire_gen.go" -type f -delete
	find . -name "*.mock.gen.go" -type f -delete

generate_api: api/openapi/api.yaml
	npx openapi-format api/openapi/api.yaml -s api/openapi/openapi-sort.json -f api/openapi/openapi-filter.json -o api/openapi/api.yaml
	mkdir -p generated/api
	go tool oapi-codegen --package api -generate types $< > generated/api/api-types.gen.go

generate_wire:
	go tool wire ./...

preview_open_api:
	redocly preview-docs api/openapi/api.yaml

run-api:
	cd cmd && go run . rest-api