# syntax=docker/dockerfile:1
FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

# Install Mockery to generate the mocks needed for the unit test
RUN go install github.com/vektra/mockery/v2@latest && go generate ./...

# Run unit test
RUN go test ./pkg/... && go test ./internal/...

# Run integration tests
CMD [ "go" , "test", "-v", "./integration/..." ]

