# IoT HUE

## Usage
```
Usage of ./iot-hue:
  -cert_path string
    	Path to certificate file
  -hue_host string
    	Hue bridge host
  -hue_username string
    	Hue bridge API username
  -key_path string
    	Path to certificate file
  -mqtt_host string
    	MQTT host and port
```

## Message Payload
This service expects messages to the following payload format:

### Group ID
Turn on all lights in group 3
```
3:on
```

Turn off all lights in group 3
```
3:off
```

### Group Name
When `hue-iot` starts it connects to your HUE bridge and makes a mapping of group names, in lowercase, to group IDs. You can send the group name string in the payload instead of the ID.

Turn on all lights in 'kitchen' group
```
kitchen:on
```

Turn off all lights in 'kitchen' group
```
kitchen:off
```
