package main

import (
	"log"
	"fmt"
	"math"
)

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	fmt.Println("START\n")
	mw := new(MyMainWindow)

	if _, err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Test",
		MinSize:  Size{320, 240},
		Size:     Size{800, 600},
		Layout:   VBox{MarginsZero: true},
		Children: []Widget{
			CustomWidget{
				AssignTo:            &mw.paintWidget,
				ClearsBackground:    true,
				InvalidatesOnResize: true,
				Paint:               mw.drawStuff,
			},
		},
	}).Run(); err != nil {
		log.Fatal(err)
	}
}

type MyMainWindow struct {
	*walk.MainWindow
	paintWidget *walk.CustomWidget
}

func (mw *MyMainWindow) drawStuff(canvas *walk.Canvas, updateBounds walk.Rectangle) error {
	bounds := mw.paintWidget.ClientBounds()
	err := drawBgBlack(canvas, bounds)
	if err != nil {
		return err
	}
	err = drawFeature(canvas, bounds)
	if err != nil {
		return err
	}
	return nil
}

func drawBgBlack(canvas *walk.Canvas, bounds walk.Rectangle) error {
	bgBrush, err := walk.NewSolidColorBrush(walk.RGB(0, 0, 0))
	if err != nil {
		return err
	}
	defer bgBrush.Dispose()
	bgPen, err := walk.NewGeometricPen(walk.PenSolid, bounds.Width * 2, bgBrush)
	if err != nil {
		return err
	}
	defer bgPen.Dispose()
	err = canvas.DrawLine(bgPen, walk.Point{bounds.X, bounds.Y}, walk.Point{bounds.X, bounds.Height})
	if err != nil {
		return err
	}
	return nil
}

func drawFeature(canvas *walk.Canvas, bounds walk.Rectangle) error {
	length := 100
	center := walk.Point{bounds.Width / 2 - 50, bounds.Height / 2 - 50}
	inc := 6
	cnt := int(360 / inc)
	for i := 0; i < cnt; i++ {
		rad := math.Pi * float64(i) / float64(180 / inc)
		newPoint := walk.Point{
			center.X + int(math.Cos(rad) * 100),
			center.Y + int(math.Sin(rad) * 100)}
		newRect := walk.Rectangle{
			newPoint.X, newPoint.Y, length, length}
		color := getColor(i * inc)
		linesPen, err := walk.NewCosmeticPen(walk.PenSolid, color)
		if err != nil {
			return err
		}
		err = canvas.DrawEllipse(linesPen, newRect)
		if err != nil {
			return err
		}
	}
	return nil
}

func getColor(angle int) walk.Color {
	red := byte(0)
	green := byte(0)
	blue := byte(0)
	div := int(angle / 60)
	val := byte(angle % 60 * 4)
	switch div {
	case 0:
		red = 240
		green = val
		break
	case 1:
		red = 240 - val
		green = 240
		break
	case 2:
		green = 240
		blue = val
		break
	case 3:
		green = 240 - val
		blue = 240
		break
	case 4:
		blue = 240
		red = val
		break
	default:
		blue = 240 - val
		red = 240
		break
	}
	color := walk.RGB(red, green, blue)
	return color
}
