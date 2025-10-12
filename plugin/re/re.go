package re

import (
	"regexp"

	"github.com/r0kyi/glua/core"
)

type Re struct {
	regexp  *regexp.Regexp
	pattern string
	src     string
	repl    string
}

func (r *Re) compile() error {
	compiled, err := regexp.Compile(r.pattern)
	if err != nil {
		return err
	}
	r.regexp = compiled
	return nil
}

func (r *Re) matchString() bool {
	return r.regexp.Match(core.S2B(r.src))
}

func (r *Re) findString() string {
	return r.regexp.FindString(r.src)
}

func (r *Re) findAllString() []string {
	return r.regexp.FindAllString(r.src, -1)
}

func (r *Re) replaceAllString() string {
	return r.regexp.ReplaceAllString(r.src, r.repl)
}

func (r *Re) split() []string {
	return r.regexp.Split(r.src, -1)
}
