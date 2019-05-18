# guff [![Build Status](https://travis-ci.org/toothrot/guff.svg?branch=master)](https://travis-ci.org/toothrot/guff)

## Development

### Testing

#### Setup

```bash
sudo -u postgres createuser -s "$USER"
createdb guff_test
```

#### Running tests

```bash
make test
```
