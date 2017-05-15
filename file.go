package main

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

const TAG_ROOT_DIR string = `__ROOT_DIR__`

type File struct {
	origin    *os.File
	fileInfo  os.FileInfo
	localPath string
}

func NewFile(filename string) *File {
	origin, err := os.Open(filename)
	if err != nil {
		LogError("file[%s] error[%v]", filename, err)
	}

	file := &File{
		origin:    origin,
		localPath: filename,
	}

	file.fileInfo, _ = origin.Stat()

	return file
}

//
func (f *File) LocalPath() string {
	return f.localPath
}

func (f *File) CloudPath(cloudDir string) string {
	if !f.IsAbsolutePath() {
		return path.Join(cloudDir, f.LocalPath())
	}

	return path.Join(cloudDir, TAG_ROOT_DIR, strings.TrimLeft(f.localPath, `/`))
}

func (f *File) IsAbsolutePath() bool {
	return f.localPath[0] == '/'
}

func (f *File) IsDir() bool {
	return f.fileInfo.IsDir()
}

func (f *File) BackupIn(cloudDir string) (err error) {
	src := f.LocalPath()
	dst := f.CloudPath(cloudDir)

	//	@todo	根据实情建权限
	os.MkdirAll(filepath.Dir(dst), 0755)

	err = removeFile(dst)
	err = moveFile(src, dst)
	LogWarning(`%v`, err)
	err = linkFile(dst, src)

	return
}

func (f *File) RestoreFrom(cloudDir string) (err error) {
	return
}

func (f *File) UninstallFrom(cloudDir string) (err error) {
	local := f.LocalPath()
	cloud := f.CloudPath(cloudDir)

	err = removeFile(local)
	err = moveFile(cloud, local)

	return
}

func removeFile(src string) error {
	src = absolutePath(src)
	LogWarning(`remove file from "%s"`, src)
	err := os.Remove(src)
	return err
}

func moveFile(src, dst string) error {
	src = absolutePath(src)
	dst = absolutePath(dst)
	LogWarning(`move file from "%s" to "%s"`, src, dst)
	return os.Rename(src, dst)
}

func linkFile(src, link string) (err error) {
	src = absolutePath(src)
	link = absolutePath(link)
	LogWarning(`link "%s" to "%s"`, src, link)

	err = os.Link(src, link)

	return nil
}

func absolutePath(filename string) string {
	abs, err := filepath.Abs(filename)
	if err != nil {
		LogError(``, err)
	}

	return abs
}
