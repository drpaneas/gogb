package main

import (
	//"log"
	//"github.com/hajimehoshi/ebiten/v2"
	//"github.com/drpaneas/gogb/internal/gameboy"
	//"github.com/drpaneas/gogb/internal/cartridge"
	"github.com/drpaneas/gogb/internal/cpu"
	"fmt"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func main() {

	cpu := &cpu.CPU{}

	flagByte := cpu.SetZ(1);
	fmt.Printf("%08b\n", flagByte)

	flagByte = cpu.SetH(1);
	fmt.Printf("%08b\n", flagByte)

	flagByte = cpu.SetC(1);
	fmt.Printf("%08b\n", flagByte)

	flagByte = cpu.SetZ(0);
	fmt.Printf("%08b\n", flagByte)
	// cart := &cartridge.Cartridge{}
	// cart.ReadRom()
	// romTitle := cart.GetCartridgeTitle()
	
	

	// gb := &gameboy.Gameboy{}
	// gb.MMU.Memory[65351] = 1 //change the memory address which is responsible for pallete manually for now
	
    // game := &Game{
    //     gb: gb,            
    // }

	// w,h := game.Layout(0,0)
	// ebiten.SetWindowSize(w,h)
	// ebiten.SetWindowTitle(romTitle)
	// if err := ebiten.RunGame(game); err != nil {
	// 	log.Fatal(err)
	// }


}
