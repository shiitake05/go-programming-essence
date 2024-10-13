// 端末を制御して見応えのあるUIを提供する
package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	data := [][]string{
		{"A", "The Good", "500"},
		{"B", "The Very very Bad Man", "288"},
		{"C", "The Ugly", "120"},
		{"D", "The Gopher", "800"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Description", "Number"})

	// このようにすることで、カラムごとのアライメントを変更できる
	table.SetColumnAlignment([]int{
		tablewriter.ALIGN_CENTER,
		tablewriter.ALIGN_DEFAULT,
		tablewriter.ALIGN_DEFAULT,
	})

	// このように変更することでヘッダの色を変更する
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.FgRedColor, tablewriter.Bold, tablewriter.BgBlackColor},
		tablewriter.Colors{tablewriter.BgRedColor, tablewriter.FgWhiteColor},
	)

	// 合計欄を追加する
	table.SetFooterAlignment(tablewriter.ALIGN_RIGHT)
	table.SetFooter([]string{
		"", "", "427.0",
	})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
