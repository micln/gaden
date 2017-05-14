package main

import (
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

type File struct {
	origin    *os.File
	fileInfo  os.FileInfo
	inputPath string
}

func NewFile(filename string) *File {

	origin, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}

	file := &File{
		origin:    origin,
		inputPath: filename,
	}

	file.fileInfo, _ = origin.Stat()

	return file
}

//
func (f *File) LocalPath() string {
	return f.inputPath
}

func (f *File) CloudPath(cloudDir string) string {
	if !f.IsAbsolutePath() {
		return path.Join(cloudDir, f.LocalPath())
	}

	return path.Join(cloudDir, strings.TrimLeft(f.inputPath, `/`))
}

func (f *File) IsAbsolutePath() bool {
	return f.inputPath[0] == '/'
}

func (f *File) IsDir() bool {
	return f.fileInfo.IsDir()
}

func (f *File) MoveTo(dst string) error {
	cmd := exec.Command(`mv`, f.inputPath, dst)
	log.Fatalln(cmd)
	return nil
}

func (f *File) CopyTo(dst string) error {
	return nil
}

func (f *File) LinkTo(link string) error {
	return nil
}

func (f *File) Remove() error {
	return nil
}

func (f *File) ForceRemove() error {
	return nil
}
