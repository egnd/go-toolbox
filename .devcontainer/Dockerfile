FROM mcr.microsoft.com/devcontainers/go:0-1.20

RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get -y install --no-install-recommends \
        make grep graphviz

RUN wget https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh -O | sh -s -- -b $(go env GOPATH)/bin v1.50.1 && \
    golangci-lint --version

RUN go install github.com/vektra/mockery/v2@v2.16.0
