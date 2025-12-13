package time

import (
	"fmt"
	"time"

	lua "github.com/r0kyi/gopher-lua"
)

func (t *Time) String() string {
	return fmt.Sprintf("glua.time: %p", t)
}

func (t *Time) Type() lua.LValueType {
	return lua.LTObject
}

func (t *Time) AssertFunction() (*lua.LFunction, bool) {
	return nil, false
}

func (t *Time) yearL(L *lua.LState) int {
	L.Push(lua.LNumber(t.getYear()))

	return 1
}

func (t *Time) monthL(L *lua.LState) int {
	L.Push(lua.LNumber(t.getMonth()))

	return 1
}

func (t *Time) dayL(L *lua.LState) int {
	L.Push(lua.LNumber(t.getDay()))

	return 1
}

func (t *Time) hourL(L *lua.LState) int {
	L.Push(lua.LNumber(t.getHour()))

	return 1
}

func (t *Time) minuteL(L *lua.LState) int {
	L.Push(lua.LNumber(t.getMinute()))

	return 1
}

func (t *Time) secondL(L *lua.LState) int {
	L.Push(lua.LNumber(t.getSecond()))

	return 1
}

func (t *Time) nanosecondL(L *lua.LState) int {
	L.Push(lua.LNumber(t.getNanosecond()))

	return 1
}

func (t *Time) formatL(L *lua.LState) int {
	layout := L.CheckString(1)

	L.Push(lua.LString(t.format(layout)))

	return 1
}

func (t *Time) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "now":
		return L.NewFunction(nowL)
	case "date":
		return L.NewFunction(dateL)
	case "unix":
		return L.NewFunction(unixL)
	case "parse":
		return L.NewFunction(parseL)
	case "parse_in_location":
		return L.NewFunction(parseInLocationL)
	case "year":
		return L.NewFunction(t.yearL)
	case "month":
		return L.NewFunction(t.monthL)
	case "day":
		return L.NewFunction(t.dayL)
	case "hour":
		return L.NewFunction(t.hourL)
	case "minute":
		return L.NewFunction(t.minuteL)
	case "second":
		return L.NewFunction(t.secondL)
	case "nanosecond":
		return L.NewFunction(t.nanosecondL)
	case "format":
		return L.NewFunction(t.formatL)
	default:
		return lua.LNil
	}
}

func nowL(L *lua.LState) int {
	newT := time.Now()
	t := &Time{
		time: &newT,
	}

	L.Push(t)

	return 1
}

func dateL(L *lua.LState) int {
	year := L.CheckNumber(1)
	month := L.CheckNumber(2)
	day := L.CheckNumber(3)
	hour := L.CheckNumber(4)
	min_ := L.CheckNumber(5)
	sec := L.CheckNumber(6)
	nsec := L.CheckNumber(7)
	loc := L.CheckString(8)
	lo, err := time.LoadLocation(loc)

	if err != nil {
		lo = time.Local
	}

	newT := time.Date(int(year), time.Month(month), int(day), int(hour), int(min_), int(sec), int(nsec), lo)
	t := &Time{
		time: &newT,
	}

	L.Push(t)

	return 1
}

func unixL(L *lua.LState) int {
	sec := L.CheckNumber(1)
	nsec := L.CheckNumber(2)

	newT := time.Unix(int64(sec), int64(nsec))
	t := &Time{
		time: &newT,
	}

	L.Push(t)

	return 1
}

func parseL(L *lua.LState) int {
	layout := L.CheckString(1)
	value := L.CheckString(2)

	newT, err := time.Parse(layout, value)
	if err != nil {
		L.Push(lua.LNil)
		return 1
	}

	t := &Time{
		time: &newT,
	}

	L.Push(t)

	return 1
}

func parseInLocationL(L *lua.LState) int {
	layout := L.CheckString(1)
	value := L.CheckString(2)
	loc := L.CheckString(3)

	lo, err := time.LoadLocation(loc)
	if err != nil {
		lo = time.Local
	}

	newT, err := time.ParseInLocation(layout, value, lo)
	if err != nil {
		L.Push(lua.LNil)
		return 1
	}

	t := &Time{
		time: &newT,
	}

	L.Push(t)

	return 1

}

func Preload() lua.LValue {
	t := &Time{}

	return t
}
