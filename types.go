package types

import (
	"encoding/json"
	"strconv"
	"time"
)

type Duration time.Duration

func (d *Duration) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	parsed, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	*d = Duration(parsed)
	return nil
}

func (d *Duration) Value() time.Duration {
	return time.Duration(*d)
}

type EtcdInt int

func (i *EtcdInt) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	value, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*i = EtcdInt(value)
	return nil
}

func (i *EtcdInt) Value() int {
	return int(*i)
}
