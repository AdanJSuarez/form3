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

.PHONE: rmmock
rmmock:
	@echo "==> Removing all mock files  <=="
	find . -name 'mock_*' -type f -delete

.PHONY: test
test:
	@echo "==> Running Unit Tests 🎮 <=="
	go test ./pkg/... -cover
	go test ./internal/... -cover

.PHONY: testmock
testmock:
	@echo "==> Generating mocks and then run unit tests 🏀 <=="
	make mock
	make test

.PHONY: integration
integration:
	@echo "==> Running Integration Test 🎵 <=="
	go test -v ./integration

.PHONY: clean
clean:
	@echo "==> Cleaning ..."
	docker-compose down
	docker image rm form3-test
	docker builder prune

.PHONY: cover
cover:
	@echo "==> Visual coverage for $(FOLDER)"
	go test $(FOLDER) -coverprofile=coverage/lastCoverage.out
	go tool cover -html=coverage/lastCoverage.out
