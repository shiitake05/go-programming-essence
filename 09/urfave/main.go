// ./urfaveで実行
// package main

// import (
// 	"os"

// 	"github.com/urfave/cli/v2"
// )

// func main() {
// 	app := cli.NewApp()
// 	app.Flags = []cli.Flag{
// 		&cli.StringFlag{
// 			Name:    "config",
// 			Aliases: []string{"c"},
// 			Usage:   "Load configuration from `FILE`",
// 		},
// 	}
// 	app.Name = "score"
// 	app.Usage = "Show student's score"
// 	app.Run(os.Args)
// }

// サブコマンド対応
// ./urfave list -jsonで新しく追加したサブコマンドlistに対するフラグ -jsonを扱うことができる
package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func listStudents() error {
	return nil
}

func listStudentsAsJSON() error {
	return nil
}

func cmdList(c *cli.Context) error {
	if c.Bool("jsom") {
		return listStudentsAsJSON()
	}
	return listStudents()
}

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Load configuration from `FILE`",
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:  "list",
			Usage: "list students",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "json",
					Usage: "output as JSON",
					Value: false,
				},
			},
			Action: cmdList,
		},
	}
	app.Name = "score"
	app.Usage = "Show student's score"
	app.Run(os.Args)
}
