package iot

import (
	"testing"
)

func TestMessageChannel(t *testing.T) {
	mockClient := NewMockMqttClient()

	_, err := MessageChannel(mockClient, "test")
	if err != nil {
		t.Error()
	}

}
