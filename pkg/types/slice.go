package types

import (
	"encoding/json"
	"strconv"
	"strings"
)

type FloatSlice []float64

func (o *FloatSlice) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	for _, v := range strings.Split(s, "_") {
		if v == "" {
			continue
		}

		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return err
		}

		*o = append(*o, f)
	}
	return nil
}

type IntSlice []int64

func (o *IntSlice) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	for _, v := range strings.Split(s, "_") {
		if v == "" {
			continue
		}

		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return err
		}

		*o = append(*o, i)
	}
	return nil
}
