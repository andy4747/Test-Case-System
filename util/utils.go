package util

import (
	"fmt"
	"time"
)

func GetCurrentDate() string {
	currentTime := time.Now()
	date := fmt.Sprintf(currentTime.Format("2006-01-02 15:04:05"))
	return date
}
