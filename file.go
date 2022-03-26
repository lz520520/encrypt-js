package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

func WriteFile(filename string, writeBytes []byte) (err error) {
	MkdirFromFile(filename)
	f, err := os.Create(filename)
	if err != nil {
		return
	}
	defer f.Close()
	f.Write(writeBytes)
	return nil
}
func MkdirFromFile(src string) error {
	dstDir := filepath.Dir(src)
	_, err := os.Stat(dstDir)
	if err != nil {
		err = os.MkdirAll(dstDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
func ReadFile(filename string) (resultSlice []string, err error) {
	resultSlice = make([]string, 0)
	resultBytes, err := ReadFileBytes(filename)
	if err != nil {
		return
	}
	resultSlice = strings.Split(string(resultBytes), "\n")
	for i, c := range resultSlice {
		resultSlice[i] = strings.TrimSpace(c)
	}
	return resultSlice, nil

}
func ReadFileBytes(filename string) (resultBytes []byte, err error) {
	resultBytes = make([]byte, 0)
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			return resultBytes, err
		}
		if n == 0 {
			break
		}

		resultBytes = append(resultBytes, buf[:n]...)
	}
	return resultBytes, nil
}

func GetBaseName(name string) string {
	filenameWithSuffix := filepath.Base(name)
	//fileSuffix := filepath.Ext(filenameWithSuffix)
	return filenameWithSuffix
}
