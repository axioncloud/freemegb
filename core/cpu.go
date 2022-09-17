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
var INTERRUPTS INTERRUPTSType = INTERRUPTSType{
	master: 0x00,
	enable: 0x00,
	flags:  0x00,
}

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
}

// CPU is the exported object used in the system
// CPU is exported to become a shared variable in the System object
var CPU = CPUType{
	INSTRUCTIONS: INSTRUCTIONS,
	REGISTERS:    &REGISTERS,
	DEBUG:        false,
	STEP:         false,
}

// Run is the thread loop function for the CPU
func (cpu *CPUType) Run(debug bool, registerTreeView *gtk.TreeView, registerListStore *gtk.ListStore) {
	// TODO: Proper CPU control flow with stepping

	// TODO: Proper Breakpoint insertion using an array of addresses
	cpu.DEBUG = debug
	for {
		if cpu.INSTRUCTIONS[ROM.data[cpu.REGISTERS.PC]].Name == "UNKNOWN" {
			var PCString = cpu.REGISTERS.Register16toString(cpu.REGISTERS.PC)
			Logger.Logf(LogTypes.ERROR, "UNKNOWN INSTRUCTION:\n\t\t\t\tINSTRUCTION: 0x%02X\n\t\t\t\tAt ROM Offset: %s\n",
				cpu.INSTRUCTIONS[ROM.data[cpu.REGISTERS.PC]].Opcode, PCString)
			Notify(fmt.Sprintf("INSTRUCTION: 0x%02X\nAt ROM Offset: %s",
				cpu.INSTRUCTIONS[ROM.data[cpu.REGISTERS.PC]].Opcode, PCString))
			break
		}
		Logger.Logf(LogTypes.INFO, "Instruction: %s\n", cpu.INSTRUCTIONS[ROM.data[cpu.REGISTERS.PC]].Name)
		if cpu.DEBUG {
			for !cpu.STEP {
				time.Sleep(150 * time.Millisecond)
			}

			cpu.REGISTERS.Print()
			time.Sleep(1 * time.Second)
			cpu.STEP = false
		} else {
			time.Sleep(4 * time.Microsecond)
		}
		cpu.INSTRUCTIONS[ROM.data[cpu.REGISTERS.PC]].Exec()
		cpu.REGISTERS.UpdateRegisterTable(registerTreeView, registerListStore)
		cpu.REGISTERS.PC++

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
	cpu.REGISTERS.FLAG_CLEAR(cpu.REGISTERS.FLAGS.ZERO)
}
