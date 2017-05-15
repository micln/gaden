package main

import (
	"log"

	"github.com/micln/go-utils"
)

type Command struct {
	Action    string
	Overwrite bool
	File      *File
	Conf      *Config

	ff *Flag
}

const ACTION_BACKUP = `backup`
const ACTION_REVERT = `revert`
const ACTION_UNINSTALL = `uninstall`

func NewCommandFromFlag(ff *Flag) *Command {

	cmd := &Command{}

	if ff.IsForce {
		cmd.Overwrite = true
	}

	if ff.IsConfigFile {
		LogError(`config file is not support yet !`)
	} else {
		cmd.File = NewFile(ff.LocalFile)
	}

	switch {
	case ff.IsRestore:
		cmd.Action = ACTION_REVERT
	case ff.IsUninstall:
		cmd.Action = ACTION_UNINSTALL
	default:
		cmd.Action = ACTION_BACKUP
	}

	cmd.ff = ff

	return cmd
}

func ListActions() []string {
	return []string{
		ACTION_BACKUP,
		ACTION_REVERT,
		ACTION_UNINSTALL,
	}
}

func (cmd *Command) Run() {
	handle := cmd.getActionFunc()

	log.Println(cmd.File.LocalPath())
	log.Println(go_utils.JsonEncode(cmd))

	handle(cmd.File)
}

func (cmd *Command) getActionFunc() func(*File) {

	if cmd.Action == ACTION_BACKUP {
		return func(f *File) {
			f.BackupIn(cmd.ff.CloudDir)
		}
	}

	if cmd.Action == ACTION_UNINSTALL {
		return func(f *File) {
			f.UninstallFrom(cmd.ff.CloudDir)
		}
	}

	panic(`unknown action`)

}

func (cmd *Command) Each(files []*File, do func(*File)) {
	for i := range files {
		do(files[i])
	}
}
