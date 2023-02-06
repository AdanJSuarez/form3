# Makefile for Form3 client #
#############################

.PHONY: vendor
vendor:
	@echo "==> Go vendor. Gathering all exteranl dependencies in vendor folder <=="
	go mod vendor -v

.PHONY: mock
mock:
	@echo "==> Generating mocks for unit test <=="
	# TODO: Generate mocks

.PHONY: test
unittest:
	@echo "==> Running Unit Tests <=="
	go test ./...

.PHONY: testmock
	@echo "==> Generating mocks and then run unit test altogether <=="
	make mock
	make test

.PHONY: coverage
ut-coverage:
	@echo "==> Running Unit Test with coverage <=="
	# TODO: run unit test with coverage

.PHONY: integration
integrationtest:
	@echo "==> Running Integration Test <=="
	# TODO: run integration test. Include dependencies for docker-compose.
