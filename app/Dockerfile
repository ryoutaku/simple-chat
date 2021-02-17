FROM golang:1.15.3 AS builder
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /go/bin/app -ldflags '-s -w'

FROM scratch
COPY --from=builder /go/bin/app /go/bin/app
ENTRYPOINT ["/go/bin/app"]
