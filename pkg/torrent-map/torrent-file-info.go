package torrent_map

import (
	"github.com/anacrolix/torrent/metainfo"
	"math"
)

// TorrentFileInfo 种子内文件信息
type TorrentFileInfo struct {
	Info        *metainfo.Info
	TorrentName string
	FilePath    string
	FileSize    int64
	FileOffset  int64
}

// ContrastFileHash 比较这两个文件是否可能相同
func (tFileInfo *TorrentFileInfo) ContrastFileHash(sFilePath string) bool {
	// 获取下一个完整文件块的hash
	pieceLength := tFileInfo.Info.PieceLength
	pieceIndexThis := int64(math.Ceil(float64(tFileInfo.FileOffset) / float64(pieceLength)))
	pieceOffsetStart := pieceIndexThis * pieceLength
	pieceOffsetEnd := pieceOffsetStart + pieceLength

	// 如果下一块 hash 的范围以超出本文件范围则不匹配
	fileStart := tFileInfo.FileOffset
	fileEnd := fileStart + tFileInfo.FileSize

	// 不在范围内就直接不相同了
	if (fileStart > pieceOffsetStart) || (fileEnd < pieceOffsetEnd) {
		return false
	}

	// 计算该 chunk 的 位置
	fileChunkStart := pieceOffsetStart - fileStart

	// 取出文件该位置内容并计算hash
	fileChunkHash, _ := HashFileSection(sFilePath, fileChunkStart, pieceLength)

	// 对比 hash key
	pieceHash := tFileInfo.Info.Piece(int(pieceIndexThis)).Hash().String()

	// fmt.Printf("pieceHash: %s, fileChunkHash: %s\n", pieceHash, fileChunkHash)

	return pieceHash == fileChunkHash
}
