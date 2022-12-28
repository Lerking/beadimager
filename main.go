package main

import (
	"log"
	"math"
	"os/user"

	"github.com/lxn/walk"
	"github.com/otiai10/copy"
)

type MyMainWindow struct {
	MainWindow     *walk.MainWindow
	content        *walk.Composite
	leftPanel      *walk.Composite
	rightPanel     *walk.Composite
	colors         *walk.ScrollView
	canvasScroll   *walk.ScrollView
	canvas         *Canvas
	drawWidget     *walk.CustomWidget
	propScroll     *walk.ScrollView
	pallette       Pallette
	beads          []*BeadColor
	brand_combo    *walk.ComboBox
	brand_model    []string
	serie_combo    *walk.ComboBox
	serie_model    []string
	pegboard_combo *walk.ComboBox
	Pegboards      *Pegboards
	properties     *properties
}

const (
	AppName   string = "BeadImager"
	Version   string = "0.4.4"
	CopyRight string = "©2022 Jan Lerking"
	STD_MESS  string = "Ready"
	LogFile   string = "BeadImager.log"
	Sep       string = "\\"
)

var (
	UserPath              string
	ConfigBrand           string
	ConfigSerie           string
	ConfigPegboard        string
	ConfigScale           string
	ConfigShowGrid        string
	ConfogGridColor       string
	ConfigShowBeads       string
	ConfigBackgroundColor string
)

func SetupMainWindow(mw *MyMainWindow) {
	mw.MainWindow, _ = walk.NewMainWindow()
	mw.MainWindow.SetTitle(AppName + " " + Version)
	mw.MainWindow.SetMinMaxSize(walk.Size{Width: 800, Height: 600}, walk.Size{Width: math.MaxInt32, Height: math.MaxInt32})
	SetupMenu(mw)
	//mw.MainWindow.SetSize(walk.Size{Width: 800, Height: 600})
	vb := walk.NewVBoxLayout()
	mw.MainWindow.SetLayout(vb)
	vb.SetMargins(walk.Margins{5, 0, 5, 5})
	mw.content, _ = walk.NewComposite(mw.MainWindow)
	hb := walk.NewHBoxLayout()
	mw.content.SetLayout(hb)
	hb.SetMargins(walk.Margins{0, 0, 0, 0})
}

func main() {
	// Get current user
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	homeDir := currentUser.HomeDir
	UserPath = homeDir + Sep + "BeadImager"
	InitLogFile()
	if !CheckConfigFile() {
		CreateDefaultConfig()
		ReadConfig()
		log.Println("Config file created")
	} else {
		ReadConfig()
		log.Println("Brand: ", ConfigBrand)
		log.Println("Serie: ", ConfigSerie)
		log.Println("Pegboard: ", ConfigPegboard)
	}
	if !CheckPalletteFile() {
		// Copy files from network directory to local directory
		err := copy.Copy("pallettes\\pallette.xml", UserPath+Sep+"pallette.xml")
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Pallette file created")
	}
	walk.AppendToWalkInit(func() {
		walk.FocusEffect, _ = walk.NewBorderGlowEffect(walk.RGB(0, 63, 255))
		//walk.InteractionEffect, _ = walk.NewDropShadowEffect(walk.RGB(63, 63, 63))
		walk.ValidationErrorEffect, _ = walk.NewBorderGlowEffect(walk.RGB(255, 0, 0))
	})
	mw := &MyMainWindow{}
	SetupMainWindow(mw)
	log.Println("MainWindow created")
	CreatePallette(mw)
	mw.brand_model = CreateBrandsList(mw)
	CreatePalletteGroup(mw)
	CreateBeadsGroup(mw)
	CreateCanvasGroup(mw)
	CreateProperties(mw)

	mw.MainWindow.Show()
	mw.MainWindow.Run()
}

func (mv *MyMainWindow) clearBackground(canvas *walk.Canvas, updateBounds walk.Rectangle) error {
	brush, err := walk.NewSolidColorBrush(walk.RGB(0, 0, 0))
	if err != nil {
		return err
	}
	defer brush.Dispose()

	return canvas.FillRectangle(brush, updateBounds)
}

func (mw *MyMainWindow) drawStuff(canvas *walk.Canvas, updateBounds walk.Rectangle) error {
	bmp, err := createBitmap()
	if err != nil {
		return err
	}
	defer bmp.Dispose()

	bounds := mw.drawWidget.ClientBounds()

	rectPen, err := walk.NewCosmeticPen(walk.PenSolid, walk.RGB(255, 0, 0))
	if err != nil {
		return err
	}
	defer rectPen.Dispose()

	if err := canvas.DrawRectangle(rectPen, bounds); err != nil {
		return err
	}

	ellipseBrush, err := walk.NewHatchBrush(walk.RGB(0, 255, 0), walk.HatchCross)
	if err != nil {
		return err
	}
	defer ellipseBrush.Dispose()

	if err := canvas.FillEllipse(ellipseBrush, bounds); err != nil {
		return err
	}

	linesBrush, err := walk.NewSolidColorBrush(walk.RGB(0, 0, 255))
	if err != nil {
		return err
	}
	defer linesBrush.Dispose()

	linesPen, err := walk.NewGeometricPen(walk.PenDash, 8, linesBrush)
	if err != nil {
		return err
	}
	defer linesPen.Dispose()

	if err := canvas.DrawLine(linesPen, walk.Point{bounds.X, bounds.Y}, walk.Point{bounds.Width, bounds.Height}); err != nil {
		return err
	}
	if err := canvas.DrawLine(linesPen, walk.Point{bounds.X, bounds.Height}, walk.Point{bounds.Width, bounds.Y}); err != nil {
		return err
	}

	points := make([]walk.Point, 10)
	dx := bounds.Width / (len(points) - 1)
	for i := range points {
		points[i].X = i * dx
		points[i].Y = int(float64(bounds.Height) / math.Pow(float64(bounds.Width/2), 2) * math.Pow(float64(i*dx-bounds.Width/2), 2))
	}
	if err := canvas.DrawPolyline(linesPen, points); err != nil {
		return err
	}

	bmpSize := bmp.Size()
	if err := canvas.DrawImage(bmp, walk.Point{(bounds.Width - bmpSize.Width) / 2, (bounds.Height - bmpSize.Height) / 2}); err != nil {
		return err
	}

	return nil
}

func createBitmap() (*walk.Bitmap, error) {
	bounds := walk.Rectangle{Width: 200, Height: 200}

	bmp, err := walk.NewBitmap(bounds.Size())
	if err != nil {
		return nil, err
	}

	succeeded := false
	defer func() {
		if !succeeded {
			bmp.Dispose()
		}
	}()

	canvas, err := walk.NewCanvasFromImage(bmp)
	if err != nil {
		return nil, err
	}
	defer canvas.Dispose()

	brushBmp, err := walk.NewBitmapFromFile("images/plus.png")
	if err != nil {
		return nil, err
	}
	defer brushBmp.Dispose()

	brush, err := walk.NewBitmapBrush(brushBmp)
	if err != nil {
		return nil, err
	}
	defer brush.Dispose()

	if err := canvas.FillRectangle(brush, bounds); err != nil {
		return nil, err
	}

	font, err := walk.NewFont("Times New Roman", 40, walk.FontBold|walk.FontItalic)
	if err != nil {
		return nil, err
	}
	defer font.Dispose()

	if err := canvas.DrawText("Walk Drawing Example", font, walk.RGB(0, 0, 0), bounds, walk.TextWordbreak); err != nil {
		return nil, err
	}

	succeeded = true

	return bmp, nil
}
