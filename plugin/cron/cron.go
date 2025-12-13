package cron

import (
	"github.com/go-co-op/gocron"
	lua "github.com/r0kyi/gopher-lua"
)

type Cron struct {
	scheduler *gocron.Scheduler
}

func (c *Cron) job(cronExpression string, fn func()) error {
	_, err := c.scheduler.Cron(cronExpression).Do(fn)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cron) seconds(interval int, fn func()) error {
	_, err := c.scheduler.Every(interval).Seconds().Do(fn)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cron) minutes(interval int, fn func()) error {
	_, err := c.scheduler.Every(interval).Minutes().Do(fn)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cron) hours(interval int, fn func()) error {
	_, err := c.scheduler.Every(interval).Hours().Do(fn)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cron) days(interval int, fn func()) error {
	_, err := c.scheduler.Every(interval).Days().Do(fn)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cron) weeks(interval int, fn func()) error {
	_, err := c.scheduler.Every(interval).Weeks().Do(fn)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cron) startBlock() {
	c.scheduler.StartBlocking()
}

func (c *Cron) startAsync() {
	c.scheduler.StartAsync()
}

func toJobFun(L *lua.LState, fn *lua.LFunction) func() {
	return func() {
		_ = L.CallByParam(lua.P{
			Fn:      fn,
			NRet:    0,
			Protect: true,
		})
	}
}
