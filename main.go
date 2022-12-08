package main

import (
	"log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type MyMainWindow struct {
	*walk.MainWindow
	te *walk.TextEdit
}

const (
	AppName   string = "BeadImager"
	Version   string = "0.0.2"
	CopyRight string = "©2022 Jan Lerking"
	STD_MESS  string = "Ready"
)

func main() {
	mw := &MyMainWindow{}

	if _, err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    AppName + " " + Version,
		MinSize:  Size{800, 600},

		Layout: VBox{MarginsZero: true},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					PushButton{
						Text: "Edit Animal",
					},
					TextEdit{
						AssignTo: &mw.te,
						ReadOnly: true,
					},
				},
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}
}
