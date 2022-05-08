package timeex

import "time"

var timestamp time.Duration = 0

func Duration() time.Duration {
	tmp := timestamp
	timestamp = time.Duration(time.Now().UnixNano())
	return tmp - timestamp
}

func DurationSec() time.Duration {
	tmp := timestamp
	timestamp = time.Duration(time.Now().UnixNano())
	return (tmp - timestamp) / time.Second
}
