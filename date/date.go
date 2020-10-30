package date

import (
	"time"
)

var (
	TimeZone, _ = time.LoadLocation("Local")
)

func Today() time.Time {
	y, m, d := time.Now().Date()
	return time.Date(y, m, d, 0, 0, 0, 0, TimeZone)
}

func TodayEnd() time.Time {
	return Today().Add((86400 - 1) * time.Second)
}

func Tomorrow() time.Time {
	return Today().Add(time.Hour * 24)
}

func Yesterday() time.Time {
	return Today().Add(-(time.Hour * 24))
}
