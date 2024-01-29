FROM alpine:edge AS builder

WORKDIR /app

COPY . ./

RUN sed -i 's#https\?://dl-cdn.alpinelinux.org/alpine#https://mirrors.xtom.com.hk/alpine#g' /etc/apk/repositories \
 && apk add --no-cache go git ca-certificates build-base \
 && go build -ldflags="-s -w" -trimpath -o summaly main.go

FROM alpine:edge

WORKDIR /app

COPY --from=builder /app/summaly /app/summaly

RUN sed -i 's#https\?://dl-cdn.alpinelinux.org/alpine#https://mirrors.xtom.com.hk/alpine#g' /etc/apk/repositories \
 && apk add --no-cache ca-certificates tini \
 && addgroup -g 749 app \
 && adduser -u 749 -G app -D -h /app app \
 && chmod +x /app/summaly \
 && chown -R app:app /app

USER app
CMD [ "tini", "--", "/app/summaly" ]
