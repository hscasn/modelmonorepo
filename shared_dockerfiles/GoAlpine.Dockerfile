FROM golang:latest as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o app internal/main.go

# --

FROM alpine

# Timezones
RUN apk --no-cache add tzdata

COPY --from=builder /app/app /app
COPY /public /public

CMD ["/app"]
