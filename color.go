package main

import (
	"fmt"
	"log"

	"github.com/lxn/walk"
)

type (
	BeadColor struct {
		Color           *walk.Composite
		Checkbox        *walk.CheckBox
		backgroundColor walk.Brush
		InfoTooltip     *walk.ToolTip
		WarningTooltip  *walk.ToolTip
		ErrorTooltip    *walk.ToolTip
		add             *walk.ImageView
		info            *walk.ImageView
		warning         *walk.ImageView
		error           *walk.ImageView
		Brand           string
		Series          []*Serie
		Weight          []int
		Name            string
		ColorID         int
		Red             byte
		Green           byte
		Blue            byte
	}

	Serie struct {
		Name    string
		inStock bool
		onHand  int
	}

	Retval struct {
		Grams  int
		Number int
	}
)

func ShowBeads(mw *MyMainWindow, serie string) {
	for _, bead := range mw.beads {
		bead.Color.SetVisible(false)
		for _, s := range bead.Series {
			if s.Name == serie {
				bead.Color.SetVisible(true)
				bead.InfoTooltip.SetText(bead.info, "Approx. "+fmt.Sprint(s.onHand)+" left on hand")
				bead.WarningTooltip.SetText(bead.warning, "Only "+fmt.Sprint(s.onHand)+" left on hand")
				bead.ErrorTooltip.SetText(bead.error, "Out of stock")
				if s.onHand <= 200 {
					bead.warning.SetVisible(true)
					bead.info.SetVisible(false)
					bead.error.SetVisible(false)
				} else {
					bead.warning.SetVisible(false)
					bead.info.SetVisible(true)
					bead.error.SetVisible(false)
				}
				if s.onHand <= 0 {
					bead.warning.SetVisible(false)
					bead.info.SetVisible(false)
					bead.error.SetVisible(true)
				}
			}
		}
	}
}

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
	ShowBeads(mw, mw.serie_combo.Text())
}

func LoadBeads(mw *MyMainWindow) {
	for _, brand := range mw.pallette.Brand {
		if brand.BrandName == mw.brand_combo.Text() {
			for _, bead := range brand.Colors {
				if !bead.Disabled {
					bc := NewBeadColor(mw, bead.ColorName, bead.ColorIndex, bead.Red, bead.Green, bead.Blue)
					for _, s := range bead.Series.Serie {
						se := new(Serie)
						se.Name = s.Name
						se.inStock = s.InStock
						se.onHand = s.OnHand
						bc.Series = append(bc.Series, se)
					}
					bc.Brand = brand.BrandName
					bc.Name = bead.ColorName
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

func NewBeadColor(mw *MyMainWindow, name string, id int, red byte, green byte, blue byte) *BeadColor {
	var err error
	cm, _ := walk.NewComposite(mw.colors)
	cm.SetAlignment(walk.AlignHNearVCenter)
	hb := walk.NewHBoxLayout()
	hb.SetMargins(walk.Margins{5, 0, 20, 0})
	cm.SetLayout(hb)
	color := new(BeadColor)
	color.Color = cm
	color.SetBackgroundColor(walk.RGB(red, green, blue))
	color.Checkbox, err = walk.NewCheckBox(cm)
	if err != nil {
		log.Panic(err)
	}
	err = color.Checkbox.SetText(name)
	if err != nil {
		log.Panic(err)
	}
	walk.NewHSpacer(cm)
	color.add, err = walk.NewImageView(cm)
	if err != nil {
		log.Println("Error creating add image view: ", err)
	}
	img, _ := walk.NewImageFromFileForDPI("images/plus_alpha.png", 96)
	color.add.SetImage(img)
	color.add.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
		var (
			data Serie
			ret  = new(Retval)
		)
		if button == walk.LeftButton {
			for _, s := range color.Series {
				if s.Name == mw.serie_combo.Text() {
					data = *s
				}
			}
			val := mw.addBeads(name, data, color.ColorID, color.backgroundColor, ret)
			if val == 1 {
				for _, b := range mw.beads {
					if b.ColorID == color.ColorID {
						for _, s := range b.Series {
							if s.Name == mw.serie_combo.Text() {
								s.onHand += ret.Number
								log.Println("Added ", ret.Number, " beads of ", name, " to ", mw.serie_combo.Text(), " (", s.onHand, " on hand)")
							}
						}
					}
				}
				for ib, p := range mw.pallette.Brand {
					if p.BrandName == mw.brand_combo.Text() {
						for ic, c := range p.Colors {
							if c.ColorIndex == color.ColorID {
								for is, s := range c.Series.Serie {
									if s.Name == mw.serie_combo.Text() {
										mw.pallette.Brand[ib].Colors[ic].Series.Serie[is].OnHand += ret.Number
										log.Println("Added ", ret.Number, " beads of ", name, " to ", mw.serie_combo.Text(), " (", s.OnHand, " on hand)")
										break
									}
								}
								break
							}
						}
						break
					}
				}
				//mw.pallette.Brand[0].Colors[0].Series.Serie[1].OnHand += ret.Number
				//log.Println("Onhand updated: ", mw.pallette.Brand[0].Colors[0].Series.Serie[1].OnHand)
				log.Println("Pallette updated: ", mw.pallette)
				log.Println("Saving palette file...")
				WritePaletteFile(mw)
				ShowBeads(mw, mw.serie_combo.Text())
			} else if val == 2 {
				log.Println("Canceled")
			} else {
				log.Println("Returned value: ", val)
			}
		}
	})
	color.info, err = walk.NewImageView(cm)
	if err != nil {
		log.Println("Error creating info image view: ", err)
	}
	//Setup info tooltip
	color.InfoTooltip, _ = walk.NewToolTip()
	color.InfoTooltip.SetInfoTitle(name + " - " + fmt.Sprint(id))
	color.info.SetImage(walk.IconInformation())
	color.InfoTooltip.AddTool(color.info)
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
	color.warning.SetVisible(false)
	color.error, err = walk.NewImageView(cm)
	if err != nil {
		log.Println("Error creating error image view: ", err)
	}
	//Setup error tooltip
	color.ErrorTooltip, _ = walk.NewToolTip()
	color.ErrorTooltip.SetErrorTitle(name + " - " + fmt.Sprint(id))
	color.error.SetImage(walk.IconError())
	color.ErrorTooltip.AddTool(color.error)
	color.error.SetVisible(false)

	lbl, _ := walk.NewLabel(cm)
	lbl.SetText(fmt.Sprint("Color ID: ", id))
	cm.SetBackground(color.backgroundColor)

	return color
}

func (bc *BeadColor) SetBackgroundColor(col walk.Color) {
	bc.backgroundColor, _ = walk.NewSolidColorBrush(col)
}

func (bc *BeadColor) GetOnHand(serie string) int {
	for _, s := range bc.Series {
		if s.Name == serie {
			return s.onHand
		}
	}
	return 0
}

func (bc *BeadColor) GetInStock(serie string) bool {
	for _, s := range bc.Series {
		if s.Name == serie {
			return s.inStock
		}
	}
	return false
}

func (bc *BeadColor) SetInStock(serie string, inStock bool) {
	for _, s := range bc.Series {
		if s.Name == serie {
			s.inStock = inStock
		}
	}
}

func (bc *BeadColor) GetColorID() int {
	return bc.ColorID
}
