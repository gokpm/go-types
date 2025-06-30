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

type StringFloat64 int

func (s *StringFloat64) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	value, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return err
	}
	*s = StringFloat64(value)
	return nil
}

func (s *StringFloat64) Value() float64 {
	return float64(*s)
}

type StringBinaryByteSize float64

func (s *StringBinaryByteSize) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	parsed, err := parseSize(v, binaryByteSizeMap)
	if err != nil {
		return err
	}
	*s = StringBinaryByteSize(parsed)
	return nil
}

func (s *StringBinaryByteSize) Value() float64 {
	return float64(*s)
}

var binaryByteSizeMap = map[string]float64{
	"B": 1,       // 1 B = 1 byte
	"K": 1 << 10, // 1 KiB = 1024 bytes
	"M": 1 << 20, // 1 MiB = 1,048,576 bytes
	"G": 1 << 30, // 1 GiB = 1,073,741,824 bytes
	"T": 1 << 40, // 1 TiB = 1,099,511,627,776 bytes
	"P": 1 << 50, // 1 PiB = 1,125,899,906,842,624 bytes
	"E": 1 << 60, // 1 EiB = 1,152,921,504,606,846,976 bytes
}

var decimalSizeMap = map[string]float64{
	"K": 1000,
	"M": 1000000,
	"G": 1000000000,
	"T": 1000000000000,
	"P": 1000000000000000,
	"E": 1000000000000000000,
}

func parseSize(v string, m map[string]float64) (float64, error) {
	for unit, size := range m {
		if strings.HasSuffix(v, unit) {
			n := strings.TrimSuffix(v, unit)
			f, err := strconv.ParseFloat(n, 64)
			if err != nil {
				return 0, err
			}
			return f * size, nil
		}
	}
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

type StringDecimalSize float64

func (s *StringDecimalSize) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	parsed, err := parseSize(v, decimalSizeMap)
	if err != nil {
		return err
	}
	*s = StringDecimalSize(parsed)
	return nil
}

func (s *StringDecimalSize) Value() float64 {
	return float64(*s)
}

type StringBool bool

func (s *StringBool) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	parsed, err := strconv.ParseBool(v)
	if err != nil {
		return err
	}
	*s = StringBool(parsed)
	return nil
}

func (s *StringBool) Value() bool {
	return bool(*s)
}

type StringArray []string

func (s *StringArray) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	v = strings.Trim(v, "[]")
	parts := strings.Split(v, ",")
	*s = []string{}
	for _, part := range parts {
		part = strings.TrimSpace(part)
		*s = append(*s, part)
	}
	return nil
}

func (s *StringArray) Value() []string {
	return *s
}
