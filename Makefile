qa: analyze test

analyze:
	@go vet ./...
	@go run honnef.co/go/tools/cmd/staticcheck@latest --checks=all ./...

build: qa
	@go build -o ./build/ncg .

build-docker:
	@docker build -t thenativeweb/ncg .

coverage: test
	@mkdir -p ./coverage
	@go test -coverprofile=./coverage/coverage.out ./...
	@go tool cover -html=./coverage/coverage.out -o ./coverage/coverage.html
	@open ./coverage/coverage.html

test:
	@go test -cover ./...

.PHONY: analyze \
	build \
	build-docker \
	coverage \
	qa \
	test
