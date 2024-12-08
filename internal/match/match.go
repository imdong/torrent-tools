package match

import (
	"fmt"
	torrentMap "github.com/imdong/torrent-tools/pkg/torrent-map"
	"github.com/mgutz/ansi"
	"log"
	"path/filepath"
)

// Option 选项
type Option struct {
	OnlyMatch         bool
	IgnoreTorrentName bool
}

// 本地文件大小索引
var sSizeMap map[int64][]string

// 目标文件目录
var targetDir string

// 额外选项
var matchOption = Option{
	OnlyMatch:         false,
	IgnoreTorrentName: false,
}

// Start 文件匹配入口
func Start(torrentPath, storageDir, symlinkTargetDir string, option Option) {
	targetDir = symlinkTargetDir
	matchOption = option

	torrentFiles := scanTorrentFile(torrentPath)

	tMap := torrentMap.New(torrentFiles)
	sSizeMap = scanFileSizeMap(storageDir)

	// 逐个匹配
	tMap.MapFile(matchFile)
}

// 匹配文件是否有相同的
func matchFile(tFile *torrentMap.TorrentFileInfo) {
	// 目标文件是否已经存在
	var tFileFullPath string
	if matchOption.IgnoreTorrentName {
		tFileFullPath = filepath.Join(targetDir, tFile.FilePath)
	} else {
		tFileFullPath = filepath.Join(targetDir, tFile.TorrentName, tFile.FilePath)
	}

	tFilePath := filepath.Join(tFile.TorrentName, tFile.FilePath)
	if fileExists(tFileFullPath) {
		failMsg(tFilePath, "文件已存在")
		return
	}

	// 是否有与种子的文件大小相同的文件
	sFileSizeList, exists := sSizeMap[tFile.FileSize]
	if exists == false {
		failMsg(tFilePath, "没有相同大小的文件")
		return
	}

	// 文件是否存在
	for _, sFile := range sFileSizeList {
		if tFile.ContrastFileHash(sFile) == false {
			continue
		}

		// 文件相同
		successMag(sFile, tFileFullPath)

		// 创建文件软连接
		_ = createSymlink(sFile, tFileFullPath)

		return
	}

	// 没有相同的文件
	failMsg(tFilePath, "没有匹配到文件")

	return
}

var failColor = ansi.ColorFunc("8")
var successColor = ansi.ColorFunc("2")

// 输出未匹配成功的消息
func failMsg(file, msg string) {
	if matchOption.OnlyMatch {
		return
	}

	log.Printf(failColor("跳过, %s: %s\n"), msg, file)
}

// 匹配成功的输出
func successMag(sFile, tFile string) {
	fmt.Printf(successColor("文件相同: \n\tstorage File: %s\n\ttorrent File: %s\n"), sFile, tFile)
}
