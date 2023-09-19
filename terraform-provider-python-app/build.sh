#!/usr/bin/env bash

PROVIDER_NAME='terraform-provider-python-app_v0.0.1'

go build -o "${PROVIDER_NAME}"

mkdir -p "${HOME}/.terraform.d/plugins/my.local/my/python-app/0.0.1/darwin_arm64/"
mv "${PROVIDER_NAME}" "${HOME}/.terraform.d/plugins/my.local/my/python-app/0.0.1/darwin_arm64/${PROVIDER_NAME}"
