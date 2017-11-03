package asciicast

import (
	"fmt"
	"strconv"
	"time"
)

type JSONTimestamp time.Time

func (t JSONTimestamp) MarshalJSON() ([]byte, error) {
	ts := t.Time().Unix()
	stamp := fmt.Sprint(ts)
	return []byte(stamp), nil
}

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

func (t JSONTimestamp) Time() time.Time {
	return time.Time(t)
}
