package types

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

// StringDuration represents a time.Duration that can be unmarshaled from a JSON string
// Example JSON: "5m30s" -> 5 minutes 30 seconds
type StringDuration time.Duration

// UnmarshalJSON implements json.Unmarshaler interface for StringDuration
// Converts JSON string duration (e.g., "1h30m") to time.Duration
func (s *StringDuration) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	// Parse string using Go's time.ParseDuration
	parsed, err := time.ParseDuration(v)
	if err != nil {
		return err
	}
	*s = StringDuration(parsed)
	return nil
}

// Value returns the underlying time.Duration value
func (s *StringDuration) Value() time.Duration {
	return time.Duration(*s)
}

// StringInt represents an integer that can be unmarshaled from a JSON string
// Example JSON: "42" -> 42
type StringInt int

// UnmarshalJSON implements json.Unmarshaler interface for StringInt
// Converts JSON string number to int
func (s *StringInt) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	// Convert string to integer
	value, err := strconv.Atoi(v)
	if err != nil {
		return err
	}
	*s = StringInt(value)
	return nil
}

// Value returns the underlying int value
func (s *StringInt) Value() int {
	return int(*s)
}

// StringFloat64 represents a float64 that can be unmarshaled from a JSON string
// Note: The underlying type should be float64, not int (appears to be a typo)
// Example JSON: "3.14159" -> 3.14159
type StringFloat64 int // TODO: This should probably be float64

// UnmarshalJSON implements json.Unmarshaler interface for StringFloat64
// Converts JSON string number to float64
func (s *StringFloat64) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	// Parse string as 64-bit float
	value, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return err
	}
	*s = StringFloat64(value)
	return nil
}

// Value returns the underlying float64 value
func (s *StringFloat64) Value() float64 {
	return float64(*s)
}

// StringBinaryByteSize represents a byte size using binary units (1024-based)
// Example JSON: "1.5G" -> 1610612736 (1.5 * 1024^3)
type StringBinaryByteSize float64

// UnmarshalJSON implements json.Unmarshaler interface for StringBinaryByteSize
// Converts JSON string size with binary units (K, M, G, T, P, E) to float64 bytes
func (s *StringBinaryByteSize) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	// Parse size string using binary byte size map
	parsed, err := parseSize(v, binaryByteSizeMap)
	if err != nil {
		return err
	}
	*s = StringBinaryByteSize(parsed)
	return nil
}

// Value returns the underlying float64 value representing bytes
func (s *StringBinaryByteSize) Value() float64 {
	return float64(*s)
}

// binaryByteSizeMap defines binary (base-2) size multipliers
// Uses powers of 2 (1024-based) as per IEC binary prefixes
var binaryByteSizeMap = map[string]float64{
	"B": 1,       // 1 B = 1 byte
	"K": 1 << 10, // 1 KiB = 1024 bytes
	"M": 1 << 20, // 1 MiB = 1,048,576 bytes
	"G": 1 << 30, // 1 GiB = 1,073,741,824 bytes
	"T": 1 << 40, // 1 TiB = 1,099,511,627,776 bytes
	"P": 1 << 50, // 1 PiB = 1,125,899,906,842,624 bytes
	"E": 1 << 60, // 1 EiB = 1,152,921,504,606,846,976 bytes
}

// decimalSizeMap defines decimal (base-10) size multipliers
// Uses powers of 10 (1000-based) as per SI decimal prefixes
var decimalSizeMap = map[string]float64{
	"K": 1000,                // 1 KB = 1,000 bytes
	"M": 1000000,             // 1 MB = 1,000,000 bytes
	"G": 1000000000,          // 1 GB = 1,000,000,000 bytes
	"T": 1000000000000,       // 1 TB = 1,000,000,000,000 bytes
	"P": 1000000000000000,    // 1 PB = 1,000,000,000,000,000 bytes
	"E": 1000000000000000000, // 1 EB = 1,000,000,000,000,000,000 bytes
}

// parseSize parses a size string (e.g., "1.5G") using the provided unit map
// Returns the size in bytes as float64
// If no unit suffix is found, treats the value as raw bytes
func parseSize(v string, m map[string]float64) (float64, error) {
	// Check each unit suffix in the map
	for unit, size := range m {
		if strings.HasSuffix(v, unit) {
			// Extract numeric part by removing unit suffix
			n := strings.TrimSuffix(v, unit)
			f, err := strconv.ParseFloat(n, 64)
			if err != nil {
				return 0, err
			}
			// Multiply by unit size
			return f * size, nil
		}
	}
	// No unit found, parse as raw number (assumed to be bytes)
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}

// StringDecimalSize represents a byte size using decimal units (1000-based)
// Example JSON: "1.5G" -> 1500000000 (1.5 * 1000^3)
type StringDecimalSize float64

// UnmarshalJSON implements json.Unmarshaler interface for StringDecimalSize
// Converts JSON string size with decimal units (K, M, G, T, P, E) to float64 bytes
func (s *StringDecimalSize) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	// Parse size string using decimal size map
	parsed, err := parseSize(v, decimalSizeMap)
	if err != nil {
		return err
	}
	*s = StringDecimalSize(parsed)
	return nil
}

// Value returns the underlying float64 value representing bytes
func (s *StringDecimalSize) Value() float64 {
	return float64(*s)
}

// StringBool represents a boolean that can be unmarshaled from a JSON string
// Example JSON: "true" -> true, "false" -> false, "1" -> true, "0" -> false
type StringBool bool

// UnmarshalJSON implements json.Unmarshaler interface for StringBool
// Converts JSON string boolean to bool using Go's strconv.ParseBool
func (s *StringBool) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	// ParseBool accepts: "1", "t", "T", "TRUE", "true", "True", "0", "f", "F", "FALSE", "false", "False"
	parsed, err := strconv.ParseBool(v)
	if err != nil {
		return err
	}
	*s = StringBool(parsed)
	return nil
}

// Value returns the underlying bool value
func (s *StringBool) Value() bool {
	return bool(*s)
}

// StringArray represents a string slice that can be unmarshaled from a JSON string
// Supports both comma-separated values and array-like strings
// Example JSON: "[\"item1\", \"item2\", \"item3\"]" or "item1,item2,item3"
type StringArray []string

// UnmarshalJSON implements json.Unmarshaler interface for StringArray
// Parses comma-separated string values, handling optional brackets and quotes
func (s *StringArray) UnmarshalJSON(b []byte) error {
	var v string
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	// Remove optional surrounding brackets
	v = strings.Trim(v, "[]")
	// Split on commas
	parts := strings.Split(v, ",")
	*s = []string{}
	// Process each part: trim whitespace and quotes
	for _, part := range parts {
		part = strings.TrimSpace(part)  // Remove leading/trailing whitespace
		part = strings.Trim(part, "\"") // Remove surrounding quotes
		*s = append(*s, part)
	}
	return nil
}

// Value returns the underlying string slice
func (s *StringArray) Value() []string {
	return *s
}
