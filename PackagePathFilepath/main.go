package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	files := []string{}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	err = filepath.WalkDir(cwd, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		// fs.SkipDirを返すと、そのディレクトリの中身をスキップする
		// 以下はドットファイルをスキップする例
		if info.Name()[0] == '.' {
			return fs.SkipDir
		}

		files = append(files, path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(files) // カレンとディレクトリ配下のファイル一覧
}
