// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ulb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type LbIpSet struct {
	// Type of Elastic IP routes.
	InternetType *string `pulumi:"internetType"`
	// Elastic IP address.
	Ip *string `pulumi:"ip"`
}

type LbIpSetInput interface {
	pulumi.Input

	ToLbIpSetOutput() LbIpSetOutput
	ToLbIpSetOutputWithContext(context.Context) LbIpSetOutput
}

type LbIpSetArgs struct {
	// Type of Elastic IP routes.
	InternetType pulumi.StringPtrInput `pulumi:"internetType"`
	// Elastic IP address.
	Ip pulumi.StringPtrInput `pulumi:"ip"`
}

func (LbIpSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LbIpSet)(nil)).Elem()
}

func (i LbIpSetArgs) ToLbIpSetOutput() LbIpSetOutput {
	return i.ToLbIpSetOutputWithContext(context.Background())
}

func (i LbIpSetArgs) ToLbIpSetOutputWithContext(ctx context.Context) LbIpSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbIpSetOutput)
}

type LbIpSetArrayInput interface {
	pulumi.Input

	ToLbIpSetArrayOutput() LbIpSetArrayOutput
	ToLbIpSetArrayOutputWithContext(context.Context) LbIpSetArrayOutput
}

type LbIpSetArray []LbIpSetInput

func (LbIpSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LbIpSet)(nil)).Elem()
}

func (i LbIpSetArray) ToLbIpSetArrayOutput() LbIpSetArrayOutput {
	return i.ToLbIpSetArrayOutputWithContext(context.Background())
}

func (i LbIpSetArray) ToLbIpSetArrayOutputWithContext(ctx context.Context) LbIpSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbIpSetArrayOutput)
}

type LbIpSetOutput struct{ *pulumi.OutputState }

func (LbIpSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LbIpSet)(nil)).Elem()
}

func (o LbIpSetOutput) ToLbIpSetOutput() LbIpSetOutput {
	return o
}

func (o LbIpSetOutput) ToLbIpSetOutputWithContext(ctx context.Context) LbIpSetOutput {
	return o
}

// Type of Elastic IP routes.
func (o LbIpSetOutput) InternetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LbIpSet) *string { return v.InternetType }).(pulumi.StringPtrOutput)
}

// Elastic IP address.
func (o LbIpSetOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LbIpSet) *string { return v.Ip }).(pulumi.StringPtrOutput)
}

type LbIpSetArrayOutput struct{ *pulumi.OutputState }

func (LbIpSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LbIpSet)(nil)).Elem()
}

func (o LbIpSetArrayOutput) ToLbIpSetArrayOutput() LbIpSetArrayOutput {
	return o
}

func (o LbIpSetArrayOutput) ToLbIpSetArrayOutputWithContext(ctx context.Context) LbIpSetArrayOutput {
	return o
}

func (o LbIpSetArrayOutput) Index(i pulumi.IntInput) LbIpSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LbIpSet {
		return vs[0].([]LbIpSet)[vs[1].(int)]
	}).(LbIpSetOutput)
}

type LookupLbAttachmentsLbAttachment struct {
	// The ID of LB Attachment.
	Id string `pulumi:"id"`
	// Port opened on the backend server to receive requests, range: 1-65535.
	Port int `pulumi:"port"`
	// The private ip address for backend servers.
	PrivateIp string `pulumi:"privateIp"`
	// The ID of a backend server.
	ResourceId string `pulumi:"resourceId"`
	// The status of backend servers. Possible values are: `normalRunning`, `exceptionRunning`.
	Status string `pulumi:"status"`
}

type LookupLbAttachmentsLbAttachmentInput interface {
	pulumi.Input

	ToLookupLbAttachmentsLbAttachmentOutput() LookupLbAttachmentsLbAttachmentOutput
	ToLookupLbAttachmentsLbAttachmentOutputWithContext(context.Context) LookupLbAttachmentsLbAttachmentOutput
}

type LookupLbAttachmentsLbAttachmentArgs struct {
	// The ID of LB Attachment.
	Id pulumi.StringInput `pulumi:"id"`
	// Port opened on the backend server to receive requests, range: 1-65535.
	Port pulumi.IntInput `pulumi:"port"`
	// The private ip address for backend servers.
	PrivateIp pulumi.StringInput `pulumi:"privateIp"`
	// The ID of a backend server.
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	// The status of backend servers. Possible values are: `normalRunning`, `exceptionRunning`.
	Status pulumi.StringInput `pulumi:"status"`
}

func (LookupLbAttachmentsLbAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbAttachmentsLbAttachment)(nil)).Elem()
}

func (i LookupLbAttachmentsLbAttachmentArgs) ToLookupLbAttachmentsLbAttachmentOutput() LookupLbAttachmentsLbAttachmentOutput {
	return i.ToLookupLbAttachmentsLbAttachmentOutputWithContext(context.Background())
}

func (i LookupLbAttachmentsLbAttachmentArgs) ToLookupLbAttachmentsLbAttachmentOutputWithContext(ctx context.Context) LookupLbAttachmentsLbAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupLbAttachmentsLbAttachmentOutput)
}

type LookupLbAttachmentsLbAttachmentArrayInput interface {
	pulumi.Input

	ToLookupLbAttachmentsLbAttachmentArrayOutput() LookupLbAttachmentsLbAttachmentArrayOutput
	ToLookupLbAttachmentsLbAttachmentArrayOutputWithContext(context.Context) LookupLbAttachmentsLbAttachmentArrayOutput
}

type LookupLbAttachmentsLbAttachmentArray []LookupLbAttachmentsLbAttachmentInput

func (LookupLbAttachmentsLbAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LookupLbAttachmentsLbAttachment)(nil)).Elem()
}

func (i LookupLbAttachmentsLbAttachmentArray) ToLookupLbAttachmentsLbAttachmentArrayOutput() LookupLbAttachmentsLbAttachmentArrayOutput {
	return i.ToLookupLbAttachmentsLbAttachmentArrayOutputWithContext(context.Background())
}

func (i LookupLbAttachmentsLbAttachmentArray) ToLookupLbAttachmentsLbAttachmentArrayOutputWithContext(ctx context.Context) LookupLbAttachmentsLbAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupLbAttachmentsLbAttachmentArrayOutput)
}

type LookupLbAttachmentsLbAttachmentOutput struct{ *pulumi.OutputState }

func (LookupLbAttachmentsLbAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbAttachmentsLbAttachment)(nil)).Elem()
}

func (o LookupLbAttachmentsLbAttachmentOutput) ToLookupLbAttachmentsLbAttachmentOutput() LookupLbAttachmentsLbAttachmentOutput {
	return o
}

func (o LookupLbAttachmentsLbAttachmentOutput) ToLookupLbAttachmentsLbAttachmentOutputWithContext(ctx context.Context) LookupLbAttachmentsLbAttachmentOutput {
	return o
}

// The ID of LB Attachment.
func (o LookupLbAttachmentsLbAttachmentOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbAttachmentsLbAttachment) string { return v.Id }).(pulumi.StringOutput)
}

// Port opened on the backend server to receive requests, range: 1-65535.
func (o LookupLbAttachmentsLbAttachmentOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupLbAttachmentsLbAttachment) int { return v.Port }).(pulumi.IntOutput)
}

// The private ip address for backend servers.
func (o LookupLbAttachmentsLbAttachmentOutput) PrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbAttachmentsLbAttachment) string { return v.PrivateIp }).(pulumi.StringOutput)
}

// The ID of a backend server.
func (o LookupLbAttachmentsLbAttachmentOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbAttachmentsLbAttachment) string { return v.ResourceId }).(pulumi.StringOutput)
}

// The status of backend servers. Possible values are: `normalRunning`, `exceptionRunning`.
func (o LookupLbAttachmentsLbAttachmentOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbAttachmentsLbAttachment) string { return v.Status }).(pulumi.StringOutput)
}

type LookupLbAttachmentsLbAttachmentArrayOutput struct{ *pulumi.OutputState }

func (LookupLbAttachmentsLbAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LookupLbAttachmentsLbAttachment)(nil)).Elem()
}

func (o LookupLbAttachmentsLbAttachmentArrayOutput) ToLookupLbAttachmentsLbAttachmentArrayOutput() LookupLbAttachmentsLbAttachmentArrayOutput {
	return o
}

func (o LookupLbAttachmentsLbAttachmentArrayOutput) ToLookupLbAttachmentsLbAttachmentArrayOutputWithContext(ctx context.Context) LookupLbAttachmentsLbAttachmentArrayOutput {
	return o
}

func (o LookupLbAttachmentsLbAttachmentArrayOutput) Index(i pulumi.IntInput) LookupLbAttachmentsLbAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LookupLbAttachmentsLbAttachment {
		return vs[0].([]LookupLbAttachmentsLbAttachment)[vs[1].(int)]
	}).(LookupLbAttachmentsLbAttachmentOutput)
}

type LookupLbListenersLbListener struct {
	// Health check domain checking.
	Domain string `pulumi:"domain"`
	// Health check method. Possible values are `port` as port checking and `path` as http checking.
	HealthCheckType string `pulumi:"healthCheckType"`
	// The ID of LB Listener.
	Id string `pulumi:"id"`
	// Amount of time in seconds to wait for the response for in between two sessions if `listenType` is `requestProxy`, range: 0-86400. Amount of time in seconds to wait for one session if `listenType` is `packetsTransmit`, range: 60-900. The session will be closed as soon as no response if it is `0`.
	IdleTimeout int `pulumi:"idleTimeout"`
	// The type of LB Listener. Possible values are `requestProxy` and `packetsTransmit`.
	ListenType string `pulumi:"listenType"`
	// The load balancer method in which the listener is. Possible values are: `roundrobin`, `source`, `consistentHash`, `sourcePort` , `consistentHashPort`, `weightRoundrobin` and `leastconn`.
	// - The `consistentHash`, `sourcePort` , `consistentHashPort`, `roundrobin`, `source` and `weightRoundrobin` are valid if `listenType` is `packetsTransmit`.
	// - The `rundrobin`, `source` and `weightRoundrobin` and `leastconn` are vaild if `listenType` is `requestProxy`.
	Method string `pulumi:"method"`
	// The name of LB Listener.
	Name string `pulumi:"name"`
	// Health check path checking.
	Path string `pulumi:"path"`
	// Indicate whether the persistence session is enabled, it is invaild if `persistenceType` is `none`, an auto-generated string will be exported if `persistenceType` is `serverInsert`, a custom string will be exported if `persistenceType` is `userDefined`.
	Persistence string `pulumi:"persistence"`
	// The type of session persistence of LB Listener. Possible values are: `none` as disabled, `serverInsert` as auto-generated string and `userDefined` as cutom string. (Default: `none`).
	PersistenceType string `pulumi:"persistenceType"`
	// Port opened on the LB Listener to receive requests, range: 1-65535.
	Port int `pulumi:"port"`
	// LB Listener protocol. Possible values: `http`, `https` if `listenType` is `requestProxy`, `tcp` and `udp` if `listenType` is `packetsTransmit`.
	Protocol string `pulumi:"protocol"`
	// LB Listener status. Possible values are: `allNormal` for all resource functioning well, `partNormal` for partial resource functioning well and `allException` for all resource functioning exceptional.
	Status string `pulumi:"status"`
}

type LookupLbListenersLbListenerInput interface {
	pulumi.Input

	ToLookupLbListenersLbListenerOutput() LookupLbListenersLbListenerOutput
	ToLookupLbListenersLbListenerOutputWithContext(context.Context) LookupLbListenersLbListenerOutput
}

type LookupLbListenersLbListenerArgs struct {
	// Health check domain checking.
	Domain pulumi.StringInput `pulumi:"domain"`
	// Health check method. Possible values are `port` as port checking and `path` as http checking.
	HealthCheckType pulumi.StringInput `pulumi:"healthCheckType"`
	// The ID of LB Listener.
	Id pulumi.StringInput `pulumi:"id"`
	// Amount of time in seconds to wait for the response for in between two sessions if `listenType` is `requestProxy`, range: 0-86400. Amount of time in seconds to wait for one session if `listenType` is `packetsTransmit`, range: 60-900. The session will be closed as soon as no response if it is `0`.
	IdleTimeout pulumi.IntInput `pulumi:"idleTimeout"`
	// The type of LB Listener. Possible values are `requestProxy` and `packetsTransmit`.
	ListenType pulumi.StringInput `pulumi:"listenType"`
	// The load balancer method in which the listener is. Possible values are: `roundrobin`, `source`, `consistentHash`, `sourcePort` , `consistentHashPort`, `weightRoundrobin` and `leastconn`.
	// - The `consistentHash`, `sourcePort` , `consistentHashPort`, `roundrobin`, `source` and `weightRoundrobin` are valid if `listenType` is `packetsTransmit`.
	// - The `rundrobin`, `source` and `weightRoundrobin` and `leastconn` are vaild if `listenType` is `requestProxy`.
	Method pulumi.StringInput `pulumi:"method"`
	// The name of LB Listener.
	Name pulumi.StringInput `pulumi:"name"`
	// Health check path checking.
	Path pulumi.StringInput `pulumi:"path"`
	// Indicate whether the persistence session is enabled, it is invaild if `persistenceType` is `none`, an auto-generated string will be exported if `persistenceType` is `serverInsert`, a custom string will be exported if `persistenceType` is `userDefined`.
	Persistence pulumi.StringInput `pulumi:"persistence"`
	// The type of session persistence of LB Listener. Possible values are: `none` as disabled, `serverInsert` as auto-generated string and `userDefined` as cutom string. (Default: `none`).
	PersistenceType pulumi.StringInput `pulumi:"persistenceType"`
	// Port opened on the LB Listener to receive requests, range: 1-65535.
	Port pulumi.IntInput `pulumi:"port"`
	// LB Listener protocol. Possible values: `http`, `https` if `listenType` is `requestProxy`, `tcp` and `udp` if `listenType` is `packetsTransmit`.
	Protocol pulumi.StringInput `pulumi:"protocol"`
	// LB Listener status. Possible values are: `allNormal` for all resource functioning well, `partNormal` for partial resource functioning well and `allException` for all resource functioning exceptional.
	Status pulumi.StringInput `pulumi:"status"`
}

func (LookupLbListenersLbListenerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbListenersLbListener)(nil)).Elem()
}

func (i LookupLbListenersLbListenerArgs) ToLookupLbListenersLbListenerOutput() LookupLbListenersLbListenerOutput {
	return i.ToLookupLbListenersLbListenerOutputWithContext(context.Background())
}

func (i LookupLbListenersLbListenerArgs) ToLookupLbListenersLbListenerOutputWithContext(ctx context.Context) LookupLbListenersLbListenerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupLbListenersLbListenerOutput)
}

type LookupLbListenersLbListenerArrayInput interface {
	pulumi.Input

	ToLookupLbListenersLbListenerArrayOutput() LookupLbListenersLbListenerArrayOutput
	ToLookupLbListenersLbListenerArrayOutputWithContext(context.Context) LookupLbListenersLbListenerArrayOutput
}

type LookupLbListenersLbListenerArray []LookupLbListenersLbListenerInput

func (LookupLbListenersLbListenerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LookupLbListenersLbListener)(nil)).Elem()
}

func (i LookupLbListenersLbListenerArray) ToLookupLbListenersLbListenerArrayOutput() LookupLbListenersLbListenerArrayOutput {
	return i.ToLookupLbListenersLbListenerArrayOutputWithContext(context.Background())
}

func (i LookupLbListenersLbListenerArray) ToLookupLbListenersLbListenerArrayOutputWithContext(ctx context.Context) LookupLbListenersLbListenerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupLbListenersLbListenerArrayOutput)
}

type LookupLbListenersLbListenerOutput struct{ *pulumi.OutputState }

func (LookupLbListenersLbListenerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbListenersLbListener)(nil)).Elem()
}

func (o LookupLbListenersLbListenerOutput) ToLookupLbListenersLbListenerOutput() LookupLbListenersLbListenerOutput {
	return o
}

func (o LookupLbListenersLbListenerOutput) ToLookupLbListenersLbListenerOutputWithContext(ctx context.Context) LookupLbListenersLbListenerOutput {
	return o
}

// Health check domain checking.
func (o LookupLbListenersLbListenerOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbListenersLbListener) string { return v.Domain }).(pulumi.StringOutput)
}

// Health check method. Possible values are `port` as port checking and `path` as http checking.
func (o LookupLbListenersLbListenerOutput) HealthCheckType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbListenersLbListener) string { return v.HealthCheckType }).(pulumi.StringOutput)
}

// The ID of LB Listener.
func (o LookupLbListenersLbListenerOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbListenersLbListener) string { return v.Id }).(pulumi.StringOutput)
}

// Amount of time in seconds to wait for the response for in between two sessions if `listenType` is `requestProxy`, range: 0-86400. Amount of time in seconds to wait for one session if `listenType` is `packetsTransmit`, range: 60-900. The session will be closed as soon as no response if it is `0`.
func (o LookupLbListenersLbListenerOutput) IdleTimeout() pulumi.IntOutput {
	return o.ApplyT(func(v LookupLbListenersLbListener) int { return v.IdleTimeout }).(pulumi.IntOutput)
}

// The type of LB Listener. Possible values are `requestProxy` and `packetsTransmit`.
func (o LookupLbListenersLbListenerOutput) ListenType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbListenersLbListener) string { return v.ListenType }).(pulumi.StringOutput)
}

// The load balancer method in which the listener is. Possible values are: `roundrobin`, `source`, `consistentHash`, `sourcePort` , `consistentHashPort`, `weightRoundrobin` and `leastconn`.
// - The `consistentHash`, `sourcePort` , `consistentHashPort`, `roundrobin`, `source` and `weightRoundrobin` are valid if `listenType` is `packetsTransmit`.
// - The `rundrobin`, `source` and `weightRoundrobin` and `leastconn` are vaild if `listenType` is `requestProxy`.
func (o LookupLbListenersLbListenerOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbListenersLbListener) string { return v.Method }).(pulumi.StringOutput)
}

// The name of LB Listener.
func (o LookupLbListenersLbListenerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbListenersLbListener) string { return v.Name }).(pulumi.StringOutput)
}

// Health check path checking.
func (o LookupLbListenersLbListenerOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbListenersLbListener) string { return v.Path }).(pulumi.StringOutput)
}

// Indicate whether the persistence session is enabled, it is invaild if `persistenceType` is `none`, an auto-generated string will be exported if `persistenceType` is `serverInsert`, a custom string will be exported if `persistenceType` is `userDefined`.
func (o LookupLbListenersLbListenerOutput) Persistence() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbListenersLbListener) string { return v.Persistence }).(pulumi.StringOutput)
}

// The type of session persistence of LB Listener. Possible values are: `none` as disabled, `serverInsert` as auto-generated string and `userDefined` as cutom string. (Default: `none`).
func (o LookupLbListenersLbListenerOutput) PersistenceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbListenersLbListener) string { return v.PersistenceType }).(pulumi.StringOutput)
}

// Port opened on the LB Listener to receive requests, range: 1-65535.
func (o LookupLbListenersLbListenerOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupLbListenersLbListener) int { return v.Port }).(pulumi.IntOutput)
}

// LB Listener protocol. Possible values: `http`, `https` if `listenType` is `requestProxy`, `tcp` and `udp` if `listenType` is `packetsTransmit`.
func (o LookupLbListenersLbListenerOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbListenersLbListener) string { return v.Protocol }).(pulumi.StringOutput)
}

// LB Listener status. Possible values are: `allNormal` for all resource functioning well, `partNormal` for partial resource functioning well and `allException` for all resource functioning exceptional.
func (o LookupLbListenersLbListenerOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbListenersLbListener) string { return v.Status }).(pulumi.StringOutput)
}

type LookupLbListenersLbListenerArrayOutput struct{ *pulumi.OutputState }

func (LookupLbListenersLbListenerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LookupLbListenersLbListener)(nil)).Elem()
}

func (o LookupLbListenersLbListenerArrayOutput) ToLookupLbListenersLbListenerArrayOutput() LookupLbListenersLbListenerArrayOutput {
	return o
}

func (o LookupLbListenersLbListenerArrayOutput) ToLookupLbListenersLbListenerArrayOutputWithContext(ctx context.Context) LookupLbListenersLbListenerArrayOutput {
	return o
}

func (o LookupLbListenersLbListenerArrayOutput) Index(i pulumi.IntInput) LookupLbListenersLbListenerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LookupLbListenersLbListener {
		return vs[0].([]LookupLbListenersLbListener)[vs[1].(int)]
	}).(LookupLbListenersLbListenerOutput)
}

type LookupLbRulesLbRule struct {
	// (Optional) The domain of content forward matching fields. `path` and `domain` cannot coexist.
	Domain string `pulumi:"domain"`
	// The ID of LB Rule.
	Id string `pulumi:"id"`
	// (Optional) The path of Content forward matching fields. `path` and `domain` cannot coexist.
	Path string `pulumi:"path"`
}

type LookupLbRulesLbRuleInput interface {
	pulumi.Input

	ToLookupLbRulesLbRuleOutput() LookupLbRulesLbRuleOutput
	ToLookupLbRulesLbRuleOutputWithContext(context.Context) LookupLbRulesLbRuleOutput
}

type LookupLbRulesLbRuleArgs struct {
	// (Optional) The domain of content forward matching fields. `path` and `domain` cannot coexist.
	Domain pulumi.StringInput `pulumi:"domain"`
	// The ID of LB Rule.
	Id pulumi.StringInput `pulumi:"id"`
	// (Optional) The path of Content forward matching fields. `path` and `domain` cannot coexist.
	Path pulumi.StringInput `pulumi:"path"`
}

func (LookupLbRulesLbRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbRulesLbRule)(nil)).Elem()
}

func (i LookupLbRulesLbRuleArgs) ToLookupLbRulesLbRuleOutput() LookupLbRulesLbRuleOutput {
	return i.ToLookupLbRulesLbRuleOutputWithContext(context.Background())
}

func (i LookupLbRulesLbRuleArgs) ToLookupLbRulesLbRuleOutputWithContext(ctx context.Context) LookupLbRulesLbRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupLbRulesLbRuleOutput)
}

type LookupLbRulesLbRuleArrayInput interface {
	pulumi.Input

	ToLookupLbRulesLbRuleArrayOutput() LookupLbRulesLbRuleArrayOutput
	ToLookupLbRulesLbRuleArrayOutputWithContext(context.Context) LookupLbRulesLbRuleArrayOutput
}

type LookupLbRulesLbRuleArray []LookupLbRulesLbRuleInput

func (LookupLbRulesLbRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LookupLbRulesLbRule)(nil)).Elem()
}

func (i LookupLbRulesLbRuleArray) ToLookupLbRulesLbRuleArrayOutput() LookupLbRulesLbRuleArrayOutput {
	return i.ToLookupLbRulesLbRuleArrayOutputWithContext(context.Background())
}

func (i LookupLbRulesLbRuleArray) ToLookupLbRulesLbRuleArrayOutputWithContext(ctx context.Context) LookupLbRulesLbRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupLbRulesLbRuleArrayOutput)
}

type LookupLbRulesLbRuleOutput struct{ *pulumi.OutputState }

func (LookupLbRulesLbRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbRulesLbRule)(nil)).Elem()
}

func (o LookupLbRulesLbRuleOutput) ToLookupLbRulesLbRuleOutput() LookupLbRulesLbRuleOutput {
	return o
}

func (o LookupLbRulesLbRuleOutput) ToLookupLbRulesLbRuleOutputWithContext(ctx context.Context) LookupLbRulesLbRuleOutput {
	return o
}

// (Optional) The domain of content forward matching fields. `path` and `domain` cannot coexist.
func (o LookupLbRulesLbRuleOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbRulesLbRule) string { return v.Domain }).(pulumi.StringOutput)
}

// The ID of LB Rule.
func (o LookupLbRulesLbRuleOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbRulesLbRule) string { return v.Id }).(pulumi.StringOutput)
}

// (Optional) The path of Content forward matching fields. `path` and `domain` cannot coexist.
func (o LookupLbRulesLbRuleOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbRulesLbRule) string { return v.Path }).(pulumi.StringOutput)
}

type LookupLbRulesLbRuleArrayOutput struct{ *pulumi.OutputState }

func (LookupLbRulesLbRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LookupLbRulesLbRule)(nil)).Elem()
}

func (o LookupLbRulesLbRuleArrayOutput) ToLookupLbRulesLbRuleArrayOutput() LookupLbRulesLbRuleArrayOutput {
	return o
}

func (o LookupLbRulesLbRuleArrayOutput) ToLookupLbRulesLbRuleArrayOutputWithContext(ctx context.Context) LookupLbRulesLbRuleArrayOutput {
	return o
}

func (o LookupLbRulesLbRuleArrayOutput) Index(i pulumi.IntInput) LookupLbRulesLbRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LookupLbRulesLbRule {
		return vs[0].([]LookupLbRulesLbRule)[vs[1].(int)]
	}).(LookupLbRulesLbRuleOutput)
}

type LookupLbSslsLbSsl struct {
	// The time of creation for lb ssl, formatted in RFC3339 time string.
	CreateTime string `pulumi:"createTime"`
	// The ID of LB SSL certificate resource.
	Id string `pulumi:"id"`
	// The name of LB SSL certificate resource.
	Name string `pulumi:"name"`
}

type LookupLbSslsLbSslInput interface {
	pulumi.Input

	ToLookupLbSslsLbSslOutput() LookupLbSslsLbSslOutput
	ToLookupLbSslsLbSslOutputWithContext(context.Context) LookupLbSslsLbSslOutput
}

type LookupLbSslsLbSslArgs struct {
	// The time of creation for lb ssl, formatted in RFC3339 time string.
	CreateTime pulumi.StringInput `pulumi:"createTime"`
	// The ID of LB SSL certificate resource.
	Id pulumi.StringInput `pulumi:"id"`
	// The name of LB SSL certificate resource.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupLbSslsLbSslArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbSslsLbSsl)(nil)).Elem()
}

func (i LookupLbSslsLbSslArgs) ToLookupLbSslsLbSslOutput() LookupLbSslsLbSslOutput {
	return i.ToLookupLbSslsLbSslOutputWithContext(context.Background())
}

func (i LookupLbSslsLbSslArgs) ToLookupLbSslsLbSslOutputWithContext(ctx context.Context) LookupLbSslsLbSslOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupLbSslsLbSslOutput)
}

type LookupLbSslsLbSslArrayInput interface {
	pulumi.Input

	ToLookupLbSslsLbSslArrayOutput() LookupLbSslsLbSslArrayOutput
	ToLookupLbSslsLbSslArrayOutputWithContext(context.Context) LookupLbSslsLbSslArrayOutput
}

type LookupLbSslsLbSslArray []LookupLbSslsLbSslInput

func (LookupLbSslsLbSslArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LookupLbSslsLbSsl)(nil)).Elem()
}

func (i LookupLbSslsLbSslArray) ToLookupLbSslsLbSslArrayOutput() LookupLbSslsLbSslArrayOutput {
	return i.ToLookupLbSslsLbSslArrayOutputWithContext(context.Background())
}

func (i LookupLbSslsLbSslArray) ToLookupLbSslsLbSslArrayOutputWithContext(ctx context.Context) LookupLbSslsLbSslArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupLbSslsLbSslArrayOutput)
}

type LookupLbSslsLbSslOutput struct{ *pulumi.OutputState }

func (LookupLbSslsLbSslOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbSslsLbSsl)(nil)).Elem()
}

func (o LookupLbSslsLbSslOutput) ToLookupLbSslsLbSslOutput() LookupLbSslsLbSslOutput {
	return o
}

func (o LookupLbSslsLbSslOutput) ToLookupLbSslsLbSslOutputWithContext(ctx context.Context) LookupLbSslsLbSslOutput {
	return o
}

// The time of creation for lb ssl, formatted in RFC3339 time string.
func (o LookupLbSslsLbSslOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbSslsLbSsl) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The ID of LB SSL certificate resource.
func (o LookupLbSslsLbSslOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbSslsLbSsl) string { return v.Id }).(pulumi.StringOutput)
}

// The name of LB SSL certificate resource.
func (o LookupLbSslsLbSslOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbSslsLbSsl) string { return v.Name }).(pulumi.StringOutput)
}

type LookupLbSslsLbSslArrayOutput struct{ *pulumi.OutputState }

func (LookupLbSslsLbSslArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LookupLbSslsLbSsl)(nil)).Elem()
}

func (o LookupLbSslsLbSslArrayOutput) ToLookupLbSslsLbSslArrayOutput() LookupLbSslsLbSslArrayOutput {
	return o
}

func (o LookupLbSslsLbSslArrayOutput) ToLookupLbSslsLbSslArrayOutputWithContext(ctx context.Context) LookupLbSslsLbSslArrayOutput {
	return o
}

func (o LookupLbSslsLbSslArrayOutput) Index(i pulumi.IntInput) LookupLbSslsLbSslOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LookupLbSslsLbSsl {
		return vs[0].([]LookupLbSslsLbSsl)[vs[1].(int)]
	}).(LookupLbSslsLbSslOutput)
}

type LookupLbsLb struct {
	// The creation time of Load Balancer, formatted in RFC3339 time string.
	CreateTime string `pulumi:"createTime"`
	// The ID of Load Balancer.
	Id string `pulumi:"id"`
	// Indicate whether the load balancer is intranet.
	Internal bool               `pulumi:"internal"`
	IpSets   []LookupLbsLbIpSet `pulumi:"ipSets"`
	// The name of Load Balancer.
	Name string `pulumi:"name"`
	// The IP address of intranet IP.
	PrivateIp string `pulumi:"privateIp"`
	// The remarks of Load Balancer.
	Remark string `pulumi:"remark"`
	// The ID of subnet that intrant load balancer belongs to.
	SubnetId string `pulumi:"subnetId"`
	// A tag assigned to Load Balancer.
	Tag string `pulumi:"tag"`
	// The ID of the VPC linked to the Load Balancers.
	VpcId string `pulumi:"vpcId"`
}

type LookupLbsLbInput interface {
	pulumi.Input

	ToLookupLbsLbOutput() LookupLbsLbOutput
	ToLookupLbsLbOutputWithContext(context.Context) LookupLbsLbOutput
}

type LookupLbsLbArgs struct {
	// The creation time of Load Balancer, formatted in RFC3339 time string.
	CreateTime pulumi.StringInput `pulumi:"createTime"`
	// The ID of Load Balancer.
	Id pulumi.StringInput `pulumi:"id"`
	// Indicate whether the load balancer is intranet.
	Internal pulumi.BoolInput           `pulumi:"internal"`
	IpSets   LookupLbsLbIpSetArrayInput `pulumi:"ipSets"`
	// The name of Load Balancer.
	Name pulumi.StringInput `pulumi:"name"`
	// The IP address of intranet IP.
	PrivateIp pulumi.StringInput `pulumi:"privateIp"`
	// The remarks of Load Balancer.
	Remark pulumi.StringInput `pulumi:"remark"`
	// The ID of subnet that intrant load balancer belongs to.
	SubnetId pulumi.StringInput `pulumi:"subnetId"`
	// A tag assigned to Load Balancer.
	Tag pulumi.StringInput `pulumi:"tag"`
	// The ID of the VPC linked to the Load Balancers.
	VpcId pulumi.StringInput `pulumi:"vpcId"`
}

func (LookupLbsLbArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbsLb)(nil)).Elem()
}

func (i LookupLbsLbArgs) ToLookupLbsLbOutput() LookupLbsLbOutput {
	return i.ToLookupLbsLbOutputWithContext(context.Background())
}

func (i LookupLbsLbArgs) ToLookupLbsLbOutputWithContext(ctx context.Context) LookupLbsLbOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupLbsLbOutput)
}

type LookupLbsLbArrayInput interface {
	pulumi.Input

	ToLookupLbsLbArrayOutput() LookupLbsLbArrayOutput
	ToLookupLbsLbArrayOutputWithContext(context.Context) LookupLbsLbArrayOutput
}

type LookupLbsLbArray []LookupLbsLbInput

func (LookupLbsLbArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LookupLbsLb)(nil)).Elem()
}

func (i LookupLbsLbArray) ToLookupLbsLbArrayOutput() LookupLbsLbArrayOutput {
	return i.ToLookupLbsLbArrayOutputWithContext(context.Background())
}

func (i LookupLbsLbArray) ToLookupLbsLbArrayOutputWithContext(ctx context.Context) LookupLbsLbArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupLbsLbArrayOutput)
}

type LookupLbsLbOutput struct{ *pulumi.OutputState }

func (LookupLbsLbOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbsLb)(nil)).Elem()
}

func (o LookupLbsLbOutput) ToLookupLbsLbOutput() LookupLbsLbOutput {
	return o
}

func (o LookupLbsLbOutput) ToLookupLbsLbOutputWithContext(ctx context.Context) LookupLbsLbOutput {
	return o
}

// The creation time of Load Balancer, formatted in RFC3339 time string.
func (o LookupLbsLbOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbsLb) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The ID of Load Balancer.
func (o LookupLbsLbOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbsLb) string { return v.Id }).(pulumi.StringOutput)
}

// Indicate whether the load balancer is intranet.
func (o LookupLbsLbOutput) Internal() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLbsLb) bool { return v.Internal }).(pulumi.BoolOutput)
}

func (o LookupLbsLbOutput) IpSets() LookupLbsLbIpSetArrayOutput {
	return o.ApplyT(func(v LookupLbsLb) []LookupLbsLbIpSet { return v.IpSets }).(LookupLbsLbIpSetArrayOutput)
}

// The name of Load Balancer.
func (o LookupLbsLbOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbsLb) string { return v.Name }).(pulumi.StringOutput)
}

// The IP address of intranet IP.
func (o LookupLbsLbOutput) PrivateIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbsLb) string { return v.PrivateIp }).(pulumi.StringOutput)
}

// The remarks of Load Balancer.
func (o LookupLbsLbOutput) Remark() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbsLb) string { return v.Remark }).(pulumi.StringOutput)
}

// The ID of subnet that intrant load balancer belongs to.
func (o LookupLbsLbOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbsLb) string { return v.SubnetId }).(pulumi.StringOutput)
}

// A tag assigned to Load Balancer.
func (o LookupLbsLbOutput) Tag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbsLb) string { return v.Tag }).(pulumi.StringOutput)
}

// The ID of the VPC linked to the Load Balancers.
func (o LookupLbsLbOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbsLb) string { return v.VpcId }).(pulumi.StringOutput)
}

type LookupLbsLbArrayOutput struct{ *pulumi.OutputState }

func (LookupLbsLbArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LookupLbsLb)(nil)).Elem()
}

func (o LookupLbsLbArrayOutput) ToLookupLbsLbArrayOutput() LookupLbsLbArrayOutput {
	return o
}

func (o LookupLbsLbArrayOutput) ToLookupLbsLbArrayOutputWithContext(ctx context.Context) LookupLbsLbArrayOutput {
	return o
}

func (o LookupLbsLbArrayOutput) Index(i pulumi.IntInput) LookupLbsLbOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LookupLbsLb {
		return vs[0].([]LookupLbsLb)[vs[1].(int)]
	}).(LookupLbsLbOutput)
}

type LookupLbsLbIpSet struct {
	// Type of Load Balancer routes.
	InternetType string `pulumi:"internetType"`
	// Load Balancer address.
	Ip string `pulumi:"ip"`
}

type LookupLbsLbIpSetInput interface {
	pulumi.Input

	ToLookupLbsLbIpSetOutput() LookupLbsLbIpSetOutput
	ToLookupLbsLbIpSetOutputWithContext(context.Context) LookupLbsLbIpSetOutput
}

type LookupLbsLbIpSetArgs struct {
	// Type of Load Balancer routes.
	InternetType pulumi.StringInput `pulumi:"internetType"`
	// Load Balancer address.
	Ip pulumi.StringInput `pulumi:"ip"`
}

func (LookupLbsLbIpSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbsLbIpSet)(nil)).Elem()
}

func (i LookupLbsLbIpSetArgs) ToLookupLbsLbIpSetOutput() LookupLbsLbIpSetOutput {
	return i.ToLookupLbsLbIpSetOutputWithContext(context.Background())
}

func (i LookupLbsLbIpSetArgs) ToLookupLbsLbIpSetOutputWithContext(ctx context.Context) LookupLbsLbIpSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupLbsLbIpSetOutput)
}

type LookupLbsLbIpSetArrayInput interface {
	pulumi.Input

	ToLookupLbsLbIpSetArrayOutput() LookupLbsLbIpSetArrayOutput
	ToLookupLbsLbIpSetArrayOutputWithContext(context.Context) LookupLbsLbIpSetArrayOutput
}

type LookupLbsLbIpSetArray []LookupLbsLbIpSetInput

func (LookupLbsLbIpSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LookupLbsLbIpSet)(nil)).Elem()
}

func (i LookupLbsLbIpSetArray) ToLookupLbsLbIpSetArrayOutput() LookupLbsLbIpSetArrayOutput {
	return i.ToLookupLbsLbIpSetArrayOutputWithContext(context.Background())
}

func (i LookupLbsLbIpSetArray) ToLookupLbsLbIpSetArrayOutputWithContext(ctx context.Context) LookupLbsLbIpSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupLbsLbIpSetArrayOutput)
}

type LookupLbsLbIpSetOutput struct{ *pulumi.OutputState }

func (LookupLbsLbIpSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbsLbIpSet)(nil)).Elem()
}

func (o LookupLbsLbIpSetOutput) ToLookupLbsLbIpSetOutput() LookupLbsLbIpSetOutput {
	return o
}

func (o LookupLbsLbIpSetOutput) ToLookupLbsLbIpSetOutputWithContext(ctx context.Context) LookupLbsLbIpSetOutput {
	return o
}

// Type of Load Balancer routes.
func (o LookupLbsLbIpSetOutput) InternetType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbsLbIpSet) string { return v.InternetType }).(pulumi.StringOutput)
}

// Load Balancer address.
func (o LookupLbsLbIpSetOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbsLbIpSet) string { return v.Ip }).(pulumi.StringOutput)
}

type LookupLbsLbIpSetArrayOutput struct{ *pulumi.OutputState }

func (LookupLbsLbIpSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LookupLbsLbIpSet)(nil)).Elem()
}

func (o LookupLbsLbIpSetArrayOutput) ToLookupLbsLbIpSetArrayOutput() LookupLbsLbIpSetArrayOutput {
	return o
}

func (o LookupLbsLbIpSetArrayOutput) ToLookupLbsLbIpSetArrayOutputWithContext(ctx context.Context) LookupLbsLbIpSetArrayOutput {
	return o
}

func (o LookupLbsLbIpSetArrayOutput) Index(i pulumi.IntInput) LookupLbsLbIpSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LookupLbsLbIpSet {
		return vs[0].([]LookupLbsLbIpSet)[vs[1].(int)]
	}).(LookupLbsLbIpSetOutput)
}

func init() {
	pulumi.RegisterOutputType(LbIpSetOutput{})
	pulumi.RegisterOutputType(LbIpSetArrayOutput{})
	pulumi.RegisterOutputType(LookupLbAttachmentsLbAttachmentOutput{})
	pulumi.RegisterOutputType(LookupLbAttachmentsLbAttachmentArrayOutput{})
	pulumi.RegisterOutputType(LookupLbListenersLbListenerOutput{})
	pulumi.RegisterOutputType(LookupLbListenersLbListenerArrayOutput{})
	pulumi.RegisterOutputType(LookupLbRulesLbRuleOutput{})
	pulumi.RegisterOutputType(LookupLbRulesLbRuleArrayOutput{})
	pulumi.RegisterOutputType(LookupLbSslsLbSslOutput{})
	pulumi.RegisterOutputType(LookupLbSslsLbSslArrayOutput{})
	pulumi.RegisterOutputType(LookupLbsLbOutput{})
	pulumi.RegisterOutputType(LookupLbsLbArrayOutput{})
	pulumi.RegisterOutputType(LookupLbsLbIpSetOutput{})
	pulumi.RegisterOutputType(LookupLbsLbIpSetArrayOutput{})
}