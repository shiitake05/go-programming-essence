// spf13/cobraはボイラーテンプレートが生成できるライブラリ
//「go install github.com/spf13/cobra-cli@latest」
// 「cobra-cli init -l MIT -a ""」でボイラーテンプレートを生成
// 詳細はもう少し勉強した後にやる

package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cobraapp",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.PersistentFlags().String("foo", "", "A help for foo")
	updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
