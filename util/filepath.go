package util

import (
	"os"
	"path/filepath"
	"regexp"
)

// 路径是否存在
func ExistPath(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// 文件是否存在
func ExistFile(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	}

	return !f.IsDir()
}

// 文件夹是否存在
func ExistDir(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	}

	return f.IsDir()
}

// GetFileListByExt 获取某个目录下ext扩展名的所有文件
func GetFileListByExt(dir string, ext string) ([]string, error) {
	var fileLists []string
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ext {
			fileLists = append(fileLists, path)
		}
		return nil
	})
	return fileLists, err
}

func ClearDirAndCreateNew(path string) error {
	if ExistPath(path) {
		err := os.RemoveAll(path)
		if err != nil {
			return err
		}
	}
	err := os.MkdirAll(path, os.ModePerm)
	return err
}

func InitDirAndClearFile(path string, removePattern string) error {
	if !ExistPath(path) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	err := filepath.Walk(path, func(fileName string, f os.FileInfo, err error) error {
		if ok, _ := regexp.MatchString(removePattern, fileName); !ok {
			return nil
		}
		err = os.Remove(fileName)
		return err
	})
	return err
}
