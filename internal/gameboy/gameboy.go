package gameboy

import (
	"fmt"
	"github.com/drpaneas/gogb/internal/cpu"
	"github.com/drpaneas/gogb/internal/mmu"
)

type Gameboy struct {
	CPU cpu.CPU
	MMU mmu.MMU
}

func (gb *Gameboy) GetBackgroundColor() string {
	flag := 0

	if flag == 0 {
		flag = 1
		return "red"
	} else if flag == 1 {
		flag = 2
		return "blue"
	} else if flag == 2 {
		flag = 0
		return "green"
	}
	return "error"
}

func (gb *Gameboy) Step() int {
	opcode := gb.MMU.Read(gb.CPU.PC)
	gb.CPU.PC++

	switch opcode {

	// --- 0x0X ---
	case 0x00: // NOP
		return 4
	case 0x01: // LD BC, d16
		lo := gb.MMU.Read(gb.CPU.PC)
		hi := gb.MMU.Read(gb.CPU.PC + 1)
		gb.CPU.PC += 2
		gb.CPU.C = lo
		gb.CPU.B = hi
		return 12
	case 0x02: // LD (BC), A
		gb.MMU.Write(gb.CPU.BC(), gb.CPU.A)
		return 8
	case 0x03: // INC BC
		return 8
	case 0x04: // INC B
		return 4
	case 0x05: // DEC B
		return 4
	case 0x06: // LD B, d8
		gb.CPU.B = gb.MMU.Read(gb.CPU.PC)
		gb.CPU.PC++
		return 8
	case 0x07: // RLCA
		return 4
	case 0x08: // LD (a16), SP
		return 20
	case 0x09: // ADD HL, BC
		return 8
	case 0x0A: // LD A, (BC)
		gb.CPU.A = gb.MMU.Read(gb.CPU.BC())
		return 8
	case 0x0B: // DEC BC
		return 8
	case 0x0C: // INC C
		return 4
	case 0x0D: // DEC C
		return 4
	case 0x0E: // LD C, d8
		d8byte := gb.MMU.Read(gb.CPU.PC)
		gb.CPU.C = d8byte
		gb.CPU.PC++
		return 8
	case 0x0F: // RRCA
		return 4

	// --- 0x1X ---
	case 0x10: // STOP
		gb.CPU.Halted = true //not to be confused with halt, will fix later... probably...
		gb.CPU.PC ++
		return 4

	case 0x11: // LD DE, d16
		lo := gb.MMU.Read(gb.CPU.PC)
		hi := gb.MMU.Read(gb.CPU.PC + 1)
		gb.CPU.D = hi
		gb.CPU.E = lo
		gb.CPU.PC += 2
		return 12

	case 0x12: // LD (DE), A
		gb.MMU.Write(gb.CPU.DE(),gb.CPU.A)
		return 8
	case 0x13: // INC DE
		return 8
	case 0x14: // INC D
		return 4
	case 0x15: // DEC D
		return 4
	case 0x16: // LD D, d8
		gb.CPU.D = gb.MMU.Read(gb.CPU.PC)
		gb.CPU.PC++
		return 8
	case 0x17: // RLA
		return 4
	case 0x18: // JR r8
		return 12
	case 0x19: // ADD HL, DE
		return 8
	case 0x1A: // LD A, (DE)
		gb.CPU.A = gb.MMU.Read(gb.CPU.DE())
		return 8
	case 0x1B: // DEC DE
		return 8
	case 0x1C: // INC E
		return 4
	case 0x1D: // DEC E
		return 4
	case 0x1E: // LD E, d8
		gb.CPU.E = gb.MMU.Read(gb.CPU.PC)
		gb.CPU.PC++
		return 8
	case 0x1F: // RRA
		return 4

	// --- 0x2X ---
	case 0x20: // JR NZ, r8
		return 8
	case 0x21: // LD HL, d16
		lo := gb.MMU.Read(gb.CPU.PC)
		hi := gb.MMU.Read(gb.CPU.PC+1)
		gb.CPU.H = hi
		gb.CPU.L = lo
		gb.CPU.PC+=2
		return 12
	case 0x22: // LD (HL+), A
		HL := gb.CPU.HL()
		gb.MMU.Write(HL,gb.CPU.A)
		HL++
		gb.CPU.H = uint8(HL>>8)
		gb.CPU.L = uint8(HL)
		return 8
	case 0x23: // INC HL
		HL := gb.CPU.HL()
		HL++
		gb.CPU.H = uint8(HL>>8)
		gb.CPU.L = uint8(HL)
		return 8
	case 0x24: // INC H FLAGS MISSING MALAKA
		gb.CPU.H++
		return 4
	case 0x25: // DEC H FLAGS MISSING MALAKA
		gb.CPU.H--
		return 4
	case 0x26: // LD H, d8
		gb.CPU.H = gb.MMU.Read(gb.CPU.PC)
		gb.CPU.PC++
		return 8
	case 0x27: // DAA
		return 4
	case 0x28: // JR Z, r8
		return 8
	case 0x29: // ADD HL, HL
		return 8
	case 0x2A: // LD A, (HL+)
		HL:= gb.CPU.HL()
		gb.CPU.A = gb.MMU.Read(HL)
		HL++
		gb.CPU.H = uint8(HL>>8)
		gb.CPU.L = uint8(HL)
		return 8
	case 0x2B: // DEC HL
		return 8
	case 0x2C: // INC L
		return 4
	case 0x2D: // DEC L
		return 4
	case 0x2E: // LD L, d8
		gb.CPU.L = gb.MMU.Read(gb.CPU.PC)
		gb.CPU.PC++
		return 8
	case 0x2F: // CPL
		return 4

	// --- 0x3X ---
	case 0x30: // JR NC, r8
		return 8
	case 0x31: // LD SP, d16
		return 12
	case 0x32: // LD (HL-), A
		HL := gb.CPU.HL()
		gb.MMU.Write( HL , gb.CPU.A)
		HL--
		gb.CPU.H = uint8(HL>>8)
		gb.CPU.L = uint8(HL)
		return 8
	case 0x33: // INC SP
		return 8
	case 0x34: // INC (HL)
		return 12
	case 0x35: // DEC (HL)
		return 12
	case 0x36: // LD (HL), d8
		gb.MMU.Write(gb.CPU.HL(), gb.MMU.Read(gb.CPU.PC))
		gb.CPU.PC++
		return 12
	case 0x37: // SCF
		return 4
	case 0x38: // JR C, r8
		return 8
	case 0x39: // ADD HL, SP
		return 8
	case 0x3A: // LD A, (HL-)
		HL := gb.CPU.HL()
		gb.CPU.A = gb.MMU.Read(HL)
		HL--
		gb.CPU.H = uint8(HL>>8)
		gb.CPU.L = uint8(HL)
		return 8
	case 0x3B: // DEC SP
		return 8
	case 0x3C: // INC A
		return 4
	case 0x3D: // DEC A
		return 4
	case 0x3E: // LD A, d8
		gb.CPU.A = gb.MMU.Read(gb.CPU.PC)
		gb.CPU.PC++
		return 8
	case 0x3F: // CCF
		return 4

	// --- 0x4X ---
	case 0x40: // LD B, B
		gb.CPU.B = gb.CPU.B
		return 4
	case 0x41: // LD B, C
		gb.CPU.B = gb.CPU.C
		return 4
	case 0x42: // LD B, D
		gb.CPU.B = gb.CPU.D
		return 4
	case 0x43: // LD B, E
		gb.CPU.B = gb.CPU.E
		return 4
	case 0x44: // LD B, H
		gb.CPU.B = gb.CPU.H
		return 4
	case 0x45: // LD B, L
		gb.CPU.B = gb.CPU.L
		return 4
	case 0x46: // LD B, (HL)
		gb.CPU.B = gb.MMU.Read(gb.CPU.HL())
		return 8
	case 0x47: // LD B, A
		gb.CPU.B = gb.CPU.A
		return 4
	case 0x48: // LD C, B
		gb.CPU.C = gb.CPU.B
		return 4
	case 0x49: // LD C, C
		gb.CPU.C = gb.CPU.C
		return 4
	case 0x4A: // LD C, D
		gb.CPU.C = gb.CPU.D
		return 4
	case 0x4B: // LD C, E
		gb.CPU.C = gb.CPU.E
		return 4
	case 0x4C: // LD C, H
		gb.CPU.C = gb.CPU.H
		return 4
	case 0x4D: // LD C, L
		gb.CPU.C = gb.CPU.L
		return 4
	case 0x4E: // LD C, (HL)
		gb.CPU.C = gb.MMU.Read(gb.CPU.HL())
		return 8
	case 0x4F: // LD C, A
		gb.CPU.C = gb.CPU.A
		return 4

	// --- 0x5X ---
	case 0x50: // LD D, B
		gb.CPU.D = gb.CPU.B
		return 4
	case 0x51: // LD D, C
		gb.CPU.D = gb.CPU.C
		return 4
	case 0x52: // LD D, D
		
		gb.CPU.D = gb.CPU.D
		return 4
	case 0x53: // LD D, E
		gb.CPU.D = gb.CPU.E
		return 4
	case 0x54: // LD D, H
		gb.CPU.D = gb.CPU.H
		return 4
	case 0x55: // LD D, L
		gb.CPU.D = gb.CPU.L
		return 4
	case 0x56: // LD D, (HL)
		gb.CPU.D = gb.MMU.Read(gb.CPU.HL())
		return 8
	case 0x57: // LD D, A
		gb.CPU.D = gb.CPU.A
		return 4
	case 0x58: // LD E, B
		gb.CPU.E = gb.CPU.B
		return 4
	case 0x59: // LD E, C
		gb.CPU.E = gb.CPU.C
		return 4
	case 0x5A: // LD E, D
		gb.CPU.E = gb.CPU.D
		return 4
	case 0x5B: // LD E, E
		gb.CPU.E = gb.CPU.E
		return 4
	case 0x5C: // LD E, H
		gb.CPU.E = gb.CPU.H
		return 4
	case 0x5D: // LD E, L
		gb.CPU.E = gb.CPU.L
		return 4
	case 0x5E: // LD E, (HL)
		gb.CPU.E = gb.MMU.Read(gb.CPU.HL())
		return 8
	case 0x5F: // LD E, A
		gb.CPU.E = gb.CPU.A
		return 4

	// --- 0x6X ---
	case 0x60: // LD H, B
		gb.CPU.H = gb.CPU.B
		return 4
	case 0x61: // LD H, C
		gb.CPU.H = gb.CPU.C
		return 4
	case 0x62: // LD H, D
		gb.CPU.H = gb.CPU.D
		return 4
	case 0x63: // LD H, E
		gb.CPU.H = gb.CPU.E
		return 4
	case 0x64: // LD H, H
		gb.CPU.H = gb.CPU.H
		return 4
	case 0x65: // LD H, L
		gb.CPU.H = gb.CPU.L
		return 4
	case 0x66: // LD H, (HL)
		gb.CPU.H = gb.MMU.Read(gb.CPU.HL())
		return 8
	case 0x67: // LD H, A
		gb.CPU.H = gb.CPU.A
		return 4
	case 0x68: // LD L, B
		gb.CPU.L = gb.CPU.B
		return 4
	case 0x69: // LD L, C
		gb.CPU.L = gb.CPU.C
		return 4
	case 0x6A: // LD L, D
		gb.CPU.L = gb.CPU.D
		return 4
	case 0x6B: // LD L, E
		gb.CPU.L = gb.CPU.E
		return 4
	case 0x6C: // LD L, H
		gb.CPU.L = gb.CPU.H
		return 4
	case 0x6D: // LD L, L
		gb.CPU.L = gb.CPU.L
		return 4
	case 0x6E: // LD L, (HL)
		gb.CPU.L = gb.MMU.Read(gb.CPU.HL())
		return 8
	case 0x6F: // LD L, A
		gb.CPU.L = gb.CPU.A
		return 4

	// --- 0x7X ---
	case 0x70: // LD (HL), B
		gb.MMU.Write(gb.CPU.HL(), gb.CPU.B)
		return 8
	case 0x71: // LD (HL), C
		gb.MMU.Write(gb.CPU.HL(), gb.CPU.C)
		return 8
	case 0x72: // LD (HL), D
		gb.MMU.Write(gb.CPU.HL(), gb.CPU.D)
		return 8
	case 0x73: // LD (HL), E
		gb.MMU.Write(gb.CPU.HL(), gb.CPU.E)
		return 8
	case 0x74: // LD (HL), H
		gb.MMU.Write(gb.CPU.HL(), gb.CPU.H)
		return 8
	case 0x75: // LD (HL), L
		gb.MMU.Write(gb.CPU.HL(), gb.CPU.L)
		return 8
	case 0x76: // HALT
		gb.CPU.Halted = true
		return 4
	case 0x77: // LD (HL), A
		gb.MMU.Write(gb.CPU.HL(), gb.CPU.A)
		return 8
	case 0x78: // LD A, B
		gb.CPU.A = gb.CPU.B
		return 4
	case 0x79: // LD A, C
		gb.CPU.A = gb.CPU.C
		return 4
	case 0x7A: // LD A, D
		gb.CPU.A = gb.CPU.D
		return 4
	case 0x7B: // LD A, E
		gb.CPU.A = gb.CPU.E
		return 4
	case 0x7C: // LD A, H
		gb.CPU.A = gb.CPU.H
		return 4
	case 0x7D: // LD A, L
		gb.CPU.A = gb.CPU.L
		return 4
	case 0x7E: // LD A, (HL)
		gb.CPU.A = gb.MMU.Read(gb.CPU.HL())
		return 8
	case 0x7F: // LD A, A
		gb.CPU.A = gb.CPU.A
		return 4

	// --- 0x8X ---
	case 0x80: // ADD A, B
		gb.CPU.A += gb.CPU.B
		return 4
	case 0x81: // ADD A, C
		return 4
	case 0x82: // ADD A, D
		return 4
	case 0x83: // ADD A, E
		return 4
	case 0x84: // ADD A, H
		return 4
	case 0x85: // ADD A, L
		return 4
	case 0x86: // ADD A, (HL)
		return 8
	case 0x87: // ADD A, A
		return 4
	case 0x88: // ADC A, B
		return 4
	case 0x89: // ADC A, C
		return 4
	case 0x8A: // ADC A, D
		return 4
	case 0x8B: // ADC A, E
		return 4
	case 0x8C: // ADC A, H
		return 4
	case 0x8D: // ADC A, L
		return 4
	case 0x8E: // ADC A, (HL)
		return 8
	case 0x8F: // ADC A, A
		return 4

	// --- 0x9X ---
	case 0x90: // SUB B
		return 4
	case 0x91: // SUB C
		return 4
	case 0x92: // SUB D
		return 4
	case 0x93: // SUB E
		return 4
	case 0x94: // SUB H
		return 4
	case 0x95: // SUB L
		return 4
	case 0x96: // SUB (HL)
		return 8
	case 0x97: // SUB A
		return 4
	case 0x98: // SBC A, B
		return 4
	case 0x99: // SBC A, C
		return 4
	case 0x9A: // SBC A, D
		return 4
	case 0x9B: // SBC A, E
		return 4
	case 0x9C: // SBC A, H
		return 4
	case 0x9D: // SBC A, L
		return 4
	case 0x9E: // SBC A, (HL)
		return 8
	case 0x9F: // SBC A, A
		return 4

	// --- 0xAX ---
	case 0xA0: // AND B
		return 4
	case 0xA1: // AND C
		return 4
	case 0xA2: // AND D
		return 4
	case 0xA3: // AND E
		return 4
	case 0xA4: // AND H
		return 4
	case 0xA5: // AND L
		return 4
	case 0xA6: // AND (HL)
		return 8
	case 0xA7: // AND A
		return 4
	case 0xA8: // XOR B
		return 4
	case 0xA9: // XOR C
		return 4
	case 0xAA: // XOR D
		return 4
	case 0xAB: // XOR E
		return 4
	case 0xAC: // XOR H
		return 4
	case 0xAD: // XOR L
		return 4
	case 0xAE: // XOR (HL)
		return 8
	case 0xAF: // XOR A
		return 4

	// --- 0xBX ---
	case 0xB0: // OR B
		return 4
	case 0xB1: // OR C
		return 4
	case 0xB2: // OR D
		return 4
	case 0xB3: // OR E
		return 4
	case 0xB4: // OR H
		return 4
	case 0xB5: // OR L
		return 4
	case 0xB6: // OR (HL)
		return 8
	case 0xB7: // OR A
		return 4
	case 0xB8: // CP B
		return 4
	case 0xB9: // CP C
		return 4
	case 0xBA: // CP D
		return 4
	case 0xBB: // CP E
		return 4
	case 0xBC: // CP H
		return 4
	case 0xBD: // CP L
		return 4
	case 0xBE: // CP (HL)
		return 8
	case 0xBF: // CP A
		return 4

	// --- 0xCX ---
	case 0xC0: // RET NZ
		return 8
	case 0xC1: // POP BC
		return 12
	case 0xC2: // JP NZ, a16
		return 12
	case 0xC3: // JP a16
		lo := gb.MMU.Read(gb.CPU.PC)
		hi := gb.MMU.Read(gb.CPU.PC + 1)
		gb.CPU.PC = uint16(hi)<<8 | uint16(lo)
		return 16
	case 0xC4: // CALL NZ, a16
		return 12
	case 0xC5: // PUSH BC
		return 16
	case 0xC6: // ADD A, d8
		return 8
	case 0xC7: // RST 00H
		return 16
	case 0xC8: // RET Z
		return 8
	case 0xC9: // RET
		return 16
	case 0xCA: // JP Z, a16
		return 12
	case 0xCB: // PREFIX CB
		return gb.stepCB()
	case 0xCC: // CALL Z, a16
		return 12
	case 0xCD: // CALL a16
		return 24
	case 0xCE: // ADC A, d8
		return 8
	case 0xCF: // RST 08H
		return 16

	// --- 0xDX ---
	case 0xD0: // RET NC
		return 8
	case 0xD1: // POP DE
		return 12
	case 0xD2: // JP NC, a16
		return 12
	case 0xD4: // CALL NC, a16
		return 12
	case 0xD5: // PUSH DE
		return 16
	case 0xD6: // SUB d8
		return 8
	case 0xD7: // RST 10H
		return 16
	case 0xD8: // RET C
		return 8
	case 0xD9: // RETI
		return 16
	case 0xDA: // JP C, a16
		return 12
	case 0xDC: // CALL C, a16
		return 12
	case 0xDE: // SBC A, d8
		return 8
	case 0xDF: // RST 18H
		return 16

	// --- 0xEX ---
	case 0xE0: // LDH (a8), A
		offset := gb.MMU.Read(gb.CPU.PC)
		gb.CPU.PC++
		gb.MMU.Write(0xFF00+uint16(offset), gb.CPU.A)
		return 12
	case 0xE1: // POP HL
		return 12
	case 0xE2: // LD (C), A
		return 8
	case 0xE5: // PUSH HL
		return 16
	case 0xE6: // AND d8
		return 8
	case 0xE7: // RST 20H
		return 16
	case 0xE8: // ADD SP, r8
		return 16
	case 0xE9: // JP (HL)
		return 4
	case 0xEA: // LD (a16), A
		return 16
	case 0xEE: // XOR d8
		return 8
	case 0xEF: // RST 28H
		return 16

	// --- 0xFX ---
	case 0xF0: // LDH A, (a8)
		return 12
	case 0xF1: // POP AF
		return 12
	case 0xF2: // LD A, (C)
		return 8
	case 0xF3: // DI
		return 4
	case 0xF5: // PUSH AF
		return 16
	case 0xF6: // OR d8
		return 8
	case 0xF7: // RST 30H
		return 16
	case 0xF8: // LD HL, SP+r8
		return 12
	case 0xF9: // LD SP, HL
		return 8
	case 0xFA: // LD A, (a16)
		return 16
	case 0xFB: // EI
		return 4
	case 0xFE: // CP d8
		return 8
	case 0xFF: // RST 38H
		return 16

	default:
		panic(fmt.Sprintf("Unknown opcode: 0x%02X", opcode))
	}
}

func (gb *Gameboy) stepCB() int {
	opcode := gb.MMU.Read(gb.CPU.PC)
	gb.CPU.PC++

	switch opcode {

	// --- RLC ---
	case 0x00: // RLC B
		return 8
	case 0x01: // RLC C
		return 8
	case 0x02: // RLC D
		return 8
	case 0x03: // RLC E
		return 8
	case 0x04: // RLC H
		return 8
	case 0x05: // RLC L
		return 8
	case 0x06: // RLC (HL)
		return 16
	case 0x07: // RLC A
		return 8

	// --- RRC ---
	case 0x08: // RRC B
		return 8
	case 0x09: // RRC C
		return 8
	case 0x0A: // RRC D
		return 8
	case 0x0B: // RRC E
		return 8
	case 0x0C: // RRC H
		return 8
	case 0x0D: // RRC L
		return 8
	case 0x0E: // RRC (HL)
		return 16
	case 0x0F: // RRC A
		return 8

	// --- RL ---
	case 0x10: // RL B
		return 8
	case 0x11: // RL C
		return 8
	case 0x12: // RL D
		return 8
	case 0x13: // RL E
		return 8
	case 0x14: // RL H
		return 8
	case 0x15: // RL L
		return 8
	case 0x16: // RL (HL)
		return 16
	case 0x17: // RL A
		return 8

	// --- RR ---
	case 0x18: // RR B
		return 8
	case 0x19: // RR C
		return 8
	case 0x1A: // RR D
		return 8
	case 0x1B: // RR E
		return 8
	case 0x1C: // RR H
		return 8
	case 0x1D: // RR L
		return 8
	case 0x1E: // RR (HL)
		return 16
	case 0x1F: // RR A
		return 8

	// --- SLA ---
	case 0x20: // SLA B
		return 8
	case 0x21: // SLA C
		return 8
	case 0x22: // SLA D
		return 8
	case 0x23: // SLA E
		return 8
	case 0x24: // SLA H
		return 8
	case 0x25: // SLA L
		return 8
	case 0x26: // SLA (HL)
		return 16
	case 0x27: // SLA A
		return 8

	// --- SRA ---
	case 0x28: // SRA B
		return 8
	case 0x29: // SRA C
		return 8
	case 0x2A: // SRA D
		return 8
	case 0x2B: // SRA E
		return 8
	case 0x2C: // SRA H
		return 8
	case 0x2D: // SRA L
		return 8
	case 0x2E: // SRA (HL)
		return 16
	case 0x2F: // SRA A
		return 8

	// --- SWAP ---
	case 0x30: // SWAP B
		return 8
	case 0x31: // SWAP C
		return 8
	case 0x32: // SWAP D
		return 8
	case 0x33: // SWAP E
		return 8
	case 0x34: // SWAP H
		return 8
	case 0x35: // SWAP L
		return 8
	case 0x36: // SWAP (HL)
		return 16
	case 0x37: // SWAP A
		return 8

	// --- SRL ---
	case 0x38: // SRL B
		return 8
	case 0x39: // SRL C
		return 8
	case 0x3A: // SRL D
		return 8
	case 0x3B: // SRL E
		return 8
	case 0x3C: // SRL H
		return 8
	case 0x3D: // SRL L
		return 8
	case 0x3E: // SRL (HL)
		return 16
	case 0x3F: // SRL A
		return 8

	// --- BIT 0 ---
	case 0x40: // BIT 0, B
		return 8
	case 0x41: // BIT 0, C
		return 8
	case 0x42: // BIT 0, D
		return 8
	case 0x43: // BIT 0, E
		return 8
	case 0x44: // BIT 0, H
		return 8
	case 0x45: // BIT 0, L
		return 8
	case 0x46: // BIT 0, (HL)
		return 12
	case 0x47: // BIT 0, A
		return 8

	// --- BIT 1 ---
	case 0x48: // BIT 1, B
		return 8
	case 0x49: // BIT 1, C
		return 8
	case 0x4A: // BIT 1, D
		return 8
	case 0x4B: // BIT 1, E
		return 8
	case 0x4C: // BIT 1, H
		return 8
	case 0x4D: // BIT 1, L
		return 8
	case 0x4E: // BIT 1, (HL)
		return 12
	case 0x4F: // BIT 1, A
		return 8

	// --- BIT 2 ---
	case 0x50: // BIT 2, B
		return 8
	case 0x51: // BIT 2, C
		return 8
	case 0x52: // BIT 2, D
		return 8
	case 0x53: // BIT 2, E
		return 8
	case 0x54: // BIT 2, H
		return 8
	case 0x55: // BIT 2, L
		return 8
	case 0x56: // BIT 2, (HL)
		return 12
	case 0x57: // BIT 2, A
		return 8

	// --- BIT 3 ---
	case 0x58: // BIT 3, B
		return 8
	case 0x59: // BIT 3, C
		return 8
	case 0x5A: // BIT 3, D
		return 8
	case 0x5B: // BIT 3, E
		return 8
	case 0x5C: // BIT 3, H
		return 8
	case 0x5D: // BIT 3, L
		return 8
	case 0x5E: // BIT 3, (HL)
		return 12
	case 0x5F: // BIT 3, A
		return 8

	// --- BIT 4 ---
	case 0x60: // BIT 4, B
		return 8
	case 0x61: // BIT 4, C
		return 8
	case 0x62: // BIT 4, D
		return 8
	case 0x63: // BIT 4, E
		return 8
	case 0x64: // BIT 4, H
		return 8
	case 0x65: // BIT 4, L
		return 8
	case 0x66: // BIT 4, (HL)
		return 12
	case 0x67: // BIT 4, A
		return 8

	// --- BIT 5 ---
	case 0x68: // BIT 5, B
		return 8
	case 0x69: // BIT 5, C
		return 8
	case 0x6A: // BIT 5, D
		return 8
	case 0x6B: // BIT 5, E
		return 8
	case 0x6C: // BIT 5, H
		return 8
	case 0x6D: // BIT 5, L
		return 8
	case 0x6E: // BIT 5, (HL)
		return 12
	case 0x6F: // BIT 5, A
		return 8

	// --- BIT 6 ---
	case 0x70: // BIT 6, B
		return 8
	case 0x71: // BIT 6, C
		return 8
	case 0x72: // BIT 6, D
		return 8
	case 0x73: // BIT 6, E
		return 8
	case 0x74: // BIT 6, H
		return 8
	case 0x75: // BIT 6, L
		return 8
	case 0x76: // BIT 6, (HL)
		return 12
	case 0x77: // BIT 6, A
		return 8

	// --- BIT 7 ---
	case 0x78: // BIT 7, B
		return 8
	case 0x79: // BIT 7, C
		return 8
	case 0x7A: // BIT 7, D
		return 8
	case 0x7B: // BIT 7, E
		return 8
	case 0x7C: // BIT 7, H
		return 8
	case 0x7D: // BIT 7, L
		return 8
	case 0x7E: // BIT 7, (HL)
		return 12
	case 0x7F: // BIT 7, A
		return 8

	// --- RES 0 ---
	case 0x80: // RES 0, B
		return 8
	case 0x81: // RES 0, C
		return 8
	case 0x82: // RES 0, D
		return 8
	case 0x83: // RES 0, E
		return 8
	case 0x84: // RES 0, H
		return 8
	case 0x85: // RES 0, L
		return 8
	case 0x86: // RES 0, (HL)
		return 16
	case 0x87: // RES 0, A
		return 8

	// --- RES 1 ---
	case 0x88: // RES 1, B
		return 8
	case 0x89: // RES 1, C
		return 8
	case 0x8A: // RES 1, D
		return 8
	case 0x8B: // RES 1, E
		return 8
	case 0x8C: // RES 1, H
		return 8
	case 0x8D: // RES 1, L
		return 8
	case 0x8E: // RES 1, (HL)
		return 16
	case 0x8F: // RES 1, A
		return 8

	// --- RES 2 ---
	case 0x90: // RES 2, B
		return 8
	case 0x91: // RES 2, C
		return 8
	case 0x92: // RES 2, D
		return 8
	case 0x93: // RES 2, E
		return 8
	case 0x94: // RES 2, H
		return 8
	case 0x95: // RES 2, L
		return 8
	case 0x96: // RES 2, (HL)
		return 16
	case 0x97: // RES 2, A
		return 8

	// --- RES 3 ---
	case 0x98: // RES 3, B
		return 8
	case 0x99: // RES 3, C
		return 8
	case 0x9A: // RES 3, D
		return 8
	case 0x9B: // RES 3, E
		return 8
	case 0x9C: // RES 3, H
		return 8
	case 0x9D: // RES 3, L
		return 8
	case 0x9E: // RES 3, (HL)
		return 16
	case 0x9F: // RES 3, A
		return 8

	// --- RES 4 ---
	case 0xA0: // RES 4, B
		return 8
	case 0xA1: // RES 4, C
		return 8
	case 0xA2: // RES 4, D
		return 8
	case 0xA3: // RES 4, E
		return 8
	case 0xA4: // RES 4, H
		return 8
	case 0xA5: // RES 4, L
		return 8
	case 0xA6: // RES 4, (HL)
		return 16
	case 0xA7: // RES 4, A
		return 8

	// --- RES 5 ---
	case 0xA8: // RES 5, B
		return 8
	case 0xA9: // RES 5, C
		return 8
	case 0xAA: // RES 5, D
		return 8
	case 0xAB: // RES 5, E
		return 8
	case 0xAC: // RES 5, H
		return 8
	case 0xAD: // RES 5, L
		return 8
	case 0xAE: // RES 5, (HL)
		return 16
	case 0xAF: // RES 5, A
		return 8

	// --- RES 6 ---
	case 0xB0: // RES 6, B
		return 8
	case 0xB1: // RES 6, C
		return 8
	case 0xB2: // RES 6, D
		return 8
	case 0xB3: // RES 6, E
		return 8
	case 0xB4: // RES 6, H
		return 8
	case 0xB5: // RES 6, L
		return 8
	case 0xB6: // RES 6, (HL)
		return 16
	case 0xB7: // RES 6, A
		return 8

	// --- RES 7 ---
	case 0xB8: // RES 7, B
		return 8
	case 0xB9: // RES 7, C
		return 8
	case 0xBA: // RES 7, D
		return 8
	case 0xBB: // RES 7, E
		return 8
	case 0xBC: // RES 7, H
		return 8
	case 0xBD: // RES 7, L
		return 8
	case 0xBE: // RES 7, (HL)
		return 16
	case 0xBF: // RES 7, A
		return 8

	// --- SET 0 ---
	case 0xC0: // SET 0, B
		return 8
	case 0xC1: // SET 0, C
		return 8
	case 0xC2: // SET 0, D
		return 8
	case 0xC3: // SET 0, E
		return 8
	case 0xC4: // SET 0, H
		return 8
	case 0xC5: // SET 0, L
		return 8
	case 0xC6: // SET 0, (HL)
		return 16
	case 0xC7: // SET 0, A
		return 8

	// --- SET 1 ---
	case 0xC8: // SET 1, B
		return 8
	case 0xC9: // SET 1, C
		return 8
	case 0xCA: // SET 1, D
		return 8
	case 0xCB: // SET 1, E
		return 8
	case 0xCC: // SET 1, H
		return 8
	case 0xCD: // SET 1, L
		return 8
	case 0xCE: // SET 1, (HL)
		return 16
	case 0xCF: // SET 1, A
		return 8

	// --- SET 2 ---
	case 0xD0: // SET 2, B
		return 8
	case 0xD1: // SET 2, C
		return 8
	case 0xD2: // SET 2, D
		return 8
	case 0xD3: // SET 2, E
		return 8
	case 0xD4: // SET 2, H
		return 8
	case 0xD5: // SET 2, L
		return 8
	case 0xD6: // SET 2, (HL)
		return 16
	case 0xD7: // SET 2, A
		return 8

	// --- SET 3 ---
	case 0xD8: // SET 3, B
		return 8
	case 0xD9: // SET 3, C
		return 8
	case 0xDA: // SET 3, D
		return 8
	case 0xDB: // SET 3, E
		return 8
	case 0xDC: // SET 3, H
		return 8
	case 0xDD: // SET 3, L
		return 8
	case 0xDE: // SET 3, (HL)
		return 16
	case 0xDF: // SET 3, A
		return 8

	// --- SET 4 ---
	case 0xE0: // SET 4, B
		return 8
	case 0xE1: // SET 4, C
		return 8
	case 0xE2: // SET 4, D
		return 8
	case 0xE3: // SET 4, E
		return 8
	case 0xE4: // SET 4, H
		return 8
	case 0xE5: // SET 4, L
		return 8
	case 0xE6: // SET 4, (HL)
		return 16
	case 0xE7: // SET 4, A
		return 8

	// --- SET 5 ---
	case 0xE8: // SET 5, B
		return 8
	case 0xE9: // SET 5, C
		return 8
	case 0xEA: // SET 5, D
		return 8
	case 0xEB: // SET 5, E
		return 8
	case 0xEC: // SET 5, H
		return 8
	case 0xED: // SET 5, L
		return 8
	case 0xEE: // SET 5, (HL)
		return 16
	case 0xEF: // SET 5, A
		return 8

	// --- SET 6 ---
	case 0xF0: // SET 6, B
		return 8
	case 0xF1: // SET 6, C
		return 8
	case 0xF2: // SET 6, D
		return 8
	case 0xF3: // SET 6, E
		return 8
	case 0xF4: // SET 6, H
		return 8
	case 0xF5: // SET 6, L
		return 8
	case 0xF6: // SET 6, (HL)
		return 16
	case 0xF7: // SET 6, A
		return 8

	// --- SET 7 ---
	case 0xF8: // SET 7, B
		return 8
	case 0xF9: // SET 7, C
		return 8
	case 0xFA: // SET 7, D
		return 8
	case 0xFB: // SET 7, E
		return 8
	case 0xFC: // SET 7, H
		return 8
	case 0xFD: // SET 7, L
		return 8
	case 0xFE: // SET 7, (HL)
		return 16
	case 0xFF: // SET 7, A
		return 8

	default:
		panic(fmt.Sprintf("Unknown CB opcode: 0xCB 0x%02X", opcode))
	}
}
