package main

import (
	"testing"
)

func InitTest(opcode Opcode) *CPU {
	memory[0x0100] = byte(opcode)
	cpu := InitCPU()
	return cpu
}

func TestNOP(t *testing.T) {
	cpu := InitTest(NOP)

	cycles := Step(cpu)

	if cpu.PC != 0x0102 {
		t.Errorf("cpu.PC = 0x%04x, expected 0x0102", cpu.PC)
	} else if cycles != 4 {
		t.Errorf("Step({A: 0x%04x, PC: 0x%04x, Halted: %t}) = %d, expected 4", 0x00, 0x0100, false, cycles)
	}
}

func TestJP_a16(t *testing.T) {
	cpu := InitTest(JPa16)

	memory[cpu.PC+1] = 0x10
	memory[cpu.PC+2] = 0x01
	cycles := Step(cpu)

	if cpu.PC != 0x0110 {
		t.Errorf("cpu.PC = 0x%04x, expected 0x0110", cpu.PC)
	} else if cycles != 16 {
		t.Errorf("Step({A: 0x%04x, PC: 0x%04x, Halted: %t}) = %d, expected 16", 0x00, 0x0100, false, cycles)
	}
}

func TestJP_a16_LittleEndian(t *testing.T) {
	cpu := InitTest(JPa16)

	memory[cpu.PC+1] = 0x10
	memory[cpu.PC+2] = 0x01
	cycles := Step(cpu)

	if cpu.PC != 0x0110 {
		t.Errorf("cpu.PC = 0x%04x, expected 0x0110", cpu.PC)
	} else if cycles != 16 {
		t.Errorf("Step({A: 0x%04x, PC: 0x%04x, Halted: %t}) = %d, expected 16", 0x00, 0x0100, false, cycles)
	}
}

func TestLD_A_d8(t *testing.T) {
	cpu := InitTest(LDAn8)

	memory[cpu.PC+1] = 0xFF

	cycles := Step(cpu)
	if cpu.PC != 0x0102 {
		t.Errorf("cpu.PC = 0x%04x, expected 0x0102", cpu.PC)
	} else if cycles != 8 {
		t.Errorf("Step({A: 0x%04x, PC: 0x%04x, Halted: %t}) = %d, expected 8", 0x00, 0x0100, false, cycles)
	} else if cpu.A != byte(0xFF) {
		t.Errorf("cpu.A = 0x%02x, expected 0xFF", cpu.A)
	}
}

func TestLDH_a8_A(t *testing.T) {
	cpu := InitTest(LDHa8A)

	memory[cpu.PC+1] = 0x0F
	cpu.A = 0xFF

	cycles := Step(cpu)
	if cpu.PC != 0x0102 {
		t.Errorf("cpu.PC = 0x%04x, expected 0x0102", cpu.PC)
	} else if cycles != 12 {
		t.Errorf("Step({A: 0x%04x, PC: 0x%04x, Halted: %t}) = %d, expected 12", 0x00, 0x0100, false, cycles)
	} else if memory[IORegisters+uint16(0x0F)] != byte(0xFF) {
		t.Errorf("RAM[0xFF00 + 0x0F]= 0x%02x, expected 0xFF", memory[IORegisters+uint16(0x0F)])
	}
}

func TestHALT(t *testing.T) {
	cpu := InitTest(HALT)

	cycles := Step(cpu)

	if cpu.PC != 0x0101 {
		t.Errorf("cpu.PC = 0x%04x, expected 0x0101", cpu.PC)
	} else if cycles != 4 {
		t.Errorf("Step({A: 0x%04x, PC: 0x%04x, Halted: %t}) = %d, expected 4", 0x00, 0x0100, false, cycles)
	} else if cpu.Halted != true {
		t.Errorf("cpu.Halted = %t, expected false", cpu.Halted)
	}
}

func TestUnknownOpcode(t *testing.T) {
	cpu := InitTest(UNKNOWN)

	defer func() {
		err := recover()
		if err == nil {
			t.Fatalf("Expected panic with invalid address, received %v", err)
		}
	}()

	Step(cpu)
}
