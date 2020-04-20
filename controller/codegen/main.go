package main

import (
	controllergen "github.com/rancher/wrangler/pkg/controller-gen"
	"github.com/rancher/wrangler/pkg/controller-gen/args"
	v1 "github.com/shopnado/shopify-controller/controller/apis/shopnado.xyz/v1"
)

func main() {
	controllergen.Run(args.Options{
		OutputPackage: "github.com/shopnado/shopify-controller/controller/generated",
		Boilerplate:   "controller/hack/boilerplate.go.txt",
		Groups: map[string]args.Group{
			"shopnado.xyz": {
				Types: []interface{}{
					v1.Event{},
				},
				GenerateTypes:   true,
				GenerateClients: true,
			},
		},
	})
}
