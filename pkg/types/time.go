package types

import (
	"encoding/json"
	"strconv"
	"time"
)

type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}

	*t = Time(time.Unix(v/1000, 0))
	return nil
}
