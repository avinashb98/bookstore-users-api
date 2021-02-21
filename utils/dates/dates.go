package dates

import (
	"time"
)

func GetNowString() string {
	now := time.Now().UTC().String()
	return now
}
