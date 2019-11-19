terraform {
  backend "pg" {
    workspaces {
      name = "go-apis"
    }
  }
}

provider "kubernetes" {
}

module "go_api_deployment_project_template" {
  source          = "github.com/blockchain-abstraction-middleware/deployment/modules/deployment"
  namespace       = "go-apis"
  deployment_name = "project-template"
  docker_image    = "bamdockerhub/project-template"
  config_file     = "review.yml"
  config_path     = "/config/"
}