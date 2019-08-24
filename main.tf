terraform {
  backend "pg" {}
}
module "go_api_deployment" { 
  source          = "git::https://github.com/blockchain-abstraction-middleware/deployment.git" 
  namespace       = "go-apis" 
  deployment_name = "project-template" 
  docker_image    = "bamdockerhub/project-template" 
}