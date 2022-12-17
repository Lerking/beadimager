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
	log.Println("Creating canvas group...")
	cg, _ := walk.NewGroupBox(mw.content)
	cg.SetTitle("Canvas")
	cg.SetAlignment(walk.AlignHNearVNear)
	cg.SetLayout(walk.NewVBoxLayout())
	log.Println("Creating canvas...")
	mw.canvasScroll, _ = walk.NewScrollView(cg)
	vb := walk.NewVBoxLayout()
	mw.canvasScroll.SetLayout(vb)
	mw.canvasScroll.SetAlignment(walk.AlignHNearVNear)
	//vb.SetMargins(walk.Margins{0, 0, 0, 0})
	//dw, _ := walk.NewCustomWidgetPixels(mw.canvasScroll, 0, mw.drawStuff)
	//dw.SetClearsBackground(true)
	//dw.SetInvalidatesOnResize(true)
}
