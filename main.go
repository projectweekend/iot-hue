package main

import (
	"github.com/projectweekend/iot-hue/config"
	"github.com/projectweekend/iot-hue/hue"
	"github.com/projectweekend/iot-hue/iot"
	"log"
)

func main() {
	c := config.FromCLI()

	mqttClient, err := iot.NewClient(c.MqttHost, c.CertPath, c.KeyPath)
	if err != nil {
		log.Fatal(err)
	}

	commands, err := iot.MessageChannel(mqttClient, "iot-hue")
	if err != nil {
		log.Fatal(err)
	}

	lightController := hue.NewLightController(c.HueUsername, c.HueHost)

	for cmd := range commands {
		lightController.HandleCommand(cmd)
	}
}
