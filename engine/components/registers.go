package components

import (
	"fmt"
)

type FlagsType struct {
	ZERO       byte
	SUBTRACT   byte
	HALF_CARRY byte
	CARRY      byte
}

// TODO: Finish Flags Structure and functions
type RegistersType struct {
	AF    uint16
	BC    uint16
	DE    uint16
	HL    uint16
	SP    uint16
	PC    uint16
	FLAGS FlagsType
}

/*
	REGISTER A
*/
func (r *RegistersType) SetA(value byte) {
	var F = r.AF & 0x00FF
	r.AF = uint16(value)<<8 | F
}
func (r *RegistersType) A() byte {
	return byte(r.AF >> 8)
}

/*
	REGISTER F
*/
func (r *RegistersType) SetF(value byte) {
	var A = r.AF & 0xFF00
	r.AF = A<<8 | uint16(value)
}
func (r *RegistersType) F() byte {
	return byte(r.AF >> 8)
}

/*
	REGISTER B
*/
func (r *RegistersType) SetB(value byte) {
	var C = r.BC & 0x00FF
	r.BC = uint16(value)<<8 | C
}
func (r *RegistersType) B() byte {
	return byte(r.BC >> 8)
}

/*
	REGISTER C
*/
func (r *RegistersType) SetC(value byte) {
	var B = r.BC & 0xFF00
	r.BC = B<<8 | uint16(value)
}
func (r *RegistersType) C() byte {
	return byte(r.BC & 0x00FF)
}

/*
	REGISTER D
*/
func (r *RegistersType) SetD(value byte) {
	var E = r.BC & 0x00FF
	r.DE = uint16(value)<<8 | E
}
func (r *RegistersType) D() byte {
	return byte(r.BC >> 8)
}

/*
	REGISTER E
*/
func (r *RegistersType) SetE(value byte) {
	var D = r.DE & 0xFF00
	r.DE = D<<8 | uint16(value)
}
func (r *RegistersType) E() byte {
	return byte(r.BC & 0x00FF)
}

/*
	REGISTER H
*/
func (r *RegistersType) SetH(value byte) {
	var L = r.HL & 0x00FF
	r.HL = uint16(value)<<8 | L
}
func (r *RegistersType) H() byte {
	return byte(r.HL >> 8)
}

/*
	REGISTER L
*/
func (r *RegistersType) SetL(value byte) {
	var H = r.DE & 0xFF00
	r.HL = H<<8 | uint16(value)
}
func (r *RegistersType) L() byte {
	return byte(r.HL & 0x00FF)
}

func (r *RegistersType) Print() {
	Logger.Printf("REGISTERS:\n\tAF: 0x%04X\n\tBC: 0x%04X\n\tDE: 0x%04X\n\tHL: 0x%04X\n\tSP: 0x%04X\n\tPC: 0x%04X\n", r.AF, r.BC, r.DE, r.HL, r.SP, r.PC)
}

func (r *RegistersType) Register16toString(register uint16) string {
	return fmt.Sprintf("16-bit Register: 0x%04X\n", register)
}

func (r *RegistersType) Register8toString(register byte) string {
	return fmt.Sprintf("8-bit Register: 0x%02X\n", register)
}

func (r *RegistersType) CombineTo16(lower byte, upper byte) uint16 {
	return uint16(uint16(lower)<<8 | uint16(upper))
}

func (r *RegistersType) Add8(destination byte, source byte) byte {
	result := destination + source

	if result&0xFF != 0 {
		r.FLAG_SET(r.FLAGS.CARRY)
	} else {
		r.FLAG_CLEAR(r.FLAGS.CARRY)
	}

	destination = result & 0xFF

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

	return result
}

func (r *RegistersType) Add16(destination uint16, source uint16) {
	// destination + source
}

func (r *RegistersType) FLAG_SET(flag byte) {
	r.SetF(r.F() | flag)
}

func (r *RegistersType) FLAG_ISSET(flag byte) bool {
	return (r.F() & flag) != 0
}

func (r *RegistersType) FLAG_CLEAR(flag byte) {
	r.SetF(r.F() & ^flag)
}

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