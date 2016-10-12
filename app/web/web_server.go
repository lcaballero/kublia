package web

import (
	"fmt"
	"github.com/lcaballero/kublai/app/handlers"
	"github.com/lcaballero/kublai/app/mux"
	"github.com/lcaballero/kublai/app/settings"
	"net/http"
)

type WebServer struct {
	settings settings.Settings
}

func NewWebServer(s settings.Settings) (*WebServer, error) {
	ws := &WebServer{
		settings: s,
	}
	return ws, nil
}

func (ws *WebServer) Start() {
	r := mux.NewMux()
	r.Handle("/", ws.settings, handlers.IndexHandler)
	r.Handle("/pub", ws.settings, handlers.PubEventHandler).Methods("POST")

	e := ws.settings.Endpoint()

	fmt.Printf("binding webserver to: %s\n", e)

	err := http.ListenAndServe(e, r.ToHandler())
	if err != nil {
		panic(err)
	}
}
