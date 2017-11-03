package asciicast

import (
	"time"
	"strconv"
	"fmt"
)
type JSONTimestamp time.Time

// MarshalJSON defines how encoding/json marshals the object to JSON,
// the result is a string of the UNIX timestamp
func (t JSONTimestamp) MarshalJSON() ([]byte, error) {
	ts := t.Time().Unix()
	stamp := fmt.Sprint(ts)
	return []byte(stamp), nil
}

// UnmarshalJSON defines how encoding/json unmarshals the object from JSON,
// a UNIX timestamp string is converted to int which is used for the JSONTimestamp
// object value
func (t *JSONTimestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}

	*t = JSONTimestamp(time.Unix(int64(ts), 0))

	return nil
}
func (t JSONTimestamp) String() string {
	return t.Time().String()
}

// Time returns a time.Time object with the same time value as the JSONTimestamp
// object
func (t JSONTimestamp) Time() time.Time {
	if time.Time(t).IsZero() {
		return time.Unix(0, 0)
	}

	return time.Time(t)
}
