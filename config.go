package main

import (
	"log"
	"os"

	"github.com/bigkevmcd/go-configparser"
)

var (
	Config     *configparser.ConfigParser
	ConfigFile string = "BeadImager.conf"
)

func SetConfigBrand(s string) {
	log.Printf("Setting brand to: %s\n", s)
	Config, _ := configparser.Parse(UserPath + Sep + ConfigFile)
	Config.Set("pallette", "brand", s)
	Config.SaveWithDelimiter(UserPath+Sep+ConfigFile, "=")
}

func SetConfigSerie(s string) {
	log.Printf("Setting serie to: %s\n", s)
	Config, _ := configparser.Parse(UserPath + Sep + ConfigFile)
	Config.Set("pallette", "serie", s)
	Config.SaveWithDelimiter(UserPath+Sep+ConfigFile, "=")
}

func SetConfigPegboard(s string) {
	log.Printf("Setting pegboard to: %s\n", s)
	Config, _ := configparser.Parse(UserPath + Sep + ConfigFile)
	Config.Set("pallette", "pegboard", s)
	Config.SaveWithDelimiter(UserPath+Sep+ConfigFile, "=")
}

func SetConfigShowBeads(s string) {
	log.Printf("Setting showbeads to: %s\n", s)
	Config, _ := configparser.Parse(UserPath + Sep + ConfigFile)
	Config.Set("canvas", "showbeads", s)
	Config.SaveWithDelimiter(UserPath+Sep+ConfigFile, "=")
}

func SetConfigShowGrid(s string) {
	log.Printf("Setting showgrid to: %s\n", s)
	Config, _ := configparser.Parse(UserPath + Sep + ConfigFile)
	Config.Set("canvas", "showgrid", s)
	Config.SaveWithDelimiter(UserPath+Sep+ConfigFile, "=")
}

func SetConfigScale(s string) {
	log.Printf("Setting scale to: %s\n", s)
	Config, _ := configparser.Parse(UserPath + Sep + ConfigFile)
	Config.Set("canvas", "scale", s)
	Config.SaveWithDelimiter(UserPath+Sep+ConfigFile, "=")
}

func SetConfigShowAll(s string) {
	log.Printf("Setting showall to: %s\n", s)
	Config, _ := configparser.Parse(UserPath + Sep + ConfigFile)
	Config.Set("beads", "showall", s)
	Config.SaveWithDelimiter(UserPath+Sep+ConfigFile, "=")
}

func SetConfigGreyscale(s string) {
	log.Printf("Setting greyscale to: %s\n", s)
	Config, _ := configparser.Parse(UserPath + Sep + ConfigFile)
	Config.Set("beads", "greyscale", s)
	Config.SaveWithDelimiter(UserPath+Sep+ConfigFile, "=")
}

func SetConfigInStock(s string) {
	log.Printf("Setting instock to: %s\n", s)
	Config, _ := configparser.Parse(UserPath + Sep + ConfigFile)
	Config.Set("beads", "instock", s)
	Config.SaveWithDelimiter(UserPath+Sep+ConfigFile, "=")
}

func ReadConfig() {
	log.Printf("Reading config file: %s\n", ConfigFile)
	Config, _ = configparser.Parse(UserPath + Sep + ConfigFile)
	ConfigBrand, _ = Config.Get("pallette", "brand")
	ConfigSerie, _ = Config.Get("pallette", "serie")
	ConfigPegboard, _ = Config.Get("pallette", "pegboard")
	ConfigScale, _ = Config.Get("canvas", "scale")
	ConfigShowGrid, _ = Config.Get("canvas", "showgrid")
	ConfogGridColor, _ = Config.Get("canvas", "gridcolor")
	ConfigShowBeads, _ = Config.Get("canvas", "showbeads")
	ConfigBackgroundColor, _ = Config.Get("canvas", "backgroundcolor")
	ConfigShowAll, _ = Config.Get("beads", "showall")
	ConfigGreyscale, _ = Config.Get("beads", "greyscale")
	ConfigInStock, _ = Config.Get("beads", "instock")
}

func CheckConfigFile() bool {
	log.Printf("Checking for config file: %s\n", ConfigFile)
	if _, err := os.Stat(UserPath + Sep + ConfigFile); os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateDefaultConfig() {
	log.Printf("Creating default config file: %s\n", ConfigFile)
	os.Create(UserPath + Sep + ConfigFile)
	Config = configparser.New()
	Config.AddSection("pallette")
	Config.Set("pallette", "brand", "Hama")
	Config.Set("pallette", "serie", "Midi")
	Config.Set("pallette", "pegboard", "Square 29x29")
	Config.AddSection("canvas")
	Config.Set("canvas", "scale", "100")
	Config.Set("canvas", "showgrid", "true")
	Config.Set("canvas", "gridcolor", "#00ff00")
	Config.Set("canvas", "showbeads", "false")
	Config.Set("canvas", "backgroundcolor", "#ffffff")
	Config.AddSection("beads")
	Config.Set("beads", "showall", "true")
	Config.Set("beads", "greyscale", "false")
	Config.Set("beads", "instock", "false")
	Config.SaveWithDelimiter(UserPath+Sep+ConfigFile, "=")
}
