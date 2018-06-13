package iot

import (
	"time"
)

// NewMockMqttClient returns a mock MQTT client for testing
func NewMockMqttClient() MqttClient {
	return mockMqttClient{}
}

type mockMqttClient struct{}

// Connect is a mock for testing
func (c mockMqttClient) Connect() MqttClientToken {
	return mockClientToken{}
}

// Subscribe is a mock for testing
func (c mockMqttClient) Subscribe(s string, qos byte, m MqttMessageHandler) MqttClientToken {
	return mockClientToken{}
}

// MockClientToken is used for testing
type mockClientToken struct{}

// WaitTimeout is a mock for testing
func (t mockClientToken) WaitTimeout(d time.Duration) bool {
	return true
}

// Error is a mock for testing
func (t mockClientToken) Error() error {
	return nil
}
