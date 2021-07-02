package date

import (
	"time"
)

// @title Date
// @descripion 输出指定格式字符串时间，默认当前时间
// @param format string "格式化 2006/01/02 15:03:04"
// @param timestamps int64 "需要转化的时间戳"
// @return string "格式化后的时间字符串"
func Date(format string, timestamps ...int64) string {
	var timestamp int64
	if len(timestamps) == 0 {
		timestamp = time.Now().Unix()
	} else {
		timestamp = timestamps[0]
	}

	return time.Unix(timestamp, 0).Format(format)
}
