package main

import (
	"context"
	"errors"
	"fmt"
	"html"
	"net/http"
)

// ClientRegisterer is the symbol the plugin loader will try to load. It must implement the RegisterClient interface
var ClientRegisterer = clientRegisterer(PluginName)

type clientRegisterer string

func (r clientRegisterer) RegisterClients(f func(
	name string,
	handler func(context.Context, map[string]interface{}) (http.Handler, error),
)) {
	f(string(r), r.registerClients)
}

func (r clientRegisterer) registerClients(ctx context.Context, extra map[string]interface{}) (http.Handler, error) {
	openapiLogger.Debug("openapi's plugin registerClients called")
	// check the passed configuration and initialize the plugin
	name, ok := extra["name"].(string)
	if !ok {
		return nil, errors.New("openapi plugin: client: no name provided in config")
	}
	if name != string(r) {
		return nil, fmt.Errorf("openapi plugin: client: name provided not matching %s", name)
	}
	// return the actual handler wrapping or your custom logic so it can be used as a replacement for the default http client
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		openapiLogger.Debug("openapi plugin: client: processed", html.EscapeString(req.URL.Path))
	}), nil
}
