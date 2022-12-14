package main

import (
	"encoding/xml"
	"io"
	"log"
	"os"
)

type (
	Pallette struct {
		XMLName xml.Name     `xml:"pallette"`
		Brands  Brandsstruct `xml:"brands"`
	}

	Brandsstruct struct {
		XMLName xml.Name      `xml:"brands"`
		Brand   []Brandstruct `xml:"brand"`
	}

	Brandstruct struct {
		XMLName   xml.Name     `xml:"brand"`
		BrandName string       `xml:"brandname"`
		Series    Seriesstruct `xml:"series"`
	}

	Seriesstruct struct {
		XMLName xml.Name      `xml:"series"`
		Serie   []Seriestruct `xml:"serie"`
	}

	Seriestruct struct {
		XMLName   xml.Name        `xml:"serie"`
		SerieName string          `xml:"seriename"`
		Pegboards Pegboardsstruct `xml:"pegboards"`
		Beads     Beadsstruct     `xml:"beads"`
	}

	Pegboardsstruct struct {
		XMLName  xml.Name         `xml:"pegboards"`
		Pegboard []Pegboardstruct `xml:"pegboard"`
	}

	Pegboardstruct struct {
		XMLName xml.Name `xml:"pegboard"`
		Type    string   `xml:"type"`
		Width   string   `xml:"width"`
		Height  string   `xml:"height"`
	}

	Beadsstruct struct {
		XMLName xml.Name      `xml:"beads"`
		Color   []Colorstruct `xml:"color"`
	}

	Colorstruct struct {
		XMLName       xml.Name `xml:"color"`
		ColorIndex    int      `xml:"colorIndex,attr"`
		ColorName     string   `xml:"colorname"`
		ProductCode   string   `xml:"productCode"`
		Brand         string   `xml:"brand"`
		Red           byte     `xml:"red"`
		Green         byte     `xml:"green"`
		Blue          byte     `xml:"blue"`
		IsPearl       bool     `xml:"isPearl"`
		IsTranslucent bool     `xml:"isTranslucent"`
		IsNeutral     bool     `xml:"isNeutral"`
		IsGrayscale   bool     `xml:"isGrayscale"`
		Disabled      bool     `xml:"disabled"`
		InStock       bool     `xml:"inStock"`
		OnHand        int      `xml:"onHand"`
	}
)

func CreatePegboardsList(mw *MyMainWindow) []string {
	pegboards := make([]string, 0)
	for _, brand := range mw.pallette.Brands.Brand {
		if brand.BrandName == mw.brand_combo.Text() {
			for _, serie := range brand.Series.Serie {
				if serie.SerieName == mw.serie_combo.Text() {
					for _, pegboard := range serie.Pegboards.Pegboard {
						pegboards = append(pegboards, pegboard.Type)
					}
				}
			}
		}
	}
	return pegboards
}

func CreateSeriesList(mw *MyMainWindow) []string {
	series := make([]string, 0)
	for _, brand := range mw.pallette.Brands.Brand {
		if brand.BrandName == mw.brand_combo.Text() {
			for _, serie := range brand.Series.Serie {
				series = append(series, serie.SerieName)
			}
		}
	}
	return series
}

func CreateBrandsList(mw *MyMainWindow) []string {
	brands := make([]string, 0)
	for _, brand := range mw.pallette.Brands.Brand {
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
