package main

import "fmt"

type Opcode byte

const (
	NOP     Opcode = 0x00
	JPa16   Opcode = 0xC3
	LDAn8   Opcode = 0x3E
	LDHa8A  Opcode = 0xE0
	HALT    Opcode = 0x76
	UNKNOWN Opcode = 0xFF
)

const KB = 1024

var memory [64 * KB]byte // 64 KB Game Bay Memory

const IORegisters = 0xFF00

type CPU struct {
	A      byte
	PC     uint16
	Halted bool
}

func InitCPU() *CPU {
	cpu := CPU{A: 0x00, PC: 0x0100, Halted: false}

	return &cpu
}

func Step(cpu *CPU) uint {
	opcode := Opcode(memory[cpu.PC])
	cpu.PC++

	switch opcode {
	case NOP:
		cpu.PC++
		return 4
	case JPa16:
		lo := memory[cpu.PC]
		hi := memory[cpu.PC+1]
		cpu.PC = uint16(hi)<<8 | uint16(lo)
		return 16
	case LDAn8:
		cpu.A = memory[cpu.PC]
		cpu.PC++
		return 8
	case LDHa8A:
		offset := memory[cpu.PC]
		memory[IORegisters+uint16(offset)] = byte(cpu.A)
		cpu.PC++
		return 12
	case HALT:
		cpu.Halted = true
		return 4
	default:
		panic(fmt.Sprintf("Unknown opcode: 0x%02X", opcode))
	}
}

func main() {

}
