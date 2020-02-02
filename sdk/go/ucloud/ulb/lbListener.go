// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ulb

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Load Balancer Listener resource.
//
// > **Note** This `listenType` only support when `protocol` is `tcp` in the extranet mode and the default value is `requestProxy`. In addition, in the extranet mode, the `listenType` is `requestProxy` if `protocol`is `http` or `https`, the `listenType` is `packetsTransmit` if `protocol`is `udp`. In the intranet mode, the `listenType` is `packetsTransmit`.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-ucloud/blob/master/website/docs/r/lb_listener.html.markdown.
type LbListener struct {
	pulumi.CustomResourceState

	// Health check domain checking.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Health check method. Possible values are `port` as port checking and `path` as http checking.
	HealthCheckType pulumi.StringOutput `pulumi:"healthCheckType"`
	// Amount of time in seconds to wait for the response for in between two sessions if `listenType` is `requestProxy`, range: 0-86400. (Default: `60`). Amount of time in seconds to wait for one session if `listenType` is `packetsTransmit`, range: 60-900. The session will be closed as soon as no response if it is `0`.
	IdleTimeout pulumi.IntOutput `pulumi:"idleTimeout"`
	// The type of listener. Possible values are `requestProxy` and `packetsTransmit`. When `packetsTransmit` was specified, you need to config the instances by yourself if the instances attach to the load balancer. You may refer to [configuration instruction](https://docs.ucloud.cn/network/ulb/guide/fu-wu-jie-dian-xiang-guan-cao-zuo/editrealserver).
	ListenType pulumi.StringOutput `pulumi:"listenType"`
	// The ID of load balancer instance.
	LoadBalancerId pulumi.StringOutput `pulumi:"loadBalancerId"`
	// The load balancer method in which the listener is. Possible values are: `roundrobin`, `source`, `consistentHash`, `sourcePort` , `consistentHashPort`, `weightRoundrobin` and `leastconn`. (Default: `roundrobin`).
	// - The `consistentHash`, `sourcePort` , `consistentHashPort`, `roundrobin`, `source` and `weightRoundrobin` are valid if `listenType` is `packetsTransmit`.
	// - The `roundrobin`, `source` and `weightRoundrobin` and `leastconn` are valid if `listenType` is `requestProxy`.
	Method pulumi.StringPtrOutput `pulumi:"method"`
	Name   pulumi.StringOutput    `pulumi:"name"`
	// Health check path checking.
	Path pulumi.StringOutput `pulumi:"path"`
	// Indicate whether the persistence session is enabled, it is invalid if `persistenceType` is `none`, an auto-generated string will be exported if `persistenceType` is `serverInsert`, a custom string will be exported if `persistenceType` is `userDefined`.
	Persistence pulumi.StringOutput `pulumi:"persistence"`
	// The type of session persistence of listener. Possible values are: `none` as disabled, `serverInsert` as auto-generated key and `userDefined` as customized key. (Default: `none`).
	PersistenceType pulumi.StringPtrOutput `pulumi:"persistenceType"`
	// Port opened on the listeners to receive requests, range: 1-65535. The default value: `80` as `protocol` is `http`, `443` as `protocol` is `https`, `1024` as `protocol` is `tcp` or `udp`.
	Port pulumi.IntOutput `pulumi:"port"`
	// Listener protocol. Possible values: `http`, `https`, `tcp` if `listenType` is `requestProxy`, `tcp` and `udp` if `listenType` is `packetsTransmit`.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Listener status. Possible values are: `allNormal` for all resource functioning well, `partNormal` for partial resource functioning well and `allException` for all resource functioning exceptional.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewLbListener registers a new resource with the given unique name, arguments, and options.
func NewLbListener(ctx *pulumi.Context,
	name string, args *LbListenerArgs, opts ...pulumi.ResourceOption) (*LbListener, error) {
	if args == nil || args.LoadBalancerId == nil {
		return nil, errors.New("missing required argument 'LoadBalancerId'")
	}
	if args == nil || args.Protocol == nil {
		return nil, errors.New("missing required argument 'Protocol'")
	}
	if args == nil {
		args = &LbListenerArgs{}
	}
	var resource LbListener
	err := ctx.RegisterResource("ucloud:ulb/lbListener:LbListener", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLbListener gets an existing LbListener resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLbListener(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LbListenerState, opts ...pulumi.ResourceOption) (*LbListener, error) {
	var resource LbListener
	err := ctx.ReadResource("ucloud:ulb/lbListener:LbListener", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LbListener resources.
type lbListenerState struct {
	// Health check domain checking.
	Domain *string `pulumi:"domain"`
	// Health check method. Possible values are `port` as port checking and `path` as http checking.
	HealthCheckType *string `pulumi:"healthCheckType"`
	// Amount of time in seconds to wait for the response for in between two sessions if `listenType` is `requestProxy`, range: 0-86400. (Default: `60`). Amount of time in seconds to wait for one session if `listenType` is `packetsTransmit`, range: 60-900. The session will be closed as soon as no response if it is `0`.
	IdleTimeout *int `pulumi:"idleTimeout"`
	// The type of listener. Possible values are `requestProxy` and `packetsTransmit`. When `packetsTransmit` was specified, you need to config the instances by yourself if the instances attach to the load balancer. You may refer to [configuration instruction](https://docs.ucloud.cn/network/ulb/guide/fu-wu-jie-dian-xiang-guan-cao-zuo/editrealserver).
	ListenType *string `pulumi:"listenType"`
	// The ID of load balancer instance.
	LoadBalancerId *string `pulumi:"loadBalancerId"`
	// The load balancer method in which the listener is. Possible values are: `roundrobin`, `source`, `consistentHash`, `sourcePort` , `consistentHashPort`, `weightRoundrobin` and `leastconn`. (Default: `roundrobin`).
	// - The `consistentHash`, `sourcePort` , `consistentHashPort`, `roundrobin`, `source` and `weightRoundrobin` are valid if `listenType` is `packetsTransmit`.
	// - The `roundrobin`, `source` and `weightRoundrobin` and `leastconn` are valid if `listenType` is `requestProxy`.
	Method *string `pulumi:"method"`
	Name   *string `pulumi:"name"`
	// Health check path checking.
	Path *string `pulumi:"path"`
	// Indicate whether the persistence session is enabled, it is invalid if `persistenceType` is `none`, an auto-generated string will be exported if `persistenceType` is `serverInsert`, a custom string will be exported if `persistenceType` is `userDefined`.
	Persistence *string `pulumi:"persistence"`
	// The type of session persistence of listener. Possible values are: `none` as disabled, `serverInsert` as auto-generated key and `userDefined` as customized key. (Default: `none`).
	PersistenceType *string `pulumi:"persistenceType"`
	// Port opened on the listeners to receive requests, range: 1-65535. The default value: `80` as `protocol` is `http`, `443` as `protocol` is `https`, `1024` as `protocol` is `tcp` or `udp`.
	Port *int `pulumi:"port"`
	// Listener protocol. Possible values: `http`, `https`, `tcp` if `listenType` is `requestProxy`, `tcp` and `udp` if `listenType` is `packetsTransmit`.
	Protocol *string `pulumi:"protocol"`
	// Listener status. Possible values are: `allNormal` for all resource functioning well, `partNormal` for partial resource functioning well and `allException` for all resource functioning exceptional.
	Status *string `pulumi:"status"`
}

type LbListenerState struct {
	// Health check domain checking.
	Domain pulumi.StringPtrInput
	// Health check method. Possible values are `port` as port checking and `path` as http checking.
	HealthCheckType pulumi.StringPtrInput
	// Amount of time in seconds to wait for the response for in between two sessions if `listenType` is `requestProxy`, range: 0-86400. (Default: `60`). Amount of time in seconds to wait for one session if `listenType` is `packetsTransmit`, range: 60-900. The session will be closed as soon as no response if it is `0`.
	IdleTimeout pulumi.IntPtrInput
	// The type of listener. Possible values are `requestProxy` and `packetsTransmit`. When `packetsTransmit` was specified, you need to config the instances by yourself if the instances attach to the load balancer. You may refer to [configuration instruction](https://docs.ucloud.cn/network/ulb/guide/fu-wu-jie-dian-xiang-guan-cao-zuo/editrealserver).
	ListenType pulumi.StringPtrInput
	// The ID of load balancer instance.
	LoadBalancerId pulumi.StringPtrInput
	// The load balancer method in which the listener is. Possible values are: `roundrobin`, `source`, `consistentHash`, `sourcePort` , `consistentHashPort`, `weightRoundrobin` and `leastconn`. (Default: `roundrobin`).
	// - The `consistentHash`, `sourcePort` , `consistentHashPort`, `roundrobin`, `source` and `weightRoundrobin` are valid if `listenType` is `packetsTransmit`.
	// - The `roundrobin`, `source` and `weightRoundrobin` and `leastconn` are valid if `listenType` is `requestProxy`.
	Method pulumi.StringPtrInput
	Name   pulumi.StringPtrInput
	// Health check path checking.
	Path pulumi.StringPtrInput
	// Indicate whether the persistence session is enabled, it is invalid if `persistenceType` is `none`, an auto-generated string will be exported if `persistenceType` is `serverInsert`, a custom string will be exported if `persistenceType` is `userDefined`.
	Persistence pulumi.StringPtrInput
	// The type of session persistence of listener. Possible values are: `none` as disabled, `serverInsert` as auto-generated key and `userDefined` as customized key. (Default: `none`).
	PersistenceType pulumi.StringPtrInput
	// Port opened on the listeners to receive requests, range: 1-65535. The default value: `80` as `protocol` is `http`, `443` as `protocol` is `https`, `1024` as `protocol` is `tcp` or `udp`.
	Port pulumi.IntPtrInput
	// Listener protocol. Possible values: `http`, `https`, `tcp` if `listenType` is `requestProxy`, `tcp` and `udp` if `listenType` is `packetsTransmit`.
	Protocol pulumi.StringPtrInput
	// Listener status. Possible values are: `allNormal` for all resource functioning well, `partNormal` for partial resource functioning well and `allException` for all resource functioning exceptional.
	Status pulumi.StringPtrInput
}

func (LbListenerState) ElementType() reflect.Type {
	return reflect.TypeOf((*lbListenerState)(nil)).Elem()
}

type lbListenerArgs struct {
	// Health check domain checking.
	Domain *string `pulumi:"domain"`
	// Health check method. Possible values are `port` as port checking and `path` as http checking.
	HealthCheckType *string `pulumi:"healthCheckType"`
	// Amount of time in seconds to wait for the response for in between two sessions if `listenType` is `requestProxy`, range: 0-86400. (Default: `60`). Amount of time in seconds to wait for one session if `listenType` is `packetsTransmit`, range: 60-900. The session will be closed as soon as no response if it is `0`.
	IdleTimeout *int `pulumi:"idleTimeout"`
	// The type of listener. Possible values are `requestProxy` and `packetsTransmit`. When `packetsTransmit` was specified, you need to config the instances by yourself if the instances attach to the load balancer. You may refer to [configuration instruction](https://docs.ucloud.cn/network/ulb/guide/fu-wu-jie-dian-xiang-guan-cao-zuo/editrealserver).
	ListenType *string `pulumi:"listenType"`
	// The ID of load balancer instance.
	LoadBalancerId string `pulumi:"loadBalancerId"`
	// The load balancer method in which the listener is. Possible values are: `roundrobin`, `source`, `consistentHash`, `sourcePort` , `consistentHashPort`, `weightRoundrobin` and `leastconn`. (Default: `roundrobin`).
	// - The `consistentHash`, `sourcePort` , `consistentHashPort`, `roundrobin`, `source` and `weightRoundrobin` are valid if `listenType` is `packetsTransmit`.
	// - The `roundrobin`, `source` and `weightRoundrobin` and `leastconn` are valid if `listenType` is `requestProxy`.
	Method *string `pulumi:"method"`
	Name   *string `pulumi:"name"`
	// Health check path checking.
	Path *string `pulumi:"path"`
	// Indicate whether the persistence session is enabled, it is invalid if `persistenceType` is `none`, an auto-generated string will be exported if `persistenceType` is `serverInsert`, a custom string will be exported if `persistenceType` is `userDefined`.
	Persistence *string `pulumi:"persistence"`
	// The type of session persistence of listener. Possible values are: `none` as disabled, `serverInsert` as auto-generated key and `userDefined` as customized key. (Default: `none`).
	PersistenceType *string `pulumi:"persistenceType"`
	// Port opened on the listeners to receive requests, range: 1-65535. The default value: `80` as `protocol` is `http`, `443` as `protocol` is `https`, `1024` as `protocol` is `tcp` or `udp`.
	Port *int `pulumi:"port"`
	// Listener protocol. Possible values: `http`, `https`, `tcp` if `listenType` is `requestProxy`, `tcp` and `udp` if `listenType` is `packetsTransmit`.
	Protocol string `pulumi:"protocol"`
}

// The set of arguments for constructing a LbListener resource.
type LbListenerArgs struct {
	// Health check domain checking.
	Domain pulumi.StringPtrInput
	// Health check method. Possible values are `port` as port checking and `path` as http checking.
	HealthCheckType pulumi.StringPtrInput
	// Amount of time in seconds to wait for the response for in between two sessions if `listenType` is `requestProxy`, range: 0-86400. (Default: `60`). Amount of time in seconds to wait for one session if `listenType` is `packetsTransmit`, range: 60-900. The session will be closed as soon as no response if it is `0`.
	IdleTimeout pulumi.IntPtrInput
	// The type of listener. Possible values are `requestProxy` and `packetsTransmit`. When `packetsTransmit` was specified, you need to config the instances by yourself if the instances attach to the load balancer. You may refer to [configuration instruction](https://docs.ucloud.cn/network/ulb/guide/fu-wu-jie-dian-xiang-guan-cao-zuo/editrealserver).
	ListenType pulumi.StringPtrInput
	// The ID of load balancer instance.
	LoadBalancerId pulumi.StringInput
	// The load balancer method in which the listener is. Possible values are: `roundrobin`, `source`, `consistentHash`, `sourcePort` , `consistentHashPort`, `weightRoundrobin` and `leastconn`. (Default: `roundrobin`).
	// - The `consistentHash`, `sourcePort` , `consistentHashPort`, `roundrobin`, `source` and `weightRoundrobin` are valid if `listenType` is `packetsTransmit`.
	// - The `roundrobin`, `source` and `weightRoundrobin` and `leastconn` are valid if `listenType` is `requestProxy`.
	Method pulumi.StringPtrInput
	Name   pulumi.StringPtrInput
	// Health check path checking.
	Path pulumi.StringPtrInput
	// Indicate whether the persistence session is enabled, it is invalid if `persistenceType` is `none`, an auto-generated string will be exported if `persistenceType` is `serverInsert`, a custom string will be exported if `persistenceType` is `userDefined`.
	Persistence pulumi.StringPtrInput
	// The type of session persistence of listener. Possible values are: `none` as disabled, `serverInsert` as auto-generated key and `userDefined` as customized key. (Default: `none`).
	PersistenceType pulumi.StringPtrInput
	// Port opened on the listeners to receive requests, range: 1-65535. The default value: `80` as `protocol` is `http`, `443` as `protocol` is `https`, `1024` as `protocol` is `tcp` or `udp`.
	Port pulumi.IntPtrInput
	// Listener protocol. Possible values: `http`, `https`, `tcp` if `listenType` is `requestProxy`, `tcp` and `udp` if `listenType` is `packetsTransmit`.
	Protocol pulumi.StringInput
}

func (LbListenerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lbListenerArgs)(nil)).Elem()
}
