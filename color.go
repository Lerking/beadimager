package main

import (
	"fmt"
	"log"

	"github.com/lxn/walk"
)

type (
	BeadColor struct {
		Checkbox        *walk.CheckBox
		backgroundColor walk.Brush
		InfoTooltip     *walk.ToolTip
		WarningTooltip  *walk.ToolTip
		info            *walk.ImageView
		warning         *walk.ImageView
		Brand           string
		Series          string
		Weight          int
		Name            string
		ColorID         int
		Red             byte
		Green           byte
		Blue            byte
		inStock         bool
		onHand          int
	}
)

func CreateBeadsGroup(mw *MyMainWindow) {
	gb, _ := walk.NewGroupBox(mw.leftPanel)
	gb.SetTitle("Beads")
	gb.SetLayout(walk.NewVBoxLayout())
	btn, _ := walk.NewPushButton(gb)
	btn.SetText("Select all colors")
	btn.Clicked().Attach(func() {
		for _, bead := range mw.beads {
			bead.Checkbox.SetChecked(true)
		}
	})
	mw.colors, _ = walk.NewScrollView(gb)
	mw.colors.SetLayout(walk.NewVBoxLayout())
	LoadBeads(mw)
}

func LoadBeads(mw *MyMainWindow) {
	var located bool

	for _, brand := range mw.pallette.Brand {
		if brand.BrandName == mw.brand_combo.Text() {
			log.Println("Loading beads for brand: " + brand.BrandName + " ...")
			log.Println("Loading beads for serie: " + mw.serie_combo.Text() + " ...")
			for _, bead := range brand.Colors {
				for _, serie := range bead.Series.Serie {
					if serie == mw.serie_combo.Text() {
						located = true
					}
				}
				if located {
					log.Println("Loading bead: " + bead.ColorName + " ...")
					if !bead.Disabled {
						bc := NewBeadColor(mw, bead.ColorName, bead.ColorIndex, bead.OnHand, bead.Red, bead.Green, bead.Blue)
						for _, series := range brand.Series {
							if series.Serie == mw.serie_combo.Text() {
								bc.Series = series.Serie
								bc.Weight = series.Weight
							}
						}
						bc.Brand = brand.BrandName
						bc.Name = bead.ColorName
						bc.ColorID = bead.ColorIndex
						bc.Red = bead.Red
						bc.Green = bead.Green
						bc.Blue = bead.Blue
						bc.inStock = bead.InStock
						mw.beads = append(mw.beads, bc)
						if bead.OnHand <= 200 {
							bc.warning.SetVisible(true)
							bc.info.SetVisible(false)
						} else {
							bc.warning.SetVisible(false)
							bc.info.SetVisible(true)
						}
					}
				}
			}
		}
	}
}

func NewBeadColor(mw *MyMainWindow, name string, id int, onhand int, red byte, green byte, blue byte) *BeadColor {
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
	color.onHand = onhand
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
	color.info, err = walk.NewImageView(cm)
	if err != nil {
		log.Println("Error creating info image view: ", err)
	}
	//Setup info tooltip
	color.InfoTooltip, _ = walk.NewToolTip()
	color.InfoTooltip.SetInfoTitle(name + " - " + fmt.Sprint(id))
	color.info.SetImage(walk.IconInformation())
	color.InfoTooltip.AddTool(color.info)
	color.InfoTooltip.SetText(color.info, "Approx. "+fmt.Sprint(color.onHand)+" left on hand")
	color.info.SetVisible(false)
	color.warning, err = walk.NewImageView(cm)
	if err != nil {
		log.Println("Error creating warning image view: ", err)
	}
	//Setup warning tooltip
	color.WarningTooltip, _ = walk.NewToolTip()
	color.WarningTooltip.SetWarningTitle(name + " - " + fmt.Sprint(id))
	color.warning.SetImage(walk.IconWarning())
	color.WarningTooltip.AddTool(color.warning)
	color.WarningTooltip.SetText(color.warning, "Only "+fmt.Sprint(color.onHand)+" left on hand")
	color.warning.SetVisible(false)

	lbl, _ := walk.NewLabel(cm)
	lbl.SetText(fmt.Sprint("Color ID: ", id))
	cm.SetBackground(color.backgroundColor)

	return color
}

func (bc *BeadColor) SetBackgroundColor(col walk.Color) {
	bc.backgroundColor, _ = walk.NewSolidColorBrush(col)
}

func (bc *BeadColor) GetOnHand() int {
	return bc.onHand
}

func (bc *BeadColor) GetInStock() bool {
	return bc.inStock
}

func (bc *BeadColor) SetInStock(inStock bool) {
	bc.inStock = inStock
}

func (bc *BeadColor) GetColorID() int {
	return bc.ColorID
}

func (bc *BeadColor) AddOnHand(grams int) {
	addbeads := grams / bc.Weight * 1000
	bc.onHand += int(addbeads)
}
