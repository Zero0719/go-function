package convert

import "strconv"

func StringToInt(param string) int {
	num, err := strconv.Atoi(param)
	if err != nil {
		return 0
	}
	return num
}

func StringToInt64(param string) int64 {
	return int64(StringToInt(param))
}

func IntToString(number int) string {
	return strconv.Itoa(number)
}

func Int64ToString(number int64) string {
	return IntToString(int(number))
}
