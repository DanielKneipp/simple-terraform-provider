# Sample Terraform Provider

This repo should provide a simple custom Terraform provider with the bare minimum
to exemplify the declarative approach Terraform allows us to use when managing
systems

## The python app

This a simple stateful app that holds a list of integers. Its api only allows
individual changes on the integers held by the app. Also the order of the
numbers will be defined by the sequence used to add them.

So, if the user wants a specific list of numbers on the app, APIs need to
be called in a determined way for the desired outcome be achieved.

The available API calls are present on the Postman collection inside
`/python-app` directory

## The Terraform provider

All the relevant code is in the `/terraform-provider-python-app` directory. Here is a small description of each subdirectory:

- `api/` has the http client to interact with the service
- `provider/` contains the code to define the terraform provider and
its resource. Here is where the terraform resource `numbers` is defined
- `terraform/` has a the terraform code to test the provider

The `build.sh` builds the provider and moves it to a specific local path so that
terraform can find it. Use it before trying to play with the terraform code.

> You can check a more extensive tutorial from Hashicorp [here](https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework).
> And there are also good tutorials out there, like [this one](https://www.infracloud.io/blogs/developing-terraform-custom-provider/)
