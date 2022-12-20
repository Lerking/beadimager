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
	cmp.SetLayout(walk.NewVBoxLayout())
	cmp.Layout().SetMargins(walk.Margins{0, 0, 0, 0})
	lbl, _ := walk.NewTextLabel(cmp)
	lbl.SetText(name + " - " + strconv.Itoa(id))
	lbl.SetAlignment(walk.AlignHCenterVCenter)
	cmp.SetBackground(bg)

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
