package NATS

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

type RequestNats struct {
	Msg         []byte
	To          string
	RequestName string
	From        string
}

func (req *RequestNats) SendRequestToNats(nats *nats.Conn) ([]byte, error) {

	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(req)
	if err != nil {
		return nil, err
	}

	resp, err := nats.Request(req.To, buff.Bytes(), time.Second)
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}
func toBytes(message interface{}) []byte {

	js, err := json.Marshal(message)
	if err != nil {
		log.Printf("Unable to marshal: %+v: %s", message, err)
		return nil
	}

	return js
}
