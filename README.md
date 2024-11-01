# Instruction

An common structure for golang project

# structure

- etc: setting file
- docs: document
- global: global variables
- internal (internal module):
 <!-- TODO: -->
- dao: data access object
- middleware
- model: database model control
- routers: api routes
- service: process business logic
- pkg: package
- storage: temp file
- scripts: build, install, analysis scripts
- third_party: third_party tools

# Go generate

```sh
go generate github.com/hoops_dynasty/internal/dao/dbversion/mysql
```

# Build

```sh
go build github.com/hoops_dynasty/cmd/${PROJECT_NAME}
```

### args

```sh
Usage of ./db_lock:
  -config string
        assgin the path of config file (default "etc/")
  -mode string
        running level (info, debug)
```

# Integration Testing

```sh
go build github.com/hoops_dynasty/testing/logger
```
