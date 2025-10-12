package time

import (
	"time"

	"github.com/r0kyi/glua/core"
	lua "github.com/yuin/gopher-lua"
)

func (t *Time) String() string {
	return "time"
}

func (t *Time) AssertFunction() lua.LGFunction {
	return nil
}

func (t *Time) MetatableName() string {
	return "lua.table.time"
}

func (t *Time) yearL(L *lua.LState) int {
	t.getYear()
	L.Push(lua.LNumber(t.year))

	return 1
}

func (t *Time) monthL(L *lua.LState) int {
	t.getMonth()
	L.Push(lua.LNumber(t.month))

	return 1
}

func (t *Time) dayL(L *lua.LState) int {
	t.getDay()
	L.Push(lua.LNumber(t.day))

	return 1
}

func (t *Time) hourL(L *lua.LState) int {
	t.getHour()
	L.Push(lua.LNumber(t.hour))

	return 1
}

func (t *Time) minuteL(L *lua.LState) int {
	t.getMinute()
	L.Push(lua.LNumber(t.min))

	return 1
}

func (t *Time) secondL(L *lua.LState) int {
	t.getSecond()
	L.Push(lua.LNumber(t.sec))

	return 1
}

func (t *Time) nanosecondL(L *lua.LState) int {
	t.getNanosecond()
	L.Push(lua.LNumber(t.nsec))

	return 1
}

func (t *Time) formatL(L *lua.LState) int {
	layout := L.CheckString(1)
	t.layout = layout
	t.format()

	L.Push(lua.LString(t.formatted))

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
	ud := core.NewUserData(L, t)

	L.Push(ud)

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
	ud := core.NewUserData(L, t)

	L.Push(ud)

	return 1
}

func unixL(L *lua.LState) int {
	sec := L.CheckNumber(1)
	nsec := L.CheckNumber(2)

	newT := time.Unix(int64(sec), int64(nsec))
	t := &Time{
		time: &newT,
	}
	ud := core.NewUserData(L, t)

	L.Push(ud)

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
	ud := core.NewUserData(L, t)

	L.Push(ud)

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
	ud := core.NewUserData(L, t)

	L.Push(ud)

	return 1

}

func Preload(L *lua.LState) lua.LValue {
	t := &Time{}
	ud := core.NewUserData(L, t)

	return ud
}
