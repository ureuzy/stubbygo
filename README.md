# stubbygo

Can build a stub server by defining endpoints in yaml beforehand

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
      - type: "POST"
        response:
          status_code: 201
          headers:
            Content-Type: "text/plain"
          body: "ok"
```

Start stub server

```
$ go run cmd/main.go
2022/07/07 18:10:55 server listen :8080 ...
```

Request

```
$ curl -i -X GET 'http://localhost:8080'
HTTP/1.1 200 OK
Content-Type: text/plain
Date: Thu, 07 Jul 2022 09:13:55 GMT
Content-Length: 11

hello,world
```

```
$ curl -i -X POST -H "Content-Type:application/json" 'http://localhost:8080'
HTTP/1.1 201 Created
Content-Type: text/plain
Date: Thu, 07 Jul 2022 09:39:46 GMT
Content-Length: 2

ok
```

## Embed Query Parameters

The values of query parameters can be embedded in the response. The key is to include `query_keys`. 
This allows you to use the template syntax in the body to embed them.

```yaml
endpoints:
  - path: "/"
    methods:
      - type: "GET"
        query_keys: ["foo", "bar"]
        response:
          status_code: 200
          headers:
            Content-Type: "application/json"
          body: '
            {"data":
              [
                {"key": "foo", "value": "{{.foo}}"},
                {"key": "bar", "value": "{{.bar}}"}
              ]
            }
          '
```

```
$ curl -i -X GET 'http://localhost:8080?foo=foo_value&bar=bar_value'
HTTP/1.1 200 OK
Access-Control-Allow-Origin: *
Content-Type: application/json
Date: Thu, 07 Jul 2022 09:31:13 GMT
Content-Length: 91

{"data": [ {"key": "foo", "value": "foo_value"}, {"key": "bar", "value": "bar_value"} ] }
```
