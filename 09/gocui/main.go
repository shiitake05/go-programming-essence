// Goで端末向けのアプリケーションを作成するためのライブラリ
package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.ASCII = true
	g.SetManagerFunc(layout)

	// CTRL-Cで終了
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	s := `Go is an open source programming launguage that makes it easy to build secure, scalable systems.`
	maxX, maxY := g.Size()
	if v, err := g.SetView("hello", maxX/2-21, maxY/2-2, maxX/2+21, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, s)
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
