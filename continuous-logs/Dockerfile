FROM golang:latest
WORKDIR /go/src/continuous-logs
COPY . .

RUN go install -v ./...

# debug tools
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        curl \
        jq \
        netcat \
    && rm -rf /var/lib/apt/lists/*

CMD ["continuous-logs"]

