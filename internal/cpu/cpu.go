import(
	"github.com/drpaneas/gogb/internal/mmu"
)

package cpu

type CPU struct {
	A, F byte //tutorial said uint8 , you can use byte for type alias consistency
	B, C byte
	D, E byte
	H, L byte
	PC uint16
	SP uint16
	Halted bool
	MMU mmu.MMU
}


func (cpu *CPU) BC()uint16{
	return uint16(cpu.B)<<8 | uint16(cpu.C)
}

func (cpu *CPU) DE()uint16{
	return uint16(cpu.D)<<8 | uint16(cpu.E)
}

func (cpu *CPU) HL()uint16{
	return uint16(cpu.H)<<8 | uint16(cpu.L)
}

func (cpu *CPU) SetF( Z uint8, N uint8 ,H uint8 , C uint8 )uint8{
	F := [3]uint8{Z,N,H,C}

	for i:=1 ; i <= 4; i++ {

	}


}