package asciicast

import (
	"time"
)

type CastMetadata struct {
	Version   uint
	Width     uint
	Height    uint
	Timestamp time.Time
	Duration  time.Duration
	Command   string
	Title     string
	Env       map[string]string
}
