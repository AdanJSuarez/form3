# Makefile for Form3 client #
#############################

.PHONY: vendor
vendor:
	@echo "==> Go vendor. Gathering all exteranl dependencies in vendor folder <=="
	go mod vendor -v

.PHONY: mock
mock:
	@echo "==> Generating mocks for unit test <=="
	go generate ./...

.PHONY: test
test:
	@echo "==> Running Unit Tests <=="
	go test ./...

.PHONY: testmock
testmock:
	@echo "==> Generating mocks and then run unit test altogether <=="
	make mock
	make test

.PHONY: coverage
ut-coverage:
	@echo "==> Running Unit Test with coverage <=="
	# TODO: run unit test with coverage

.PHONY: integration
integration:
	@echo "==> Running Integration Test <=="
	# TODO: run integration test. Include dependencies for docker-compose.
	go run integration/integration.go
