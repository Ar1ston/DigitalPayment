package NATS

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"time"
)

type ConnectNATS struct {
	Host string
	Port string
}

func (dns *ConnectNATS) dnsToString() (string, error) {
	if dns == nil {
		return "", fmt.Errorf("%s", "Connection data is null")
	}
	rpl := "nats://"
	rpl += dns.Host + ":"
	rpl += dns.Port
	return rpl, nil
}
func (dns *ConnectNATS) ConnectToNATS() (*nats.Conn, error) {
	dnsString, err := dns.dnsToString()
	if err != nil {
		return nil, err
	}
	conn, err := nats.Connect(dnsString, nats.MaxReconnects(-1), nats.ReconnectWait(time.Second*5))
	if err != nil {
		return nil, fmt.Errorf("ERROR CONNECT TO NATS")
	}
	return conn, nil
}
