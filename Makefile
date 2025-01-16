generate:
	@oapi-codegen -generate "types" -package gen -o api/gen/types.go api/spec/api.yaml
	@oapi-codegen -generate "chi-server" -package gen -o api/gen/handlers_gen.go api/spec/api.yaml

run:
	@go run main.go