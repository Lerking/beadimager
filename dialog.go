package main

import (
	"log"
	"strconv"

	"github.com/lxn/walk"
)

func (mv *MyMainWindow) addBeads(name string, data Serie, id int, bg walk.Brush) error {
	log.Println("Adding beads...")
	dlg, err := walk.NewDialog(mv.MainWindow)
	if err != nil {
		return err
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
	lbl, _ := walk.NewTextLabel(cmp)
	lbl.SetText(name + " - " + strconv.Itoa(id))
	cmp.SetBackground(bg)
	gr, _ := walk.NewGroupBox(dlg)
	gr.SetTitle("Beads")
	fm, _ := walk.NewComposite(gr)
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
	lbl, _ = walk.NewTextLabel(bgr)
	lbl.SetText("Number:")
	lbl.SetAlignment(walk.AlignHCenterVNear)
	walk.NewHSpacer(bnr)
	ne, _ := walk.NewNumberEdit(bnr)
	ne.SetDecimals(0)
	ne.SetRange(0, 100000)
	bc, _ := walk.NewComposite(dlg)
	bc.SetLayout(walk.NewHBoxLayout())
	ab, _ := walk.NewPushButton(bc)
	ab.SetText("Add")
	dlg.SetDefaultButton(ab)
	walk.NewHSpacer(bc)
	cb, _ := walk.NewPushButton(bc)
	cb.SetText("Cancel")
	dlg.SetCancelButton(cb)

	dlg.Show()
	return nil
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
