# syntax=docker/dockerfile:1
FROM golang:1.19

WORKDIR /app

COPY . .

# Install Mockery to generate the mocks needed for the unit test
RUN if go install github.com/vektra/mockery/v2@latest ; \
then go generate ./... && \
go test ./pkg/... -cover && \
go test ./internal/... -cover ; \
fi


# Run integration tests
CMD [ "go" , "test", "-v", "./integration/..." ]

