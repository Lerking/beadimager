package main

import (
	"log"
	"strconv"
	"time"

	"github.com/lxn/walk"
)

func delay(edit *walk.NumberEdit, val string, ret *Retval) {
	delay := time.NewTimer(time.Second * 1)

	<-delay.C
	switch val {
	case "grams":
		edit.SetValue(float64(ret.Grams))
	case "number":
		edit.SetValue(float64(ret.Number))
	}
}

func (mw *MyMainWindow) addBeads(name string, data Serie, id int, bg walk.Brush, ret *Retval) int {
	var (
		grams_edit  *walk.NumberEdit
		number_edit *walk.NumberEdit
	)
	dlg, err := walk.NewDialog(mw.MainWindow)
	if err != nil {
		log.Println(err)
	}
	dlg.SetTitle("Add Beads")
	dlg.SetLayout(walk.NewVBoxLayout())
	dlg.SetSize(walk.Size{Width: 300, Height: 200})
	dlg.SetMinMaxSize(walk.Size{Width: 300, Height: 200}, walk.Size{Width: 300, Height: 200})
	dlg.SetX(mw.MainWindow.X() + 100)
	dlg.SetY(mw.MainWindow.Y() + 100)
	cmp, _ := walk.NewComposite(dlg)
	cmp.SetLayout(walk.NewHBoxLayout())
	cmp.Layout().SetMargins(walk.Margins{0, 0, 0, 0})
	cmp.SetAlignment(walk.AlignHCenterVCenter)
	walk.NewHSpacer(cmp)
	lbl, _ := walk.NewTextLabel(cmp)
	lbl.SetText("#" + strconv.Itoa(id) + " " + name + " - " + " Onhand: " + strconv.Itoa(data.onHand))
	lbl.SetAlignment(walk.AlignHCenterVCenter)
	walk.NewHSpacer(cmp)
	cmp.SetBackground(bg)
	fm, _ := walk.NewComposite(dlg)
	fm.SetLayout(walk.NewVBoxLayout())
	fm.Layout().SetMargins(walk.Margins{0, 0, 0, 0})
	br, _ := walk.NewSolidColorBrush(walk.RGB(255, 255, 255))
	fm.SetBackground(br)
	bgr, _ := walk.NewComposite(fm)
	bgr.SetLayout(walk.NewHBoxLayout())
	lbl, _ = walk.NewTextLabel(bgr)
	lbl.SetText("Grams:")
	lbl.SetAlignment(walk.AlignHCenterVNear)
	walk.NewHSpacer(bgr)
	grams_edit, _ = walk.NewNumberEdit(bgr)
	grams_edit.SetMinMaxSize(walk.Size{Width: 50, Height: 0}, walk.Size{Width: 100, Height: 0})
	grams_edit.SetDecimals(0)
	grams_edit.SetRange(0, 1000)
	grams_edit.ValueChanged().Attach(func() {
		ret.Grams = int(grams_edit.Value())
		switch data.Name {
		case "Mini":
			ret.Number = ret.Grams * 1000 / 30
		case "Midi":
			ret.Number = ret.Grams * 1000 / 60
		case "Maxi":
			ret.Number = ret.Grams * 1000 / 120
		}
		go delay(number_edit, "number", ret)
	})
	bnr, _ := walk.NewComposite(fm)
	bnr.SetLayout(walk.NewHBoxLayout())
	lbl, _ = walk.NewTextLabel(bnr)
	lbl.SetText("Number:")
	lbl.SetAlignment(walk.AlignHCenterVNear)
	walk.NewHSpacer(bnr)
	number_edit, _ = walk.NewNumberEdit(bnr)
	number_edit.SetMinMaxSize(walk.Size{Width: 50, Height: 0}, walk.Size{Width: 100, Height: 0})
	number_edit.SetDecimals(0)
	number_edit.SetRange(0, 100000)
	number_edit.ValueChanged().Attach(func() {
		ret.Number = int(number_edit.Value())
		switch data.Name {
		case "Mini":
			ret.Grams = ret.Number * 30 / 1000
		case "Midi":
			ret.Grams = ret.Number * 60 / 1000
		case "Maxi":
			ret.Grams = ret.Number * 120 / 1000
		}
		go delay(grams_edit, "grams", ret)
	})
	bc, _ := walk.NewComposite(dlg)
	bc.SetLayout(walk.NewHBoxLayout())
	bc.Layout().SetMargins(walk.Margins{0, 0, 0, 0})
	ab, _ := walk.NewPushButton(bc)
	ab.SetText("Add")
	dlg.SetDefaultButton(ab)
	ab.Clicked().Attach(func() {
		//ret.Grams = int(grams_edit.Value())
		//ret.Number = int(number_edit.Value())
		log.Println("grams:", ret.Grams)
		log.Println("number:", ret.Number)
		ret.Clear = false
		dlg.Accept()
	})
	walk.NewHSpacer(bc)
	cl, _ := walk.NewPushButton(bc)
	cl.SetText("Clear")
	cl.Clicked().Attach(func() {
		grams_edit.SetValue(0)
		number_edit.SetValue(0)
		ret.Clear = true
		dlg.Accept()
	})
	cb, _ := walk.NewPushButton(bc)
	cb.SetText("Cancel")
	dlg.SetCancelButton(cb)
	cb.Clicked().Attach(func() {
		dlg.Cancel()
	})

	return dlg.Run()
}

func (mw *MyMainWindow) openImage() error {
	dlg := new(walk.FileDialog)

	dlg.FilePath = UserPath
	dlg.Filter = "Image Files (*.emf;*.bmp;*.exif;*.gif;*.jpeg;*.jpg;*.png;*.tiff)|*.emf;*.bmp;*.exif;*.gif;*.jpeg;*.jpg;*.png;*.tiff"
	dlg.Title = "选择图片"

	if ok, err := dlg.ShowOpen(mw.MainWindow); err != nil {
		return err
	} else if !ok {
		return nil
	}

	prevFilePath := dlg.FilePath
	log.Println("prevFilePath:", prevFilePath)

	return nil
}
