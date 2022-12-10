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
	Version   string = "0.0.9"
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
	mw.brand_model = CreateBrandsList(mw)
	pallette_trigged := false
	serie_trigged := false
	pegboard_trigged := false

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
					if !pallette_trigged {
						log.Println("Pallette changed to: ", mw.pallette_combo.Text())
						mw.serie_model = CreateSeriesList(mw)
						mw.serie_combo.SetModel(mw.serie_model)
						mw.serie_combo.SetEnabled(true)
					}
					pallette_trigged = true
				},
			},
			HSpacer{},
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
					if !serie_trigged {
						log.Println("Serie changed to: ", mw.serie_combo.Text())
						LoadBeads(mw)
						log.Println("Beads loaded: ", mw.beads)
						mw.pegboard_model = CreatePegboardsList(mw)
						mw.pegboard_combo.SetModel(mw.pegboard_model)
						mw.pegboard_combo.SetEnabled(true)
					}
					serie_trigged = true
				},
			},
			HSpacer{},
		},
	}

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
					if !pegboard_trigged {
						log.Println("Pegboard changed to: ", mw.pegboard_combo.Text())
					}
					pegboard_trigged = true
				},
			},
			HSpacer{},
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
								Text: "Select all colors",
								OnClicked: func() {
									for _, c := range mw.beads {
										c.Checkbox.SetChecked(true)
									}
								},
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
