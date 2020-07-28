# goservice

Base web service project in Go, currently aiming for basic http api implementation.

End goal is to have a single project that could be run using either gRPC or REST api.

## Requirements

- Go v1.14
- make

## Environment Variables

- `HTTP_PORT`: 80
- `HTTP_HOST`: localhost
- `LOG_LEVEL`: debug

## Usage

Only need to clone the repo and run

```bash
make
```

## License

[MIT](https://github.com/sergionunezgo/gorest/LICENSE)
