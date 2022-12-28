package main

import "github.com/lxn/walk"

func SetupMenu(mw *MyMainWindow) {
	// MenuBar
	var tool *walk.Action
	var menutool *walk.Menu
	menutool, _ = walk.NewMenu()
	tool = walk.NewMenuAction(menutool)
	tool.SetText("File")
	open := walk.NewAction()
	open.SetText("Open")
	exit := walk.NewAction()
	exit.SetText("Exit")

	menutool.Actions().Add(open)
	menutool.Actions().Add(exit)

	men2, _ := walk.NewMenu()
	too2 := walk.NewMenuAction(men2)
	too2.SetText("Help")

	mw.MainWindow.Menu().Actions().Add(tool)
	mw.MainWindow.Menu().Actions().Add(too2)
}
