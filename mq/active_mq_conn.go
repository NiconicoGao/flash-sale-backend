package mq

import (
	"crypto/tls"
	"flash-sale-backend/utils"
	"fmt"

	"github.com/go-stomp/stomp/v3"
)

var connMQ *stomp.Conn
var config *utils.ActiveMQConfig

func ReceiveActiveMQ(m *utils.ActiveMQConfig) {
	netConn, err := tls.Dial("tcp", m.Endpoint, &tls.Config{})
	if err != nil {
		println("cannot connect to server", err.Error())
		return
	}
	conn, err := stomp.Connect(netConn, stomp.ConnOpt.Login(m.Username, m.Password))
	if err != nil {
		println("cannot connect to server", err.Error())
		return
	}

	sub, err := conn.Subscribe(m.Queue, stomp.AckAuto)
	if err != nil {
		fmt.Println("subscribe topic error: ", err.Error())
		return
	}

	for m := range sub.C {
		fmt.Println("Received ", string(m.Body))
		messageHandler(m.Body)
	}

}

func ProduceActiveMQ(m []byte) error {
	netConn, err := tls.Dial("tcp", config.Endpoint, &tls.Config{})
	if err != nil {
		return err
	}
	connMQ, err = stomp.Connect(netConn, stomp.ConnOpt.Login(config.Username, config.Password))
	if err != nil {
		return err
	}

	err = connMQ.Send(config.Queue, "text/plain", m, nil)
	if err != nil {
		return err
	}

	connMQ.Disconnect()

	return nil
}

func InitActiveMQ(m *utils.ActiveMQConfig) error {
	go func() {
		for {
			ReceiveActiveMQ(m)
		}
	}()
	config = m
	return nil
}
