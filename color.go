package main

import (
	"fmt"
	"log"

	"github.com/lxn/walk"
)

type (
	BeadColor struct {
		Brand           string
		Series          string
		ColorID         int
		Checkbox        *walk.CheckBox
		backgroundColor walk.Brush
		Red             byte
		Green           byte
		Blue            byte
	}
)

func LoadBeads(mw *MyMainWindow) {
	for _, brand := range mw.pallette.Brands.Brand {
		if brand.BrandName == mw.pallette_combo.Text() {
			log.Println("Loading beads for brand: " + brand.BrandName + " ...")
			for _, series := range brand.Series.Serie {
				if series.SerieName == mw.serie_combo.Text() {
					log.Println("Loading beads for serie: " + series.SerieName + " ...")
					for _, bead := range series.Beads.Color {
						log.Println("Loading bead: " + bead.ColorName + " ...")
						if !bead.Disabled {
							bc := NewBeadColor(mw, bead.ColorName, bead.ColorIndex, bead.Red, bead.Green, bead.Blue)
							bc.Brand = brand.BrandName
							bc.Series = series.SerieName
							bc.ColorID = bead.ColorIndex
							bc.Red = bead.Red
							bc.Green = bead.Green
							bc.Blue = bead.Blue
							mw.beads = append(mw.beads, bc)
						}
					}
				}
			}
		}
	}
}

func NewBeadColor(mw *MyMainWindow, name string, id int, red byte, green byte, blue byte) *BeadColor {
	var err error
	log.Println("Creating bead color: " + name + " ...")
	cm, _ := walk.NewComposite(mw.colors)
	cm.SetAlignment(walk.AlignHNearVCenter)
	hb := walk.NewHBoxLayout()
	hb.SetMargins(walk.Margins{5, 0, 20, 0})
	cm.SetLayout(hb)
	color := new(BeadColor)
	log.Println("Bead color struct: ", color)
	color.SetBackgroundColor(walk.RGB(red, green, blue))
	log.Println("Creating checkbox")
	color.Checkbox, err = walk.NewCheckBox(cm)
	if err != nil {
		log.Panic(err)
	}
	log.Println("Checkbox created")
	log.Println("Setting checkbox name")
	color.Checkbox.SetText(name)
	log.Println("Checkbox name set")
	walk.NewHSpacer(cm)
	lbl, _ := walk.NewLabel(cm)
	lbl.SetText(fmt.Sprint("Color ID: ", id))
	cm.SetBackground(color.backgroundColor)

	return color
}

func (bc *BeadColor) SetBackgroundColor(col walk.Color) {
	bc.backgroundColor, _ = walk.NewSolidColorBrush(col)
}
