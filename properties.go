package main

import (
	"log"

	"github.com/lxn/walk"
)

type (
	properties struct {
		propColor  *PropColor
		propScale  *PropScale
		propCanvas *PropCanvas
	}

	PropColor struct {
		*walk.Composite
	}

	PropScale struct {
		*walk.Composite
	}

	PropCanvas struct {
		property *walk.Composite
	}
)

func (cp *PropCanvas) newPropCanvas(mw *MyMainWindow) {
	var err error
	log.Println("Creating canvas properties...")
	cp.property, err = walk.NewComposite(mw.propScroll)
	if err != nil {
		log.Println("Error creating canvas properties: ", err)
	}
	cp.property.SetAlignment(walk.AlignHNearVNear)
	vb := walk.NewVBoxLayout()
	vb.SetMargins(walk.Margins{5, 0, 20, 0})
	cp.property.SetLayout(vb)
	grcom, _ := walk.NewComposite(cp.property)
	hb := walk.NewHBoxLayout()
	hb.SetMargins(walk.Margins{5, 0, 20, 0})
	grcom.SetLayout(hb)
	log.Println("Creating grid checkbox")
	cb, _ := walk.NewCheckBox(grcom)
	cb.SetText("Show grid")
	log.Println("Grid checkbox created")
	walk.NewHSpacer(grcom)
	log.Println("Creating grid color button")
	grcolb, _ := walk.NewPushButton(grcom)
	grcolb.SetText("Grid color")
	log.Println("Grid color button created")
	log.Println("Creating pixels checkbox")
	cb, _ = walk.NewCheckBox(cp.property)
	cb.SetText("Show pixels as beads")
	log.Println("Pixels checkbox created")
	log.Println("Creating canvas background color button")
	grcolb, _ = walk.NewPushButton(cp.property)
	grcolb.SetText("Grid color")
	log.Println("Grid color button created")
}

func NewCanvasProperties() *properties {
	return new(properties)
}

func newPropScale() *PropScale {
	return new(PropScale)
}

func newPropColor() *PropColor {
	return new(PropColor)
}

func newProperties() *properties {
	return new(properties)
}
