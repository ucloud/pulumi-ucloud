// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package udisk

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Cloud Disk Attachment resource for attaching Cloud Disk to UHost Instance.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-ucloud/sdk/go/ucloud/uaccount"
// 	"github.com/pulumi/pulumi-ucloud/sdk/go/ucloud/udisk"
// 	"github.com/pulumi/pulumi-ucloud/sdk/go/ucloud/uhost"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		defaultZone, err := uaccount.GetZone(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		opt0 := defaultZone.Zones[0].Id
// 		opt1 := "^CentOS 7.[1-2] 64"
// 		opt2 := "base"
// 		defaultImage, err := uhost.GetImage(ctx, &uhost.GetImageArgs{
// 			AvailabilityZone: &opt0,
// 			NameRegex:        &opt1,
// 			ImageType:        &opt2,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		defaultDisk, err := udisk.NewDisk(ctx, "defaultDisk", &udisk.DiskArgs{
// 			AvailabilityZone: pulumi.String(defaultZone.Zones[0].Id),
// 			DiskSize:         pulumi.Int(10),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		web, err := uhost.NewInstance(ctx, "web", &uhost.InstanceArgs{
// 			AvailabilityZone: pulumi.String(defaultZone.Zones[0].Id),
// 			InstanceType:     pulumi.String("n-basic-2"),
// 			ImageId:          pulumi.String(defaultImage.Images[0].Id),
// 			RootPassword:     pulumi.String("wA1234567"),
// 			Tag:              pulumi.String("tf-example"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = udisk.NewDiskAttachment(ctx, "defaultDiskAttachment", &udisk.DiskAttachmentArgs{
// 			AvailabilityZone: pulumi.String(defaultZone.Zones[0].Id),
// 			DiskId:           defaultDisk.ID(),
// 			InstanceId:       web.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DiskAttachment struct {
	pulumi.CustomResourceState

	// The Zone to attach the disk in.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// The ID of disk that needs to be attached
	DiskId pulumi.StringOutput `pulumi:"diskId"`
	// The ID of instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewDiskAttachment registers a new resource with the given unique name, arguments, and options.
func NewDiskAttachment(ctx *pulumi.Context,
	name string, args *DiskAttachmentArgs, opts ...pulumi.ResourceOption) (*DiskAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilityZone == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityZone'")
	}
	if args.DiskId == nil {
		return nil, errors.New("invalid value for required argument 'DiskId'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	var resource DiskAttachment
	err := ctx.RegisterResource("ucloud:udisk/diskAttachment:DiskAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDiskAttachment gets an existing DiskAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDiskAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskAttachmentState, opts ...pulumi.ResourceOption) (*DiskAttachment, error) {
	var resource DiskAttachment
	err := ctx.ReadResource("ucloud:udisk/diskAttachment:DiskAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DiskAttachment resources.
type diskAttachmentState struct {
	// The Zone to attach the disk in.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The ID of disk that needs to be attached
	DiskId *string `pulumi:"diskId"`
	// The ID of instance.
	InstanceId *string `pulumi:"instanceId"`
}

type DiskAttachmentState struct {
	// The Zone to attach the disk in.
	AvailabilityZone pulumi.StringPtrInput
	// The ID of disk that needs to be attached
	DiskId pulumi.StringPtrInput
	// The ID of instance.
	InstanceId pulumi.StringPtrInput
}

func (DiskAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*diskAttachmentState)(nil)).Elem()
}

type diskAttachmentArgs struct {
	// The Zone to attach the disk in.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// The ID of disk that needs to be attached
	DiskId string `pulumi:"diskId"`
	// The ID of instance.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a DiskAttachment resource.
type DiskAttachmentArgs struct {
	// The Zone to attach the disk in.
	AvailabilityZone pulumi.StringInput
	// The ID of disk that needs to be attached
	DiskId pulumi.StringInput
	// The ID of instance.
	InstanceId pulumi.StringInput
}

func (DiskAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diskAttachmentArgs)(nil)).Elem()
}

type DiskAttachmentInput interface {
	pulumi.Input

	ToDiskAttachmentOutput() DiskAttachmentOutput
	ToDiskAttachmentOutputWithContext(ctx context.Context) DiskAttachmentOutput
}

func (*DiskAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskAttachment)(nil))
}

func (i *DiskAttachment) ToDiskAttachmentOutput() DiskAttachmentOutput {
	return i.ToDiskAttachmentOutputWithContext(context.Background())
}

func (i *DiskAttachment) ToDiskAttachmentOutputWithContext(ctx context.Context) DiskAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskAttachmentOutput)
}

func (i *DiskAttachment) ToDiskAttachmentPtrOutput() DiskAttachmentPtrOutput {
	return i.ToDiskAttachmentPtrOutputWithContext(context.Background())
}

func (i *DiskAttachment) ToDiskAttachmentPtrOutputWithContext(ctx context.Context) DiskAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskAttachmentPtrOutput)
}

type DiskAttachmentPtrInput interface {
	pulumi.Input

	ToDiskAttachmentPtrOutput() DiskAttachmentPtrOutput
	ToDiskAttachmentPtrOutputWithContext(ctx context.Context) DiskAttachmentPtrOutput
}

type diskAttachmentPtrType DiskAttachmentArgs

func (*diskAttachmentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskAttachment)(nil))
}

func (i *diskAttachmentPtrType) ToDiskAttachmentPtrOutput() DiskAttachmentPtrOutput {
	return i.ToDiskAttachmentPtrOutputWithContext(context.Background())
}

func (i *diskAttachmentPtrType) ToDiskAttachmentPtrOutputWithContext(ctx context.Context) DiskAttachmentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskAttachmentPtrOutput)
}

// DiskAttachmentArrayInput is an input type that accepts DiskAttachmentArray and DiskAttachmentArrayOutput values.
// You can construct a concrete instance of `DiskAttachmentArrayInput` via:
//
//          DiskAttachmentArray{ DiskAttachmentArgs{...} }
type DiskAttachmentArrayInput interface {
	pulumi.Input

	ToDiskAttachmentArrayOutput() DiskAttachmentArrayOutput
	ToDiskAttachmentArrayOutputWithContext(context.Context) DiskAttachmentArrayOutput
}

type DiskAttachmentArray []DiskAttachmentInput

func (DiskAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*DiskAttachment)(nil))
}

func (i DiskAttachmentArray) ToDiskAttachmentArrayOutput() DiskAttachmentArrayOutput {
	return i.ToDiskAttachmentArrayOutputWithContext(context.Background())
}

func (i DiskAttachmentArray) ToDiskAttachmentArrayOutputWithContext(ctx context.Context) DiskAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskAttachmentArrayOutput)
}

// DiskAttachmentMapInput is an input type that accepts DiskAttachmentMap and DiskAttachmentMapOutput values.
// You can construct a concrete instance of `DiskAttachmentMapInput` via:
//
//          DiskAttachmentMap{ "key": DiskAttachmentArgs{...} }
type DiskAttachmentMapInput interface {
	pulumi.Input

	ToDiskAttachmentMapOutput() DiskAttachmentMapOutput
	ToDiskAttachmentMapOutputWithContext(context.Context) DiskAttachmentMapOutput
}

type DiskAttachmentMap map[string]DiskAttachmentInput

func (DiskAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*DiskAttachment)(nil))
}

func (i DiskAttachmentMap) ToDiskAttachmentMapOutput() DiskAttachmentMapOutput {
	return i.ToDiskAttachmentMapOutputWithContext(context.Background())
}

func (i DiskAttachmentMap) ToDiskAttachmentMapOutputWithContext(ctx context.Context) DiskAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskAttachmentMapOutput)
}

type DiskAttachmentOutput struct {
	*pulumi.OutputState
}

func (DiskAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskAttachment)(nil))
}

func (o DiskAttachmentOutput) ToDiskAttachmentOutput() DiskAttachmentOutput {
	return o
}

func (o DiskAttachmentOutput) ToDiskAttachmentOutputWithContext(ctx context.Context) DiskAttachmentOutput {
	return o
}

func (o DiskAttachmentOutput) ToDiskAttachmentPtrOutput() DiskAttachmentPtrOutput {
	return o.ToDiskAttachmentPtrOutputWithContext(context.Background())
}

func (o DiskAttachmentOutput) ToDiskAttachmentPtrOutputWithContext(ctx context.Context) DiskAttachmentPtrOutput {
	return o.ApplyT(func(v DiskAttachment) *DiskAttachment {
		return &v
	}).(DiskAttachmentPtrOutput)
}

type DiskAttachmentPtrOutput struct {
	*pulumi.OutputState
}

func (DiskAttachmentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskAttachment)(nil))
}

func (o DiskAttachmentPtrOutput) ToDiskAttachmentPtrOutput() DiskAttachmentPtrOutput {
	return o
}

func (o DiskAttachmentPtrOutput) ToDiskAttachmentPtrOutputWithContext(ctx context.Context) DiskAttachmentPtrOutput {
	return o
}

type DiskAttachmentArrayOutput struct{ *pulumi.OutputState }

func (DiskAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiskAttachment)(nil))
}

func (o DiskAttachmentArrayOutput) ToDiskAttachmentArrayOutput() DiskAttachmentArrayOutput {
	return o
}

func (o DiskAttachmentArrayOutput) ToDiskAttachmentArrayOutputWithContext(ctx context.Context) DiskAttachmentArrayOutput {
	return o
}

func (o DiskAttachmentArrayOutput) Index(i pulumi.IntInput) DiskAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DiskAttachment {
		return vs[0].([]DiskAttachment)[vs[1].(int)]
	}).(DiskAttachmentOutput)
}

type DiskAttachmentMapOutput struct{ *pulumi.OutputState }

func (DiskAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DiskAttachment)(nil))
}

func (o DiskAttachmentMapOutput) ToDiskAttachmentMapOutput() DiskAttachmentMapOutput {
	return o
}

func (o DiskAttachmentMapOutput) ToDiskAttachmentMapOutputWithContext(ctx context.Context) DiskAttachmentMapOutput {
	return o
}

func (o DiskAttachmentMapOutput) MapIndex(k pulumi.StringInput) DiskAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DiskAttachment {
		return vs[0].(map[string]DiskAttachment)[vs[1].(string)]
	}).(DiskAttachmentOutput)
}

func init() {
	pulumi.RegisterOutputType(DiskAttachmentOutput{})
	pulumi.RegisterOutputType(DiskAttachmentPtrOutput{})
	pulumi.RegisterOutputType(DiskAttachmentArrayOutput{})
	pulumi.RegisterOutputType(DiskAttachmentMapOutput{})
}
