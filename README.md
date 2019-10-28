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

### Swagger UI
Swagger ui will be exposed at:
```
0.0.0.0:8080/api/v1/swagger/docs/swaggerui/
```

### How to use
- `mkdir new-prog && cd new-prog`
- `git clone git@github.com:blockchain-abstraction-middleware/project-template.git .`
- `sudo rm -r .git`
- :thumbs_up:

### Things to update 

`main.tf`
- Line 12: update module "name"
- Line 15: update deployment_name
- Line 16: update docker_image

`.circleci/config.yml`
- Line 20: update TF_VAR_docker_image to === the name in `main.tf`
- Line 37: update  `terraform taint module.'your_module_name'.kubernetes_deployment.main` to === name in `main.tf`