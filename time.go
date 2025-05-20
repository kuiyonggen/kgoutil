package kgoutil

import "time"

func CurrentTime() int64 {
	return time.Now().Unix()
}
