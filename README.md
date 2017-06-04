# date_interval
Package implementing parse function for date interval, such as "1 y 3 m 2 d"

Units can be any of:

* days, day, d
* months, month, mo
* years, year, yr, y

Units are defined as string slices, so can be easily modified:
```go
        d := []string{"days", "day", "d"}
        m := []string{"months", "month", "mo"}
        y := []string{"years", "year", "yr", "y"}
```

## Usage

This package is designed to work with the **time.AddDate**

### func (Time) AddDate

```go
func (t Time) AddDate(years, months, days int) Time
```

Hence it returns the parameters you need to pass to the above

### func Parse
```go
func Parse(interval string) (years, months, days int, err error)
```

## Example

```go
package main

import (
	"fmt"
	"github.com/DavidSantia/date_interval"
)

func main() {

	sample := "2y 1month"

	y, m, d, err := date_interval.Parse(sample)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Parse(%q) returns: %d y, %d m, %d d\n", sample, y, m, d)
}
```

Run: **go run main.go**
```
Parse("2y 1month") returns: 2 y, 1 m, 0 d
```

