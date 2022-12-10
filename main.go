package main

import (
	"log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type MyMainWindow struct {
	*walk.MainWindow
	colors     *walk.ScrollView
	canvas     *walk.ScrollView
	properties *walk.ScrollView
	pallette   Pallette
	beads      []*BeadColor
}

const (
	AppName   string = "BeadImager"
	Version   string = "0.0.7"
	CopyRight string = "©2022 Jan Lerking"
	STD_MESS  string = "Ready"
	UserPath  string = "C:\\Users\\janle\\BeadImager"
	LogFile   string = "BeadImager.log"
	Sep       string = "\\"
)

func main() {
	InitLogFile()

	walk.AppendToWalkInit(func() {
		walk.FocusEffect, _ = walk.NewBorderGlowEffect(walk.RGB(0, 63, 255))
		walk.InteractionEffect, _ = walk.NewDropShadowEffect(walk.RGB(63, 63, 63))
		walk.ValidationErrorEffect, _ = walk.NewBorderGlowEffect(walk.RGB(255, 0, 0))
	})
	mw := &MyMainWindow{}
	log.Println("MainWindow created")
	CreatePallette(mw)
	log.Println("Pallette created: ", mw.pallette)
	//LoadBeads(mw)
	//log.Println("Beads loaded: ", mw.beads)
	brand_model := CreateBrandsList(mw)
	pallette_combo := new(walk.ComboBox)

	DD_Pallette := Composite{
		Layout: HBox{MarginsZero: true},
		Children: []Widget{
			Label{
				Text: "Pallette:",
			},
			ComboBox{
				AssignTo: &pallette_combo,
				Model:    brand_model,
				OnCurrentIndexChanged: func() {
					log.Println("Pallette changed to: ", pallette_combo.Text())
				},
			},
		},
	}

	serie_model := CreateSeriesList(mw)
	serie_combo := new(walk.ComboBox)

	DD_Serie := Composite{
		Layout: HBox{MarginsZero: true},
		Children: []Widget{
			Label{
				Text: "Serie:",
			},
			ComboBox{
				AssignTo: &serie_combo,
				Model:    serie_model,
				OnCurrentIndexChanged: func() {
					log.Println("Serie changed to: ", serie_combo.Text())
				},
			},
		},
	}

	if _, err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    AppName + " " + Version,
		MinSize:  Size{800, 600},

		Layout: VBox{MarginsZero: true},
		Children: []Widget{
			Composite{
				Layout: HBox{MarginsZero: true},
				Children: []Widget{
					Composite{
						Layout: VBox{MarginsZero: true},
						Children: []Widget{
							DD_Pallette,
							DD_Serie,
							PushButton{
								Text:      "Edit Animal",
								OnClicked: func() {},
							},
							ScrollView{
								AssignTo: &mw.colors,
								Layout:   VBox{MarginsZero: true},
							},
						},
					},
					ScrollView{
						AssignTo: &mw.canvas,
						Layout:   VBox{MarginsZero: true},
					},
					ScrollView{
						AssignTo: &mw.properties,
						Layout:   VBox{MarginsZero: true},
					},
				},
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}
}
