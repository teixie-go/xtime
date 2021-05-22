package xtime

import (
	"fmt"
	"regexp"
	"time"
)

const (
	Fmts      = "05"
	Fmti      = "04"
	FmtH      = "15"
	Fmtd      = "02"
	Fmtm      = "01"
	FmtY      = "2006"
	FmtHi     = "15:04"
	Fmtmd     = "01-02"
	FmtHis    = "15:04:05"
	FmtYmd    = "2006-01-02"
	FmtYmdHi  = "2006-01-02 15:04"
	FmtYmdHis = "2006-01-02 15:04:05"
)

var (
	location *time.Location
)

// set location
func SetLocation(loc *time.Location) {
	location = loc
}

// get location
func GetLocation() *time.Location {
	if location != nil {
		return location
	}
	return time.Local
}

// now
func Now() time.Time {
	return time.Now().In(GetLocation())
}

// start of today
func Today() time.Time {
	return StartOfDay(Now())
}

// end of today
func EndOfToday() time.Time {
	return EndOfDay(Now())
}

// start of tomorrow
func Tomorrow() time.Time {
	return Today().Add(24 * time.Hour)
}

// start of yesterday
func Yesterday() time.Time {
	return Today().Add(-24 * time.Hour)
}

// start of day
func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, GetLocation())
}

// end of day
func EndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, GetLocation())
}

// unix zero
func UnixZero() time.Time {
	return time.Date(1970, 1, 1, 0, 0, 0, 0, GetLocation())
}

// parse
func Parse(t interface{}) (time.Time, error) {
	if t == nil {
		return time.Time{}, nil
	}

	switch t.(type) {
	case time.Time:
		return t.(time.Time), nil
	case string:
		return time.Parse(FmtYmdHis, t.(string))
	case int:
		return time.Unix(int64(t.(int)), 0), nil
	case int64:
		return time.Unix(t.(int64), 0), nil
	}

	return time.Time{}, fmt.Errorf("supported input type:time.Time,string,int,int64")
}

// format, e.g. YYYY-mm-dd HH:ii:ss
func Format(t time.Time, fmtStr string) string {
	exists, err := regexp.Match("[YymdHis]+", []byte(fmtStr))
	if err == nil && !exists {
		return t.Format(fmtStr)
	}

	timeStr := t.String()
	o := map[string]string{
		"Y+": timeStr[0:4],
		"y+": timeStr[2:4],
		"m+": timeStr[5:7],
		"d+": timeStr[8:10],
		"H+": timeStr[11:13],
		"i+": timeStr[14:16],
		"s+": timeStr[17:19],
	}
	for k, v := range o {
		re, _ := regexp.Compile(k)
		fmtStr = re.ReplaceAllString(fmtStr, v)
	}
	return fmtStr
}
