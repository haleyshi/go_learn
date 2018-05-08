package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

// Server host and port
type Server struct {
	host string
	port int
}

func (server *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if proxyURL, err := url.Parse("http://" + server.host + ":" + strconv.Itoa(server.port)); err != nil {
		panic(err)
	} else {
		proxy := httputil.NewSingleHostReverseProxy(proxyURL)
		proxy.ServeHTTP(w, r)
	}
}

func startServer() {
	// Redirect :8081 to 127.0.0.1:8080
	h := &Server{host: "127.0.0.1", port: 8080}
	if err := http.ListenAndServe(":8081", h); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() {
	startServer()
}
