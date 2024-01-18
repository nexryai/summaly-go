FROM golang:alpine AS builder

WORKDIR /app

COPY . ./

RUN apk add --no-cache git ca-certificates build-base \
 && go build -ldflags="-s -w" -trimpath -o summaly main.go

FROM alpine:3.12

WORKDIR /app

COPY --from=builder /app/summaly /app/summaly

RUN apk add --no-cache ca-certificates tini \
 && addgroup -g 749 app \
 && adduser -u 749 -G app -D -h /app app \
 && chmod +x /app/summaly \
 && chown -R app:app /app

USER app
CMD [ "tini", "--", "/app/summaly" ]