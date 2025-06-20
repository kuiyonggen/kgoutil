package kgoutil

import (
    "fmt"
    "time"
)

func CurrentTime() int64 {
	return time.Now().Unix()
}

// YYYYMMDD -> YYYY-MM-DD
func FormatDate(s string) string {
    if len(s) == 8 {
        return fmt.Sprintf("%s-%s-%s", s[0:4], s[4:6], s[6:])
    } else {
        return s
    }
}
