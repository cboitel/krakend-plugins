package main

import (
	"fmt"

	gologging "github.com/devopsfaith/krakend-gologging"
	"github.com/devopsfaith/krakend/logging"
)

// PluginVersion defined version of plugin
var PluginVersion = "undefined"

var defaultLogger logging.Logger

func init() {
	defaultExtra := map[string]interface{}{
		"github_com/devopsfaith/krakend-gologging": map[string]interface{}{
			"level":  "INFO",
			"prefix": "[OPENAPI]",
			"syslog": false,
			"stdout": true,
		},
	}
	initLogger, err := gologging.NewLogger(defaultExtra)
	if err != nil {
		panic(fmt.Sprintf("error initializing default logger for openapi plugin: %v", err))
	}
	defaultLogger = initLogger
	defaultLogger.Info("krakend-openapi handler plugin version %s loaded !", PluginVersion)
}

func main() {}
