from pulumi import export
import pulumi_ucloud as ucloud

images = ucloud.uhost.getImage(
    availability_zone="cn-bj2-04",
    name_regex="^CentOS 8.2 64",
    image_type="base",
)

instance = ucloud.uhost.Instance(
    availability_zone="cn-bj2-04",
    image_id=images.images("0").id,
    instance_type="n-basic-2",
    root_password="wA1234567",
    name="cdktf-example-instance",
    tag="tf-example",
    boot_disk_type="cloud_ssd",
)

export("instance_id", instance.id)
