# Web deps
FROM node:10.15 as web-deps

WORKDIR /app

RUN wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add -
RUN echo "deb http://dl.google.com/linux/chrome/deb/ stable main" > /etc/apt/sources.list.d/google.list
RUN apt-get update && apt-get install -y google-chrome-stable xvfb procps

COPY ./web/package*.json /app/
RUN npm install

# Web src
FROM web-deps as web-src

WORKDIR /app

COPY ./web /app/

# Web build
FROM web-src as web

ARG configuration=production
RUN npm run build -- --output-path=./dist/out --configuration $configuration

# Web test
FROM web-src as web-test

WORKDIR /app

CMD npm run -- ng test --watch=false --browsers=ChromeHeadless

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
