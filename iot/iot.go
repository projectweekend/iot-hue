package iot

import (
	"crypto/tls"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/satori/go.uuid"
	"time"
)

// NewClient returns a configured mqtt.Client
func NewClient(host, certPath, keyPath string) (mqtt.Client, error) {
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, err
	}

	u, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	connOptions := mqtt.NewClientOptions()
	connOptions.SetClientID(u.String())
	connOptions.SetMaxReconnectInterval(1 * time.Second)
	connOptions.SetTLSConfig(&tls.Config{Certificates: []tls.Certificate{cert}})
	connOptions.AddBroker(host)

	return mqtt.NewClient(connOptions), nil
}

// MessageChannel subscribes to a topic and returns a channel of strings.
// Messages received on the topic deliver its string payload to the channel.
func MessageChannel(client mqtt.Client, topic string) (<-chan string, error) {
	messageChan := make(chan string)

	token := client.Connect()
	token.WaitTimeout(3 * time.Second)
	err := token.Error()
	if err != nil {
		fmt.Println("Iot connect error")
		return messageChan, err
	}

	callback := func(c mqtt.Client, m mqtt.Message) {
		messageChan <- string(m.Payload())
	}

	token = client.Subscribe(topic, 0, callback)
	token.WaitTimeout(3 * time.Second)
	err = token.Error()
	if err != nil {
		fmt.Println("Iot subscribe error")
		return messageChan, err
	}

	return messageChan, nil
}
