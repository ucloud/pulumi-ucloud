# UCloud Pulumi Provider

[WIP] Work in progress, DON'T USE IT IN PRODUCTION!

This repository contains boilerplate code for building a new Pulumi provider which wraps an existing
Terraform provider, if the existing provider uses _Go Modules_.

Modify this README to describe:

- The type of resources the provider manages
- Add a build status image from Travis at the top of the README
- Update package names in the information below
- Add any important documentation of concepts (e.g. the "serverless" components in the AWS provider).

### Add dependencies

In order to properly build the sdks, the following tools are expected:
- `pulumictl` (See the project's README for installation instructions: https://github.com/pulumi/pulumictl)

In the root of the repository, run:

- `(cd provider && go get github.com/terraform-providers/terraform-provider-foo)`  (where `foo` is the name of the provider - note the parenthesis to run this in a subshell)
- `(cd provider && go mod download)`

### Build the provider:

- Edit `provider/resources.go` to map each resource, and specify provider information
- `make build_sdks`

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

## Reference

For detailed reference documentation, please visit [the API docs][1].


[1]: https://www.pulumi.com/docs/reference/pkg/x/
