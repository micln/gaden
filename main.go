package main

import (
	"os"
)

var VERSION = `0.1`

var USER = os.Getenv(`USER`)
var HOME = os.Getenv(`HOME`)

func init() {
	os.Chdir(HOME)
}

func main() {
	ff := GetFlag()

	cmd := NewCommandFromFlag(ff)
	cmd.Run()
}
