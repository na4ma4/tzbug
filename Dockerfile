ARG GO_VERSION
ARG ALPINE_VERSION
FROM golang:${GO_VERSION} as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o bin/test

FROM alpine:${ALPINE_VERSION}
RUN apk --no-cache add tzdata
COPY --from=builder /app/bin/test /test
CMD ["/test"]