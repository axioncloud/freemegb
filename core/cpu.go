package core

import (
	"fmt"
	"time"

	"github.com/gotk3/gotk3/gtk"
	// "log"
	// "github.com/gotk3/gotk3/glib"
)

// INTERRUPTSType is the structure to define constant values used to identify an interrupt
type INTERRUPTSType struct {
	master byte
	enable byte
	flags  byte
}

// INTERRUPTS is the exported object used in the system
//
// INTERRUPTS is exported for value setting in other files
var INTERRUPTS INTERRUPTSType

// CPUType is the structure to define what's inside a CPU
//
//	CPU Structure
//	================
//	---> Instructions Array
//	---> Registers Structure
//	---> DEBUG boolean value set with CPU.Run()
//	================
type CPUType struct {
	INSTRUCTIONS []InstructionType
	REGISTERS    *RegistersType
	DEBUG        bool
	STEP         bool
	KEEP_STEP    bool
	RUNNING      bool
	PAUSED       bool
	BREAKPOINTS  map[uint16]bool
}

// CPU is the exported object used in the system
// CPU is exported to become a shared variable in the System object
var CPU CPUType

func init() {
	CPU = CPUType{
		INSTRUCTIONS: INSTRUCTIONS,
		REGISTERS:    &REGISTERS,
		DEBUG:        false,
		STEP:         false,
		RUNNING:      false,
		PAUSED:       false,
		BREAKPOINTS:  make(map[uint16]bool),
	}
	INTERRUPTS = INTERRUPTSType{
		master: 0x00,
		enable: 0x00,
		flags:  0x00,
	}
}

// Run is the thread loop function for the CPU
func (cpu *CPUType) Run(debug bool, registerTreeView *gtk.TreeView, registerListStore *gtk.ListStore) {
	// TODO: Proper CPU control flow with stepping
	cpu.BREAKPOINTS[0x101] = true

	// TODO: Proper Breakpoint insertion using an array of addresses
	cpu.DEBUG = debug
	cpu.RUNNING = true
	for {
		if !cpu.RUNNING {
			break
		}
		bp_enabled, bp_exists := cpu.BREAKPOINTS[cpu.REGISTERS.PC]
		if bp_exists && bp_enabled {
			cpu.DEBUG = true
			cpu.STEP = true
			cpu.KEEP_STEP = true
		}
		if cpu.KEEP_STEP {
			cpu.STEP = true
		}

		if cpu.INSTRUCTIONS[ROM.data[cpu.REGISTERS.PC]].Name == "UNKNOWN" {
			var PCString = cpu.REGISTERS.Register16toString(cpu.REGISTERS.PC)
			Logger.Logf(LogTypes.ERROR, "UNKNOWN INSTRUCTION:\n\t\t\t\tINSTRUCTION: 0x%02X\n\t\t\t\tAt ROM Offset: %s\n",
				cpu.INSTRUCTIONS[ROM.data[cpu.REGISTERS.PC]].Opcode, PCString)
			Notify(fmt.Sprintf("INSTRUCTION: 0x%02X\nAt ROM Offset: %s",
				cpu.INSTRUCTIONS[ROM.data[cpu.REGISTERS.PC]].Opcode, PCString))
			break
		}
		Logger.Logf(LogTypes.INFO, "Instruction: %s\n", cpu.INSTRUCTIONS[ROM.data[cpu.REGISTERS.PC]].Name)

		for cpu.PAUSED {
			time.Sleep(400 * time.Millisecond)
		}

		if cpu.DEBUG {
			for cpu.STEP && cpu.RUNNING {
				time.Sleep(150 * time.Millisecond)
			}

			cpu.REGISTERS.Print()
			time.Sleep(500 * time.Millisecond)
		} else {
			time.Sleep(80 * time.Millisecond)
		}
		var instruction = cpu.INSTRUCTIONS[ROM.data[cpu.REGISTERS.PC]]
		cpu.REGISTERS.PC++
		if instruction.NumOperands != 0 {
			cpu.REGISTERS.ReadOperand(&instruction, &ROM)
		}
		cpu.REGISTERS.PC += uint16(instruction.NumOperands)
		instruction.Exec(instruction.Operand)
		instruction.Operand = nil
		cpu.REGISTERS.UpdateRegisterTable(registerTreeView, registerListStore)

	}
	//	finished <- true
}

// Reset will reset the CPU, INTERRUPTS and REGISTERS to their default values
func (cpu *CPUType) Reset() {
	cpu.REGISTERS.AF = 0x01B0
	cpu.REGISTERS.BC = 0x0013
	cpu.REGISTERS.DE = 0x00D8
	cpu.REGISTERS.HL = 0x01B0
	cpu.REGISTERS.SP = 0xFFFE
	cpu.REGISTERS.PC = 0x0100
	cpu.RUNNING = false
	cpu.REGISTERS.FLAG_CLEAR(cpu.REGISTERS.FLAGS.ZERO)
	for breakpoint := range cpu.BREAKPOINTS {
		delete(cpu.BREAKPOINTS, breakpoint)
	}
}
