# stubbygo

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Docker](https://img.shields.io/docker/v/ureuzy/stubbygo/v0.1.0?color=blue&logo=docker)](https://hub.docker.com/repository/docker/ureuzy/stubbygo)
[![report](https://goreportcard.com/badge/github.com/ureuzy/stubbygo)](https://goreportcard.com/report/github.com/ureuzy/stubbygo)

Can build a stub server by defining endpoints in yaml beforehand

- Dockerized

# Quick Start

config.yaml

```yaml
endpoints:
  - path: "/"
    methods:
      - type: "GET"
        response:
          status_code: 200
          headers:
            Content-Type: "text/plain"
          body: 'hello,world'
```

Start stub server

```
$ go run cmd/main.go -c config.yaml
2022/07/07 18:10:55 server listen :8080 ...
```

Request / Response

```
$ curl -i -X GET 'http://localhost:8080'
HTTP/1.1 200 OK
Content-Type: text/plain
Date: Thu, 07 Jul 2022 09:13:55 GMT
Content-Length: 11

hello,world
```

## Embed Query Parameters

The values of query parameters can be embedded in the response. The key is to include `queries`. 
This allows you to use the template syntax in the body to embed them.

```yaml
endpoints:
  - path: "/"
    methods:
      - type: "GET"
        queries: ["foo", "bar"]
        response:
          status_code: 200
          headers:
            Content-Type: "text/plain"
          body: "{{.foo}},{{.bar}}"
```

```
$ curl -i -X GET 'http://localhost:8080?foo=hello&bar=world'
HTTP/1.1 200 OK
Access-Control-Allow-Origin: *
Content-Type: application/json
Date: Thu, 07 Jul 2022 09:31:13 GMT
Content-Length: 91

hello,world
```

# Docker

```
docker run -d -p 8080:8080 -v $(pwd)/endpoints.yaml:/config/endpoints.yaml  ureuzy/stubbygo
```
