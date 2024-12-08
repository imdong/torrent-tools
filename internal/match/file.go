package match

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// 扫描指定目录下的文件列表
func scanFileSizeMap(storageDir string) map[int64][]string {
	sizeMap := make(map[int64][]string)

	err := filepath.Walk(storageDir, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 跳过目录
		if info.IsDir() {
			return nil
		}

		sizeMap[info.Size()] = append(sizeMap[info.Size()], filePath)

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return sizeMap
}

// 判断文件是否存在
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// 创建软连接
func createSymlink(target, link string) error {
	linkDir := filepath.Dir(link)
	if !fileExists(linkDir) {
		err := os.MkdirAll(linkDir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory for symlink: %w", err)
		}
	}

	err := os.Symlink(target, link)
	if err != nil {
		return fmt.Errorf("failed to create symlink from %s to %s: %w", target, link, err)
	}
	return nil
}
