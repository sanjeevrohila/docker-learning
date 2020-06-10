FROM golang:latest
WORKDIR /go/src/continious-log
COPY . .

RUN go install -v ./...

# debug tools
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        curl \
        jq \
        netcat \
    && rm -rf /var/lib/apt/lists/*

CMD ["continious-log"]

