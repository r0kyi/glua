package time

import (
	"time"
)

type Time struct {
	time *time.Time
}

func (t *Time) getYear() int {
	return t.time.Year()
}

func (t *Time) getMonth() int {
	return int(t.time.Month())
}

func (t *Time) getDay() int {
	return t.time.Day()
}

func (t *Time) getHour() int {
	return t.time.Hour()
}

func (t *Time) getMinute() int {
	return t.time.Minute()
}

func (t *Time) getSecond() int {
	return t.time.Second()
}

func (t *Time) getNanosecond() int {
	return t.time.Nanosecond()
}

func (t *Time) format(layout string) string {
	return t.time.Format(layout)
}
