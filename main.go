// Copyright (c) Citrix, Inc.

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/citrix/packer-plugin-citrix/builder/citrix/applayering"
	applayeringProv "github.com/citrix/packer-plugin-citrix/provisioner/citrix/applayering"
	applayeringVersion "github.com/citrix/packer-plugin-citrix/version"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	pps := plugin.NewSet()
	pps.RegisterBuilder("applayering", &applayering.Builder{
		StrategyFactory: applayering.NewBuilderStrategy,
	})
	pps.RegisterProvisioner("layerartifact", new(applayeringProv.Provisioner))
	pps.SetVersion(applayeringVersion.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
