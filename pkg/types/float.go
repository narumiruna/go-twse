package types

import (
	"encoding/json"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type Float float64

func (f *Float) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		*f = 0
		log.WithError(err).Warnf("failed to parse float: %s", s)
	}

	*f = Float(v)
	return nil
}
