package re

import (
	"regexp"

	"github.com/r0kyi/glua/core"
)

type Re struct {
	regexp *regexp.Regexp
}

func (r *Re) matchString(src string) bool {
	return r.regexp.Match(core.S2B(src))
}

func (r *Re) findString(src string) string {
	return r.regexp.FindString(src)
}

func (r *Re) findAllString(src string) []string {
	return r.regexp.FindAllString(src, -1)
}

func (r *Re) replaceAllString(src string, repl string) string {
	return r.regexp.ReplaceAllString(src, repl)
}

func (r *Re) split(src string) []string {
	return r.regexp.Split(src, -1)
}
