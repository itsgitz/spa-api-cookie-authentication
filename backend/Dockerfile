FROM golang:alpine AS builder

LABEL maintainer="Anggit M Ginanjar <anggit.ginanjar.dev@gmail.com>"

WORKDIR /srv

COPY . .

RUN GOOS=linux go build
RUN ls -lah

FROM alpine
COPY --from=builder /srv/go-api /srv

ENTRYPOINT ["/srv/go-api"]
