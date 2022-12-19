package main

import (
	"encoding/xml"
	"io"
	"log"
	"os"

	"github.com/lxn/walk"
)

type (
	Pallette struct {
		XMLName xml.Name      `xml:"pallette"`
		Brand   []Brandstruct `xml:"brand"`
	}

	Brandstruct struct {
		BrandName string           `xml:"name,attr"`
		Series    []Seriestruct    `xml:"serie"`
		Pegboards []Pegboardstruct `xml:"pegboard"`
		Colors    []Colorstruct    `xml:"color"`
	}

	Seriestruct struct {
		Serie  string `xml:"name,attr"`
		Weight int    `xml:"weightPerThousand"`
	}

	Pegboardstruct struct {
		Serie string `xml:"serie,attr"`
		Type  string `xml:"type"`
		Size  string `xml:"size"`
	}

	Colorstruct struct {
		Series struct {
			XMLName xml.Name `xml:"series"`
			Serie   []struct {
				XMLName xml.Name `xml:"serie"`
				Name    string   `xml:"name,attr"`
				InStock bool     `xml:"inStock"`
				OnHand  int      `xml:"onHand"`
			}
		}
		ColorIndex    int    `xml:"colorIndex,attr"`
		ColorName     string `xml:"colorname"`
		ProductCode   string `xml:"productCode"`
		Brand         string `xml:"brand"`
		Red           byte   `xml:"red"`
		Green         byte   `xml:"green"`
		Blue          byte   `xml:"blue"`
		IsPearl       bool   `xml:"isPearl"`
		IsTranslucent bool   `xml:"isTranslucent"`
		IsNeutral     bool   `xml:"isNeutral"`
		IsGrayscale   bool   `xml:"isGrayscale"`
		Disabled      bool   `xml:"disabled"`
	}

	Pegboards struct {
		Boards []Pegboard
	}

	Pegboard struct {
		brand string
		serie string
		model []string
	}
)

var (
	Serie_triggered bool = false
	Disable         bool
)

func CreatePalletteGroup(mw *MyMainWindow) *walk.GroupBox {
	mw.Pegboards = new(Pegboards)
	mw.leftPanel, _ = walk.NewComposite(mw.content)
	vb := walk.NewVBoxLayout()
	mw.leftPanel.SetLayout(vb)
	vb.SetMargins(walk.Margins{0, 0, 0, 0})
	mw.leftPanel.SetMinMaxSize(walk.Size{Width: 280, Height: 0}, walk.Size{Width: 280, Height: 0})
	pallette_group, _ := walk.NewGroupBox(mw.leftPanel)
	pallette_group.SetTitle("Pallette")
	vb = walk.NewVBoxLayout()
	pallette_group.SetLayout(vb)
	comp, _ := walk.NewComposite(pallette_group)
	hb := walk.NewHBoxLayout()
	comp.SetLayout(hb)
	comp.Layout().SetMargins(walk.Margins{0, 0, 0, 0})
	lbl, _ := walk.NewLabel(comp)
	lbl.SetText("Brand:")
	mw.brand_combo, _ = walk.NewComboBox(comp)
	mw.brand_combo.SetModel(CreateBrandsList(mw))
	mw.brand_combo.SetText(ConfigBrand)
	comp, _ = walk.NewComposite(pallette_group)
	comp.SetLayout(walk.NewHBoxLayout())
	comp.Layout().SetMargins(walk.Margins{0, 0, 0, 0})
	lbl, _ = walk.NewLabel(comp)
	lbl.SetText("Serie:")
	mw.serie_combo, _ = walk.NewComboBox(comp)
	mw.serie_combo.SetModel(CreateSeriesList(mw))
	mw.serie_combo.SetText(ConfigSerie)
	mw.serie_combo.CurrentIndexChanged().Attach(func() {
		log.Println("Serie triggered: ", mw.serie_combo.Text())
		ShowBeads(mw, mw.serie_combo.Text())
		for _, model := range mw.Pegboards.Boards {
			if model.brand == mw.brand_combo.Text() && model.serie == mw.serie_combo.Text() {
				mw.pegboard_combo.SetText("")
				mw.pegboard_combo.SetModel(model.model)
				log.Println("Pegboard model set to: ", model.model)
			}
		}
	})
	comp, _ = walk.NewComposite(pallette_group)
	comp.SetLayout(walk.NewHBoxLayout())
	comp.Layout().SetMargins(walk.Margins{0, 0, 0, 0})
	lbl, _ = walk.NewLabel(comp)
	lbl.SetText("Pegboard:")
	mw.pegboard_combo, _ = walk.NewComboBox(comp)
	CreatePegboardsList(mw)
	for _, model := range mw.Pegboards.Boards {
		if model.brand == mw.brand_combo.Text() && model.serie == mw.serie_combo.Text() {
			mw.pegboard_combo.SetModel(model.model)
		}
	}
	mw.pegboard_combo.SetText(ConfigPegboard)
	ShowBeads(mw, mw.serie_combo.Text())
	return pallette_group
}

func CreatePegboardsList(mw *MyMainWindow) {
	var pb *Pegboard
	for _, brand := range mw.pallette.Brand {
		for _, serie := range brand.Series {
			pb = new(Pegboard)
			pb.brand = brand.BrandName
			pb.serie = serie.Serie
			for _, pegboard := range brand.Pegboards {
				if pegboard.Serie == pb.serie {
					pb.model = append(pb.model, pegboard.Type)
				}
			}
			mw.Pegboards.Boards = append(mw.Pegboards.Boards, *pb)
		}
	}
	log.Println("Pegboard: ", mw.Pegboards)
}

func CreateSeriesList(mw *MyMainWindow) []string {
	series := make([]string, 0)
	for _, brand := range mw.pallette.Brand {
		if brand.BrandName == mw.brand_combo.Text() {
			for _, serie := range brand.Series {
				series = append(series, serie.Serie)
			}
		}
	}
	return series
}

func CreateBrandsList(mw *MyMainWindow) []string {
	brands := make([]string, 0)
	for _, brand := range mw.pallette.Brand {
		brands = append(brands, brand.BrandName)
	}
	return brands
}

func CreatePallette(mw *MyMainWindow) {
	// Open our xmlFile
	XMLFile, err := os.Open("pallettes\\pallette.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Print("Failed to open pallette.xml")
		log.Panic(err)
	}

	log.Println("Successfully Opened pallette.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer XMLFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := io.ReadAll(XMLFile)

	er := xml.Unmarshal(byteValue, &mw.pallette)
	if er != nil {
		log.Printf("Failed to unmarshal: %v", er)
	}
}
