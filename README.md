# 時雨 (Shigure)

Yet another chatbot SDK that satisfies multiple chatbot backend specifications.

## Usage

> TBD

## Supported bot specifications

### [OneBot V11](https://github.com/botuniverse/onebot-11/)

Currently, we support part of OneBot V11 API, which can be known by examining the source code.

For the connection, we support following:

- [HTTP](https://github.com/botuniverse/onebot-11/blob/master/communication/http.md)
- [HTTP-Post](https://github.com/botuniverse/onebot-11/blob/master/communication/http-post.md)

To configure a Shigure-Bot for an OneBot backend, we need to provide the configuration in following format(if one of which was not configured, it won't be invoked):

```json
{
  "http_post": {
    "host": "example.com",
    "port": 11451
  },
  "http_server": {
    "port": 19198
  }
}
```

You can refer to [example/onebot-v11.go](example/onebot-v11.go) for an example usage.

## Author

arttnba3 <arttnba3@outlook.com>

## License

GPL V2
