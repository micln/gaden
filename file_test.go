package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"runtime"
	"testing"
	"time"
)

func TestRemoveFile(t *testing.T) {
	a := randomFile()
	assertTrue(isExist(a))

	removeFile(a)
	assertTrue(isNotExist(a))
}

func TestMoveFile(t *testing.T) {
	a := randomFile()
	b := randomFile()
	removeFile(b)

	assertTrue(isExist(a))
	assertTrue(isNotExist(b))

	moveFile(a, b)

	assertTrue(isNotExist(a))
	assertTrue(isExist(b))

	removeFile(b)
}

func TestLinkFile(t *testing.T) {
	a := randomFile()
	b := randomFile()
	removeFile(b)

	assertTrue(isExist(a))
	assertTrue(isNotExist(b))

	linkFile(a, b)

	assertTrue(isExist(a))
	assertTrue(isExist(b))

	removeFile(a)
	removeFile(b)
}

func randomFile() string {
	filename := fmt.Sprintf("t%v_r%v", time.Now().UnixNano(), rand.Int63())
	err := ioutil.WriteFile(filename, []byte(filename), 0755)
	//err := os.Mkdir(filename, 0755)
	if err != nil {
		panic(err)
	}
	return filename
}

func isExist(filename string) bool {
	return !isNotExist(filename)
}

func isNotExist(filename string) bool {
	_, err := os.Stat(filename)
	return err != nil && os.IsNotExist(err)
}

func assertTrue(b bool) {
	if !b {
		failure()
	}
}

func failure() {
	_, file, line, _ := runtime.Caller(2)
	panic(fmt.Sprintf("failure in %s:%v", file, line))
}
