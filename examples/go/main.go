package main

import (
	"github.com/rebelopsio/pulumi-synadia-cloud/sdk/go/synadia-cloud"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myRandomResource, err := synadia-cloud.NewRandom(ctx, "myRandomResource", &synadia-cloud.RandomArgs{
			Length: pulumi.Int(24),
		})
		if err != nil {
			return err
		}
		ctx.Export("output", map[string]interface{}{
			"value": myRandomResource.Result,
		})
		return nil
	})
}
