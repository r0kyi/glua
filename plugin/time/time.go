package time

import (
	"time"
)

type Time struct {
	time *time.Time

	year  int
	month int
	day   int
	hour  int
	min   int
	sec   int
	nsec  int
	loc   string

	layout    string
	formatted string
}

func (t *Time) getYear() {
	t.year = t.time.Year()
}

func (t *Time) getMonth() {
	t.month = int(t.time.Month())
}

func (t *Time) getDay() {
	t.day = t.time.Day()
}

func (t *Time) getHour() {
	t.hour = t.time.Hour()
}

func (t *Time) getMinute() {
	t.min = t.time.Minute()
}

func (t *Time) getSecond() {
	t.sec = t.time.Second()
}

func (t *Time) getNanosecond() {
	t.nsec = t.time.Nanosecond()
}

func (t *Time) format() {
	t.formatted = t.time.Format(t.layout)
}
