FROM golang:latest as builder

WORKDIR /app

COPY . .

ARG SSH_PRIVATE_KEY
RUN git config --global url.ssh://git@github.com/.insteadOf https://github.com/
RUN mkdir /root/.ssh/
RUN touch /root/.ssh/known_hosts
RUN ssh-keyscan github.com >> /root/.ssh/known_hosts
RUN echo "${SSH_PRIVATE_KEY}" | base64 -d > /root/.ssh/id_rsa
RUN chmod 600 /root/.ssh/id_rsa

RUN CGO_ENABLED=0 GOOS=linux go build -v -o app internal/main.go

# --

FROM alpine

# Timezones
RUN apk --no-cache add tzdata

COPY --from=builder /app/app /app
COPY /public /public

CMD ["/app"]
