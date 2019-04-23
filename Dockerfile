# Build UI
FROM node:10.15-alpine as ui-builder

WORKDIR /app

COPY ./web/package*.json /app/
RUN npm install

COPY ./web /app/

ARG configuration=production

RUN npm run build -- --output-path=./dist/out --configuration $configuration

# Build backend
FROM golang:1.11-alpine as go-builder

RUN apk add --no-cache git

WORKDIR /app

COPY backend /app
COPY backend /app
RUN go mod download

COPY backend /app/

RUN CGO_ENABLED=0 go build -o guff

# Run
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app/

COPY --from=ui-builder /app/dist/out ./web

COPY --from=go-builder /app/guff .

CMD ["./guff", "--logtostderr", "--web_root=./web"]
