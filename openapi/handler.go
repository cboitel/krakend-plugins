package main

import (
	"context"
	"errors"
	"fmt"
	"html"
	"net/http"
)

// HandlerRegisterer is the symbol the plugin loader will try to load. It must implement the Registerer interface
var HandlerRegisterer = registerer("github.com/cboitel/krakend-plugins/openapi")

type registerer string

func (r registerer) RegisterHandlers(f func(
	name string,
	handler func(context.Context, map[string]interface{}, http.Handler) (http.Handler, error),
)) {
	f(string(r), r.registerHandlers)
}

func (r registerer) registerHandlers(ctx context.Context, extra map[string]interface{}, originalHandler http.Handler) (http.Handler, error) {
	openapiLogger.Debug("openapi's plugin registerHandlers called")
	// check the passed configuration and initialize the plugin
	name, ok := extra["name"].(string)
	if !ok {
		return nil, errors.New("openapi plugin: no name provided in config")
	}
	if name != string(r) {
		return nil, fmt.Errorf("openapi plugin: name provided not matching %s", name)
	}
	// return the actual handler wrapping or your custom logic so it can be used as a replacement for the default http handler
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		originalHandler.ServeHTTP(w, req)
		openapiLogger.Debug("openapi plugin: processed", html.EscapeString(req.URL.Path))
	}), nil
}
