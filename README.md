![](https://img.shields.io/badge/version-0.0.1-brightgreen)

# go-ngx-config
A nginx config parser


## Basic Usage
Usage:
```sh
# -f          file path location nginx configm, e,g: ./examples/basic/nginx.conf
# -o          output json file path location, e.g: ./examples/basic/output


go-ngx-config parse -f <NGINX_CONF_FILE> -o <OUTPUT_JSON_FILE_DUMP>
```

<details>
<summary>Result</summary>

```json
{
  "Directives": [
    {
      "Block": null,
      "Name": "daemon",
      "Parameters": [
        "off"
      ]
    },
    {
      "Block": null,
      "Name": "worker_processes",
      "Parameters": [
        "2"
      ]
    },
    {
      "Block": null,
      "Name": "user",
      "Parameters": [
        "www-data"
      ]
    },
    {
      "Block": {
        "Directives": [
          {
            "Block": null,
            "Name": "use",
            "Parameters": [
              "epoll"
            ]
          },
          {
            "Block": null,
            "Name": "worker_connections",
            "Parameters": [
              "128"
            ]
          }
        ]
      },
      "Name": "events",
      "Parameters": null
    },
    {
      "Block": null,
      "Name": "error_log",
      "Parameters": [
        "logs/error.log",
        "info"
      ]
    },
    {
      "Servers": [
        {
          "Block": {
            "Directives": [
              {
                "Block": null,
                "Name": "server_name",
                "Parameters": [
                  "localhost"
                ]
              },
              {
                "Block": null,
                "Name": "listen",
                "Parameters": [
                  "127.0.0.1:80"
                ]
              },
              {
                "Block": null,
                "Name": "error_page",
                "Parameters": [
                  "500",
                  "502",
                  "503",
                  "504",
                  "/50x.html"
                ]
              },
              {
                "Name": "location",
                "Modifier": "=",
                "Match": "/favicon.ico"
              },
              {
                "Name": "location",
                "Modifier": "",
                "Match": "/"
              }
            ]
          }
        }
      ],
      "Name": "http",
      "Directives": [
        {
          "Block": null,
          "Name": "server_tokens",
          "Parameters": [
            "off"
          ]
        },
        {
          "Block": null,
          "Name": "include",
          "Parameters": [
            "mime.types"
          ]
        },
        {
          "Block": null,
          "Name": "charset",
          "Parameters": [
            "utf-8"
          ]
        },
        {
          "Block": null,
          "Name": "access_log",
          "Parameters": [
            "logs/access.log",
            "combined"
          ]
        }
      ]
    }
  ],
  "Filepath": "./examples/basic/nginx.conf"
}
```
</details>

<br/>

## TODO(s):
- [ ] Include directive and reads the glob (?)
- [ ] HTTP Server for see the config on UI browser
- [ ] And lot more...