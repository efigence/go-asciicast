package asciicast

import (
	"io"
	"encoding/json"
)

const V2OutputEvent = "o"
const V2InputEvent = "i"
const V2SizeEvent = "size"

type CastV2Header struct {
	Version uint `json:"version"`
	Width uint `json:"width"`
	Height uint `json:"height"`
	Timestamp *JSONTimestamp `json:"timestamp,omitempty"`
	Duration float64 `json:"duration,omitempty"`
	Title string  `json:"title,omitempty"`
	Command string  `json:"command,omitempty"`
	Env *map[string]string `json:"env,omitempty"`
	// this is a pointer only because that's the easiest way to force Golang's
	// JSON marshaller to not emit it if empty
	Theme *CastV2Theme `json:"theme,omitempty"`
	outputStream *json.Encoder
}

type CastV2Theme struct {
	Fg string `json:"fg"`
	Bg string `json:"bg"`
	Palette string `json:"palette"`
}

// Version 2 cast uses streaming, therefore you need to pass IO stream to it
// Header will be automatically emitted after cast is created
func NewCastV2(meta CastMetadata, fd io.Writer) (*CastV2Header, error) {
	var c CastV2Header
	c.Version = 2
	c.Width = meta.Width
	c.Height = meta.Height
	c.Title = meta.Title
	c.Command = meta.Command
	// drop TS if it is null or before unix time
	if meta.Timestamp.Unix() > 0 {
		ts := JSONTimestamp(meta.Timestamp)
		c.Timestamp = &ts
	}
	if meta.Duration.Seconds() > 0 {
		c.Duration = meta.Duration.Seconds()
	}
	if meta.Env != nil {
		c.Env = &meta.Env
	}
	c.outputStream = json.NewEncoder(fd)
	c.outputStream.Encode(c)
	return &c, nil
}

func (c *CastV2Header) PushFrame(ts float64, data []byte) {
	out := make([]interface{},3)
	out[0] = ts
	out[1] = V2OutputEvent
	out[2] = string(data)
	c.outputStream.Encode(out)
}
