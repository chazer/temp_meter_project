FROM alpine

RUN apk add --no-cache curl jq

COPY cli /usr/local/bin

ENTRYPOINT ["/bin/sh", "-c", "exec \"$@\"", "-"]

CMD ["true"]
