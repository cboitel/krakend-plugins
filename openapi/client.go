package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	guuid "github.com/google/uuid"
)

type clientRegisterer string

// ClientRegisterer is the symbol the plugin loader will try to load. It must implement the RegisterClient interface
var ClientRegisterer = clientRegisterer(PluginName)
var clientLogPrefix = "[" + PluginName + "-client]: "

func (r clientRegisterer) RegisterClients(f func(
	name string,
	handler func(context.Context, map[string]interface{}) (http.Handler, error),
)) {
	f(string(r), r.registerClients)
}

func (r clientRegisterer) registerClients(ctx context.Context, extra map[string]interface{}) (http.Handler, error) {
	openapiLogger.Debug(clientLogPrefix + fmt.Sprintf("registerClients called with extra config %+v", extra))
	// check the passed configuration and initialize the plugin
	name, ok := extra["name"].(string)
	if !ok {
		return nil, errors.New(clientLogPrefix + "no name provided in config")
	}
	if name != string(r) {
		return nil, fmt.Errorf(clientLogPrefix+"name provided not matching %s", name)
	}
	pluginConfig, okConfig := extra[name].(map[string]interface{})
	if !okConfig {
		return nil, fmt.Errorf(clientLogPrefix+"no \"%s\" property found in config", name)
	}
	urlPattern, urlConfig := pluginConfig["url"].(string)
	if !urlConfig {
		return nil, errors.New(clientLogPrefix + "no url property set")
	}
	urlRegexp, regexpCompileErr := regexp.Compile(urlPattern)
	if regexpCompileErr != nil {
		return nil, fmt.Errorf(clientLogPrefix+"url '%s' not a regex: %s", urlPattern, regexpCompileErr)
	}
	openapiLogger.Info(clientLogPrefix + "enabled to intercept " + urlPattern)
	// return the actual handler wrapping or your custom logic so it can be used as a replacement for the default http client
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		clientFuncLogPrefix := clientLogPrefix + "[" + guuid.New().String() + "] "
		openapiLogger.Debug(clientFuncLogPrefix + fmt.Sprintf("start of processing %+v", req))
		if urlRegexp.MatchString(req.URL.RequestURI()) {
			openapiLogger.Debug(clientFuncLogPrefix + "intercepting")
			w.Header().Set("content-type", "text/plain")
			w.Write([]byte("Hello from " + clientFuncLogPrefix + " !"))
		} else {
			openapiLogger.Debug(clientFuncLogPrefix + "do nothing")
		}
		openapiLogger.Debug(clientFuncLogPrefix + "end of processing")
	}), nil
}
