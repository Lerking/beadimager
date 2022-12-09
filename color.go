package main

import (
	"log"

	"github.com/lxn/walk"
)

type (
	BeadColor struct {
		Brand           string
		Checkbox        *walk.CheckBox
		backgroundColor walk.Brush
		Red             byte
		Green           byte
		Blue            byte
	}
)

func LoadBeads(mw *MyMainWindow) {
	for _, brand := range mw.pallette.Brands.Brand {
		log.Println("Loading beads for brand: " + brand.BrandName + " ...")
		for _, series := range brand.Series.Serie {
			for _, bead := range series.Beads.Color {
				log.Println("Loading bead: " + bead.ColorName + " ...")
				if !bead.Disabled {
					bc := NewBeadColor(mw, bead.ColorName, bead.Red, bead.Green, bead.Blue)
					bc.Brand = brand.BrandName
					bc.Red = bead.Red
					bc.Green = bead.Green
					bc.Blue = bead.Blue
					mw.beads = append(mw.beads, bc)
				}
			}
		}
	}
}

func NewBeadColor(mw *MyMainWindow, name string, red byte, green byte, blue byte) *BeadColor {
	color := new(BeadColor)
	//color.SetBackgroundColor(walk.RGB(red, green, blue))
	color.Checkbox, _ = walk.NewCheckBox(mw.colors)
	color.Checkbox.SetName(name)
	//color.Checkbox.SetBackground(color.backgroundColor)

	return color
}

func (bc *BeadColor) SetBackgroundColor(col walk.Color) {
	bc.backgroundColor, _ = walk.NewSolidColorBrush(col)
}
