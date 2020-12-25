package util

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
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

// 获取某个目录下ext扩展名的所有文件
func GetFileListByExt(dir string, ext string) ([]string, error) {
	var fileLists []string
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			log.Fatalln(err)
		}
		if f.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".toml" {
			fileLists = append(fileLists, path)
		}
		return nil
	})
	return fileLists, err
}
