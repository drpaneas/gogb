package main

import (
	"fmt"
	"image/color"
	"github.com/drpaneas/gogb/internal/gameboy"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{
	gb *gameboy.Gameboy
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func (g *Game) Update() error {
	g.gb.Step()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	palette := g.gb.MMU.Read(0xFF47)
	fmt.Sprintf("%d",palette)



	switch(palette){
	case 0:
		screen.Fill(color.RGBA{0xff, 0, 0, 0xff})
	case 1:
		screen.Fill(color.RGBA{0, 0xff, 0, 0xff})
	case 2:
		screen.Fill(color.RGBA{0, 0, 0xff, 0xff})
	case 3:
		screen.Fill(color.RGBA{0, 0, 0xff, 0xff})
	case 4:
		panic(fmt.Sprintf("Invalid pallete code"))
	}
}
