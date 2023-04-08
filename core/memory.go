package core

import (
	"math/rand"
)

type MMUType struct {
}

var MMU = MMUType{}

var sRAM [0x2000]byte
var io [0x100]byte
var vRAM [0x2000]byte
var oam [0x100]byte
var wRAM [0x2000]byte
var hRAM [0x80]byte

const OFFSETsRAM uint16 = 0xA000
const OFFSETvRAM uint16 = 0x8000
const OFFSETwRAMlower uint16 = 0xC000
const OFFSETwRAMupper uint16 = 0xE000
const OFFSEToam uint16 = 0xFE00
const OFFSEThRAM uint16 = 0xFF80
const OFFSETio uint16 = 0xFF00

func (mmu *MMUType) ReadByte(address uint16) byte {
	if address <= 0x7FFF {
		return ROM.data[address]
	} else if address >= 0xA000 && address <= 0xBFFF {
		return sRAM[address-OFFSETsRAM]
	} else if address >= 0x8000 && address <= 0x9FFF {
		return vRAM[address-OFFSETvRAM]
	} else if address >= 0xC000 && address <= 0xDFFF {
		return wRAM[address-OFFSETwRAMlower]
	} else if address >= 0xE000 && address <= 0xFDFF {
		return wRAM[address-OFFSETwRAMupper]
	} else if address >= 0xFE00 && address <= 0xFEFF {
		return oam[address-OFFSEToam]
	} else if address == 0xFF04 {
		return byte(rand.Intn(0xFF + 1)) // generate random value range: [0, 255] (byte)
	} else if address == 0xFF40 {
		return GPU.control
	} else if address == 0xFF42 {
		return GPU.scrollY
	} else if address == 0xFF43 {
		return GPU.scrollX
	} else if address == 0xFF44 {
		return GPU.scanline
	} else if address == 0xFF00 {
		// io block
	} else if address == 0xFF0F {
		return INTERRUPTS.flags
	} else if address == 0xFFFF {
		return INTERRUPTS.enable
	} else if address >= 0xFF80 && address <= 0xFFFE {
		return hRAM[address-OFFSEThRAM]
	} else if address >= 0xFF00 && address <= 0xFF7F {
		return io[address-OFFSETio]
	}
	return 0
}

func (mmu *MMUType) WriteByte(address uint16, value byte) {

	if address <= 0x7FFF {
		ROM.data[address] = value
	} else if address >= 0xA000 && address <= 0xBFFF {
		sRAM[address-OFFSETsRAM] = value
	} else if address >= 0x8000 && address <= 0x9FFF {
		vRAM[address-OFFSETvRAM] = value
		//update tile too
	} else if address >= 0xC000 && address <= 0xDFFF {
		wRAM[address-OFFSETwRAMlower] = value
	} else if address >= 0xE000 && address <= 0xFDFF {
		wRAM[address-OFFSETwRAMupper] = value
	} else if address >= 0xFE00 && address <= 0xFEFF {
		oam[address-OFFSEToam] = value
	} else if address == 0xFF40 {
		GPU.control = value
	} else if address == 0xFF42 {
		GPU.scrollY = value
	} else if address == 0xFF43 {
		GPU.scrollX = value
	} else if address == 0xFF46 {
		//OAM DMA copy
	} else if address == 0xFF00 {
		// io block
	} else if address == 0xFF0F {
		INTERRUPTS.flags = value
	} else if address == 0xFFFF {
		INTERRUPTS.enable = value
	} else if address >= 0xFF00 && address <= 0xFF7F {
		io[address-OFFSETio] = value
	}
}

func (mmu *MMUType) ReadShort(address uint16) uint16 {
	return uint16(uint16(mmu.ReadByte(address)) | uint16(mmu.ReadByte((address+1)))<<8)
}

func (mmu *MMUType) WriteShort(address uint16, value uint16) {
	mmu.WriteByte(address, byte(value&0x00FF))
	mmu.WriteByte(address+1, byte((value&0xFF00)>>8))
}

func (mmu *MMUType) ReadShortFromStack() uint16 {
	var value = mmu.ReadShort(REGISTERS.SP)
	REGISTERS.SP += 2

	return value
}

func (mmu *MMUType) WriteShortToStack(value uint16) {
	REGISTERS.SP -= 2
	mmu.WriteShort(REGISTERS.SP, value)
}
