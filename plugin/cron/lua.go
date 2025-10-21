package cron

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/r0kyi/glua/core"
	lua "github.com/yuin/gopher-lua"
)

func (c *Cron) String() string {
	return "cron"
}

func (c *Cron) AssertFunction() lua.LGFunction {
	return NewCronL
}

func (c *Cron) MetatableName() string {
	return "lua.table.cron"
}

func (c *Cron) jobL(L *lua.LState) int {
	cronExpression := L.CheckString(1)
	fn := L.CheckFunction(2)
	c.cronExpression = cronExpression
	c.fn = toJobFun(L, fn)

	err := c.job()
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
	c.interval = int(interval)
	c.fn = toJobFun(L, fn)

	err := c.seconds()
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
	c.interval = int(interval)
	c.fn = toJobFun(L, fn)

	err := c.minutes()
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
	c.interval = int(interval)
	c.fn = toJobFun(L, fn)

	err := c.hours()
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
	c.interval = int(interval)
	c.fn = toJobFun(L, fn)

	err := c.days()
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
	c.interval = int(interval)
	c.fn = toJobFun(L, fn)

	err := c.weeks()
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

func NewCronL(L *lua.LState) int {
	location := time.Local
	var err error
	if L.GetTop() == 2 {
		loc := L.CheckString(2)
		location, err = time.LoadLocation(loc)
		if err != nil {
			location = time.Local
		}
	}

	c := &Cron{
		scheduler: gocron.NewScheduler(location),
	}
	ud := core.NewUserData(L, c)
	L.Push(ud)

	return 1
}

func Preload(L *lua.LState) lua.LValue {
	c := &Cron{}
	ud := core.NewUserData(L, c)

	return ud
}
