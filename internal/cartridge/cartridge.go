package cartridge 


import(
	"os"
	"strings"
)

type Cartridge struct {
	RomMemory [0x20000]byte

}

func (cart *Cartridge) ReadRom(){
	data, err := os.ReadFile("./rom/tetris.gb")
	if err != nil {
		panic(err)
	}
	copy(cart.RomMemory[:], data) 
	return
}


func (cart *Cartridge) GetCartridgeTitle() string{
	romTitle := ""
	for i := 0x0134; i<=0x0143 ;i++ {
		romTitle+= string(cart.RomMemory[i])
	}
	return strings.ToUpper(string(romTitle[0])) + strings.ToLower(romTitle[1:])
}
