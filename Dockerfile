# build stage
FROM golang:1.15 AS builder
WORKDIR /go/src/urlshortener
COPY . .
RUN make

# final stage
FROM ubuntu:20.04
WORKDIR /root/
COPY --from=builder /go/src/urlshortener/urlshortener .
COPY --from=builder /go/src/urlshortener/config/config.yaml ./config/config.yaml
ENTRYPOINT ["./urlshortener"]