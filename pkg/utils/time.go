package utils

import (
	"strconv"
	"time"
)

func GetNow() time.Time {
	return time.Now()
}

func GetNowUnix() int64 {
	return time.Now().UnixMicro()
}

func ConvertStrMicroUnix2Ts(s string) time.Time {
	nunix, _ := strconv.Atoi(s)
	return time.Unix(int64(nunix/1000000000), int64(nunix%1000000000))
}

// fmt: 2025-01-13T03:52:24.763500+00:00
func ConvertTs2StrMicroUnix2(s string) (int64, error) {
	nano, err := time.Parse(time.RFC3339Nano, s)
	if err != nil {
		return 0, err
	}
	return nano.UnixMicro() * 1000, nil // 桁数を19桁にするため(元々19桁だった場合は...)
}
