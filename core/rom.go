package core

import (
	"fmt"
	"math"
)

// ROM_OFFSET_NAME is the location in every ROM of the name
const ROM_OFFSET_NAME = 0x0134
const ROM_OFFSET_TYPE = 0x0147
const ROM_OFFSET_ROM_SIZE = 0x0148
const ROM_OFFSET_RAM_SIZE = 0x0149

var romTypeMap = map[uint8]string{
	0x00: "ROM_PLAIN",
	0x01: "ROM_MBC1",
	0x02: "ROM_MBC1_RAM",
	0x03: "ROM_MBC1_RAM_BATT",
	0x05: "ROM_MBC2",
	0x06: "ROM_MBC2_BATTERY",
	0x08: "ROM_RAM",
	0x09: "ROM_RAM_BATTERY",
	0x0B: "ROM_MMM01",
	0x0C: "ROM_MMM01_SRAM",
	0x0D: "ROM_MMM01_SRAM_BATT",
	0x0F: "ROM_MBC3_TIMER_BATT",
	0x10: "ROM_MBC3_TIMER_RAM_BATT",
	0x11: "ROM_MBC3",
	0x12: "ROM_MBC3_RAM",
	0x13: "ROM_MBC3_RAM_BATT",
	0x19: "ROM_MBC5",
	0x1A: "ROM_MBC5_RAM",
	0x1B: "ROM_MBC5_RAM_BATT",
	0x1C: "ROM_MBC5_RUMBLE",
	0x1D: "ROM_MBC5_RUMBLE_SRAM",
	0x1E: "ROM_MBC5_RUMBLE_SRAM_BATT",
	0x1F: "ROM_POCKET_CAMERA",
	0xFD: "ROM_BANDAI_TAMA5",
	0xFE: "ROM_HUDSON_HUC3",
	0xFF: "ROM_HUDSON_HUC1",
}

type ROMType struct {
	data         []byte
	model        []interface{}
	modelColumns []string
	romType      string
	romName      string
	romSize      int
	romRAMSize   int
}

// GetType gets the ROM Type of the ROM from within the ROM's file
func (rom *ROMType) GetType() string {
	return romTypeMap[rom.data[ROM_OFFSET_TYPE]]
}

// GetROMSize get the size of the ROM from within the ROM's file
func (rom *ROMType) GetROMSize() int {
	var romSize int = int(rom.data[ROM_OFFSET_ROM_SIZE])
	if (romSize & 0xF0) == 0x50 {
		romSize = int(math.Pow(2.0, float64((((0x52)&0xF)+1))+64))
	} else {
		romSize = int(math.Pow(2.0, float64((romSize + 1))))
	}
	return romSize
}

// GetRAMSize get the size of the ROM's RAM from within the ROM's file
func (rom *ROMType) GetRAMSize() int {
	return int(math.Pow(4.0, float64(rom.data[ROM_OFFSET_RAM_SIZE])))
}

// GetName gets the name of the ROM from within the ROM's file
func (rom *ROMType) GetName() string {
	var name [256]byte
	var returnString string
	// var charRead int = 0
	for i := 0; i < 256; i++ {
		if string(rom.data[ROM_OFFSET_NAME+i]) != "\000" {
			name[i] = rom.data[ROM_OFFSET_NAME+i]
		} else {
			returnString = string(name[:i])
			break
		}
	}
	return returnString
}

// BuildModel builds a GTK TreeModel using an instruction map
func (rom *ROMType) BuildModel() {
	model := []interface{}{}

	for i := 0x0000; i < len(rom.data); i++ {
		row := []string{}
		if CPU.INSTRUCTIONS[rom.data[i]].NumOperands == 1 {
			i += int(CPU.INSTRUCTIONS[rom.data[i]].NumOperands)
			row = append(row, fmt.Sprintf("0x%04X", i))
			row = append(row, fmt.Sprintf("%s, 0x%02X", CPU.INSTRUCTIONS[rom.data[i]].Name, rom.data[i+1]))
		} else if CPU.INSTRUCTIONS[rom.data[i]].NumOperands == 2 {
			i += int(CPU.INSTRUCTIONS[rom.data[i]].NumOperands)
			row = append(row, fmt.Sprintf("0x%04X", i))
			row = append(row, fmt.Sprintf("%s 0x%02X 0x%02X", CPU.INSTRUCTIONS[rom.data[i]].Name, rom.data[i+1], rom.data[i+2]))
		} else {
			if CPU.INSTRUCTIONS[rom.data[i]].Name == "UNKNOWN" {
				row = append(row, fmt.Sprintf("0x%04X", i))
				row = append(row, fmt.Sprintf("%s: 0x%02X", CPU.INSTRUCTIONS[rom.data[i]].Name, rom.data[i]))
			} else {
				row = append(row, fmt.Sprintf("0x%04X", i))
				row = append(row, fmt.Sprintf("%s", CPU.INSTRUCTIONS[rom.data[i]].Name))
			}
		}
		model = append(model, row)
	}
	rom.model = model
}

var ROM = ROMType{
	data:         []byte{},
	model:        nil,
	modelColumns: []string{"Offset", "Instruction"},
	romType:      "",
	romName:      "",
	romSize:      0,
	romRAMSize:   0,
}
