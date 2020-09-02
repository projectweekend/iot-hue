你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
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
