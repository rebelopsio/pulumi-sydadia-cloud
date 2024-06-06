import pulumi
import pulumi_synadia-cloud as synadia-cloud

my_random_resource = synadia-cloud.Random("myRandomResource", length=24)
pulumi.export("output", {
    "value": my_random_resource.result,
})
