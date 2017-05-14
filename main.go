package main

import (
	"log"
	"os"
	"os/exec"
)

var VERSION = `0.1`

var USER = os.Getenv(`USER`)
var HOME = os.Getenv(`HOME`)

func init() {
	os.Chdir(HOME)

	cmd := exec.Command(`mv`, `a b/c d`, `e b`)
	log.Fatalln(cmd)
}

func main() {
	ff := GetFlag()

	cmd := NewCommandFromFlag(ff)
	cmd.Run()
}
