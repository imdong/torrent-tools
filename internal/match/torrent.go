package match

import (
	"log"
	"os"
	"path/filepath"
)

// 扫描目录下的 Torrent 文件列表
func scanTorrentFile(torrentPath string) (torrentFiles []string) {
	tFileInfo, err := os.Stat(torrentPath)
	if err != nil {
		log.Fatal(err)
	}

	if !tFileInfo.IsDir() {
		torrentFiles = append(torrentFiles, torrentPath)
		return
	}

	err = filepath.Walk(torrentPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(info.Name())
		if ext == ".torrent" {
			torrentFiles = append(torrentFiles, path)
		}

		return nil
	})

	return
}
