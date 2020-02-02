module github.com/pulumi/pulumi-ucloud

go 1.12

replace (
	github.com/Nvveen/Gotty => github.com/ijc25/Gotty v0.0.0-20170406111628-a8b993ba6abd
	github.com/golang/glog => github.com/pulumi/glog v0.0.0-20180820174630-7eaa6ffb71e4
	github.com/hashicorp/terraform v0.12.7 => github.com/hashicorp/terraform v0.12.4
	github.com/uber/jaeger-lib v2.1.1+incompatible => github.com/uber/jaeger-lib v1.5.0
)

require (
	github.com/ahmetb/go-linq v3.0.0+incompatible
	github.com/hashicorp/terraform-plugin-sdk v1.4.0
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.9.1
	github.com/pulumi/pulumi-terraform v1.1.0 // indirect
	github.com/pulumi/pulumi-terraform-bridge v1.6.5
	github.com/stretchr/testify v1.4.1-0.20191106224347-f1bd0923b832
	github.com/terraform-providers/terraform-provider-ucloud v1.15.1
	github.com/uber/jaeger-lib v2.1.1+incompatible // indirect
)
