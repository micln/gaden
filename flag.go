package main

import (
	"flag"
	"fmt"
	"os"
)

var ff *Flag

type Flag struct {
	IsForce      bool
	IsConfigFile bool
	IsRestore    bool
	IsUninstall  bool
	IsVersion    bool

	LocalFile string
	CloudDir  string
}

func init() {
	force := flag.Bool(`f`, false, `force overwrite`)
	cfg := flag.Bool(`c`, false, `input file is config`)
	restore := flag.Bool(`restore`, false, `do the restore command`)
	uninstall := flag.Bool(`uninstall`, false, `do the uninstall command`)
	showVersion := flag.Bool(`V`, false, `show the version`)
	target := flag.String(`target`, ``, `backup place`)

	flag.Parse()

	ff = &Flag{
		*force,
		*cfg,
		*restore,
		*uninstall,
		*showVersion,
		flag.Arg(0),
		*target,
	}

	if ff.IsVersion {
		fmt.Println(`gaden version is`, VERSION)
		ff.Exit()
	}

	if len(ff.LocalFile) == 0 {
		LogWarning(`empty filename`)
		ff.Print()
		ff.Exit()
	}

	if len(ff.CloudDir) == 0 {
		LogError(`empty target path`)
	}

	tar := NewFile(*target)
	if !tar.IsAbsolutePath() {
		LogError(`backup place must be a absolution path`)
	}

	if !tar.IsDir() {
		LogError(`backup path must be a folder`)
	}
}

func GetFlag() *Flag {
	return ff
}

func (f *Flag) Print() {
	flag.Usage()
}

func (f *Flag) Exit() {
	os.Exit(0)
}
