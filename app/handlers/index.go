package handlers

import (
	"github.com/lcaballero/kublai/app/mux"
	. "github.com/lcaballero/gel"
	"net/http"
	"bytes"
	"github.com/lcaballero/kublai/queue"
	"encoding/binary"
	"encoding/json"
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

type PubResponse struct {
	PublisherID int32
	TopicID     int32
	PayloadSize int32
	State       queue.LogMsgState
}

type PubRequest struct {
	PublisherID int32
	TopicID     int32
	PayloadSize int32
	Payload     []byte
}

func PubEventHandler(p mux.Params) error {
	req := p.Request()
	ev := PubRequest{}

	binary.Read(req.Body, binary.LittleEndian, &ev.PublisherID)
	binary.Read(req.Body, binary.LittleEndian, &ev.TopicID)
	binary.Read(req.Body, binary.LittleEndian, &ev.PayloadSize)
	ev.Payload = make([]byte, ev.PayloadSize)
	req.Body.Read(ev.Payload)

//	fmt.Println("pub id", ev.PublisherID)
//	fmt.Println("top id", ev.TopicID)
//	fmt.Println("payload size", ev.PayloadSize)
//	fmt.Printf("%v\n", ev.Payload)

	bin, _ := json.MarshalIndent(&ev, "", "  ")

	p.Response().WriteHeader(http.StatusOK)
	p.Response().Header().Add("Content-Type", "text/html")
	p.Response().Write(bin)
	return nil
}
