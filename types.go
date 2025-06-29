package types

import (
	"encoding/json"
	"strconv"
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

func (d *StringDuration) Value() time.Duration {
	return time.Duration(*d)
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
