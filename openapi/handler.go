package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	guuid "github.com/google/uuid"
)

type handlerRegisterer string

// HandlerRegisterer is the symbol the plugin loader will try to load. It must implement the Registerer interface
var HandlerRegisterer = handlerRegisterer(PluginName)
var handlerLogPrefix = "[" + PluginName + "-handler]: "

func (r handlerRegisterer) RegisterHandlers(f func(
	name string,
	handler func(context.Context, map[string]interface{}, http.Handler) (http.Handler, error),
)) {
	f(string(r), r.registerHandlers)
}

func (r handlerRegisterer) registerHandlers(ctx context.Context, extra map[string]interface{}, originalHandler http.Handler) (http.Handler, error) {
	openapiLogger.Debug(handlerLogPrefix + fmt.Sprintf("registerHandlers called with extra config %+v", extra))
	// check the passed configuration and initialize the plugin
	name, nameOk := extra["name"].(string)
	if !nameOk {
		return nil, errors.New(handlerLogPrefix + "no name provided in config")
	}
	if name != string(r) {
		return nil, fmt.Errorf(handlerLogPrefix+"name provided not matching %s", name)
	}
	pluginConfig, okConfig := extra[name].(map[string]interface{})
	if !okConfig {
		return nil, fmt.Errorf(handlerLogPrefix+"no \"%s\" property found in config", name)
	}
	urlPattern, urlConfig := pluginConfig["url"].(string)
	if !urlConfig {
		return nil, errors.New(handlerLogPrefix + "no url property set")
	}
	urlRegexp, regexpCompileErr := regexp.Compile(urlPattern)
	if regexpCompileErr != nil {
		return nil, fmt.Errorf(handlerLogPrefix+"url '%s' not a regex: %s", urlPattern, regexpCompileErr)
	}
	openapiLogger.Info(handlerLogPrefix + "enabled to intercept " + urlPattern)

	// return the actual handler wrapping or your custom logic so it can be used as a replacement for the default http handler
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		handlerFuncLogPrefix := handlerLogPrefix + "[" + guuid.New().String() + "] "
		openapiLogger.Debug(handlerFuncLogPrefix + fmt.Sprintf("start of processing %+v", req))
		if urlRegexp.MatchString(req.URL.RequestURI()) {
			openapiLogger.Debug(handlerFuncLogPrefix + "intercepting")
			w.Header().Set("content-type", "text/plain")
			w.Write([]byte("Hello from " + handlerFuncLogPrefix + " !"))
		} else {
			openapiLogger.Debug(handlerFuncLogPrefix + "serving using original handler")
			originalHandler.ServeHTTP(w, req)
		}
		openapiLogger.Debug(handlerFuncLogPrefix + "end of processing")
	}), nil
}
