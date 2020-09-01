// START OMIT

package main

import (
	"github.com/lxn/walk/declarative"
)

func main() {
	mw := declarative.MainWindow{
		Title: "this is title",
		Size: declarative.Size{500,500},
		Layout: declarative.VBox{MarginsZero: true,},
		Children: []declarative.Widget{
			declarative.TextLabel{
				Text: "hello world",
			},
		},
	}

	mw.Run()
}
// END OMIT
