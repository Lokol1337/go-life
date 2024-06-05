package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

const width = 650

var scale int = 10

const height = 650

var blue color.Color = color.RGBA{69, 145, 196, 255}
var yellow color.Color = color.RGBA{255, 230, 120, 255}
var initGrid []cell
var shiftX, shiftY = 0, 0
var start_work = false
var duration = 1.0
var counter = 0
var counter2 = 0

type makeContext struct {
	is_ bool
	x   int
	y   int
}

var context = makeContext{is_: false, x: 0, y: 0}

func makeDart(x int, y int) {
	initGrid = append(initGrid, cell{x: x, y: y})

	initGrid = append(initGrid, cell{x: x + 1, y: y + 1})
	initGrid = append(initGrid, cell{x: x - 1, y: y + 1})

	initGrid = append(initGrid, cell{x: x + 2, y: y + 2})
	initGrid = append(initGrid, cell{x: x - 2, y: y + 2})

	initGrid = append(initGrid, cell{x: x, y: y + 3})
	initGrid = append(initGrid, cell{x: x + 1, y: y + 3})
	initGrid = append(initGrid, cell{x: x - 1, y: y + 3})

	initGrid = append(initGrid, cell{x: x + 2, y: y + 5})
	initGrid = append(initGrid, cell{x: x - 2, y: y + 5})
	initGrid = append(initGrid, cell{x: x + 3, y: y + 5})
	initGrid = append(initGrid, cell{x: x - 3, y: y + 5})

	initGrid = append(initGrid, cell{x: x + 1, y: y + 6})
	initGrid = append(initGrid, cell{x: x - 1, y: y + 6})
	initGrid = append(initGrid, cell{x: x + 5, y: y + 6})
	initGrid = append(initGrid, cell{x: x - 5, y: y + 6})

	initGrid = append(initGrid, cell{x: x + 1, y: y + 7})
	initGrid = append(initGrid, cell{x: x - 1, y: y + 7})
	initGrid = append(initGrid, cell{x: x + 5, y: y + 7})
	initGrid = append(initGrid, cell{x: x - 5, y: y + 7})
	initGrid = append(initGrid, cell{x: x + 6, y: y + 7})
	initGrid = append(initGrid, cell{x: x - 6, y: y + 7})

	initGrid = append(initGrid, cell{x: x + 1, y: y + 8})
	initGrid = append(initGrid, cell{x: x - 1, y: y + 8})
	initGrid = append(initGrid, cell{x: x + 7, y: y + 8})
	initGrid = append(initGrid, cell{x: x - 7, y: y + 8})

	initGrid = append(initGrid, cell{x: x + 1, y: y + 9})
	initGrid = append(initGrid, cell{x: x - 1, y: y + 9})
	initGrid = append(initGrid, cell{x: x + 3, y: y + 9})
	initGrid = append(initGrid, cell{x: x - 3, y: y + 9})
	initGrid = append(initGrid, cell{x: x + 4, y: y + 9})
	initGrid = append(initGrid, cell{x: x - 4, y: y + 9})
	initGrid = append(initGrid, cell{x: x - 6, y: y + 9})
	initGrid = append(initGrid, cell{x: x + 6, y: y + 9})

}

func makeClock(x int, y int) {
	initGrid = append(initGrid, cell{x: x, y: y})
	initGrid = append(initGrid, cell{x: x + 1, y: y + 1})
	initGrid = append(initGrid, cell{x: x + 2, y: y + 1})
	initGrid = append(initGrid, cell{x: x, y: y + 2})
	initGrid = append(initGrid, cell{x: x - 1, y: y + 2})
	initGrid = append(initGrid, cell{x: x + 1, y: y + 3})
}

func makeArrow(x int, y int) {
	initGrid = append(initGrid, cell{x: x, y: y})
	initGrid = append(initGrid, cell{x: x + 1, y: y + 1})
	initGrid = append(initGrid, cell{x: x - 1, y: y + 1})
	initGrid = append(initGrid, cell{x: x, y: y + 2})

	initGrid = append(initGrid, cell{x: x, y: y + 3})
	initGrid = append(initGrid, cell{x: x + 2, y: y + 3})
	initGrid = append(initGrid, cell{x: x + 3, y: y + 3})
	initGrid = append(initGrid, cell{x: x - 2, y: y + 3})
	initGrid = append(initGrid, cell{x: x - 3, y: y + 3})

	initGrid = append(initGrid, cell{x: x, y: y + 4})

	initGrid = append(initGrid, cell{x: x, y: y + 5})
	initGrid = append(initGrid, cell{x: x + 2, y: y + 5})
	initGrid = append(initGrid, cell{x: x - 2, y: y + 5})

	initGrid = append(initGrid, cell{x: x, y: y + 6})
}

func makeGosperGun(x int, y int) {
	initGrid = append(initGrid, cell{x: x, y: y})
	initGrid = append(initGrid, cell{x: x - 1, y: y})
	initGrid = append(initGrid, cell{x: x - 1, y: y + 1})
	initGrid = append(initGrid, cell{x: x - 2, y: y + 2})
	initGrid = append(initGrid, cell{x: x - 1, y: y - 1})
	initGrid = append(initGrid, cell{x: x - 2, y: y - 2})
	initGrid = append(initGrid, cell{x: x - 3, y: y})

	initGrid = append(initGrid, cell{x: x - 4, y: y - 3})
	initGrid = append(initGrid, cell{x: x - 4, y: y + 3})
	initGrid = append(initGrid, cell{x: x - 5, y: y - 3})
	initGrid = append(initGrid, cell{x: x - 5, y: y + 3})
	initGrid = append(initGrid, cell{x: x - 6, y: y - 2})
	initGrid = append(initGrid, cell{x: x - 6, y: y + 2})
	initGrid = append(initGrid, cell{x: x - 7, y: y - 1})
	initGrid = append(initGrid, cell{x: x - 7, y: y + 1})
	initGrid = append(initGrid, cell{x: x - 7, y: y})

	initGrid = append(initGrid, cell{x: x - 16, y: y})
	initGrid = append(initGrid, cell{x: x - 16, y: y - 1})
	initGrid = append(initGrid, cell{x: x - 17, y: y})
	initGrid = append(initGrid, cell{x: x - 17, y: y - 1})

	initGrid = append(initGrid, cell{x: x + 3, y: y - 1})
	initGrid = append(initGrid, cell{x: x + 3, y: y - 2})
	initGrid = append(initGrid, cell{x: x + 3, y: y - 3})
	initGrid = append(initGrid, cell{x: x + 4, y: y - 1})
	initGrid = append(initGrid, cell{x: x + 4, y: y - 2})
	initGrid = append(initGrid, cell{x: x + 4, y: y - 3})

	initGrid = append(initGrid, cell{x: x + 5, y: y - 4})
	initGrid = append(initGrid, cell{x: x + 5, y: y})
	initGrid = append(initGrid, cell{x: x + 7, y: y - 4})
	initGrid = append(initGrid, cell{x: x + 7, y: y})
	initGrid = append(initGrid, cell{x: x + 7, y: y - 5})
	initGrid = append(initGrid, cell{x: x + 7, y: y + 1})

	initGrid = append(initGrid, cell{x: x + 17, y: y - 2})
	initGrid = append(initGrid, cell{x: x + 17, y: y - 3})
	initGrid = append(initGrid, cell{x: x + 18, y: y - 2})
	initGrid = append(initGrid, cell{x: x + 18, y: y - 3})

}

// func removeDuplicates(cells []cell) []cell {
// 	mapping := make(map[string]bool)
// 	var result []cell

// 	for _, c := range cells {
// 		key := fmt.Sprintf("%d:%d", c.x, c.y)
// 		if _, ok := mapping[key]; !ok {
// 			mapping[key] = true
// 			result = append(result, c)
// 		}
// 	}

// 	return result
// 	// if len(cells) < 2 {
// 	// 	return cells
// 	// }

// 	// head := cells[0]
// 	// tail := cells[1:]

// 	// if head == tail[0] {
// 	// 	return removeDuplicates(tail)
// 	// }

// 	// return append(removeDuplicates(tail), head)
// }

// func displayField(window *ebiten.Image) error {
// 	window.Fill(blue)
// 	for _, element := range initGrid {
// 		for i := 0; i < scale; i++ {
// 			for j := 0; j < scale; j++ {
// 				window.Set((element.x)*scale+i+shiftX, (element.y)*scale+j+shiftY, yellow)
// 			}
// 		}
// 	}
// 	if start_work {
// 		if duration == 0.09 || counter == int(duration*60) {
// 			initGrid = removeDuplicates_(updateField(initGrid))
// 			counter = 0
// 		}
// 	}
// 	counter += 1
// 	curSpeed := ""
// 	switch duration {
// 	case 1:
// 		curSpeed = "1 - 1s"
// 	case 0.5:
// 		curSpeed = "2 - 0.5s"
// 	case 0.1:
// 		curSpeed = "3 - 0.1s"
// 	case 0.09:
// 		curSpeed = "4 - draw delay"
// 	}
// 	cur_scale := fmt.Sprint((scale * 10)) + "%"
// 	tutorial := fmt.Sprintf("Space: Play/Pause\nEsc - Clean \nArrow to move\n1, 2, 3, 4 - update delay speed\nCur speed: %s \n'+' - Zoom up\n'-' - Zoom down\n Zoom: %s \n\nHOTKEYS: \n D - Dart \n G - Gosper glider gun \n C - Clock \n A - Arrow", curSpeed, cur_scale)
// 	msg := fmt.Sprintf("%s", tutorial)
// 	ebitenutil.DebugPrint(window, msg)

// 	if ebiten.IsKeyPressed(ebiten.KeySpace) && inpututil.KeyPressDuration(ebiten.KeySpace) == 1 {
// 		if start_work {
// 			start_work = false
// 		} else {
// 			start_work = true
// 		}
// 	}

// 	if ebiten.IsKeyPressed(ebiten.KeyMinus) && inpututil.KeyPressDuration(ebiten.KeyMinus) == 1 && scale != 5 {
// 		scale -= 1
// 	}

// 	if ebiten.IsKeyPressed(ebiten.KeyEqual) && inpututil.KeyPressDuration(ebiten.KeyEqual) == 1 && scale != 15 {
// 		scale += 1
// 	}

// 	if ebiten.IsKeyPressed(ebiten.Key1) && inpututil.KeyPressDuration(ebiten.Key1) == 1 {
// 		duration = 1
// 		counter = 0
// 	}

// 	if ebiten.IsKeyPressed(ebiten.Key2) && inpututil.KeyPressDuration(ebiten.Key2) == 1 {
// 		duration = 0.5
// 		counter = 0
// 	}

// 	if ebiten.IsKeyPressed(ebiten.Key3) && inpututil.KeyPressDuration(ebiten.Key3) == 1 {
// 		duration = 0.1
// 		counter = 0
// 	}

// 	if ebiten.IsKeyPressed(ebiten.Key4) && inpututil.KeyPressDuration(ebiten.Key4) == 1 {
// 		duration = 0.09
// 		counter = 0
// 	}

// 	if ebiten.IsKeyPressed(ebiten.KeyEscape) && inpututil.KeyPressDuration(ebiten.KeyEscape) == 1 {
// 		initGrid = []cell{}
// 	}

// 	if ebiten.IsKeyPressed(ebiten.KeyG) && inpututil.KeyPressDuration(ebiten.KeyG) == 1 {
// 		newCell := cell{}
// 		newCell.x, newCell.y = ebiten.CursorPosition()
// 		newCell.x = (newCell.x - shiftX) / scale
// 		newCell.y = (newCell.y - shiftY) / scale
// 		makeGosperGun(newCell.x, newCell.y)
// 	}

// 	if ebiten.IsKeyPressed(ebiten.KeyA) && inpututil.KeyPressDuration(ebiten.KeyA) == 1 {
// 		newCell := cell{}
// 		newCell.x, newCell.y = ebiten.CursorPosition()
// 		newCell.x = (newCell.x - shiftX) / scale
// 		newCell.y = (newCell.y - shiftY) / scale
// 		makeArrow(newCell.x, newCell.y)
// 	}

// 	if ebiten.IsKeyPressed(ebiten.KeyC) && inpututil.KeyPressDuration(ebiten.KeyC) == 1 {
// 		newCell := cell{}
// 		newCell.x, newCell.y = ebiten.CursorPosition()
// 		newCell.x = (newCell.x - shiftX) / scale
// 		newCell.y = (newCell.y - shiftY) / scale
// 		makeClock(newCell.x, newCell.y)
// 	}

// 	if ebiten.IsKeyPressed(ebiten.KeyD) && inpututil.KeyPressDuration(ebiten.KeyD) == 1 {
// 		newCell := cell{}
// 		newCell.x, newCell.y = ebiten.CursorPosition()
// 		newCell.x = (newCell.x - shiftX) / scale
// 		newCell.y = (newCell.y - shiftY) / scale
// 		makeDart(newCell.x, newCell.y)
// 	}

// 	if ebiten.IsKeyPressed(ebiten.KeyRight) {
// 		shiftX = shiftX - int(1/(duration))
// 	}
// 	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
// 		shiftX = shiftX + int(1/(duration))
// 	}

// 	if ebiten.IsKeyPressed(ebiten.KeyUp) {
// 		shiftY = shiftY + int(1/(duration))
// 	}
// 	if ebiten.IsKeyPressed(ebiten.KeyDown) {
// 		shiftY = shiftY - int(1/(duration))
// 	}

// 	// }

// 	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft) == 1 {
// 		newCell := cell{}
// 		newCell.x, newCell.y = ebiten.CursorPosition()
// 		newCell.x = (newCell.x - shiftX) / scale
// 		newCell.y = (newCell.y - shiftY) / scale
// 		initGrid = append(initGrid, newCell)
// 	}

// 	// if !ebiten.IsDrawingSkipped() {

// 	// }
// 	return nil
// }

// func sendDataOnWindow(data []cell) {
// 	initGrid = data
// 	updateField(initGrid)
// }

// func makeWindow() {
// 	ebiten.Run(displayField, width, height, 2, "Game of Life")
// }

// func main() {
// 	makeWindow()
// }

// Game implements ebiten.Game interface.
type Game struct{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update(screen *ebiten.Image) error {

	if ebiten.IsKeyPressed(ebiten.KeySpace) && inpututil.KeyPressDuration(ebiten.KeySpace) == 1 {
		if start_work {
			start_work = false
		} else {
			start_work = true
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyMinus) && inpututil.KeyPressDuration(ebiten.KeyMinus) == 1 && scale != 5 {
		scale -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyEqual) && inpututil.KeyPressDuration(ebiten.KeyEqual) == 1 && scale != 15 {
		scale += 1
	}

	if ebiten.IsKeyPressed(ebiten.Key1) && inpututil.KeyPressDuration(ebiten.Key1) == 1 {
		duration = 1
		counter = 0
	}

	if ebiten.IsKeyPressed(ebiten.Key2) && inpututil.KeyPressDuration(ebiten.Key2) == 1 {
		duration = 0.5
		counter = 0
	}

	if ebiten.IsKeyPressed(ebiten.Key3) && inpututil.KeyPressDuration(ebiten.Key3) == 1 {
		duration = 0.1
		counter = 0
	}

	if ebiten.IsKeyPressed(ebiten.Key4) && inpututil.KeyPressDuration(ebiten.Key4) == 1 {
		duration = 0.09
		counter = 0
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) && inpututil.KeyPressDuration(ebiten.KeyEscape) == 1 {
		initGrid = []cell{}
	}

	if ebiten.IsKeyPressed(ebiten.KeyG) && inpututil.KeyPressDuration(ebiten.KeyG) == 1 {
		newCell := cell{}
		newCell.x, newCell.y = ebiten.CursorPosition()
		newCell.x = (newCell.x - shiftX) / scale
		newCell.y = (newCell.y - shiftY) / scale
		makeGosperGun(newCell.x, newCell.y)
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) && inpututil.KeyPressDuration(ebiten.KeyA) == 1 {
		newCell := cell{}
		newCell.x, newCell.y = ebiten.CursorPosition()
		newCell.x = (newCell.x - shiftX) / scale
		newCell.y = (newCell.y - shiftY) / scale
		makeArrow(newCell.x, newCell.y)
	}

	if ebiten.IsKeyPressed(ebiten.KeyC) && inpututil.KeyPressDuration(ebiten.KeyC) == 1 {
		newCell := cell{}
		newCell.x, newCell.y = ebiten.CursorPosition()
		newCell.x = (newCell.x - shiftX) / scale
		newCell.y = (newCell.y - shiftY) / scale
		makeClock(newCell.x, newCell.y)
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) && inpututil.KeyPressDuration(ebiten.KeyD) == 1 {
		newCell := cell{}
		newCell.x, newCell.y = ebiten.CursorPosition()
		newCell.x = (newCell.x - shiftX) / scale
		newCell.y = (newCell.y - shiftY) / scale
		makeDart(newCell.x, newCell.y)
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		shiftX = shiftX - int(1/(duration))
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		shiftX = shiftX + int(1/(duration))
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		shiftY = shiftY + int(1/(duration))
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		shiftY = shiftY - int(1/(duration))
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft) == 1 {
		newCell := cell{}
		newCell.x, newCell.y = ebiten.CursorPosition()
		newCell.x = (newCell.x - shiftX) / scale
		newCell.y = (newCell.y - shiftY) / scale
		initGrid = append(initGrid, newCell)
	}
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(blue)
	curSpeed := ""
	switch duration {
	case 1:
		curSpeed = "1 - 1s"
	case 0.5:
		curSpeed = "2 - 0.5s"
	case 0.1:
		curSpeed = "3 - 0.1s"
	case 0.09:
		curSpeed = "4 - draw delay"
	}
	cur_scale := fmt.Sprint((scale * 10)) + "%"
	tutorial := fmt.Sprintf("Space: Play/Pause\nEsc - Clean \nArrow to move\n1, 2, 3, 4 - update delay speed\nCur speed: %s \n'+' - Zoom up\n'-' - Zoom down\n Zoom: %s \n\nHOTKEYS: \n D - Dart \n G - Gosper glider gun \n C - Clock \n A - Arrow", curSpeed, cur_scale)
	ebitenutil.DebugPrint(screen, tutorial)
	for _, element := range initGrid {
		for i := 0; i < scale; i++ {
			for j := 0; j < scale; j++ {
				screen.Set((element.x)*scale+i+shiftX, (element.y)*scale+j+shiftY, yellow)
			}
		}
	}
	if start_work {
		if duration == 0.09 || counter == int(duration*60) {
			initGrid = removeDuplicates_(updateField(initGrid))
			counter = 0
		}
	}
	counter += 1
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func main() {
	game := &Game{}
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("go-life")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
