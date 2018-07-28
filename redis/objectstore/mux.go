package main

import "net/http"

// Option is a generic configuration option of HTTPMetdhoMux
type Option func(mux *HTTPMethodMux)

func WithHandler(method string, handler http.Handler) Option {
	return func(mux *HTTPMethodMux) {
		mux.methodMap[method] = handler
	}
}

func NewHTTPMethodMux(opts ...Option) *HTTPMethodMux {
	methodMap := make(map[string]http.Handler)
	mux := &HTTPMethodMux{methodMap: methodMap}
	for _, opt := range opts {
		opt(mux)
	}

	return mux
}

// HTTPMethodMux is a HTTP connection multiplexer,
// routes requests to handlers according to request HTTP method
type HTTPMethodMux struct {
	methodMap map[string]http.Handler
}

func (m *HTTPMethodMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h, ok := m.methodMap[r.Method]
	if !ok {
		http.Error(w, "Not implemented", http.StatusNotImplemented)
		return
	}

	h.ServeHTTP(w, r)
}
