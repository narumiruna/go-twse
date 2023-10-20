package types

import (
	"encoding/json"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type Int int64

func (i *Int) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.WithError(err).Warnf("failed to parse integer: %s", s)
		*i = 0
	}

	*i = Int(v)
	return nil
}
