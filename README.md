<div align="center">

<a href="https://smms.app/image/UZpg7XuikoI2VB9" target="_blank"><img src="https://s2.loli.net/2025/02/21/UZpg7XuikoI2VB9.jpg" width="256" height="256"></a>

# 時雨 (Shigure-Bot)

⌚Yet another chatbot SDK that compatible for multiple chatbot backend implementations.🌧

<p align="center">

<img src="https://img.shields.io/github/license/arttnba3/Shigure-Bot?style=for-the-badge" alt="license">
<img src="https://img.shields.io/badge/OneBot-11-black?style=for-the-badge&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAHAAAABwCAMAAADxPgR5AAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAAxQTFRF////29vbr6+vAAAAk1hCcwAAAAR0Uk5T////AEAqqfQAAAKcSURBVHja7NrbctswDATQXfD//zlpO7FlmwAWIOnOtNaTM5JwDMa8E+PNFz7g3waJ24fviyDPgfhz8fHP39cBcBL9KoJbQUxjA2iYqHL3FAnvzhL4GtVNUcoSZe6eSHizBcK5LL7dBr2AUZlev1ARRHCljzRALIEog6H3U6bCIyqIZdAT0eBuJYaGiJaHSjmkYIZd+qSGWAQnIaz2OArVnX6vrItQvbhZJtVGB5qX9wKqCMkb9W7aexfCO/rwQRBzsDIsYx4AOz0nhAtWu7bqkEQBO0Pr+Ftjt5fFCUEbm0Sbgdu8WSgJ5NgH2iu46R/o1UcBXJsFusWF/QUaz3RwJMEgngfaGGdSxJkE/Yg4lOBryBiMwvAhZrVMUUvwqU7F05b5WLaUIN4M4hRocQQRnEedgsn7TZB3UCpRrIJwQfqvGwsg18EnI2uSVNC8t+0QmMXogvbPg/xk+Mnw/6kW/rraUlvqgmFreAA09xW5t0AFlHrQZ3CsgvZm0FbHNKyBmheBKIF2cCA8A600aHPmFtRB1XvMsJAiza7LpPog0UJwccKdzw8rdf8MyN2ePYF896LC5hTzdZqxb6VNXInaupARLDNBWgI8spq4T0Qb5H4vWfPmHo8OyB1ito+AysNNz0oglj1U955sjUN9d41LnrX2D/u7eRwxyOaOpfyevCWbTgDEoilsOnu7zsKhjRCsnD/QzhdkYLBLXjiK4f3UWmcx2M7PO21CKVTH84638NTplt6JIQH0ZwCNuiWAfvuLhdrcOYPVO9eW3A67l7hZtgaY9GZo9AFc6cryjoeFBIWeU+npnk/nLE0OxCHL1eQsc1IciehjpJv5mqCsjeopaH6r15/MrxNnVhu7tmcslay2gO2Z1QfcfX0JMACG41/u0RrI9QAAAABJRU5ErkJggg==" alt="onebot">

</p>

</div>

## Usage

Firstly you may need to introduce this library to local as follow:

```shell
$ go get github.com/arttnba3/Shigure-Bot
```

Then just import the module into your project, and create a bot simply as follow:

```go
import "github.com/arttnba3/Shigure-Bot/bot"

//...

bot, err = shigure.NewShigureBot(botType, configJson, Logger, handlers)
```

You can refer to [example](example) directory for a detailed usage.

> Note that you project should be run together with a supported backend individually, as this is only a SDK for communicating with corresponding backend implementations.

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

To handle events reported by OneBot-compatible backend, we need to prepare a function-table-like structure as the handler to the bot constructor. This should be a map like:

```go
func HandleExampleEvent(params ...any) {
    bot, ok1 := params[0].(*onebot_v11_impl.V11Bot)
	event, ok2 := params[1].(onebot_v11_api_event.ExampleEvent)
	// replace ExampleEvent to your expected type, e.g., PrivateMessage
	// refer to the OneBot v11 specification for details
	//...
}

var handlers map[string]func(params ...any) = map[string]func(params ...any){
    "example_key": HandleExampleEvent,
    //...
}
```

For the key of the map, it consists of two parts. The first part is the type, which can be `message`, `meta`, `notice`, `request`, and the second part is the subtype. An example key of handling received private message is `message_private`. You can refer to OneBot's docs and the [bot/onebot/v11/handler.go](bot/onebot/v11/handler.go) for details.

You can refer to [example/onebot-v11.go](example/onebot-v11.go) for an example usage.

## Author

- Code: arttnba3 <arttnba3@outlook.com>
- Illustration: Sumi Kiriko <kirikosumi@gmail.com>

## License

GPL V2
