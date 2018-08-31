FROM golang:alpine AS builder

RUN	apk add --no-cache \
	ca-certificates

COPY  . /go/src/github.com/mboldysh/streaming-service

RUN set -x \
    && apk add --no-cache --virtual .build-deps \
        make \
    && cd /go/src/github.com/mboldysh/streaming-service \
    && make build \
    && mv streaming-service /usr/bin/streaming-service \
    && apk del .build-deps \
    && rm -rf /go \
    && echo "Build complete"

FROM scratch

COPY --from=builder /usr/bin/streaming-service /usr/bin/streaming-service
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/ 

EXPOSE 8080 

CMD ["streaming-service"]