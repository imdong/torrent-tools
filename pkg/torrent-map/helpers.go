package torrent_map

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/anacrolix/torrent/metainfo"
	"io"
	"log"
	"os"
)

// 从 torrentFile 获取种子详细内容
func makeTorrentInfo(torrentFile string) *metainfo.Info {
	mi, err := metainfo.LoadFromFile(torrentFile)
	if err != nil {
		log.Fatalf("Failed to load torrent file: %v", err)
	}

	// 获取 Info 部分
	info, err := mi.UnmarshalInfo()
	if err != nil {
		log.Fatalf("Failed to unmarshal torrent info: %v", err)
	}

	return &info
}

// HashFileSection 获取文件该位置的hash
func HashFileSection(filePath string, offset int64, length int64) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	// Seek to the desired offset within the file.
	_, err = file.Seek(offset, io.SeekStart)
	if err != nil {
		return "", fmt.Errorf("failed to seek to position in file: %w", err)
	}

	// 读取文件块
	buffer := make([]byte, length)
	_, err = file.Read(buffer)
	if err != nil {
		return "", fmt.Errorf("failed to read from file: %w", err)
	}

	// 计算该文件的 sha256
	harsher := sha1.New()
	_, err = harsher.Write(buffer)
	if err != nil {
		return "", fmt.Errorf("failed to calculate hash of file section: %w", err)
	}
	sha256Hash := harsher.Sum(nil)
	sha256Hex := hex.EncodeToString(sha256Hash)
	//fmt.Printf("sha256Hex: %s\n", sha256Hex)

	return sha256Hex, nil
}
