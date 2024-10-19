package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 320
	screenHeight = 240
	cellSize     = 3
	gridRows     = screenHeight / cellSize
	gridCols     = screenWidth / cellSize
)

type Game struct {
	mouseX, mouseY int
	grid           [320][240]int
}

func (g *Game) Update() error {
	g.mouseX, g.mouseY = ebiten.CursorPosition()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Falling Sand!")
	mouseCol := g.mouseX / cellSize
	mouseRow := g.mouseY / cellSize
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.grid[mouseCol][mouseRow] = 1
	}
	//reseting thr grid on keyboard press of r
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		for x := 0; x < 320; x++ {
			for y := 0; y < 240; y++ {
				g.grid[x][y] = 0
			}
		}
	}
	for y := gridRows - 1; y >= 0; y-- {
		for x := 0; x < gridCols; x++ {
			if g.grid[x][y] == 1 && y < gridRows-1 {
				if g.grid[x][y+1] == 0 {
					g.grid[x][y] = 0
					g.grid[x][y+1] = 1
				} else if x > 0 && g.grid[x-1][y+1] == 0 {
					g.grid[x][y] = 0
					g.grid[x-1][y+1] = 1
				} else if x < gridCols-1 && g.grid[x+1][y+1] == 0 {
					g.grid[x][y] = 0
					g.grid[x+1][y+1] = 1
				}
			}
		}
	}
	drawGrid(g, screen)

	drawCircle(screen, float64(g.mouseX), float64(g.mouseY), 5, color.RGBA{0xff, 0, 0, 0xff})

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
func drawCircle(img *ebiten.Image, x, y, r float64, clr color.Color) {
	for i := 0; i < 360; i++ {
		theta := float64(i) * 3.14159 / 180
		xx := x + r*math.Cos(theta)
		yy := y + r*math.Sin(theta)
		img.Set(int(xx), int(yy), clr)
	}
}
func drawGrid(g *Game, screen *ebiten.Image) {
	for y := 0; y < gridRows; y++ {
		for x := 0; x < gridCols; x++ {
			if g.grid[x][y] == 1 {
				// ebitenutil.DrawRect(screen, float64(x*cellSize), float64(y*cellSize), cellSize, cellSize, color.White)
				vector.DrawFilledRect(screen, float32(x*cellSize), float32(y*cellSize), cellSize, cellSize, color.White, true)
			} else {
				//vector.StrokeRect(screen, float32(x*cellSize), float32(y*cellSize), cellSize, 100, 1, color.White, false)
			}
		}
	}
}
