package main

import (
	"log"

	"github.com/lxn/walk"
)

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
