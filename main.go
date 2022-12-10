package main

import (
	"log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type MyMainWindow struct {
	*walk.MainWindow
	colors         *walk.ScrollView
	canvas         *walk.ScrollView
	properties     *walk.ScrollView
	pallette       Pallette
	beads          []*BeadColor
	pallette_combo *walk.ComboBox
	brand_model    []string
	serie_combo    *walk.ComboBox
	serie_model    []string
	pegboard_combo *walk.ComboBox
	pegboard_model []string
}

const (
	AppName   string = "BeadImager"
	Version   string = "0.0.8"
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
	mw.brand_model = CreateBrandsList(mw)

	DD_Pallette := Composite{
		Layout: HBox{MarginsZero: true},
		Children: []Widget{
			Label{
				Text: "Pallette:",
			},
			ComboBox{
				AssignTo: &mw.pallette_combo,
				Model:    mw.brand_model,
				OnCurrentIndexChanged: func() {
					log.Println("Pallette changed to: ", mw.pallette_combo.Text())
					mw.serie_model = CreateSeriesList(mw)
					mw.serie_combo.SetModel(mw.serie_model)
					mw.serie_combo.SetEnabled(true)
				},
			},
		},
	}

	DD_Serie := Composite{
		Layout: HBox{MarginsZero: true},
		Children: []Widget{
			Label{
				Text: "Serie:",
			},
			ComboBox{
				AssignTo: &mw.serie_combo,
				Enabled:  false,
				OnCurrentIndexChanged: func() {
					log.Println("Serie changed to: ", mw.serie_combo.Text())
					mw.pegboard_model = CreatePegboardsList(mw)
					mw.pegboard_combo.SetModel(mw.pegboard_model)
					mw.pegboard_combo.SetEnabled(true)
				},
			},
		},
	}

	//pegboard_model := CreatePegboardsList(mw)

	DD_Pegboard := Composite{
		Layout: HBox{MarginsZero: true},
		Children: []Widget{
			Label{
				Text: "Pegboard:",
			},
			ComboBox{
				AssignTo: &mw.pegboard_combo,
				Enabled:  false,
				OnCurrentIndexChanged: func() {
					log.Println("Pegboard changed to: ", mw.pegboard_combo.Text())
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
							DD_Pegboard,
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
