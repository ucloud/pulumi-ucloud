# Pulumi UCloud Provider

This repository is forked from [pulumi-tf-provider-boilerplate](https://github.com/pulumi/pulumi-tf-provider-boilerplate) 
and modified to create Pulumi UCloud Provider. UCloud is a leading public cloud mainly based in China.

UCloud also provide an official [Terraform Provider](https://www.terraform.io/docs/providers/ucloud/index.html)

## Add dependencies

In order to properly build the sdks, the following tools are expected:
- tf2pulumi (See the project's README for installation instructions: https://github.com/pulumi/tf2pulumi)
- pandoc (`brew install pandoc`)

In the root of the repository, run:

- `go get github.com/pulumi/scripts/gomod-doccopy` (Note: do not set `GO111MODULE=on` here)
- `GO111MODULE=on go get github.com/pulumi/pulumi-terraform@master`
- `GO111MODULE=on go get github.com/terraform-providers/terraform-provider-ucloud`
- `GO111MODULE=on go mod vendor`
- `make ensure`

## Build the provider:

- Edit `resources.go` to map each resource, and specify provider information (It's unnecessary if you just want to build sdk.)
- Enumerate any examples in `examples/examples_test.go` (It's unnecessary if you just want to build sdk.)
- `make`

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/ucloud

or `yarn`:

    $ yarn add @pulumi/ucloud

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_ucloud

### Go

For now, this repository is under my personal account, which module name in go.mod is different with repository url.
If you want use go sdk module from other repository, please replace module from "github.com/pulumi/pulumi-ucloud" to "github.com/{YOUR_ACCOUNT_HERE}/pulumi-ucloud" in go.mod first

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-ucloud/sdk/go/...

## Configuration

The following configuration points are **required** for the `ucloud` provider:

- `ucloud:region` (environment: `UCLOUD_REGION`) - UCloud Region
- `ucloud:public_key` (environment: `UCLOUD_PUBLIC_KEY`) - UCloud Public Key
- `ucloud:private_key` (environment: `UCLOUD_PRIVATE_KEY`) - UCloud Private Key
- `ucloud:project_id` (environment: `UCLOUD_PROJECT_ID`) - UCloud Project ID

The following configuration points are **optional** for `ucloud` provider:
- `ucloud:profile` (environment: `UCLOUD_PROFILE`) - UCloud Profile Name
- `ucloud:base_url` - UCloud base URL.(Default: https://api.ucloud.cn)
- `ucloud:max_retries` - Max retry attempts number. Default max retry attempts number is 0
- `ucloud:insecure` - Switch to disable/enable https. (Default: false, means enable https)
- `ucloud:shared_credentials_file` (environment: `UCLOUD_SHARED_CREDENTIAL_FILE`) - Path To The Shared Credentials File
