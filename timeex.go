package gotools

import "time"

var timestamp time.Duration = 0

func SetTimestamp() {
	timestamp = time.Duration(time.Now().UnixNano())
}

func GetDuration() time.Duration {
	return time.Duration(time.Now().UnixNano()) - timestamp
}

func GetDurationSec() time.Duration {
	return (time.Duration(time.Now().UnixNano()) - timestamp) / time.Second
}
