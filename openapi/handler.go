package main

import (
	"context"
	"errors"
	"fmt"
	"html"
	"net/http"
	"os"

	gologging "github.com/devopsfaith/krakend-gologging"
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

func (r registerer) registerHandlers(ctx context.Context, extra map[string]interface{}, _ http.Handler) (http.Handler, error) {
	fmt.Println("registerHandlers")
	// check the passed configuration and initialize the plugin
	name, ok := extra["name"].(string)
	if !ok {
		return nil, errors.New("wrong config")
	}
	if name != string(r) {
		return nil, fmt.Errorf("unknown register %s", name)
	}
	logger, logger_err := gologging.NewLogger(extra, os.Stdout)
	if logger_err != nil {
		return nil, logger_err
	}
	// return the actual handler wrapping or your custom logic so it can be used as a replacement for the default http handler
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(req.URL.Path))
		logger.Info("processed %q", html.EscapeString(req.URL.Path))
	}), nil
}

func init() {
	fmt.Println("krakend-openapi handler plugin loaded!!!")
}

func main() {}
