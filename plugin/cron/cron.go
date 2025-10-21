package cron

import (
	"github.com/go-co-op/gocron"
	lua "github.com/yuin/gopher-lua"
)

type Cron struct {
	cronExpression string
	fn             func()
	interval       int

	scheduler *gocron.Scheduler
}

func (c *Cron) job() error {
	_, err := c.scheduler.Cron(c.cronExpression).Do(c.fn)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cron) seconds() error {
	_, err := c.scheduler.Every(c.interval).Seconds().Do(c.fn)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cron) minutes() error {
	_, err := c.scheduler.Every(c.interval).Minutes().Do(c.fn)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cron) hours() error {
	_, err := c.scheduler.Every(c.interval).Hours().Do(c.fn)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cron) days() error {
	_, err := c.scheduler.Every(c.interval).Days().Do(c.fn)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cron) weeks() error {
	_, err := c.scheduler.Every(c.interval).Weeks().Do(c.fn)
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
