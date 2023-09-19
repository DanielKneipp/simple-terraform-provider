terraform {
  required_providers {
    python-app = {
      source  = "my.local/my/python-app"
      version = "0.0.1"
    }
  }
}

provider "python-app" {
  endpoint = "http://127.0.0.1:5000"
}

resource "numbers" "my_numbers" {
  provider = python-app

  numbers = [1, 2, 3, 4, 5, 10]
}
