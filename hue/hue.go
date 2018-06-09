package hue

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var httpClient = &http.Client{
	Timeout: time.Second * 10,
}

// LightController wraps interactions with the Philips HUE HTTP API
type LightController interface {
	GroupOn(string)
	GroupOff(string)
}

// NewLightController returns a LightController with info about the HUE bridge
func NewLightController(username, host string) LightController {
	c := controller{username, host}
	return c
}

type controller struct {
	username string
	host     string
}

func (c controller) GroupOn(name string) {
	url := c.hueGroupActionURL(name)
	body := hueActionOnJSONBody()
	c.hueMakeActionPutRequest(url, body)
}

func (c controller) GroupOff(name string) {
	url := c.hueGroupActionURL(name)
	body := hueActionOffJSONBody()
	c.hueMakeActionPutRequest(url, body)
}

func (c controller) hueURL() string {
	format := "http://%s/api/%s"
	return fmt.Sprintf(format, c.host, c.username)
}

func (c controller) hueGroupActionURL(groupID string) string {
	format := "%s/groups/%s/action"
	return fmt.Sprintf(format, c.hueURL(), groupID)
}

func (c controller) hueMakeActionPutRequest(url string, body io.Reader) {
	req, _ := http.NewRequest("PUT", url, body)
	req.Header.Add("Content-Type", "application/json")
	httpClient.Do(req)
}
