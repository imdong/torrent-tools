package torrent_map

// TorrentSizeMap 种子内各文件 Size 索引的 MapFile
type TorrentSizeMap struct {
	torrentFiles []string
	sizeMap      []TorrentFileInfo
}

// New 从种子文件列表中创建 TorrentSizeMap 实例
func New(torrentFiles []string) (t *TorrentSizeMap) {
	t = &TorrentSizeMap{
		torrentFiles: torrentFiles,
	}

	//for _, file := range torrentFiles {
	//	t.AddTorrent(file)
	//}

	return t
}

// AddTorrent 构建种子内的文件大小映射（单个）
func (t *TorrentSizeMap) AddTorrent(torrentFile string) {
	tInfo := makeTorrentInfo(torrentFile)

	// 要统计偏移量的
	var currentOffset int64 = 0
	for _, file := range tInfo.Files {
		t.sizeMap = append(t.sizeMap, TorrentFileInfo{
			Info:        tInfo,
			TorrentName: torrentFile,
			FilePath:    file.DisplayPath(tInfo),
			FileSize:    file.Length,
			FileOffset:  currentOffset,
		})

		currentOffset += file.Length
	}
}

// MapFile 遍历每个种子文件
func (t *TorrentSizeMap) MapFile(fn func(tFile *TorrentFileInfo)) {
	// 遍历每个种子
	for _, tFile := range t.torrentFiles {
		tInfo := makeTorrentInfo(tFile)
		var currentOffset int64 = 0

		// 遍历种子内的每个文件
		for _, file := range tInfo.Files {
			fInfo := &TorrentFileInfo{
				Info:        tInfo,
				TorrentName: tInfo.Name,
				FilePath:    file.DisplayPath(tInfo),
				FileSize:    file.Length,
				FileOffset:  currentOffset,
			}
			fn(fInfo)

			currentOffset += file.Length
		}
	}
}
