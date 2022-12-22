package main

import (
	"log"

	"github.com/lxn/walk"
)

type (
	Canvas struct {
		*walk.Canvas
	}
)

func CreateCanvasGroup(mw *MyMainWindow) {
	cg, err := walk.NewGroupBox(mw.content)
	if err != nil {
		log.Println("Error creating canvas group: ", err)
	}
	err = cg.SetTitle("Canvas")
	if err != nil {
		log.Println("Error setting canvas group title: ", err)
	}
	err = cg.SetAlignment(walk.AlignHNearVNear)
	if err != nil {
		log.Println("Error setting canvas group alignment: ", err)
	}
	err = cg.SetLayout(walk.NewVBoxLayout())
	if err != nil {
		log.Println("Error setting canvas group layout: ", err)
	}
	mw.canvasScroll, err = walk.NewScrollView(cg)
	if err != nil {
		log.Println("Error creating canvas scroll: ", err)
	}
	vb := walk.NewVBoxLayout()
	err = mw.canvasScroll.SetLayout(vb)
	if err != nil {
		log.Println("Error setting canvas scroll layout: ", err)
	}
	err = mw.canvasScroll.SetAlignment(walk.AlignHNearVNear)
	if err != nil {
		log.Println("Error setting canvas scroll alignment: ", err)
	}
}
