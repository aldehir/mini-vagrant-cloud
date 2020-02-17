package server

import (
	"fmt"
	"log"
	"net/http"
)

type BoxServer struct {
	server *http.Server
	serveMux *http.ServeMux
}

func CreateBoxServer(addr string) *BoxServer {
	serveMux := http.NewServeMux()
	server := &http.Server{Addr: addr, Handler: serveMux}

	boxServer := &BoxServer{server, serveMux}
	boxServer.setHandlers()
	return boxServer
}

func (b *BoxServer) ListenAndServe() error {
	log.Printf("Listening on %s\n", b.server.Addr)
	return b.server.ListenAndServe()
}

func (b *BoxServer) setHandlers() {
	b.serveMux.HandleFunc("/", b.viewBoxes)
}

func (b *BoxServer) viewBoxes(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request for %s\n", r.URL.Path)
	fmt.Fprintf(w, "Hello, world!\n")
}
