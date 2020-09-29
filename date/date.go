package date

import (
	"time"
)

func Today() time.Time {
	y, m, d := time.Now().Date()
	loc, _ := time.LoadLocation("Local")
	return time.Date(y, m, d, 0, 0, 0, 0, loc)
}

func TodayEnd() time.Time {
	return Today().Add((86400 - 1) * time.Second)
}

func Tomorrow() time.Time {
	return Today().Add(1 * time.Hour * 24)
}
