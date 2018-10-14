package util

import "net/http"

func NewCORSHandler(handler func(http.ResponseWriter, *http.Request)) http.Handler {
	return &CORSHandler{handler}
}

type CORSHandler struct {
	handler func(http.ResponseWriter, *http.Request)
}

func (c *CORSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	c.handler(w, r)
}
