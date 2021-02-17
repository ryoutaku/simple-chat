FROM golang:1.15.3 AS builder
COPY . /api
WORKDIR /api
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o server -ldflags '-s -w'

FROM scratch
COPY --from=builder /api/server /
COPY --from=builder /api/.env /
ENTRYPOINT ["/server"]