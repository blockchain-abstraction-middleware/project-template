# Project-Template

## Overview
Go scaffolding project template, this includes:
- Sample rest service code
- Sample circleci configs
- Sample Dockerfile
- Sample terraform deployment

### Running the service
```shell
make serve
```

### Running Unit Tests
```shell
make test
```

### Linting
```shell
make lint
```

### Docker build
```shell
make docker-build
```

### User Testing
```shell
curl 0.0.0.0:8080/api/v1/health/status
```

### How to use
- `mkdir new-prog && cd new-prog`
- `git clone git@github.com:blockchain-abstraction-middleware/project-template.git .`
- `sudo rm -r .git`
- :thumbs_up: