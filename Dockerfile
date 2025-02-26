FROM golang:1-alpine AS builder

WORKDIR /app

COPY . ./

ENV CGO_ENABLED=0
RUN go build -trimpath -o summaly main.go

FROM gcr.io/distroless/static-debian12

WORKDIR /app

COPY --from=builder /app/summaly /app/summaly

USER 749
ENTRYPOINT [ "/app/summaly" ]
