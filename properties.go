package main

import (
	"log"
	"strconv"

	"github.com/lxn/walk"
)

type (
	properties struct {
		propColor  *PropColor
		propScale  *PropScale
		propCanvas *PropCanvas
	}

	PropColor struct {
		property *walk.Composite
	}

	PropScale struct {
		property *walk.Composite
		visible  bool
	}

	PropCanvas struct {
		property *walk.Composite
		visible  bool
	}
)

func CreateSettingsGroup(mw *MyMainWindow) {
	log.Println("Setting up settings...")
	mw.rightPanel, _ = walk.NewComposite(mw.content)
	vb := walk.NewVBoxLayout()
	mw.rightPanel.SetLayout(vb)
	vb.SetMargins(walk.Margins{0, 0, 0, 0})
	mw.rightPanel.SetMinMaxSize(walk.Size{Width: 220, Height: 0}, walk.Size{Width: 220, Height: 0})
	sg, _ := walk.NewGroupBox(mw.rightPanel)
	sg.SetTitle("Settings")
	sg.SetAlignment(walk.AlignHNearVNear)
	vb = walk.NewVBoxLayout()
	sg.SetLayout(vb)
	mw.propScroll, _ = walk.NewScrollView(sg)
	vb = walk.NewVBoxLayout()
	mw.propScroll.SetLayout(vb)
	vb.SetMargins(walk.Margins{0, 0, 0, 0})
	CreateColorProperties(mw)
	CreateScaleProperties(mw)
	CreateCanvasProperties(mw)
}

func CreateColorProperties(mw *MyMainWindow) {
	log.Println("Creating color properties...")
	cp, _ := walk.NewComposite(mw.propScroll)
	cp.SetAlignment(walk.AlignHNearVNear)
	vb := walk.NewVBoxLayout()
	//vb.SetMargins(walk.Margins{5, 0, 5, 0})
	cp.SetLayout(vb)
	log.Println("Creating color label...")
	lbl, _ := walk.NewTextLabel(cp)
	log.Println("Setting color label text...")
	lbl.SetText("Color: ")
	log.Println("Creating color background brush...")
	bg, _ := walk.NewSolidColorBrush(walk.RGB(167, 45, 234))
	log.Println("Setting color background...")
	cp.SetBackground(bg)
}

func CreateScaleProperties(mw *MyMainWindow) {
	log.Println("Creating scale properties...")
	sp, _ := walk.NewComposite(mw.propScroll)
	sp.SetAlignment(walk.AlignHNearVNear)
	vb := walk.NewVBoxLayout()
	//vb.SetMargins(walk.Margins{5, 0, 5, 0})
	sp.SetLayout(vb)
	grcom, _ := walk.NewComposite(sp)
	grcom.SetAlignment(walk.AlignHNearVNear)
	hb := walk.NewVBoxLayout()
	hb.SetMargins(walk.Margins{0, 0, 0, 0})
	grcom.SetLayout(hb)
	log.Println("Creating scale label...")
	lbl, _ := walk.NewTextLabel(grcom)
	log.Println("Setting scale label text...")
	lbl.SetText("Scale:")
	log.Println("Creating scale slider...")
	slider, _ := walk.NewSlider(grcom)
	log.Println("Setting scale slider properties...")
	slider.SetTracking(true)
	slider.SetRange(10, 200)
	val, _ := strconv.Atoi(ConfigScale)
	slider.SetValue(val)
	sc, _ := walk.NewNumberEdit(grcom)
	slider.ValueChanged().Attach(func() {
		log.Println("Scale slider value changed")
		nn := float64(slider.Value())
		log.Println("Setting scale number edit value to: ", nn)
		sc.SetValue(nn)
	})
	sc.SetDecimals(0)
	sc.SetRange(10, 200)
	nn := float64(slider.Value())
	sc.SetValue(nn)
	sc.ValueChanged().Attach(func() {
		log.Println("Scale number edit value changed")
		nn := float64(sc.Value())
		log.Println("Setting scale slider value to: ", nn)
		slider.SetValue(int(nn))
	})
	log.Println("Setting background color...")
	bg, _ := walk.NewSolidColorBrush(walk.RGB(255, 255, 255))
	sp.SetBackground(bg)
}

func CreateCanvasProperties(mw *MyMainWindow) {
	log.Println("Creating canvas properties...")
	cp, _ := walk.NewComposite(mw.propScroll)
	cp.SetAlignment(walk.AlignHNearVNear)
	vb := walk.NewVBoxLayout()
	//vb.SetMargins(walk.Margins{5, 0, 5, 0})
	cp.SetLayout(vb)
	log.Println("Creating canvas label...")
	lbl, _ := walk.NewTextLabel(cp)
	log.Println("Setting canvas label text...")
	lbl.SetText("Canvas:")
	grcom, _ := walk.NewComposite(cp)
	grcom.SetAlignment(walk.AlignHNearVNear)
	hb := walk.NewHBoxLayout()
	hb.SetMargins(walk.Margins{0, 0, 0, 0})
	grcom.SetLayout(hb)
	log.Println("Creating grid checkbox")
	cb, _ := walk.NewCheckBox(grcom)
	cb.SetAlignment(walk.AlignHNearVNear)
	cb.SetText("Show grid")
	switch ConfigShowGrid {
	case "true":
		cb.SetChecked(true)
	case "false":
		cb.SetChecked(false)
	}
	log.Println("Grid checkbox created")
	walk.NewHSpacer(grcom)
	log.Println("Creating grid color button")
	grcolb, _ := walk.NewPushButton(grcom)
	cb.SetAlignment(walk.AlignHFarVNear)
	grcolb.SetText("Grid color")
	log.Println("Grid color button created")
	log.Println("Creating pixels checkbox")
	cb, _ = walk.NewCheckBox(cp)
	cb.SetAlignment(walk.AlignHNearVNear)
	cb.SetText("Show pixels as beads")
	switch ConfigShowBeads {
	case "true":
		cb.SetChecked(true)
	case "false":
		cb.SetChecked(false)
	}
	log.Println("Pixels checkbox created")
	log.Println("Creating canvas background color button")
	grcolb, _ = walk.NewPushButton(cp)
	grcolb.SetAlignment(walk.AlignHFarVNear)
	grcolb.SetText("Background color")
	log.Println("Background color button created")
	bg, _ := walk.NewSolidColorBrush(walk.RGB(255, 255, 255))
	cp.SetBackground(bg)
}
