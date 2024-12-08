package main

import (
	"fmt"
	"github.com/imdong/torrent-tools/internal/match"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func main() {
	torrentPath, storagePath, symlinkTargetPath, option := parseParams()

	log.Printf(
		"Torrent Path: %s\nStorage Path: %s\nSymlink Target Path: %s\nOnly Match: %t\nIgnore Torrent Name: %t\n",
		torrentPath, storagePath, symlinkTargetPath,
		option.OnlyMatch, option.IgnoreTorrentName,
	)

	match.Start(torrentPath, storagePath, symlinkTargetPath, option)
}

// 解析启动参数
func parseParams() (torrentPath, storagePath, symlinkTargetPath string, option match.Option) {
	var rootCmd = &cobra.Command{
		Use:   "torrent-match <torrent_path> <storage_path> [symlink_target_path]",
		Short: "搜索种子中已经存在于本地的文件，并创建软连接到指定目录。",
		Args:  cobra.MinimumNArgs(2), // 最少需要两个参数：<torrent_path> <storage_path>
		Run: func(cmd *cobra.Command, args []string) {
			// 获取必需的参数
			torrentPath = args[0]
			storagePath = args[1]

			// 获取可选参数（如果有）
			if len(args) > 2 {
				symlinkTargetPath = args[2]
			} else {
				symlinkTargetPath = storagePath
			}
		},
	}

	// 添加可选标志
	rootCmd.Flags().BoolVar(&option.OnlyMatch, "only-match", false, "只输出匹配成功的文件")
	rootCmd.Flags().BoolVar(&option.IgnoreTorrentName, "ignore-torrent-name", false, "创建的软连接路径中忽略种子名")

	// 执行命令
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return
}
