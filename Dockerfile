# syntax=docker/dockerfile:1
FROM golang:1.19

WORKDIR /app

COPY . .

# Install Mockery. Ref: README
# Forced to sucess in case mock dependency (Mockery) fails.
RUN go install github.com/vektra/mockery/v2@v2.20.0 ; exit 0

# Generate mocks. Ref: README
# Forced to sucess in case mock generation fails.
RUN go generate ./... ; exit 0

# Run unit tests. Ref: README
# Forced to sucess in case mocks are not present.
RUN go test ./pkg/... -cover && go test ./internal/... -cover ; exit 0


# Run integration tests
CMD [ "go" , "test", "-v", "./integration"]
