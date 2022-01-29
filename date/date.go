package date

import (
	"strings"
	"time"
)

func Date(format string) (string, error) {
	if strings.TrimSpace(format) == "" {
		format = "2006-01-02 15:04:05 -0700 MST"
	}
	currentTime := time.Now()
	return currentTime.Format(format), nil
}
