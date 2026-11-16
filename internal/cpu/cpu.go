package cpu

import(
	"github.com/drpaneas/gogb/internal/mmu"
)

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


func (cpu *CPU) SetZ(param uint8)uint8{
	cpu.F = cpu.F &^ 0x80
	mask := uint8(param << 7)
	cpu.F = cpu.F | mask
	return cpu.F
}

func (cpu *CPU) SetN(param uint8)uint8{
	cpu.F = cpu.F &^ 0x40
	mask := uint8(param << 6)
	cpu.F = cpu.F | mask
	return cpu.F
}

func (cpu *CPU) SetH(param uint8)uint8{
	cpu.F = cpu.F &^ 0x20
	mask := uint8(param << 5)
	cpu.F = cpu.F | mask
	return cpu.F
}

func (cpu *CPU) SetC(param uint8)uint8{
	cpu.F = cpu.F &^ 0x10
	mask := uint8(param << 4)
	cpu.F = cpu.F | mask
	return cpu.F
}

func (cpu *CPU) ReadC()uint8{
	carry := (cpu.F >> 4) & 1
	return carry
}

func (cpu *CPU) CalZ(param uint8)uint8{
	if param == 0  {
		cpu.F = cpu.F &^ 0x80
		mask := uint8(1 << 7)
		cpu.F = cpu.F | mask
		return cpu.F
	} else {
		cpu.F = cpu.F &^ 0x80
		mask := uint8(0 << 7)
		cpu.F = cpu.F | mask
		return cpu.F
	}
}

func (cpu *CPU) CalH(reg1 uint8 , reg2 uint8, operand string , carry_optional ...uint8)uint8{
	var carry uint8

	if len(carry_optional) > 0 {
		carry = carry_optional[0]
	}
	var res uint8
	if operand == "+" {
		res = (reg1 &^ 0xF0) + (reg2 &^ 0xF0) + carry
	} else if operand == "-" {
		res = (reg1 &^ 0xF0) - (reg2 &^ 0xF0) - carry
	}

	if res > 0x0F {
		cpu.F = cpu.F &^ 0x20
		mask := uint8(1 << 5)
		cpu.F = cpu.F | mask
		return cpu.F
	} else {
		cpu.F = cpu.F &^ 0x20
		mask := uint8(0 << 5)
		cpu.F = cpu.F | mask
		return cpu.F
	}
	

}

func (cpu *CPU) CalC(param uint16)uint8{
	if param > 0xFF {
		cpu.F = cpu.F &^ 0x10
		mask := uint8(1 << 4)
		cpu.F = cpu.F | mask
		return cpu.F
	} else {
		cpu.F = cpu.F &^ 0x10
		mask := uint8(0 << 4)
		cpu.F = cpu.F | mask
		return cpu.F
	}
}