package server

import (
	"log"
	"net/http"
)

type BoxServer struct {
	server *http.Server
	serveMux *http.ServeMux
	router *Router
}

func CreateBoxServer(addr string) *BoxServer {
	serveMux := http.NewServeMux()
	server := &http.Server{Addr: addr, Handler: serveMux}

	boxServer := &BoxServer{server, serveMux, CreateRouter()}
	boxServer.createRoutes()

	return boxServer
}

func (b *BoxServer) ListenAndServe() error {
	log.Printf("Listening on %s\n", b.server.Addr)
	return b.server.ListenAndServe()
}

func (b *BoxServer) createRoutes() {
	b.serveMux.Handle("/", b.router)
	b.router.HandleFunc("^/(?P<group>[^/]+)/(?P<name>[^/]+)/?$", b.viewBoxes)
}

func (b *BoxServer) viewBoxes(w http.ResponseWriter, r *http.Request) {
	params := r.Context().Value("params").(map[string]string)

	log.Printf("Received request for %s\n", r.URL.Path)
	log.Printf(".. group: %s\n", params["group"])
	log.Printf(".. name: %s\n", params["name"])
}
