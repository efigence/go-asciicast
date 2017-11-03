package asciicast

import (
)



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
}

type CastV2Theme struct {
	Fg string `json:"fg"`
	Bg string `json:"bg"`
	Palette string `json:"palette"`
}





func NewCastV2(meta CastMetadata) (c CastV2Header, err error) {
	c.Version = 2
	c.Width = meta.Width
	c.Height = meta.Height
	c.Title = meta.Title
	c.Command = meta.Command
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
	return
}
