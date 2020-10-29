FROM golang:1.15.3-alpine as builder

COPY tmeter /go/src

WORKDIR /go/src

ENV CGO_ENABLED=0 \
    GOOS=linux

RUN go build \
    -a -v \
    -ldflags="-s -w" \
    -gcflags=-trimpath="$(pwd)" \
    -asmflags=-trimpath="$(pwd)" \
    -o /go/bin/tmeter \
    .


FROM scratch

COPY --from=builder /go/bin/tmeter /bin/tmeter

ENV LISTEN_HOST=0.0.0.0 \
    LISTEN_PORT=8080 \
    JSON_INDENT_RESPONSES=no

EXPOSE 8080

ENTRYPOINT ["/bin/tmeter"]
