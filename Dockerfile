FROM golang:alpine AS builder

WORKDIR /app

ADD . ./
RUN go mod download
RUN go build -o /app/book-store ./cmd/api


FROM alpine:latest
RUN addgroup -g 1000 app
RUN adduser -u 1000 -G app -h /home/goapp -D goapp
USER goapp
WORKDIR /app
COPY --from=builder /app/book-store  /app/
COPY --from=builder /app/internal/book/repository/postgres/migrations  /app/internal/book/repository/postgres/migrations/

CMD ["./book-store"]
#CMD ["sh", "-c", "tail -f /dev/null"]