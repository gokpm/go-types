package types

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type StringDuration time.Duration

func (s *StringDuration) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	parsed, err := time.ParseDuration(v)
	if err != nil {
		return err
	}
	*s = StringDuration(parsed)
	return nil
}

func (s *StringDuration) Value() time.Duration {
	return time.Duration(*s)
}

type StringInt int

func (s *StringInt) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	value, err := strconv.Atoi(v)
	if err != nil {
		return err
	}
	*s = StringInt(value)
	return nil
}

func (s *StringInt) Value() int {
	return int(*s)
}

type StringSize float64

func (s *StringSize) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	parsed, err := parseSize(v)
	if err != nil {
		return err
	}
	*s = StringSize(parsed)
	return nil
}

func (s *StringSize) Value() float64 {
	return float64(*s)
}

var sizeMap = map[string]int64{
	"B": 1,
	"K": 1 << 10,
	"M": 1 << 20,
	"G": 1 << 30,
}

func parseSize(v string) (float64, error) {
	for unit, size := range sizeMap {
		if strings.HasSuffix(v, unit) {
			n := strings.TrimSuffix(v, unit)
			f, err := strconv.ParseFloat(n, 64)
			if err != nil {
				return 0, err
			}
			return f * float64(size), nil
		}
	}
	return 0, nil
}
