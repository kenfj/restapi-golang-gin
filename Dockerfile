FROM golang:1.21.3-alpine AS builder

RUN apk update \
    && apk add git \
    && apk cache clean

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN go build -o /main

FROM golang:1.21.3-alpine
COPY --from=builder /main /main

EXPOSE 8080
ENTRYPOINT ["/main"]
