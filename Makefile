# Makefile for Form3 client #
#############################

.PHONY: vendor
vendor:
	@echo "==> Go vendor. Gathering all exteranl dependencies in vendor folder 🎲 <=="
	go mod vendor -v

.PHONY: mock
mock:
	@echo "==> Generating mocks for unit test 🇪🇸 <=="
	go generate ./...

.PHONY: mockdocker
mockdocker:
	@echo "==> Generation mocks for unit test using Docker"
	docker run -v "$PWD":/src -w /src vektra/mockery --all --inpackage

.PHONE: rmmock
rmmock:
	@echo "==> Removing all mock files  <=="
	find . -name 'mock_*' -type f -delete

.PHONY: test
test:
	@echo "==> Running Unit Tests 🇨🇦 <=="
	go test ./pkg/...
	go test ./internal/...

.PHONY: testmock
testmock:
	@echo "==> Generating mocks and then run unit test altogether 🏀 <=="
	make mock
	make test

.PHONY: coverage
ut-coverage:
	@echo "==> Running Unit Test with coverage 🎮 <=="
	# TODO: run unit test with coverage

.PHONY: integration
integration:
	@echo "==> Running Integration Test 🎵 <=="
	# TODO: run integration test. Include dependencies for docker-compose.
	go test ./integration

.PHONE: clean
clean:
	@echo "==> Cleaning ..."
	docker compose down
	docker image rm form3-test
	docker builder prune
