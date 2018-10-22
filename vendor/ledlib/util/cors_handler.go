package util

import "net/http"

func NewCORSHandler(corsAllowed string, handler func(http.ResponseWriter, *http.Request)) http.Handler {
	return &CORSHandler{corsAllowed, handler}
}

type CORSHandler struct {
	corsAllowed string
	handler     func(http.ResponseWriter, *http.Request)
}

func (c *CORSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", c.corsAllowed)
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	c.handler(w, r)
}
