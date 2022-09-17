package core

import (
	"io/ioutil"

	"fmt"
	"time"

	"github.com/gotk3/gotk3/gtk"
)

// SystemType is the object definition type
//
//	System Structure
//	================
//	CPU Structure
//	 ---> Instructions Array
//	 ---> Registers Structure
//	 ---> MMU Structure
//	GPU Structure
//	 ---> Graphics Processing
//	 ---> Screen Pipeline
type SystemType struct {
	CPU *CPUType
	GPU *GPUType
	ROM *ROMType
}

var System = SystemType{
	CPU: &CPU,
	GPU: &GPU,
	ROM: &ROM,
}

func (system *SystemType) LoadROM(location string, romListStore *gtk.ListStore,
	romTreeView *gtk.TreeView, romProgressBar *gtk.ProgressBar,
	menuDebug *gtk.MenuItem, menuRun *gtk.MenuItem) {
	Logger.Log(LogTypes.INFO, "ROM: Loading")
	before := time.Now()
	rom, err := ioutil.ReadFile(location)
	if err == nil {
		ROM.data = rom
		ROM.BuildModel()

		ROM.romName = ROM.GetName()
		Logger.Log(LogTypes.INFO, "Found ROM: "+ROM.romName)

		ROM.romType = ROM.GetType()
		Logger.Log(LogTypes.INFO, "ROM TYPE: "+ROM.romType)

		ROM.romSize = ROM.GetROMSize()
		Logger.Log(LogTypes.INFO, "ROM SIZE: "+fmt.Sprintf("%dKB", ROM.romSize))

		ROM.romRAMSize = ROM.GetRAMSize()
		Logger.Log(LogTypes.INFO, "ROM RAM SIZE: "+fmt.Sprintf("%dKB", ROM.romRAMSize))

		menuDebug.SetSensitive(false)
		menuRun.SetSensitive(false)
		romTreeView.SetModel(nil)
		romListStore.Clear()
		romModelLength := len(ROM.model)
		percentStep := int(float64(0.01) * float64(romModelLength))
		for i := 0; i < romModelLength; i++ {
			row := ROM.model[i].([]string)
			iter := romListStore.Append()
			err := romListStore.Set(iter,
				[]int{0, 1},
				[]interface{}{row[0], row[1]})
			if err != nil {
				Logger.Log(LogTypes.ERROR, err)
			}
			if i%percentStep == 0 {
				romProgressBar.SetFraction(float64(i) / float64(romModelLength))
			} else if i == romModelLength-1 {
				romProgressBar.SetFraction(1)
			}
		}
		romTreeView.SetModel(romListStore)
		ROMref = ROM.data
	} else {
		Logger.Log(LogTypes.ERROR, "ROM: Error loading")
		after := time.Now()
		Logger.Log(LogTypes.INFO, after.Sub(before))
		return
	}
	after := time.Now()
	Logger.Logf(LogTypes.COMPLETED, "ROM: %s Loaded in %.2fs", ROM.romName, after.Sub(before).Seconds())
	menuDebug.SetSensitive(true)
	menuRun.SetSensitive(true)
	system.CPU.Reset()
	Logger.Logf(LogTypes.INFO, "CPU Initialized")
	return
}
