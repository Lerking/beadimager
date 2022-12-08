package main

import "encoding/xml"

type (
	Palette struct {
		xml.Name `xml:"palette"`
		Brands   []Brand `xml:"brand"`
	}

	Brand struct {
		Name   string   `xml:"name"`
		Series []Series `xml:"series"`
	}

	Series struct {
		Name      string     `xml:"name"`
		Pegboards []Pegboard `xml:"pegboard"`
		Beads     []Bead     `xml:"beads"`
	}

	Pegboard struct {
		Name   string `xml:"name"`
		Width  int    `xml:"width"`
		Height int    `xml:"height"`
	}

	Bead struct {
		Colors []Color `xml:"color;attr"`
	}

	Color struct {
		Name          string `xml:"name"`
		ProductCode   string `xml:"productCode"`
		Brand         string `xml:"brand"`
		Red           uint8  `xml:"red"`
		Green         uint8  `xml:"green"`
		Blue          uint8  `xml:"blue"`
		IsPearl       bool   `xml:"isPearl"`
		IsTranslucent bool   `xml:"isTranslucent"`
		IsNeutral     bool   `xml:"isNeutral"`
		IsGrayscale   bool   `xml:"isGrayscale"`
		Disabled      bool   `xml:"disabled"`
	}
)
