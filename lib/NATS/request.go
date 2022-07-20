package NATS

import (
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
	resp, err := nats.Request(req.To, toBytes(req), 7*time.Second)
	if err != nil {
		return nil, err
	}
	rpl := RequestNats{}

	err = json.Unmarshal(resp.Data, &rpl)
	if err != nil {
		return nil, err
	}
	//if rpl.To != req.From {
	//	return nil, fmt.Errorf("ERROR RESPONS FROM NATS (%s)", req.To)
	//}
	return rpl.Msg, nil
}
func toBytes(message interface{}) []byte {

	js, err := json.Marshal(message)
	if err != nil {
		log.Printf("Unable to marshal: %+v: %s", message, err)
		return nil
	}

	return js
}
