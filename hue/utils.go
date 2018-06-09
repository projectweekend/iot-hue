package hue

import (
	"bytes"
	"io"
)

func hueActionOnJSONBody() io.Reader {
	return bytes.NewBufferString("{\"on\": true}")
}

func hueActionOffJSONBody() io.Reader {
	return bytes.NewBufferString("{\"on\": false}")
}
