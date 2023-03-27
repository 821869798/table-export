package util

import (
	"io"
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

func CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return err
	}

	return dstFile.Sync()
}

func CopyDir(srcDir, dstDir string) error {
	if !ExistPath(dstDir) {
		err := os.MkdirAll(dstDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return filepath.Walk(srcDir, func(srcPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Ignore the source directory.
		if srcPath == srcDir {
			return nil
		}

		baseFileName := filepath.Base(srcPath)
		// Calculate the destination path.
		dstPath := filepath.Join(dstDir, baseFileName)

		// Check if it's a directory, create it if so.
		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		// It's a file, so copy it.
		return CopyFile(srcPath, dstPath)
	})
}
