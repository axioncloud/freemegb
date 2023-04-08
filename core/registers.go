package core

import (
	"fmt"
	"reflect"

	"github.com/gotk3/gotk3/gtk"
)

// FlagsType holds the FLAG values required by documentation
type FlagsType struct {
	ZERO       byte
	SUBTRACT   byte
	HALF_CARRY byte
	CARRY      byte
}

// RegistersType is the strucutre that holds the individual registers of the CPU
type RegistersType struct {
	AF    uint16
	BC    uint16
	DE    uint16
	HL    uint16
	SP    uint16
	PC    uint16
	FLAGS FlagsType
}

// SetA is the setter for REGISTER A
func (r *RegistersType) SetA(value byte) {
	var F = r.AF & 0x00FF
	r.AF = uint16(value)<<8 | F
}

// A  is the accessor for REGISTER A
func (r *RegistersType) A() byte {
	return byte(r.AF >> 8)
}

// SetF is the setter for REGISTER F
func (r *RegistersType) SetF(value byte) {
	var A = r.AF & 0xFF00
	r.AF = A<<8 | uint16(value)
}

// F is the accessor for REGISTER F
func (r *RegistersType) F() byte {
	return byte(r.AF >> 8)
}

// SetB is the setter for REGISTER B
func (r *RegistersType) SetB(value byte) {
	var C = r.BC & 0x00FF
	r.BC = uint16(value)<<8 | C
}

// B is the accessor for REGISTER B
func (r *RegistersType) B() byte {
	return byte(r.BC >> 8)
}

// SetC is the setter for REGISTER C
func (r *RegistersType) SetC(value byte) {
	var B = r.BC & 0xFF00
	r.BC = B<<8 | uint16(value)
}

// C is the accessor for REGISTER C
func (r *RegistersType) C() byte {
	return byte(r.BC & 0x00FF)
}

// SetD is the setter for REGISTER D
func (r *RegistersType) SetD(value byte) {
	var E = r.BC & 0x00FF
	r.DE = uint16(value)<<8 | E
}

// D is the accessor for REGISTER D
func (r *RegistersType) D() byte {
	return byte(r.BC >> 8)
}

// SetE is the setter for REGISTER E
func (r *RegistersType) SetE(value byte) {
	var D = r.DE & 0xFF00
	r.DE = D<<8 | uint16(value)
}

// E is the accessor for REGISTER E
func (r *RegistersType) E() byte {
	return byte(r.BC & 0x00FF)
}

// SetH is the setter for REGISTER H
func (r *RegistersType) SetH(value byte) {
	var L = r.HL & 0x00FF
	r.HL = uint16(value)<<8 | L
}

// H is the accessor for REGISTER H
func (r *RegistersType) H() byte {
	return byte(r.HL >> 8)
}

// SetL is the setter for REGISTER L
func (r *RegistersType) SetL(value byte) {
	var H = r.DE & 0xFF00
	r.HL = H<<8 | uint16(value)
}

// L is the accessor for REGISTER L
func (r *RegistersType) L() byte {
	return byte(r.HL & 0x00FF)
}

// Print will print all registers, preformatted
func (r *RegistersType) Print() {
	Logger.Logf(LogTypes.INFO, "REGISTERS:\n\t\t\t\t\tAF: 0x%04X\n\t\t\t\t\tBC: 0x%04X\n\t\t\t\t\tDE: 0x%04X\n\t\t\t\t\tHL: 0x%04X\n\t\t\t\t\tSP: 0x%04X\n\t\t\t\t\tPC: 0x%04X\n", r.AF, r.BC, r.DE, r.HL, r.SP, r.PC)
}

// Register16toString converts a register (uint16) to a string, preformatted
func (r *RegistersType) Register16toString(register uint16) string {
	return fmt.Sprintf("0x%04X\n", register)
}

// Register8toString converts a register (byte) to a string, preformatted
func (r *RegistersType) Register8toString(register byte) string {
	return fmt.Sprintf("0x%02X\n", register)
}

// CombineTo16 combines two 8-bit registers into a 16-bit register
func (r *RegistersType) CombineTo16(lower byte, upper byte) uint16 {
	return uint16(uint16(lower)<<8 | uint16(upper))
}

// ADD8 addes two 8-bit registers
func (r *RegistersType) ADD8(destination byte, source byte) byte {
	var result uint16 = uint16(destination) + uint16(source)

	if result&0xFF00 != 0 {
		r.FLAG_SET(r.FLAGS.CARRY)
	} else {
		r.FLAG_CLEAR(r.FLAGS.CARRY)
	}

	destination = byte(result & 0xFF)

	if destination != 0 {
		r.FLAG_CLEAR(r.FLAGS.ZERO)
	} else {
		r.FLAG_SET(r.FLAGS.ZERO)
	}

	if ((destination & 0x0F) + (source & 0x0F)) > 0x0F {
		r.FLAG_SET(r.FLAGS.HALF_CARRY)
	} else {
		r.FLAG_CLEAR(r.FLAGS.HALF_CARRY)
	}

	r.FLAG_CLEAR(r.FLAGS.SUBTRACT)

	return byte(destination)
}

// ADD16 addes two 16-bit registers
func (r *RegistersType) ADD16(destination uint16, source uint16) uint16 {
	// destination + source
	var result uint32 = uint32(destination) + uint32(source)

	if result&0xFFFF0000 != 0 {
		r.FLAG_SET(r.FLAGS.CARRY)
	} else {
		r.FLAG_CLEAR(r.FLAGS.CARRY)
	}

	destination = uint16(result & 0xFFFF)

	if ((destination & 0x0F) + (source & 0x0F)) > 0x0F {
		r.FLAG_SET(r.FLAGS.HALF_CARRY)
	} else {
		r.FLAG_CLEAR(r.FLAGS.HALF_CARRY)
	}

	r.FLAG_CLEAR(r.FLAGS.SUBTRACT)

	return uint16(destination)
}

// ADDC is a helper function to add-carry
func (r *RegistersType) ADDC(value byte) {
	if r.FLAG_ISSET(r.FLAGS.CARRY) {
		value++
	}

	var result uint16 = uint16(r.A()) + uint16(value)

	if (result&0xFF00)&1 != 0 {
		r.FLAG_SET(r.FLAGS.CARRY)
	} else {
		r.FLAG_CLEAR(r.FLAGS.CARRY)
	}

	if value == r.A() {
		r.FLAG_SET(r.FLAGS.ZERO)
	} else {
		r.FLAG_CLEAR(r.FLAGS.ZERO)
	}

	if (value&0x0F)+(r.A()&0x0F) > 0x0F {
		r.FLAG_SET(r.FLAGS.HALF_CARRY)
	} else {
		r.FLAG_CLEAR(r.FLAGS.HALF_CARRY)
	}

	r.FLAG_SET(r.FLAGS.SUBTRACT)

	r.SetA(byte(result & 0xFF))
}

// SUBC is a helper function to sub-carry
func (r *RegistersType) SUBC(value byte) {
	if r.FLAG_ISSET(r.FLAGS.CARRY) {
		value++
	}

	r.FLAG_SET(r.FLAGS.SUBTRACT)

	if value > r.A() {
		r.FLAG_SET(r.FLAGS.CARRY)
	} else {
		r.FLAG_CLEAR(r.FLAGS.CARRY)
	}

	if value == r.A() {
		r.FLAG_SET(r.FLAGS.ZERO)
	} else {
		r.FLAG_CLEAR(r.FLAGS.ZERO)
	}

	if value > r.A() {
		r.FLAG_SET(r.FLAGS.CARRY)
	} else {
		r.FLAG_CLEAR(r.FLAGS.CARRY)
	}

	if (value & 0x0F) > (r.A() & 0x0F) {
		r.FLAG_SET(r.FLAGS.HALF_CARRY)
	} else {
		r.FLAG_CLEAR(r.FLAGS.HALF_CARRY)
	}

	r.SetA(r.A() - value)
}

func (r *RegistersType) SUB(value byte) {
	r.FLAG_SET(r.FLAGS.SUBTRACT)

	if value > r.A() {
		r.FLAG_SET(r.FLAGS.CARRY)
	} else {
		r.FLAG_CLEAR(r.FLAGS.CARRY)
	}

	if (value & 0x0F) > (r.A() & 0x0F) {
		r.FLAG_SET(r.FLAGS.HALF_CARRY)
	} else {
		r.FLAG_CLEAR(r.FLAGS.HALF_CARRY)
	}

	r.SetA(r.A() - value)

	if r.A()&1 != 0 {
		r.FLAG_CLEAR(r.FLAGS.ZERO)
	} else {
		r.FLAG_SET(r.FLAGS.ZERO)
	}
}

// AND is a helper function to and the value to Register A
func (r *RegistersType) AND(value byte) {
	r.SetA(r.A() & value)

	if r.A()&1 != 0 {
		r.FLAG_CLEAR(r.FLAGS.ZERO)
	} else {
		r.FLAG_SET(r.FLAGS.ZERO)
	}

	r.FLAG_CLEAR(r.FLAGS.CARRY | r.FLAGS.SUBTRACT)
	r.FLAG_SET(r.FLAGS.HALF_CARRY)
}

// OR is a helper function to or the value to Register A
func (r *RegistersType) OR(value byte) {
	r.SetA(r.A() | value)

	if r.A()&1 != 0 {
		r.FLAG_CLEAR(r.FLAGS.ZERO)
	} else {
		r.FLAG_SET(r.FLAGS.ZERO)
	}

	r.FLAG_CLEAR(r.FLAGS.CARRY | r.FLAGS.SUBTRACT | r.FLAGS.HALF_CARRY)
}

// XOR is a helper function to xor the value to Register A
func (r *RegistersType) XOR(value byte) {
	r.SetA(r.A() ^ value)

	if r.A()&1 != 0 {
		r.FLAG_CLEAR(r.FLAGS.ZERO)
	} else {
		r.FLAG_SET(r.FLAGS.ZERO)
	}

	r.FLAG_CLEAR(r.FLAGS.CARRY | r.FLAGS.SUBTRACT | r.FLAGS.HALF_CARRY)
}

// INC is a helper function to increment the value
func (r *RegistersType) INC(value byte) byte {
	if (value & 0x0F) == 0x0F {
		r.FLAG_SET(r.FLAGS.HALF_CARRY)
	} else {
		r.FLAG_CLEAR(r.FLAGS.HALF_CARRY)
	}

	value++

	if value&1 != 0 {
		r.FLAG_CLEAR(r.FLAGS.ZERO)
	} else {
		r.FLAG_SET(r.FLAGS.ZERO)
	}

	r.FLAG_SET(r.FLAGS.SUBTRACT)

	return value
}

// dec is a helper function to decrement the value
func (r *RegistersType) DEC(value byte) byte {
	if (value & 0x0F) != 0 {
		r.FLAG_CLEAR(r.FLAGS.HALF_CARRY)
	} else {
		r.FLAG_SET(r.FLAGS.HALF_CARRY)
	}

	value--

	if value&1 != 0 {
		r.FLAG_CLEAR(r.FLAGS.ZERO)
	} else {
		r.FLAG_SET(r.FLAGS.ZERO)
	}

	r.FLAG_SET(r.FLAGS.SUBTRACT)

	return value
}

// dec is a helper function to decrement the value
func (r *RegistersType) ReadOperand(ins *InstructionType, rom *ROMType) {
	switch ins.NumOperands {
	case 0:
		break
	case 1:
		ins.Operand = MMU.ReadByte(CPU.REGISTERS.PC)
	case 2:
		ins.Operand = MMU.ReadShort(CPU.REGISTERS.PC)
	default:
		break
	}
}

// FLAG_SET is a helper function to set flags
func (r *RegistersType) FLAG_SET(flag byte) {
	r.SetF(r.F() | flag)
}

// FLAG_ISSET is a helper function to check if a flag is set
func (r *RegistersType) FLAG_ISSET(flag byte) bool {
	return (r.F() & flag) != 0
}

// FLAG_CLEAR is a helper function to clear flags
func (r *RegistersType) FLAG_CLEAR(flag byte) {
	r.SetF(r.F() & ^flag)
}

func (r *RegistersType) UpdateRegisterTable(registerTreeView *gtk.TreeView, registerListStore *gtk.ListStore) {
	model := []interface{}{}

	registerValues := reflect.ValueOf(*r)
	registers := registerValues.Type()

	for i := 0; i < registerValues.NumField(); i++ {
		row := []string{}

		value, err := registerValues.Field(i).Interface().(uint16)
		if !err {
		} // do nothing

		row = append(row, registers.Field(i).Name+":")
		if registers.Field(i).Name == "FLAGS" {
			flagsValues := reflect.ValueOf((*r).FLAGS)
			flags := flagsValues.Type()
			for j := 0; j < flagsValues.NumField(); j++ {
				row = append(row, flags.Field(j).Name)
			}
		} else {
			row = append(row, r.Register16toString(value))
		}
		model = append(model, row)
	}

	registerListStore.Clear()
	registerLength := len(model)
	for i := 0; i < registerLength; i++ {
		iter := registerListStore.Append()
		row := model[i].([]string)
		err := registerListStore.Set(iter,
			[]int{0, 1},
			[]interface{}{row[0], row[1]})
		if err != nil {
			Logger.Log(LogTypes.ERROR, err)
		}
	}

	registerTreeView.SetModel(registerListStore)
}

// REGISTERS is the exported object used in the CPU
// REGISTERS is exported to become a shared variable in the System object
var REGISTERS = RegistersType{
	AF: 0x01B0,
	BC: 0x0013,
	DE: 0x00D8,
	HL: 0x014D,
	SP: 0xFFFE,
	PC: 0x0100,
	FLAGS: FlagsType{
		ZERO:       0x80,
		SUBTRACT:   0x40,
		HALF_CARRY: 0x20,
		CARRY:      0x10,
	},
}
