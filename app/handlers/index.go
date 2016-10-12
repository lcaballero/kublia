package handlers

import (
	"github.com/lcaballero/kublai/app/mux"
	. "github.com/lcaballero/gel"
	"net/http"
	"bytes"
	"github.com/lcaballero/kublai/queue"
	"encoding/binary"
	"encoding/json"
	"time"
)

func IndexHandler(p mux.Params) error {
	el := Frag(
		Div.Class("container").Text("Index Handler"),
	)
	buf := bytes.NewBuffer([]byte{})
	el.ToNode().WriteTo(buf)

	p.Response().WriteHeader(http.StatusOK)
	p.Response().Header().Add("Content-Type", "text/html")
	p.Response().Write(buf.Bytes())
	return nil
}
