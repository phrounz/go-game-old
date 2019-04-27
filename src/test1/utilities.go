package main

import (
	"time"
)

func nowMs() int64 {
	return time.Now().UnixNano() / 1000000
}
