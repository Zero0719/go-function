package date

import (
	"time"
)

func Date(format string, timestamps ...int64) string {
	var timestamp int64
	if len(timestamps) == 0 {
		timestamp = time.Now().Unix()
	} else {
		timestamp = timestamps[0]
	}

	return time.Unix(timestamp, 0).Format(format)
}
