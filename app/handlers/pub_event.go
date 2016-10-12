package handlers

import (
	"net/http"
	"encoding/json"
	"time"
	"encoding/binary"
	"github.com/lcaballero/kublai/app/mux"
	"github.com/lcaballero/kublai/queue"
)


type PubEvent struct {
	PublisherID int32
	TopicID     int32
	PayloadSize int32
	Payload     []byte
	State       queue.LogMsgState
}

func PubEventHandler(p mux.Params) error {
	reader := p.Request().Body
	req := PubEvent{}

	binary.Read(reader, binary.LittleEndian, &req.PublisherID)
	binary.Read(reader, binary.LittleEndian, &req.TopicID)
	binary.Read(reader, binary.LittleEndian, &req.PayloadSize)
	req.Payload = make([]byte, req.PayloadSize)
	reader.Read(req.Payload)

	select {
	case p.Appender() <- req:
	default:
		return http.Error(p.Response(), "Enhance Your Calm", http.StatusTooManyRequests)
	}

	timeout := time.NewTicker(1*time.Second)

	select {
	case res := <-p.Appender():
		bin, _ := json.MarshalIndent(&res, "", "  ")
		p.Response().WriteHeader(http.StatusOK)
		p.Response().Header().Add("Content-Type", "text/html")
		p.Response().Write(bin)
		return nil
	case <-timeout.C:
		return http.Error(p.Response(), "Timed out persisting request", http.StatusExpectationFailed)
	}
}

