package mmu

type MMU struct {
	Memory [0x10000]byte //32kb
}

func (m *MMU) Read(addr uint16) byte{
	return m.Memory[addr]
}

func (m *MMU) Write(addr uint16 , value byte){
	if addr < 0x8000 {
		return
	}
		m.Memory[addr] = value
}