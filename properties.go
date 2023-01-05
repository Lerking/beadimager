package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/lxn/walk"
)

type (
	properties struct {
		propColor  *PropColor
		propScale  *PropScale
		propCanvas *PropCanvas
		propBeads  *PropBeads
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

	PropBeads struct {
		property  *walk.Composite
		showAll   *walk.RadioButton
		greyScale *walk.RadioButton
		inStock   *walk.CheckBox
		visible   bool
	}
)

func ShowInStock(mw *MyMainWindow) {
	for _, bead := range mw.beads {
		for _, s := range bead.Series {
			if s.Name == mw.serie_combo.Text() {
				if s.inStock {
					bead.Color.SetVisible(true)
				} else {
					bead.Color.SetVisible(false)
				}
				break
			}
			break
		}
	}
}

func ShowGreyscale(mw *MyMainWindow) {
	mw.properties.propBeads.showAll.SetChecked(false)
	mw.properties.propBeads.greyScale.SetChecked(true)
	for _, bead := range mw.beads {
		if bead.GreyScale {
			if mw.properties.propBeads.inStock.Checked() {
				for _, s := range bead.Series {
					if s.Name == mw.serie_combo.Text() {
						if s.inStock {
							bead.Color.SetVisible(true)
						} else {
							bead.Color.SetVisible(false)
						}
						break
					}
					break
				}
				break
			} else {
				bead.Color.SetVisible(true)
			}
		} else {
			bead.Color.SetVisible(false)
		}
	}
}

func ShowAll(mw *MyMainWindow) {
	mw.properties.propBeads.showAll.SetChecked(true)
	mw.properties.propBeads.greyScale.SetChecked(false)
	for _, bead := range mw.beads {
		if mw.properties.propBeads.inStock.Checked() {
			for _, s := range bead.Series {
				if s.Name == mw.serie_combo.Text() {
					if s.inStock {
						bead.Color.SetVisible(true)
					} else {
						bead.Color.SetVisible(false)
					}
					break
				}
				break
			}
			break
		} else {
			bead.Color.SetVisible(true)
		}
	}
}

func CreateProperties(mw *MyMainWindow) {
	mw.properties = new(properties)
	mw.properties.propColor = new(PropColor)
	mw.properties.propScale = new(PropScale)
	mw.properties.propCanvas = new(PropCanvas)
	mw.properties.propBeads = new(PropBeads)
	CreateSettingsGroup(mw)
}

func CreateSettingsGroup(mw *MyMainWindow) {
	var err error
	//log.Println("Setting up settings...")
	mw.rightPanel, err = walk.NewComposite(mw.content)
	if err != nil {
		log.Println("Error creating right panel: ", err)
	}
	vb := walk.NewVBoxLayout()
	err = mw.rightPanel.SetLayout(vb)
	if err != nil {
		log.Println("Error setting right panel layout: ", err)
	}
	err = vb.SetMargins(walk.Margins{0, 0, 0, 0})
	if err != nil {
		log.Println("Error setting right panel margins: ", err)
	}
	mw.rightPanel.SetMinMaxSize(walk.Size{Width: 220, Height: 0}, walk.Size{Width: 220, Height: 0})
	sg, err := walk.NewGroupBox(mw.rightPanel)
	if err != nil {
		log.Println("Error creating settings group: ", err)
	}
	err = sg.SetTitle("Settings")
	if err != nil {
		log.Println("Error setting settings group title: ", err)
	}
	err = sg.SetAlignment(walk.AlignHNearVNear)
	if err != nil {
		log.Println("Error setting settings group alignment: ", err)
	}
	vb = walk.NewVBoxLayout()
	err = sg.SetLayout(vb)
	if err != nil {
		log.Println("Error setting settings group layout: ", err)
	}
	mw.propScroll, err = walk.NewScrollView(sg)
	if err != nil {
		log.Println("Error creating property scroll: ", err)
	}
	vb = walk.NewVBoxLayout()
	err = mw.propScroll.SetLayout(vb)
	if err != nil {
		log.Println("Error setting property scroll layout: ", err)
	}
	err = vb.SetMargins(walk.Margins{0, 0, 0, 0})
	if err != nil {
		log.Println("Error setting property scroll margins: ", err)
	}
	CreateColorProperties(mw)
	CreateScaleProperties(mw)
	CreateCanvasProperties(mw)
	CreateBeadsProperties(mw)
}

func CreateColorProperties(mw *MyMainWindow) {
	var err error
	//log.Println("Creating color properties...")
	mw.properties.propColor.property, err = walk.NewComposite(mw.propScroll)
	if err != nil {
		log.Println("Error creating color property: ", err)
	}
	err = mw.properties.propColor.property.SetAlignment(walk.AlignHNearVNear)
	if err != nil {
		log.Println("Error setting color property alignment: ", err)
	}
	vb := walk.NewVBoxLayout()
	err = mw.properties.propColor.property.SetLayout(vb)
	if err != nil {
		log.Println("Error setting color property layout: ", err)
	}
	lbl, err := walk.NewTextLabel(mw.properties.propColor.property)
	if err != nil {
		log.Println("Error creating color label: ", err)
	}
	err = lbl.SetText("Color: ")
	if err != nil {
		log.Println("Error setting color label text: ", err)
	}
	bg, err := walk.NewSolidColorBrush(walk.RGB(167, 45, 234))
	if err != nil {
		log.Println("Error creating color brush: ", err)
	}
	mw.properties.propColor.property.SetBackground(bg)
}

func CreateScaleProperties(mw *MyMainWindow) {
	var err error
	mw.properties.propScale.property, err = walk.NewComposite(mw.propScroll)
	if err != nil {
		log.Println("Error creating scale property: ", err)
	}
	err = mw.properties.propScale.property.SetAlignment(walk.AlignHNearVNear)
	if err != nil {
		log.Println("Error setting scale property alignment: ", err)
	}
	vb := walk.NewVBoxLayout()
	err = mw.properties.propScale.property.SetLayout(vb)
	if err != nil {
		log.Println("Error setting scale property layout: ", err)
	}
	grcom, err := walk.NewComposite(mw.properties.propScale.property)
	if err != nil {
		log.Println("Error creating scale group: ", err)
	}
	err = grcom.SetAlignment(walk.AlignHNearVNear)
	if err != nil {
		log.Println("Error setting scale group alignment: ", err)
	}
	hb := walk.NewVBoxLayout()
	err = hb.SetMargins(walk.Margins{0, 0, 0, 0})
	if err != nil {
		log.Println("Error setting scale group margins: ", err)
	}
	err = grcom.SetLayout(hb)
	if err != nil {
		log.Println("Error setting scale group layout: ", err)
	}
	lbl, err := walk.NewTextLabel(grcom)
	if err != nil {
		log.Println("Error creating scale label: ", err)
	}
	err = lbl.SetText("Scale:")
	if err != nil {
		log.Println("Error setting scale label text: ", err)
	}
	slider, err := walk.NewSlider(grcom)
	if err != nil {
		log.Println("Error creating scale slider: ", err)
	}
	slider.SetTracking(true)
	slider.SetRange(10, 200)
	val, err := strconv.Atoi(ConfigScale)
	if err != nil {
		log.Println("Error converting scale config value to int: ", err)
	}
	slider.SetValue(val)
	sc, err := walk.NewNumberEdit(grcom)
	if err != nil {
		log.Println("Error creating scale number edit: ", err)
	}
	slider.ValueChanged().Attach(func() {
		log.Println("Scale slider value changed")
		nn := float64(slider.Value())
		err = sc.SetValue(nn)
		if err != nil {
			log.Println("Error setting scale number edit value: ", err)
		}
		SetConfigScale(fmt.Sprintf("%0.0f", nn))
	})
	err = sc.SetDecimals(0)
	if err != nil {
		log.Println("Error setting scale number edit decimals: ", err)
	}
	err = sc.SetRange(10, 200)
	if err != nil {
		log.Println("Error setting scale number edit range: ", err)
	}
	nn := float64(slider.Value())
	err = sc.SetValue(nn)
	if err != nil {
		log.Println("Error setting scale number edit value: ", err)
	}
	sc.ValueChanged().Attach(func() {
		log.Println("Scale number edit value changed")
		nn := float64(sc.Value())
		slider.SetValue(int(nn))
		SetConfigScale(fmt.Sprintf("%0.0f", nn))
	})
	bg, err := walk.NewSolidColorBrush(walk.RGB(255, 255, 255))
	if err != nil {
		log.Println("Error creating scale brush: ", err)
	}
	mw.properties.propScale.property.SetBackground(bg)
}

func CreateCanvasProperties(mw *MyMainWindow) {
	var err error
	mw.properties.propCanvas.property, err = walk.NewComposite(mw.propScroll)
	if err != nil {
		log.Println("Error creating canvas property: ", err)
	}
	err = mw.properties.propCanvas.property.SetAlignment(walk.AlignHNearVNear)
	if err != nil {
		log.Println("Error setting canvas property alignment: ", err)
	}
	vb := walk.NewVBoxLayout()
	err = mw.properties.propCanvas.property.SetLayout(vb)
	if err != nil {
		log.Println("Error setting canvas property layout: ", err)
	}
	lbl, err := walk.NewTextLabel(mw.properties.propCanvas.property)
	if err != nil {
		log.Println("Error creating canvas label: ", err)
	}
	err = lbl.SetText("Canvas:")
	if err != nil {
		log.Println("Error setting canvas label text: ", err)
	}
	grcom, err := walk.NewComposite(mw.properties.propCanvas.property)
	if err != nil {
		log.Println("Error creating canvas group: ", err)
	}
	err = grcom.SetAlignment(walk.AlignHNearVNear)
	if err != nil {
		log.Println("Error setting canvas group alignment: ", err)
	}
	hb := walk.NewHBoxLayout()
	err = hb.SetMargins(walk.Margins{0, 0, 0, 0})
	if err != nil {
		log.Println("Error setting canvas group margins: ", err)
	}
	err = grcom.SetLayout(hb)
	if err != nil {
		log.Println("Error setting canvas group layout: ", err)
	}
	cb, err := walk.NewCheckBox(grcom)
	if err != nil {
		log.Println("Error creating canvas checkbox: ", err)
	}
	err = cb.SetAlignment(walk.AlignHNearVNear)
	if err != nil {
		log.Println("Error setting canvas checkbox alignment: ", err)
	}
	err = cb.SetText("Show grid")
	if err != nil {
		log.Println("Error setting canvas checkbox text: ", err)
	}
	switch ConfigShowGrid {
	case "true":
		cb.SetChecked(true)
	case "false":
		cb.SetChecked(false)
	}
	cb.CheckedChanged().Attach(func() {
		log.Println("Grid checkbox changed")
		if !cb.Checked() {
			SetConfigShowGrid("false")
		} else {
			SetConfigShowGrid("true")
		}
	})
	walk.NewHSpacer(grcom)
	grcolb, err := walk.NewPushButton(grcom)
	if err != nil {
		log.Println("Error creating grid color button: ", err)
	}
	err = cb.SetAlignment(walk.AlignHFarVNear)
	if err != nil {
		log.Println("Error setting grid color button alignment: ", err)
	}
	err = grcolb.SetText("Grid color")
	if err != nil {
		log.Println("Error setting grid color button text: ", err)
	}
	grcolb.Clicked().Attach(func() {
		log.Println("Grid color button clicked")
		mw.openImage()
	})
	cb, err = walk.NewCheckBox(mw.properties.propCanvas.property)
	if err != nil {
		log.Println("Error creating pixels checkbox: ", err)
	}
	err = cb.SetAlignment(walk.AlignHNearVNear)
	if err != nil {
		log.Println("Error setting pixels checkbox alignment: ", err)
	}
	err = cb.SetText("Show pixels as beads")
	if err != nil {
		log.Println("Error setting pixels checkbox text: ", err)
	}
	switch ConfigShowBeads {
	case "true":
		cb.SetChecked(true)
	case "false":
		cb.SetChecked(false)
	}
	cb.CheckedChanged().Attach(func() {
		log.Println("Grid checkbox changed")
		if !cb.Checked() {
			SetConfigShowBeads("false")
		} else {
			SetConfigShowBeads("true")
		}
	})
	grcolb, err = walk.NewPushButton(mw.properties.propCanvas.property)
	if err != nil {
		log.Println("Error creating background color button: ", err)
	}
	err = grcolb.SetAlignment(walk.AlignHFarVNear)
	if err != nil {
		log.Println("Error setting background color button alignment: ", err)
	}
	err = grcolb.SetText("Background color")
	if err != nil {
		log.Println("Error setting background color button text: ", err)
	}
	grcolb.Clicked().Attach(func() {
		log.Println("Background color button clicked")
		mw.openImage()
	})
	bg, err := walk.NewSolidColorBrush(walk.RGB(255, 255, 255))
	if err != nil {
		log.Println("Error creating canvas brush: ", err)
	}
	mw.properties.propCanvas.property.SetBackground(bg)
}

func CreateBeadsProperties(mw *MyMainWindow) {
	var err error
	mw.properties.propBeads.property, err = walk.NewComposite(mw.propScroll)
	if err != nil {
		log.Println("Error creating beads property: ", err)
	}
	err = mw.properties.propBeads.property.SetAlignment(walk.AlignHNearVNear)
	if err != nil {
		log.Println("Error setting beads property alignment: ", err)
	}
	vb := walk.NewVBoxLayout()
	err = mw.properties.propBeads.property.SetLayout(vb)
	if err != nil {
		log.Println("Error setting beads property layout: ", err)
	}
	lbl, err := walk.NewTextLabel(mw.properties.propBeads.property)
	if err != nil {
		log.Println("Error creating beads label: ", err)
	}
	err = lbl.SetText("Beads: ")
	if err != nil {
		log.Println("Error setting beads label text: ", err)
	}
	grcom, err := walk.NewComposite(mw.properties.propBeads.property)
	if err != nil {
		log.Println("Error creating beads group: ", err)
	}
	err = grcom.SetAlignment(walk.AlignHNearVNear)
	if err != nil {
		log.Println("Error setting beads group alignment: ", err)
	}
	hb := walk.NewHBoxLayout()
	err = hb.SetMargins(walk.Margins{0, 0, 0, 0})
	if err != nil {
		log.Println("Error setting beads group margins: ", err)
	}
	err = grcom.SetLayout(hb)
	if err != nil {
		log.Println("Error setting beads group layout: ", err)
	}
	mw.properties.propBeads.showAll, err = walk.NewRadioButton(grcom)
	if err != nil {
		log.Println("Error creating beads checkbox: ", err)
	}
	err = mw.properties.propBeads.showAll.SetAlignment(walk.AlignHNearVNear)
	if err != nil {
		log.Println("Error setting beads checkbox alignment: ", err)
	}
	err = mw.properties.propBeads.showAll.SetText("Show all")
	if err != nil {
		log.Println("Error setting beads checkbox text: ", err)
	}
	switch ConfigShowAll {
	case "true":
		mw.properties.propBeads.showAll.SetChecked(true)
	case "false":
		mw.properties.propBeads.showAll.SetChecked(false)
	}
	mw.properties.propBeads.showAll.CheckedChanged().Attach(func() {
		log.Println("Show all checkbox changed")
		if mw.properties.propBeads.showAll.Checked() {
			SetConfigShowAll("true")
			SetConfigGreyscale("false")
			mw.properties.propBeads.greyScale.SetChecked(!mw.properties.propBeads.showAll.Checked())
			//ShowAll(mw)
		}
	})
	walk.NewHSpacer(grcom)
	mw.properties.propBeads.greyScale, err = walk.NewRadioButton(grcom)
	if err != nil {
		log.Println("Error creating pixels checkbox: ", err)
	}
	err = mw.properties.propBeads.greyScale.SetAlignment(walk.AlignHNearVNear)
	if err != nil {
		log.Println("Error setting pixels checkbox alignment: ", err)
	}
	err = mw.properties.propBeads.greyScale.SetText("Greyscale")
	if err != nil {
		log.Println("Error setting pixels checkbox text: ", err)
	}
	switch ConfigGreyscale {
	case "true":
		mw.properties.propBeads.greyScale.SetChecked(true)
	case "false":
		mw.properties.propBeads.greyScale.SetChecked(false)
	}
	mw.properties.propBeads.greyScale.CheckedChanged().Attach(func() {
		log.Println("Greyscale checkbox changed")
		if mw.properties.propBeads.greyScale.Checked() {
			SetConfigGreyscale("true")
			SetConfigShowAll("false")
			mw.properties.propBeads.showAll.SetChecked(!mw.properties.propBeads.greyScale.Checked())
			//ShowGreyscale(mw)
		}
	})
	mw.properties.propBeads.inStock, err = walk.NewCheckBox(mw.properties.propBeads.property)
	if err != nil {
		log.Println("Error creating pixels checkbox: ", err)
	}
	err = mw.properties.propBeads.inStock.SetAlignment(walk.AlignHNearVNear)
	if err != nil {
		log.Println("Error setting pixels checkbox alignment: ", err)
	}
	err = mw.properties.propBeads.inStock.SetText("Show only beads in stock")
	if err != nil {
		log.Println("Error setting pixels checkbox text: ", err)
	}
	switch ConfigInStock {
	case "true":
		mw.properties.propBeads.inStock.SetChecked(true)
	case "false":
		mw.properties.propBeads.inStock.SetChecked(false)
	}
	mw.properties.propBeads.inStock.CheckedChanged().Attach(func() {
		log.Println("In stock changed")
		if mw.properties.propBeads.showAll.Checked() && mw.properties.propBeads.greyScale.Checked() {
			mw.properties.propBeads.greyScale.SetChecked(false)
		}
		if mw.properties.propBeads.inStock.Checked() {
			//ShowInStock(mw)
			SetConfigInStock("true")
		} else {
			SetConfigInStock("false")
		}
	})
	bg, err := walk.NewSolidColorBrush(walk.RGB(255, 255, 255))
	if err != nil {
		log.Println("Error creating beads brush: ", err)
	}
	mw.properties.propBeads.property.SetBackground(bg)
}
