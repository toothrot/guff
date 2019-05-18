# guff [![Build Status](https://travis-ci.org/toothrot/guff.svg?branch=master)](https://travis-ci.org/toothrot/guff)

## Development

Starting the app locally for development:
```bash
make secrets
docker-compose up

# Alternatively:
make docker-dev-run
```

### Testing

#### Docker

Run all tests with Make and Docker Compose:
```bash
make docker-test
```

Run tests individually with docker-compose:
```bash
docker-compose run --rm backend-test
docker-compose run --rm web-test
```

#### Local

##### Setup

```bash
sudo -u postgres createuser -s "$USER"
createdb guff_test
```

##### Running tests

Run all tests:
```bash
make test
```

Run tests individually:
```bash
# Backend
cd backend; go test ./...

# Web
cd web; npm test
```
