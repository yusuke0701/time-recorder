package time

import "time"

var jst = time.FixedZone("Asia/Tokyo", 9*60*60)

var DefaultFormat = "2006-01-2"

func NowInJST() time.Time {
	return time.Now().In(jst)
}

func InJST(date time.Time) time.Time {
	return date.In(jst)
}
