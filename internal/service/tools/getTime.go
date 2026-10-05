package tools

import (
	"time"
)

func GetTime() string {
	return time.Now().Format("15:04:05")
}
