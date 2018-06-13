package iot

import (
	"time"
)

// MqttClient implements only the methods from mqtt.Client that iot package uses
type MqttClient interface {
	Connect() MqttClientToken
	Subscribe(string, byte, MqttMessageHandler) MqttClientToken
}

// MqttClientToken implements only the methods from mqtt.ClientToken that iot package uses
type MqttClientToken interface {
	WaitTimeout(time.Duration) bool
	Error() error
}

// MqttMessageHandler implements the mqtt.MessageHandler interface
type MqttMessageHandler func(MqttClient, MqttMessage)

// MqttMessage implements the interface from mqtt.Message
type MqttMessage interface {
	Duplicate() bool
	Qos() byte
	Retained() bool
	Topic() string
	MessageID() uint16
	Payload() []byte
}
