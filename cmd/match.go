package cmd

import (
	"github.com/imdong/torrent-tools/internal/match"
	"github.com/spf13/cobra"
	"log"
)

// matchCmd 匹配种子
var matchCmd = &cobra.Command{
	Use:   "match <torrent_path> <storage_path> [symlink_path]",
	Short: "根据种子内文件布局创建软链接",
	Long:  `提供种子路径与文件存放路径，自动扫描相同的文件并按照种子内文件布局与命名创建软连接到目标路径`,
	Args:  cobra.MinimumNArgs(2),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var comps []string
		if len(args) == 0 {
			comps = cobra.AppendActiveHelp(comps, "种子文件路径或存放种子文件的目录路径")
		} else if len(args) == 1 {
			comps = cobra.AppendActiveHelp(comps, "用于匹配的本地文件储存路经")
		} else if len(args) == 2 {
			comps = cobra.AppendActiveHelp(comps, "创建软连接的目录")
		} else {
			comps = cobra.AppendActiveHelp(comps, "错误: 指定的参数太多")
		}
		return comps, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		// 获取必需的参数
		torrentPath := args[0]
		storagePath := args[1]

		// 获取可选参数（如果有）
		var symlinkPath string
		if len(args) > 2 {
			symlinkPath = args[2]
		} else {
			symlinkPath = storagePath
		}

		// 获取 flag 值
		var option = match.Option{}
		var err error
		option.OnlyMatch, err = cmd.Flags().GetBool("only-match")
		if err != nil {
			log.Fatalf("Error getting only-match flag: %v", err)
		}
		option.IgnoreTorrentName, err = cmd.Flags().GetBool("ignore-torrent-name")
		if err != nil {
			log.Fatalf("Error getting ignore-torrent-name flag: %v", err)
		}

		match.Start(torrentPath, storagePath, symlinkPath, option)
	},
}

func init() {
	rootCmd.AddCommand(matchCmd)

	// 添加可选标志
	matchCmd.Flags().BoolP("only-match", "m", false, "只输出匹配成功的文件")
	matchCmd.Flags().BoolP("ignore-torrent-name", "i", false, "创建的软连接路径中忽略种子名")
}
