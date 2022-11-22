![](https://img.shields.io/badge/version-0.4.0-brightgreen)

# go-ngx-config
A nginx config parser built on top of `crossplane`


## Basic Usage
### Binary
Usage:
```sh
# Parse
# -f          file path location nginx config, e.g: ./examples/basic/nginx.conf
# -o          output json file path location, e.g: ./examples/basic/output
go-ngx-config parse -f <NGINX_CONF_FILE> -o <OUTPUT_JSON_FILE_DUMP>

# Location Matcher
# -f          file path location nginx config, e.g: ./examples/basic/nginx.conf
# -u          url target, e.g: /my-location
go-ngx-config lt -f <NGINX_CONF_FILE> -u <URL_TARGET>
```

<details>
<summary>Result</summary>

```json
{
  "status": "ok",
  "errors": [],
  "config": [
    {
      "file": "./examples/basic/nginx.conf",
      "status": "ok",
      "errors": [],
      "parsed": [
        {
          "directive": "daemon",
          "line": 1,
          "args": [
            "off"
          ]
        },
        {
          "directive": "worker_processes",
          "line": 2,
          "args": [
            "2"
          ]
        },
        {
          "directive": "user",
          "line": 3,
          "args": [
            "www-data"
          ]
        },
        {
          "directive": "events",
          "line": 5,
          "args": [],
          "block": [
            {
              "directive": "use",
              "line": 6,
              "args": [
                "epoll"
              ]
            },
            {
              "directive": "worker_connections",
              "line": 7,
              "args": [
                "128"
              ]
            }
          ]
        },
        {
          "directive": "error_log",
          "line": 10,
          "args": [
            "logs/error.log",
            "info"
          ]
        },
        {
          "directive": "http",
          "line": 12,
          "args": [],
          "block": [
            {
              "directive": "server_tokens",
              "line": 13,
              "args": [
                "off"
              ]
            },
            {
              "directive": "include",
              "line": 14,
              "args": [
                "mime.types"
              ]
            },
            {
              "directive": "charset",
              "line": 15,
              "args": [
                "utf-8"
              ]
            },
            {
              "directive": "access_log",
              "line": 17,
              "args": [
                "logs/access.log",
                "combined"
              ]
            },
            {
              "directive": "server",
              "line": 19,
              "args": [],
              "block": [
                {
                  "directive": "server_name",
                  "line": 20,
                  "args": [
                    "localhost"
                  ]
                },
                {
                  "directive": "listen",
                  "line": 21,
                  "args": [
                    "127.0.0.1:80"
                  ]
                },
                {
                  "directive": "error_page",
                  "line": 23,
                  "args": [
                    "500",
                    "502",
                    "503",
                    "504",
                    "/50x.html"
                  ]
                },
                {
                  "directive": "include",
                  "line": 25,
                  "args": [
                    "conf-includes/proxy.conf"
                  ]
                },
                {
                  "directive": "include",
                  "line": 26,
                  "args": [
                    "handlers/*.conf"
                  ]
                },
                {
                  "directive": "location",
                  "line": 28,
                  "args": [
                    "/"
                  ],
                  "block": [
                    {
                      "directive": "add_header",
                      "line": 29,
                      "args": [
                        "x-foo",
                        "x-bar"
                      ]
                    },
                    {
                      "directive": "proxy_pass",
                      "line": 30,
                      "args": [
                        "http://my-upstream"
                      ]
                    }
                  ]
                }
              ]
            }
          ]
        }
      ]
    }
  ]
}
```
</details>

<br/>


### Web Assembly

Exported Global Function
```js
// goNgxParseConfig: generate AST in JSON format
// goNgxTestLocation: to test location matcher

// Some Basic Example
await goNgxParseConfig(`
server {
  server_name my-server.domain.com;

  location = /my-location {
    proxy_pass http://lite-dev;
  }
}
`);

await goNgxTestLocation(`
server {
  server_name my-server.domain.com;

  location = /my-location {
    proxy_pass http://lite-dev;
  }
}
`, '/my-location');
```

<br/>


## TODO(s):
- [x] .wasm binary 
- [x] Location Tester
- [x] Include directive and reads the glob
- [ ] HTTP Server for see the config on UI browser
- [ ] And lot more...