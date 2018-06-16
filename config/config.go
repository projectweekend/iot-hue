package config

import (
	"flag"
)

func FromCLI() CLIArgs {
	args := CLIArgs{}

	flag.StringVar(&args.CertPath, "cert_path", "", "Path to certificate file")
	flag.StringVar(&args.KeyPath, "key_path", "", "Path to certificate file")
	flag.StringVar(&args.MqttHost, "mqtt_host", "", "MQTT host and port")
	flag.StringVar(&args.HueHost, "hue_host", "", "Hue bridge host")
	flag.StringVar(&args.HueUsername, "hue_username", "", "Hue bridge API username")

	flag.Parse()

	return args
}

type CLIArgs struct {
	CertPath    string
	KeyPath     string
	MqttHost    string
	HueHost     string
	HueUsername string
}
