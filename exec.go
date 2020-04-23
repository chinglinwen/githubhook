package main

import (
	"github.com/michaelmass/goshell"
)

var shell = goshell.New()

func exec(commands string) {
	shell.Dir = "."
	shell.Exec("bash", commands)
}
