module github.com/pulumi/pulumi-ucloud

go 1.13

replace (
	git.apache.org/thrift.git v0.12.0 => github.com/apache/thrift v0.12.0
	github.com/Nvveen/Gotty => github.com/ijc25/Gotty v0.0.0-20170406111628-a8b993ba6abd
	github.com/golang/glog => github.com/pulumi/glog v0.0.0-20180820174630-7eaa6ffb71e4
	github.com/hashicorp/terraform v0.12.7 => github.com/hashicorp/terraform v0.12.4
	github.com/uber/jaeger-lib v2.1.1+incompatible => github.com/uber/jaeger-lib v1.5.0
)

require (
	cloud.google.com/go/logging v1.0.0 // indirect
	github.com/Azure/go-autorest/autorest v0.9.1 // indirect
	github.com/Azure/go-autorest/autorest/azure/auth v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.2.0 // indirect
	github.com/ahmetb/go-linq v3.0.0+incompatible
	github.com/hashicorp/terraform v0.12.9
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.1.0
	github.com/pulumi/pulumi-terraform v0.18.4-0.20190914112649-fc132172e90d
	github.com/stretchr/testify v1.4.0
	github.com/terraform-providers/terraform-provider-ucloud v1.12.1
	github.com/uber/jaeger-lib v2.1.1+incompatible // indirect
	gocloud.dev/secrets/hashivault v0.17.0 // indirect
)
