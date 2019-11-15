package components

import (
	"fmt"
)

type RegistersType struct {
	AF	uint16
	BC	uint16
	DE	uint16
	HL	uint16
	SP	uint16
	PC	uint16
}

/*
	REGISTER A
*/
func (r *RegistersType) SetA(value byte) {
	var F = r.AF & 0x00FF
	r.AF = uint16(value) << 8 | F
}
func (r *RegistersType) A() byte {
	return byte(r.AF >> 8)
}

/*
	REGISTER F
*/
func (r *RegistersType) SetF(value byte) {
	var A = r.AF & 0xFF00
	r.AF = A << 8 | uint16(value)
}
func (r *RegistersType) F() byte {
	return byte(r.AF >> 8)
}

/*
	REGISTER B
*/
func (r *RegistersType) SetB(value byte) {
	var C = r.BC & 0x00FF
	r.BC = uint16(value) << 8 | C
}
func (r *RegistersType) B() byte {
	return byte(r.BC >> 8)
}

/*
	REGISTER C
*/
func (r *RegistersType) SetC(value byte) {
	var B = r.BC & 0xFF00
	r.BC = B << 8 | uint16(value)
}
func (r *RegistersType) C() byte {
	return byte(r.BC & 0x00FF)
}


/*
	REGISTER D
*/
func (r *RegistersType) SetD(value byte) {
	var E = r.BC & 0x00FF
	r.DE = uint16(value) << 8 | E
}
func (r *RegistersType) D() byte {
	return byte(r.BC >> 8)
}

/*
	REGISTER E
*/
func (r *RegistersType) SetE(value byte) {
	var D = r.DE & 0xFF00
	r.DE = D << 8 | uint16(value)
}
func (r *RegistersType) E() byte {
	return byte(r.BC & 0x00FF)
}

/*
	REGISTER H
*/
func (r *RegistersType) SetH(value byte) {
	var L = r.HL & 0x00FF
	r.HL = uint16(value) << 8 | L
}
func (r *RegistersType) H() byte {
	return byte(r.HL >> 8)
}

/*
	REGISTER L
*/
func (r *RegistersType) SetL(value byte) {
	var H = r.DE & 0xFF00
	r.HL = H << 8 | uint16(value)
}
func (r *RegistersType) L() byte {
	return byte(r.HL & 0x00FF)
}

func (r *RegistersType) Print() {
	fmt.Printf("AF: 0x%04X\nBC: 0x%04X\nDE: 0x%04X\nHL: 0x%04X\nSP: 0x%04X\nPC: 0x%04X\n", r.AF, r.BC, r.DE, r.HL, r.SP, r.PC)
}

func (r *RegistersType) PrintRegister16(register uint16) {
	fmt.Println(fmt.Sprintf("16-bit Register: 0x%04X\n", register))
}

func (r *RegistersType) PrintRegister8(register byte) {
	fmt.Println(fmt.Sprintf("8-bit Register: 0x%02X\n", register))
}

var REGISTERS = RegistersType {
	AF: 0x01B0,
	BC:	0x0013,
	DE:	0x00D8,
	HL:	0x014D,
	SP:	0xFFFE,
	PC:	0x0100,
}
