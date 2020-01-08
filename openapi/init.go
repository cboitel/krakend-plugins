package main

import (
	gologging "github.com/op/go-logging"
)

// PluginVersion defines version of plugin
var PluginVersion = "develop"

// PluginBuildDate defines date at which build occured
var PluginBuildDate = "N/A"

// PluginName defines name of plugin used in config
var PluginName = "github_com/cboitel/krakend-plugins/openapi"

var openapiLogger *gologging.Logger

func init() {
	openapiLogger = gologging.MustGetLogger("OPENAPI")
	openapiLogger.Info("openapi plugin: version:", PluginVersion)
	openapiLogger.Info("openapi plugin: buildDate:", PluginBuildDate)
	openapiLogger.Info("openapi plugin: name:", PluginName)
}

func main() {}
