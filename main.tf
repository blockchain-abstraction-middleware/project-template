terraform {
  backend "pg" {
    workspaces {
      name = "go-apis"
    }
  }
}

provider "kubernetes" {
}

module "go_api_deployment" { 
  source          = "git::https://github.com/blockchain-abstraction-middleware/modules/deployment.git"
  namespace       = "go-apis"
  deployment_name = "project-template"
  docker_image    = "bamdockerhub/project-template"
}