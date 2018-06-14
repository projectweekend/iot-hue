package hue

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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

func (c controller) GroupOn(id string) {
	url := c.hueGroupActionURL(id)
	body := hueActionOnJSONBody()
	c.hueMakeActionPutRequest(url, body)
}

func (c controller) GroupOff(id string) {
	url := c.hueGroupActionURL(id)
	body := hueActionOffJSONBody()
	c.hueMakeActionPutRequest(url, body)
}

func (c controller) groupNameIDMapping() map[string]string {
	outputMap := make(map[string]string)

	req, _ := http.NewRequest("GET", c.hueGroupListURL(), nil)
	req.Header.Add("Content-Type", "application/json")
	res, _ := httpClient.Do(req)
	defer res.Body.Close()

	var data map[string]map[string]interface{}
	err := json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		fmt.Println("Failed to decode JSON body")
		return outputMap
	}

	for groupID, dataMap := range data {
		name := strings.ToLower(dataMap["name"].(string))
		outputMap[name] = groupID
	}
	return outputMap
}

func (c controller) hueURL() string {
	format := "http://%s/api/%s"
	return fmt.Sprintf(format, c.host, c.username)
}

func (c controller) hueGroupListURL() string {
	format := "%s/groups"
	return fmt.Sprintf(format, c.hueURL())
}

func (c controller) hueGroupActionURL(groupID string) string {
	format := "%s/groups/%s/action"
	return fmt.Sprintf(format, c.hueURL(), groupID)
}

func (c controller) hueMakeActionPutRequest(url string, body io.Reader) {
	req, _ := http.NewRequest("PUT", url, body)
	req.Header.Add("Content-Type", "application/json")
	res, _ := httpClient.Do(req)
	res.Body.Close()
}
