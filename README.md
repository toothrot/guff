# guff

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
