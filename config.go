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
	Config.Set("pallette", "pegboard", "Large 29x29")
	Config.AddSection("canvas")
	Config.Set("canvas", "scale", "100")
	Config.Set("canvas", "showgrid", "true")
	Config.Set("canvas", "gridcolor", "#00ff00")
	Config.Set("canvas", "showbeads", "false")
	Config.Set("canvas", "backgroundcolor", "#ffffff")
	Config.SaveWithDelimiter(UserPath+Sep+ConfigFile, "=")
}

func SetConfigScale(s string) {
	log.Printf("Setting scale to: %s\n", s)
	Config, _ := configparser.Parse(UserPath + Sep + ConfigFile)
	Config.Set("canvas", "scale", s)
	Config.SaveWithDelimiter(UserPath+Sep+ConfigFile, "=")
}
