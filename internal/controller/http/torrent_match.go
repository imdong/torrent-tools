package http

import (
	"fmt"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/evrone/go-clean-template/pkg/logger"
	"github.com/gin-gonic/gin"
	"os"
)

type torrentMatchRoutes struct {
	l logger.Interface
}

func newTorrentMatchRoutes(handler *gin.RouterGroup, l logger.Interface) {
	t := &torrentMatchRoutes{
		l: l,
	}

	h := handler.Group("torrent-match")
	{
		h.GET("ping", t.ping)
	}
}

func (t torrentMatchRoutes) ping(c *gin.Context) {

	//  获取一个种子文件的信息
	torrentFile := "/Users/imdong/Downloads/[M-TEAM]Jade.The Heir to the Throne.Ep15.HDTV.1080p.H264-CNHK.torrent"
	torrentFile = "/Users/imdong/Downloads/[DBY].[庆余年 第二季].Joy.of.Life.2024.S02.2160p.WEB-DL.HEVC.DDP.2Audios-QHstudIo.torrent"

	file, err := os.Open(torrentFile)
	if err != nil {
		t.l.Fatal("Failed to open torrent file: %v", err)
	}
	defer file.Close()

	// 解析种子文件
	mi, err := metainfo.Load(file)
	if err != nil {
		t.l.Fatal("Failed to load torrent file: %v", err)
	}

	// 获取info字典
	info, err := mi.UnmarshalInfo()
	if err != nil {
		t.l.Fatal("Failed to unmarshal info: %v", err)
	}

	t.l.Debug("PieceLength: %d", info.PieceLength)

	pieceLength := info.NumPieces()
	for i := 0; i < pieceLength; i++ {
		piece := info.Piece(i)
		t.l.Debug("Index: %d, Offset: %d, Length: %d, Index: %d, Hash: %d", i, piece.Offset(), piece.Length(), piece.Index(), piece.Hash())
	}

	// 打印文件信息
	fmt.Printf("Name: %s\n", info.Name)
	for _, file := range info.Files {
		fmt.Printf("File: %s, Length: %d\n", file.Path, file.Length)
	}

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
