package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type MyMainWindow struct {
	*walk.MainWindow
	colors     *walk.ScrollView
	canvas     *walk.ScrollView
	properties *walk.ScrollView
}

const (
	AppName   string = "BeadImager"
	Version   string = "0.0.4"
	CopyRight string = "©2022 Jan Lerking"
	STD_MESS  string = "Ready"
	UserPath  string = "C:\\Users\\janle\\BeadImager"
	LogFile   string = "BeadImager.log"
	Sep       string = "\\"
)

var LoggingFile *os.File

func InitLogFile() {
	_, err := os.Stat(fmt.Sprintf(UserPath + Sep + LogFile))
	if err == nil {
		LoggingFile, err := os.OpenFile(fmt.Sprintf(UserPath+Sep+LogFile), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		log.SetOutput(LoggingFile)
	} else {
		if _, err := os.Stat(UserPath); err != nil {
			os.Mkdir(UserPath, 0755)
		}
		_, err = os.Create(fmt.Sprintf(UserPath + Sep + LogFile))
		if err != nil {
			log.Fatal(err)
		}
		LoggingFile, err := os.OpenFile(fmt.Sprintf(UserPath+Sep+LogFile), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		log.SetOutput(LoggingFile)
	}
	log.Print("Logging initialized")
}

func main() {
	InitLogFile()
	walk.AppendToWalkInit(func() {
		walk.FocusEffect, _ = walk.NewBorderGlowEffect(walk.RGB(0, 63, 255))
		walk.InteractionEffect, _ = walk.NewDropShadowEffect(walk.RGB(63, 63, 63))
		walk.ValidationErrorEffect, _ = walk.NewBorderGlowEffect(walk.RGB(255, 0, 0))
	})
	mw := &MyMainWindow{}
	log.Println("MainWindow created")
	//ss := mw.MainWindow.MaxSize()
	//log.Println(ss)

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
