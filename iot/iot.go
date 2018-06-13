package iot

import (
	"fmt"
	"time"
)

// MessageChannel subscribes to a topic and returns a channel of strings.
// Messages received on the topic deliver their JSON string payload to the channel.
func MessageChannel(client MqttClient, topic string) (chan string, error) {
	messageChan := make(chan string)

	token := client.Connect()
	token.WaitTimeout(3 * time.Second)
	err := token.Error()
	if err != nil {
		fmt.Println("Iot connect error")
		return messageChan, err
	}

	callback := func(c MqttClient, m MqttMessage) {
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
