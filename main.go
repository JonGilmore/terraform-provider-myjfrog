package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/jfrog/terraform-provider-myjfrog/pkg/myjfrog"
)

// Run the docs generation tool, check its repository for more information on how it works and how docs
// can be customized.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/jfrog/myjfrog",
		Debug:   debug,
	}

	err := providerserver.Serve(context.Background(), myjfrog.NewProvider(), opts)
	if err != nil {
		log.Fatal(err.Error())
	}
}
