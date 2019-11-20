package main

import (
	gologging "github.com/op/go-logging"
)

// PluginVersion defined version of plugin
var PluginVersion = "undefined"

var openapiLogger *gologging.Logger

func init() {
	openapiLogger = gologging.MustGetLogger("OPENAPI")
	openapiLogger.Info("openapi plugin: version", PluginVersion, "loaded !")
}

func main() {}
