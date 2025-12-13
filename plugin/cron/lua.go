package cron

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/r0kyi/glua/core"
	lua "github.com/r0kyi/gopher-lua"
)

func (c *Cron) String() string {
	return fmt.Sprintf("glua.cron: %p", c)
}

func (c *Cron) Type() lua.LValueType {
	return lua.LTObject
}

func (c *Cron) AssertFunction() (*lua.LFunction, bool) {
	return core.NewFunction(newCronL), true
}

func (c *Cron) jobL(L *lua.LState) int {
	cronExpression := L.CheckString(1)
	fn := L.CheckFunction(2)

	err := c.job(cronExpression, toJobFun(L, fn))
	if err != nil {
		L.Push(lua.LString(err.Error()))
	} else {
		L.Push(lua.LNil)
	}

	return 1
}

func (c *Cron) secondsL(L *lua.LState) int {
	interval := L.CheckNumber(1)
	fn := L.CheckFunction(2)

	err := c.seconds(int(interval), toJobFun(L, fn))
	if err != nil {
		L.Push(lua.LString(err.Error()))
	} else {
		L.Push(lua.LNil)
	}

	return 1
}

func (c *Cron) minutesL(L *lua.LState) int {
	interval := L.CheckNumber(1)
	fn := L.CheckFunction(2)

	err := c.minutes(int(interval), toJobFun(L, fn))
	if err != nil {
		L.Push(lua.LString(err.Error()))
	} else {
		L.Push(lua.LNil)
	}

	return 1
}

func (c *Cron) hoursL(L *lua.LState) int {
	interval := L.CheckNumber(1)
	fn := L.CheckFunction(2)

	err := c.hours(int(interval), toJobFun(L, fn))
	if err != nil {
		L.Push(lua.LString(err.Error()))
	} else {
		L.Push(lua.LNil)
	}

	return 1
}

func (c *Cron) daysL(L *lua.LState) int {
	interval := L.CheckNumber(1)
	fn := L.CheckFunction(2)

	err := c.days(int(interval), toJobFun(L, fn))
	if err != nil {
		L.Push(lua.LString(err.Error()))
	} else {
		L.Push(lua.LNil)
	}

	return 1
}

func (c *Cron) weeksL(L *lua.LState) int {
	interval := L.CheckNumber(1)
	fn := L.CheckFunction(2)

	err := c.weeks(int(interval), toJobFun(L, fn))
	if err != nil {
		L.Push(lua.LString(err.Error()))
	} else {
		L.Push(lua.LNil)
	}

	return 1
}

func (c *Cron) startBlockL(L *lua.LState) int {
	c.startBlock()

	return 0
}

func (c *Cron) startAsyncL(L *lua.LState) int {
	c.startAsync()

	return 0
}

func (c *Cron) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "job":
		return L.NewFunction(c.jobL)
	case "seconds":
		return L.NewFunction(c.secondsL)
	case "minutes":
		return L.NewFunction(c.minutesL)
	case "hours":
		return L.NewFunction(c.hoursL)
	case "days":
		return L.NewFunction(c.daysL)
	case "weeks":
		return L.NewFunction(c.weeksL)
	case "start_block":
		return L.NewFunction(c.startBlockL)
	case "start_async":
		return L.NewFunction(c.startAsyncL)
	default:
		return lua.LNil
	}
}

func newCronL(L *lua.LState) int {
	location := time.Local
	var err error
	if L.GetTop() == 1 {
		loc := L.CheckString(1)
		location, err = time.LoadLocation(loc)
		if err != nil {
			location = time.Local
		}
	}

	c := &Cron{
		scheduler: gocron.NewScheduler(location),
	}
	L.Push(c)

	return 1
}

func Preload() lua.LValue {
	c := &Cron{}

	return c
}
