# go-types

A Go package providing custom types that can unmarshal JSON string values into their native Go equivalents.

## Installation

```bash
go get github.com/gokpm/go-types
```

## Usage

This package provides types that parse string-encoded values from JSON:

```go
import "github.com/gokpm/go-types"

type Config struct {
    Timeout     types.StringDuration       `json:"timeout"`
    Port        types.StringInt            `json:"port"`
    Rate        types.StringFloat64        `json:"rate"`
    CacheSize   types.StringBinaryByteSize `json:"cache_size"`
    DiskSize    types.StringDecimalSize    `json:"disk_size"`
    Enabled     types.StringBool           `json:"enabled"`
    Hosts       types.StringArray          `json:"hosts"`
}
```

JSON input:
```json
{
    "timeout": "30s",
    "port": "8080",
    "rate": "1.5",
    "cache_size": "1G",
    "disk_size": "500M",
    "enabled": "true",
    "hosts": "[host1,host2,host3]"
}
```

Unmarshal and access values:
```go
var config Config
err := json.Unmarshal(jsonData, &config)
if err != nil {
    log.Fatal(err)
}

duration := config.Timeout.Value()  // time.Duration
port := config.Port.Value()         // int
```

## Types

- `StringDuration` - Parses duration strings (e.g., "30s", "5m")
- `StringInt` - Parses integer strings
- `StringFloat64` - Parses float strings
- `StringBinaryByteSize` - Parses binary byte sizes (1K = 1024 bytes)
- `StringDecimalSize` - Parses decimal sizes (1K = 1000 units)
- `StringBool` - Parses boolean strings
- `StringArray` - Parses comma-separated string arrays