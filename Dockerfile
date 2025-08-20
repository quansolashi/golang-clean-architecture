FROM golang:1.24.2-alpine as builder

ENV LANG C.UTF-8

WORKDIR /go/src/clean-architecture

RUN apk add --update --no-cache \
      git \
      make

COPY . .
RUN make build

FROM alpine

ENV LANG C.UTF-8

RUN apk add --update --no-cache \
      ca-certificates \
      tzdata

COPY --from=builder /go/src/clean-architecture/app /app

EXPOSE 8080
CMD ["./app"]