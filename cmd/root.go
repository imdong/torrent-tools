package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd 入口命令
var rootCmd = &cobra.Command{
	Use:   "torrent-tools",
	Short: "种子工具集",
	Long:  `这是一个种子相关的工具合集`,
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}
