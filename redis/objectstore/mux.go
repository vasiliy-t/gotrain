package main

import "net/http"

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
