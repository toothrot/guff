# Web deps
FROM node:10.15-alpine as web-deps

WORKDIR /app

COPY ./web/package*.json /app/
RUN npm install

# Web build
FROM web-deps as web

WORKDIR /app

COPY ./web /app/

ARG configuration=production
RUN npm run build -- --output-path=./dist/out --configuration $configuration

# Backend deps
FROM golang:1.12-alpine as backend-deps

RUN apk add --no-cache git

WORKDIR /app

COPY ./backend/go.* /app/
RUN go mod download

# Backend src
FROM backend-deps as backend-src

WORKDIR /app

COPY ./backend /app/

# Backend build
FROM backend-src as backend

RUN CGO_ENABLED=0 go build -o guff

# Backend test
FROM backend-src as backend-test

WORKDIR /app

ENV CGO_ENABLED=0
CMD ["go", "test", "./..."]

# Run
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app/

COPY --from=web /app/dist/out ./web

COPY --from=backend /app/guff .

CMD ["./guff", "--logtostderr", "--web_root=./web"]
