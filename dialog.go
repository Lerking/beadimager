package main

import (
	"log"
	"strconv"

	"github.com/lxn/walk"
)

func (mv *MyMainWindow) addBeads(name string, data Serie, id int, bg walk.Brush, ret *Retval) int {
	log.Println("Adding beads...")
	dlg, err := walk.NewDialog(mv.MainWindow)
	if err != nil {
		log.Println(err)
	}
	dlg.SetTitle("Add Beads")
	dlg.SetLayout(walk.NewVBoxLayout())
	dlg.SetSize(walk.Size{Width: 300, Height: 200})
	dlg.SetMinMaxSize(walk.Size{Width: 300, Height: 200}, walk.Size{Width: 300, Height: 200})
	dlg.SetX(mv.MainWindow.X() + 100)
	dlg.SetY(mv.MainWindow.Y() + 100)
	cmp, _ := walk.NewComposite(dlg)
	cmp.SetLayout(walk.NewHBoxLayout())
	cmp.Layout().SetMargins(walk.Margins{0, 0, 0, 0})
	cmp.SetAlignment(walk.AlignHCenterVCenter)
	walk.NewHSpacer(cmp)
	lbl, _ := walk.NewTextLabel(cmp)
	lbl.SetText(name + " - " + strconv.Itoa(id))
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
	le, _ := walk.NewNumberEdit(bgr)
	le.SetDecimals(0)
	le.SetRange(0, 1000)
	bnr, _ := walk.NewComposite(fm)
	bnr.SetLayout(walk.NewHBoxLayout())
	lbl, _ = walk.NewTextLabel(bnr)
	lbl.SetText("Number:")
	lbl.SetAlignment(walk.AlignHCenterVNear)
	walk.NewHSpacer(bnr)
	ne, _ := walk.NewNumberEdit(bnr)
	ne.SetDecimals(0)
	ne.SetRange(0, 100000)
	bc, _ := walk.NewComposite(dlg)
	bc.SetLayout(walk.NewHBoxLayout())
	bc.Layout().SetMargins(walk.Margins{0, 0, 0, 0})
	ab, _ := walk.NewPushButton(bc)
	ab.SetText("Add")
	dlg.SetDefaultButton(ab)
	ab.Clicked().Attach(func() {
		ret.Grams = int(le.Value())
		ret.Number = int(ne.Value())
		log.Println("grams:", ret.Grams)
		log.Println("number:", ret.Number)
		dlg.Accept()
	})
	walk.NewHSpacer(bc)
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
