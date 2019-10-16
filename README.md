# Project-Template

## Overview
Go scaffolding project template

- Linting
```shell
make lint
```

- Running
```shell
make serve
```

- Testing
POST
```
curl -X POST -H 'application/x-www-form-urlencoded' -d "hello=0xb4FC6774a95A86134768d81A8eE19Cfbf5A171F6" localhost:8080
```
GET
```shell
curl 0.0.0.0:8080/api/v1/health/status
```