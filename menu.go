package main

import (
	"log"

	"github.com/lxn/walk"
)

func SetupMenu(mw *MyMainWindow) {
	// MenuBar
	var tool *walk.Action
	var menutool *walk.Menu
	menutool, _ = walk.NewMenu()
	tool = walk.NewMenuAction(menutool)
	tool.SetText("File")
	open := walk.NewAction()
	open.SetText("Open")
	openImg, err := walk.NewImageFromFileForDPI("images\\16x16\\status\\folder-open.png", 144)
	if err != nil {
		log.Println("Error loading icon: ", err)
	}
	err = open.SetImage(openImg)
	if err != nil {
		log.Println("Error setting icon: ", err)
	}
	exit := walk.NewAction()
	exit.SetText("Exit")
	exit.Triggered().Attach(func() {
		mw.MainWindow.Close()
	})

	menutool.Actions().Add(open)
	menutool.Actions().Add(exit)

	men2, _ := walk.NewMenu()
	too2 := walk.NewMenuAction(men2)
	too2.SetText("Help")
	about := walk.NewAction()
	about.SetText("About")
	about.Triggered().Attach(func() {
		txt := AppName + " " + Version + "\n" + CopyRight
		walk.MsgBox(nil, "About", txt, walk.MsgBoxIconInformation)
	})

	men2.Actions().Add(about)

	mw.MainWindow.Menu().Actions().Add(tool)
	mw.MainWindow.Menu().Actions().Add(too2)
}
