package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

const (
	gridWidth  = 20
	gridHeight = 20
	cellSize   = 20
)

func main() {
	a := app.New()
	w := a.NewWindow("Game of Life")

	// 2D slice of rectangles
	grid := make([][]*canvas.Rectangle, gridHeight)
	objects := []fyne.CanvasObject{}

	for y := range gridHeight {
		grid[y] = make([]*canvas.Rectangle, gridWidth)
		for x := range gridWidth {
			rect := canvas.NewRectangle(color.White)
			rect.Move(fyne.NewPos(float32(x*cellSize), float32(y*cellSize)))
			rect.Resize(fyne.NewSize(cellSize, cellSize))

			// Make some cells alive for demo
			if (x+y)%2 == 0 {
				rect.FillColor = color.Black
			}

			grid[y][x] = rect
			objects = append(objects, rect)
		}
	}

	go func() {
		x := 0
		y := 0
		for range time.Tick(time.Second / 9) {
			fyne.Do(func() {
				if x >= gridWidth {
					x = 0
					y++
				}
				grid[y][x].FillColor = color.RGBA{255, 0, 0, 255}
				grid[y][x].Refresh()
				x += 1
			})
		}
	}()

	c := container.NewWithoutLayout(objects...)
	window_size := fyne.NewSize(gridWidth*cellSize, gridHeight*cellSize)
	c.Resize(window_size)
	w.Resize(window_size)
	w.SetFixedSize(true)
	w.SetContent(c)

	w.ShowAndRun()
}
