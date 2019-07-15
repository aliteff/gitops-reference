#=======
# Builder
#=======
FROM golang:alpine AS builder

WORKDIR /app
COPY . /app

RUN apk add ca-certificates tzdata git musl-dev
RUN GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o example .

#=======
# Runner
#=======
FROM scratch
WORKDIR /app

ARG VERSION
ARG BRANCH
ARG RELEASE

LABEL richcontext.version=$VERSION \
      richcontext.branch=$BRANCH \
      richcontext.release=$RELEASE

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/example /app/example

CMD ["/app/example"]
