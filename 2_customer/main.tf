# load the provider
terraform {
  required_providers {
    myuserprovider = {
      source  = "example.com/me/myuserprovider"
      # version = "~> 1.0"
    }
  }
}

# configure the provider
provider "myuserprovider" {
  endpoint = "http://localhost:6251/"
}

# # configure the resource
# resource "myuserprovider_user" "john_doe" {
#   id   = "1"
#   name = "John Doe"
# }
resource "myuserprovider_user" "john_doe2" {
  id   = "2"
  name = "John Doe2"
}
resource "myuserprovider_user" "john_doe3" {
  id   = "3"
  name = "John Doe3"
}