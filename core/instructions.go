package core

// InstructionType is the structure that holds the execution function,
// opcode value, the name, number of operands, operand locations, and CPU cycles
type InstructionType struct {
	Exec        func(op interface{}) // executed code
	Opcode      uint8                // opcode
	Name        string               // name
	NumOperands byte                 // number of operands
	Operand     interface{}          // operands
	Cycles      uint8                // cpu cycles

}

// INSTRUCTIONS is the array holding InstructionType elements to build a ROM execution table
var INSTRUCTIONS = []InstructionType{
	// 0x00 - NOP
	{
		Exec: func(op interface{}) {
		},
		Opcode:      0x00,
		Name:        "NOP",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x01 - LOAD BC NN
	{
		Exec: func(op interface{}) {
			REGISTERS.BC = op.(uint16)
		},
		Opcode:      0x01,         // opcode
		Name:        "LOAD BC NN", // name
		NumOperands: 2,            // number of operands
		Operand:     nil,          // operands
		Cycles:      12,           // cpu cycles
	},
	// 0x02 - LOAD BC A
	{
		Exec:        func(op interface{}) { MMU.WriteByte(REGISTERS.BC, REGISTERS.A()) }, // executed code
		Opcode:      0x02,                                                                // opcode
		Name:        "LOAD BC A",                                                         // name
		NumOperands: 0,                                                                   // number of operands
		Operand:     nil,                                                                 // operands
		Cycles:      4,                                                                   // cpu cycles
	},
	// 0x03 - INC BC
	{
		Exec:        func(op interface{}) { REGISTERS.BC++ },
		Opcode:      0x03,
		Name:        "INC BC",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x04 - INC B
	{
		Exec:        func(op interface{}) { REGISTERS.SetB(REGISTERS.INC(REGISTERS.B())) },
		Opcode:      0x04,
		Name:        "INC B",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x05 - DEC B
	{
		Exec:        func(op interface{}) { REGISTERS.SetB(REGISTERS.DEC(REGISTERS.B())) },
		Opcode:      0x05,
		Name:        "DEC B",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x06 - LOAD B N
	{
		Exec: func(op interface{}) {
			REGISTERS.SetB(op.(uint8))
		},
		Opcode:      0x06,
		Name:        "LOAD B N",
		NumOperands: 1,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x07 - RLCA
	{
		Exec: func(op interface{}) {
			var carry = (REGISTERS.A() & 0x80) >> 7
			if carry != 0 {
				REGISTERS.FLAG_SET(REGISTERS.FLAGS.CARRY)
			} else {
				REGISTERS.FLAG_CLEAR(REGISTERS.FLAGS.CARRY)
			}

			REGISTERS.SetA(REGISTERS.A() << 1)
			REGISTERS.SetA(REGISTERS.A() + carry)

			REGISTERS.FLAG_CLEAR(REGISTERS.FLAGS.SUBTRACT)
			REGISTERS.FLAG_CLEAR(REGISTERS.FLAGS.ZERO)
			REGISTERS.FLAG_CLEAR(REGISTERS.FLAGS.HALF_CARRY)
		},
		Opcode:      0x07,
		Name:        "RLCA",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x08 - LOAD NN SP
	{
		Exec: func(op interface{}) {
			MMU.WriteShort(REGISTERS.SP, op.(uint16))
		},
		Opcode:      0x08,
		Name:        "LOAD NN SP",
		NumOperands: 2,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x09 - ADD HL BC
	{
		Exec: func(op interface{}) {
			REGISTERS.HL = REGISTERS.ADD16(REGISTERS.HL, REGISTERS.BC)
		},
		Opcode:      0x09,
		Name:        "ADD HL BC",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x0A - LOAD A BC*
	{
		Exec: func(op interface{}) {
			REGISTERS.SetA(MMU.ReadByte(REGISTERS.BC))
		},
		Opcode:      0x0A,
		Name:        "LOAD A BC*",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x0B - DEC BC
	{
		Exec:        func(op interface{}) { REGISTERS.BC-- },
		Opcode:      0x0B,
		Name:        "DEC BC",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x0C - INC C
	{
		Exec:        func(op interface{}) { REGISTERS.SetC(REGISTERS.INC(REGISTERS.C())) },
		Opcode:      0x0C,
		Name:        "INC C",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x0D - DEC C
	{
		Exec:        func(op interface{}) { REGISTERS.SetC(REGISTERS.DEC(REGISTERS.C())) },
		Opcode:      0x0D,
		Name:        "DEC C",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x0E - LOAD C N
	{
		Exec:        func(op interface{}) { REGISTERS.SetC(op.(uint8)) },
		Opcode:      0x0E,
		Name:        "LOAD C N",
		NumOperands: 1,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x0F - RRCA
	{
		Exec: func(op interface{}) {
			var carry uint8 = REGISTERS.A() & 0x01
			if carry != 0 {
				REGISTERS.FLAG_SET(REGISTERS.FLAGS.CARRY)
			} else {
				REGISTERS.FLAG_CLEAR(REGISTERS.FLAGS.CARRY)
			}

			REGISTERS.SetA(REGISTERS.A() >> 1)
			if carry != 0 {
				REGISTERS.SetA(REGISTERS.A() | 0x80)
			}

			REGISTERS.FLAG_CLEAR(REGISTERS.FLAGS.SUBTRACT)
			REGISTERS.FLAG_CLEAR(REGISTERS.FLAGS.ZERO)
			REGISTERS.FLAG_CLEAR(REGISTERS.FLAGS.HALF_CARRY)
		},
		Opcode:      0x0F,
		Name:        "RRCA",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x10 - STOP
	{
		Exec: func(op interface{}) {
			CPU.RUNNING = false
		},
		Opcode:      0x10,
		Name:        "STOP",
		NumOperands: 1,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x11 - LOAD DE NN
	{
		Exec: func(op interface{}) {
			REGISTERS.DE = op.(uint16)
		},
		Opcode:      0x11,
		Name:        "LOAD DE NN",
		NumOperands: 2,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x12 - LOAD DE* A
	{
		Exec:        func(op interface{}) { MMU.WriteByte(REGISTERS.DE, REGISTERS.A()) },
		Opcode:      0x12,
		Name:        "LOAD DE* A",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x13 - INC DE
	{
		Exec:        func(op interface{}) { REGISTERS.DE++ },
		Opcode:      0x13,
		Name:        "INC DE",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x14 - INC D
	{
		Exec:        func(op interface{}) { REGISTERS.SetD(REGISTERS.INC(REGISTERS.D())) },
		Opcode:      0x14,
		Name:        "INC D",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x15 - DEC D
	{
		Exec:        func(op interface{}) { REGISTERS.SetD(REGISTERS.DEC(REGISTERS.D())) },
		Opcode:      0x15,
		Name:        "INC D",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x16 - LOAD D N
	{
		Exec: func(op interface{}) {
			REGISTERS.SetD(op.(uint8))
		},
		Opcode:      0x16,
		Name:        "LOAD D N",
		NumOperands: 1,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x17 - RLA
	{
		Exec: func(op interface{}) {
			var carry int = 0
			if REGISTERS.FLAG_ISSET(REGISTERS.FLAGS.CARRY) {
				carry = 1
			}

			if REGISTERS.A()&0x08 != 0 {
				REGISTERS.FLAG_SET(REGISTERS.FLAGS.CARRY)
			} else {
				REGISTERS.FLAG_CLEAR(REGISTERS.FLAGS.CARRY)
			}

			REGISTERS.SetA(REGISTERS.A() << 1)
			REGISTERS.SetA(REGISTERS.A() + byte(carry))

			REGISTERS.FLAG_CLEAR(REGISTERS.FLAGS.SUBTRACT)
			REGISTERS.FLAG_CLEAR(REGISTERS.FLAGS.ZERO)
			REGISTERS.FLAG_CLEAR(REGISTERS.FLAGS.HALF_CARRY)
		},
		Opcode:      0x17,
		Name:        "RLA",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x18 - JUMP PC+N
	{
		// Set PC to PC + Operand
		Exec:        func(op interface{}) { REGISTERS.PC += uint16(op.(uint8)) },
		Opcode:      0x18,
		Name:        "JUMP PC+N",
		NumOperands: 1,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x19 - ADD HL DE
	{
		Exec: func(op interface{}) {
			REGISTERS.HL += REGISTERS.ADD16(REGISTERS.HL, REGISTERS.DE)
		},
		Opcode:      0x19,
		Name:        "ADD HL DE",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x1A - LOAD A DE*
	{
		Exec: func(op interface{}) {
			REGISTERS.SetA(MMU.ReadByte(REGISTERS.DE))
		},
		Opcode:      0x1A,
		Name:        "LOAD A DE*",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x1B - DEC DE
	{
		Exec:        func(op interface{}) { REGISTERS.DE-- },
		Opcode:      0x1B,
		Name:        "DEC DE",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x1C - INC
	{
		Exec:        func(op interface{}) { REGISTERS.SetE(REGISTERS.INC(REGISTERS.E())) },
		Opcode:      0x1C,
		Name:        "INC",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x1D - DEC E
	{
		Exec:        func(op interface{}) { REGISTERS.SetE(REGISTERS.DEC(REGISTERS.E())) },
		Opcode:      0x1D,
		Name:        "DEC E",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x1E - LOAD E N
	{
		Exec:        func(op interface{}) { REGISTERS.SetE(op.(uint8)) },
		Opcode:      0x1E,
		Name:        "LOAD E N",
		NumOperands: 1,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x1F - RRA
	{
		Exec: func(op interface{}) {
			var carry int = 0
			if REGISTERS.FLAG_ISSET(REGISTERS.FLAGS.CARRY) {
				carry = 1 << 7
			}

			if REGISTERS.A()&0x01 != 0 {
				REGISTERS.FLAG_SET(REGISTERS.FLAGS.CARRY)
			} else {
				REGISTERS.FLAG_CLEAR(REGISTERS.FLAGS.CARRY)
			}

			REGISTERS.SetA(REGISTERS.A() >> 1)
			REGISTERS.SetA(REGISTERS.A() + byte(carry))

			REGISTERS.FLAG_CLEAR(REGISTERS.FLAGS.SUBTRACT)
			REGISTERS.FLAG_CLEAR(REGISTERS.FLAGS.ZERO)
			REGISTERS.FLAG_CLEAR(REGISTERS.FLAGS.HALF_CARRY)
		},
		Opcode:      0x1F,
		Name:        "RRA",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x20 - JR NZ N
	{
		Exec: func(op interface{}) {
			if REGISTERS.FLAG_ISSET(REGISTERS.FLAGS.ZERO) {

			} else {
				REGISTERS.PC += uint16(op.(uint8))
			}
		},
		Opcode:      0x20,
		Name:        "JR NZ N",
		NumOperands: 1,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x21 - LOAD NN HL
	{
		Exec: func(op interface{}) {
			REGISTERS.HL = op.(uint16)
		},
		Opcode:      0x21,
		Name:        "LOAD NN HL",
		NumOperands: 2,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x22 - LOAD HL*++ A
	{
		Exec: func(op interface{}) {
			MMU.WriteByte(REGISTERS.HL, REGISTERS.A())
			REGISTERS.HL++
		},
		Opcode:      0x22,
		Name:        "LOAD HL*++ A",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x23 - INC HL
	{
		Exec: func(op interface{}) {
			REGISTERS.HL++
		},
		Opcode:      0x23,
		Name:        "INC HL",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x24 - INC H
	{
		Exec: func(op interface{}) {
			REGISTERS.SetH(REGISTERS.INC(REGISTERS.H()))
		},
		Opcode:      0x24,
		Name:        "INC H",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x25 - DEC H
	{
		Exec: func(op interface{}) {
			REGISTERS.SetH(REGISTERS.DEC(REGISTERS.H()))
		},
		Opcode:      0x25,
		Name:        "DEC H",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x26 - LOAD N H
	{
		Exec: func(op interface{}) {
			REGISTERS.SetH(op.(uint8))
		},
		Opcode:      0x26,
		Name:        "LOAD N H",
		NumOperands: 1,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x27 - DAA
	{
		Exec: func(op interface{}) {
			var A = uint16(REGISTERS.A())

			if REGISTERS.FLAG_ISSET(REGISTERS.FLAGS.SUBTRACT) {
				if REGISTERS.FLAG_ISSET(REGISTERS.FLAGS.HALF_CARRY) {
					A = (A - 0x06) & 0xFF
				}
				if REGISTERS.FLAG_ISSET(REGISTERS.FLAGS.CARRY) {
					A -= 0x60
				}
			} else {
				if REGISTERS.FLAG_ISSET(REGISTERS.FLAGS.HALF_CARRY) || (A&0xF) > 9 {
					A += 0x06
				}
				if REGISTERS.FLAG_ISSET(REGISTERS.FLAGS.CARRY) || A > 0x9F {
					A += 0x60
				}
			}
			REGISTERS.SetA(uint8(A))

			REGISTERS.FLAG_CLEAR(REGISTERS.FLAGS.HALF_CARRY)

			if REGISTERS.A() != 0 {
				REGISTERS.FLAG_CLEAR(REGISTERS.FLAGS.ZERO)
			} else {
				REGISTERS.FLAG_SET(REGISTERS.FLAGS.ZERO)
			}

			if A >= 0x100 {
				REGISTERS.FLAG_SET(REGISTERS.FLAGS.CARRY)
			}
		},
		Opcode:      0x27,
		Name:        "DAA",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x28 - JUMP Z N
	{
		Exec: func(op interface{}) {
			if REGISTERS.FLAG_ISSET(REGISTERS.FLAGS.ZERO) {

				REGISTERS.PC += uint16(op.(uint8))
			}
		},
		Opcode:      0x28,
		Name:        "JUMP NZ N",
		NumOperands: 1,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x29 - ADD HL HL
	{
		Exec: func(op interface{}) {
			REGISTERS.HL = REGISTERS.ADD16(REGISTERS.HL, REGISTERS.HL)
		},
		Opcode:      0x29,
		Name:        "ADD HL HL",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x2A - LOAD A HL*++
	{
		Exec: func(op interface{}) {
			REGISTERS.SetA(MMU.ReadByte(REGISTERS.HL))
			REGISTERS.HL++
		},
		Opcode:      0x2A,
		Name:        "LOAD A HL*++",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x2B - DEC HL
	{
		Exec:        func(op interface{}) { REGISTERS.HL-- },
		Opcode:      0x2B,
		Name:        "DEC HL",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x2C - INC L
	{
		Exec:        func(op interface{}) { REGISTERS.SetL(REGISTERS.INC(REGISTERS.L())) },
		Opcode:      0x2C,
		Name:        "INC L",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x2D - DEC L
	{
		Exec:        func(op interface{}) { REGISTERS.SetL(REGISTERS.DEC(REGISTERS.L())) },
		Opcode:      0x2D,
		Name:        "DEC L",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x2E - LOAD L N
	{
		Exec:        func(op interface{}) { REGISTERS.SetL(op.(uint8)) },
		Opcode:      0x2E,
		Name:        "LOAD L N",
		NumOperands: 1,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x2F - CPL
	{
		Exec: func(op interface{}) {
			REGISTERS.SetA(^REGISTERS.A())
			REGISTERS.FLAG_SET(REGISTERS.FLAGS.SUBTRACT)
			REGISTERS.FLAG_SET(REGISTERS.FLAGS.HALF_CARRY)
		},
		Opcode:      0x2F,
		Name:        "CPL",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x30 - JUMP NC N
	{
		Exec: func(op interface{}) {
			if REGISTERS.FLAG_ISSET(REGISTERS.FLAGS.CARRY) {

			} else {
				REGISTERS.PC += uint16(op.(uint8))
			}
		},
		Opcode:      0x30,
		Name:        "JUMP NC N",
		NumOperands: 1,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x31 - LOAD NN SP
	{
		Exec: func(op interface{}) {
			REGISTERS.SP = op.(uint16)
		},
		Opcode:      0x31,
		Name:        "LOAD NN SP",
		NumOperands: 2,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x32 - LOAD HL*-- A
	{
		Exec: func(op interface{}) {
			MMU.WriteByte(REGISTERS.HL, REGISTERS.A())
			REGISTERS.HL--
		},
		Opcode:      0x32,
		Name:        "LOAD HL*-- A",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x33 - INC SP
	{
		Exec: func(op interface{}) {
			REGISTERS.SP++
		},
		Opcode:      0x33,
		Name:        "INC SP",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x34 - INC HL*
	{
		Exec:        func(op interface{}) { MMU.WriteByte(REGISTERS.HL, REGISTERS.INC(MMU.ReadByte(REGISTERS.HL))) },
		Opcode:      0x34,
		Name:        "INC HL*",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x35 - DEC HL*
	{
		Exec:        func(op interface{}) { MMU.WriteByte(REGISTERS.HL, REGISTERS.DEC(MMU.ReadByte(REGISTERS.HL))) },
		Opcode:      0x35,
		Name:        "DEC HL*",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x36 - LOAD HL N
	{
		Exec: func(op interface{}) {
			MMU.WriteByte(REGISTERS.HL, op.(uint8))
		},
		Opcode:      0x36,
		Name:        "LOAD HL N",
		NumOperands: 1,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x37 - SCF
	{
		Exec: func(op interface{}) {
			REGISTERS.FLAG_SET(REGISTERS.FLAGS.CARRY)
			REGISTERS.FLAG_CLEAR(REGISTERS.FLAGS.ZERO)
			REGISTERS.FLAG_CLEAR(REGISTERS.FLAGS.HALF_CARRY)
		},
		Opcode:      0x37,
		Name:        "SCF",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x38 - JUMP C N
	{
		Exec: func(op interface{}) {
			if REGISTERS.FLAG_ISSET(REGISTERS.FLAGS.CARRY) {
				REGISTERS.PC += uint16(op.(uint8))
			}
		},
		Opcode:      0x38,
		Name:        "JUMP C N",
		NumOperands: 1,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x39 - ADD HL SP
	{
		Exec:        func(op interface{}) { REGISTERS.HL = REGISTERS.ADD16(REGISTERS.HL, REGISTERS.SP) },
		Opcode:      0x39,
		Name:        "ADD HL SP",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x3A - LOAD A HL*--
	{
		Exec: func(op interface{}) {
			REGISTERS.SetA(MMU.ReadByte(REGISTERS.HL))
			REGISTERS.HL--
		},
		Opcode:      0x3A,
		Name:        "LOAD A HL*--",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x3B - DEC SP
	{
		Exec:        func(op interface{}) { REGISTERS.SP-- },
		Opcode:      0x3B,
		Name:        "DEC SP",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x3C - INC A
	{
		Exec:        func(op interface{}) { REGISTERS.SetA(REGISTERS.INC(REGISTERS.A())) },
		Opcode:      0x3C,
		Name:        "INC A",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x3D - DEC A
	{
		Exec:        func(op interface{}) { REGISTERS.SetA(REGISTERS.DEC(REGISTERS.A())) },
		Opcode:      0x3D,
		Name:        "DEC A",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x3E - LOAD A N
	{
		Exec:        func(op interface{}) { REGISTERS.SetA(op.(uint8)) },
		Opcode:      0x3E,
		Name:        "LOAD A N",
		NumOperands: 1,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x3F - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x3F,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x40 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x40,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x41 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x41,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x42 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x42,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x43 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x43,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x44 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x44,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x45 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x45,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x46 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x46,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x47 - LOAD B A
	{
		Exec:        func(op interface{}) { REGISTERS.SetB(REGISTERS.A()) },
		Opcode:      0x47,
		Name:        "LOAD B A",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x48 - LOAD C B
	{
		Exec:        func(op interface{}) { REGISTERS.SetC(REGISTERS.B()) },
		Opcode:      0x48,
		Name:        "LOAD C B",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x49 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x49,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x4A - LOAD C D
	{
		Exec:        func(op interface{}) { REGISTERS.SetC(REGISTERS.D()) },
		Opcode:      0x4A,
		Name:        "LOAD C D",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x4B - LOAD C E
	{
		Exec:        func(op interface{}) { REGISTERS.SetC(REGISTERS.E()) },
		Opcode:      0x4B,
		Name:        "LOAD C E",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x4C - LOAD C H
	{
		Exec:        func(op interface{}) { REGISTERS.SetC(REGISTERS.H()) },
		Opcode:      0x4C,
		Name:        "LOAD C H",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x4D - LOAD C L
	{
		Exec:        func(op interface{}) { REGISTERS.SetC(REGISTERS.L()) },
		Opcode:      0x4D,
		Name:        "LOAD C L",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x4E - LOAD C HL*
	{
		Exec:        func(op interface{}) { REGISTERS.SetC(MMU.ReadByte(REGISTERS.HL)) },
		Opcode:      0x4E,
		Name:        "LOAD C HL*",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x4F - LOAD C A
	{
		Exec:        func(op interface{}) { REGISTERS.SetC(REGISTERS.A()) },
		Opcode:      0x4F,
		Name:        "LOAD C A",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x50 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x50,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x51 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x51,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x52 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x52,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x53 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x53,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x54 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x54,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x55 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x55,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x56 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x56,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x57 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x57,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x58 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x58,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x59 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x59,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x5A - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x5A,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x5B - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x5B,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x5C - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x5C,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x5D - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x5D,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x5E - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x5E,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x5F - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x5F,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x60 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x60,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x61 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x61,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x62 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x62,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x63 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x63,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x64 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x64,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x65 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x65,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x66 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x66,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x67 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x67,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x68 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x68,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x69 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x69,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x6A - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x6A,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x6B - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x6B,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x6C - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x6C,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x6D - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x6D,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x6E - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x6E,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x6F - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x6F,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x70 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x70,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x71 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x71,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x72 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x72,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x73 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x73,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x74 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x74,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x75 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x75,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x76 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x76,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x77 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x77,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x78 - LOAD A B
	{
		Exec: func(op interface{}) {
			REGISTERS.SetA(REGISTERS.B())
		},
		Opcode:      0x78,
		Name:        "LOAD A B",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x79 - LOAD A C
	{
		Exec: func(op interface{}) {
			REGISTERS.SetA(REGISTERS.C())
		},
		Opcode:      0x79,
		Name:        "LOAD A C",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x7A - LOAD A D
	{
		Exec: func(op interface{}) {
			REGISTERS.SetA(REGISTERS.D())
		},
		Opcode:      0x7A,
		Name:        "LOAD A D",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x7B - LOAD A E
	{
		Exec: func(op interface{}) {
			REGISTERS.SetA(REGISTERS.E())
		},
		Opcode:      0x7B,
		Name:        "LOAD A E",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x7C - LOAD A H
	{
		Exec: func(op interface{}) {
			REGISTERS.SetA(REGISTERS.H())
		},
		Opcode:      0x7C,
		Name:        "LOAD A H",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x7D - LOAD A L
	{
		Exec: func(op interface{}) {
			REGISTERS.SetA(REGISTERS.L())
		},
		Opcode:      0x7D,
		Name:        "LOAD A L",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x7E - LOAD A HL
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x7E,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x7F - LOAD A A
	{
		Exec:        func(op interface{}) {},
		Opcode:      0x7F,
		Name:        "LOAD A A",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x80 - ADD A B
	{
		Exec: func(op interface{}) {
			REGISTERS.SetA(REGISTERS.ADD8(REGISTERS.A(), REGISTERS.B()))
		},
		Opcode:      0x80,
		Name:        "ADD A B",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x81 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x81,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x82 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x82,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x83 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x83,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x84 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x84,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x85 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x85,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x86 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x86,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x87 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x87,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x88 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x88,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x89 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x89,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x8A - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x8A,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x8B - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x8B,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x8C - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x8C,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x8D - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x8D,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x8E - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x8E,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x8F - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x8F,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x90 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x90,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x91 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x91,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x92 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x92,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x93 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x93,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x94 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x94,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x95 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x95,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x96 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x96,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x97 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x97,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x98 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x98,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x99 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x99,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x9A - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x9A,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x9B - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x9B,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x9C - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x9C,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x9D - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x9D,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x9E - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x9E,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0x9F - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0x9F,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xA0 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xA0,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xA1 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xA1,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xA2 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xA2,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xA3 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xA3,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xA4 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xA4,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xA5 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xA5,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xA6 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xA6,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xA7 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xA7,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xA8 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xA8,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xA9 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xA9,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xAA - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xAA,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xAB - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xAB,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xAC - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xAC,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xAD - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xAD,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xAE - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xAE,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xAF - XOR A
	{
		Exec:        func(op interface{}) { REGISTERS.XOR(REGISTERS.A()) },
		Opcode:      0xAF,
		Name:        "XOR A",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xB0 - OR B
	{
		Exec:        func(op interface{}) { REGISTERS.OR(REGISTERS.B()) },
		Opcode:      0xB0,
		Name:        "OR B",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xB1 - OR C
	{
		Exec:        func(op interface{}) { REGISTERS.OR(REGISTERS.C()) },
		Opcode:      0xB1,
		Name:        "OR C",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xB2 - OR D
	{
		Exec:        func(op interface{}) { REGISTERS.OR(REGISTERS.D()) },
		Opcode:      0xB2,
		Name:        "OR D",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xB3 - OR E
	{
		Exec:        func(op interface{}) { REGISTERS.OR(REGISTERS.E()) },
		Opcode:      0xB3,
		Name:        "OR E",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xB4 - OR H
	{
		Exec:        func(op interface{}) { REGISTERS.OR(REGISTERS.H()) },
		Opcode:      0xB4,
		Name:        "OR H",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xB5 - OR L
	{
		Exec:        func(op interface{}) { REGISTERS.OR(REGISTERS.L()) },
		Opcode:      0xB5,
		Name:        "OR L",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xB6 - OR HL
	{
		Exec:        func(op interface{}) { REGISTERS.OR(MMU.ReadByte(REGISTERS.HL)) },
		Opcode:      0xB6,
		Name:        "OR HL",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xB7 - OR A
	{
		Exec:        func(op interface{}) { REGISTERS.OR(REGISTERS.A()) },
		Opcode:      0xB7,
		Name:        "OR A",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xB8 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xB8,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xB9 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xB9,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xBA - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xBA,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xBB - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xBB,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xBC - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xBC,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xBD - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xBD,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xBE - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xBE,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xBF - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xBF,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xC0 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xC0,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xC1 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xC1,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xC2 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xC2,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xC3 - JUMP NN
	{
		Exec: func(op interface{}) {
			REGISTERS.PC = op.(uint16)
		},
		Opcode:      0xC3,
		Name:        "JUMP NN",
		NumOperands: 2,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xC4 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xC4,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xC5 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xC5,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xC6 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xC6,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xC7 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xC7,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xC8 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xC8,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xC9 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xC9,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xCA - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xCA,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xCB - CB N
	{
		Exec:        func(op interface{}) { Logger.Log(LogTypes.WARNING, "TODO 'CB N'") },
		Opcode:      0xCB,
		Name:        "CB N",
		NumOperands: 1,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xCC - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xCC,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xCD - CALL NN
	{
		Exec: func(op interface{}) {
			MMU.WriteShortToStack(REGISTERS.PC)
			REGISTERS.PC = op.(uint16)
		},
		Opcode:      0xCD,
		Name:        "CALL NN",
		NumOperands: 2,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xCE - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xCE,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xCF - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xCF,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xD0 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xD0,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xD1 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xD1,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xD2 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xD2,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xD3 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xD3,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xD4 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xD4,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xD5 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xD5,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xD6 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xD6,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xD7 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xD7,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xD8 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xD8,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xD9 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xD9,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xDA - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xDA,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xDB - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xDB,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xDC - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xDC,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xDD - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xDD,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xDE - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xDE,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xDF - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xDF,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xE0 - LOAD 0xFF00 N A
	{
		Exec: func(op interface{}) {
			MMU.WriteByte(0xFF00+uint16(op.(uint8)), REGISTERS.A())
		},
		Opcode:      0xE0,
		Name:        "LOAD 0xFF00 N A",
		NumOperands: 1,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xE1 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xE1,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xE2 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xE2,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xE3 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xE3,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xE4 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xE4,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xE5 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xE5,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xE6 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xE6,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xE7 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xE7,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xE8 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xE8,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xE9 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xE9,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xEA - LOAD NNP A
	{
		Exec: func(op interface{}) {
			MMU.WriteByte(op.(uint16), REGISTERS.A())
		},
		Opcode:      0xEA,
		Name:        "LOAD NNP A",
		NumOperands: 2,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xEB - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xEB,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xEC - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xEC,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xED - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xED,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xEE - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xEE,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xEF - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xEF,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xF0 - LOAD A 0xFF00 N
	{
		Exec: func(op interface{}) {
			REGISTERS.SetA(MMU.ReadByte(0xFF00 + uint16(op.(uint8))))
		},
		Opcode:      0xF0,
		Name:        "LOAD 0xFF00 N",
		NumOperands: 1,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xF1 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xF1,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xF2 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xF2,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xF3 - DISABLE INTERRUPTS
	{
		Exec:        func(op interface{}) { INTERRUPTS.master = 0 },
		Opcode:      0xF3,
		Name:        "DI",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xF4 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xF4,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xF5 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xF5,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xF6 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xF6,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xF7 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xF7,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xF8 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xF8,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xF9 - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xF9,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xFA - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xFA,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xFB - ENABLE INTERRUPTS
	{
		Exec:        func(op interface{}) { INTERRUPTS.master = 1 },
		Opcode:      0xFB,
		Name:        "EI",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xFC - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xFC,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xFD - UNKNOWN
	{
		Exec:        func(op interface{}) { Logger.Panic("INSTRUCTION NOT IMPLEMENTED") },
		Opcode:      0xFD,
		Name:        "UNKNOWN",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xFE - CP A N
	{
		Exec: func(op interface{}) {
			REGISTERS.FLAG_SET(REGISTERS.FLAGS.SUBTRACT)
			var operand = op.(uint8)

			if REGISTERS.A() == operand {
				REGISTERS.FLAG_SET(REGISTERS.FLAGS.ZERO)
			} else {
				REGISTERS.FLAG_CLEAR(REGISTERS.FLAGS.ZERO)
			}

			if REGISTERS.A() < operand {
				REGISTERS.FLAG_SET(REGISTERS.FLAGS.CARRY)
			} else {
				REGISTERS.FLAG_CLEAR(REGISTERS.FLAGS.CARRY)
			}

			if (REGISTERS.A() & 0x0f) < (operand & 0x0f) {
				REGISTERS.FLAG_SET(REGISTERS.FLAGS.HALF_CARRY)
			} else {
				REGISTERS.FLAG_CLEAR(REGISTERS.FLAGS.HALF_CARRY)
			}
		},
		Opcode:      0xFE,
		Name:        "CP A N",
		NumOperands: 1,
		Operand:     nil,
		Cycles:      4,
	},
	// 0xFF - RST 38
	{
		Exec: func(op interface{}) {
			MMU.WriteShortToStack(REGISTERS.PC)
			REGISTERS.PC = 0x0038
		},
		Opcode:      0xFF,
		Name:        "RST 38",
		NumOperands: 0,
		Operand:     nil,
		Cycles:      4,
	},
}
