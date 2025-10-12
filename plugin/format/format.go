package format

import (
	"fmt"
)

type Format struct {
	format string
	args   []any
}

func (f *Format) toString() string {
	return fmt.Sprintf(f.format, f.args...)
}
