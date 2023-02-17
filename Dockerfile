# syntax=docker/dockerfile:1
FROM golang:1.19

WORKDIR /test

COPY . .

# Install Mockery to generate the mocks needed for the unit tests
RUN if go install github.com/vektra/mockery/v2@v2.20.0 ; \
then go generate ./... && \
go test ./pkg/... -cover && \
go test ./internal/... -cover ; \
fi

# Uncomment the following line if the mocks are present already in the repository. Ref: README.md
# RUN go test ./pkg/... -cover && go test ./internal/... -cover

# Run integration tests
CMD [ "go" , "test", "-v", "./integration/..." ]

