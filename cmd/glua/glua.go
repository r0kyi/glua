package main

import (
	"os"

	"github.com/r0kyi/glua"
)

func main() {
	L := glua.NewState()
	defer L.Close()

	args := os.Args
	if len(args) < 2 {
		println("usage: glua <lua file>")
		return
	}

	if err := L.DoFile(args[1]); err != nil {
		println(err.Error())
	}
}
