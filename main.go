package main

import (
	"flag"
	"github.com/myuujinn/nesemu/emu/core"
	"strconv"
)

func main() {
	romPath := flag.String("rom", "", "ROM file to load")
	entrypoint := flag.String("entrypoint", "", "Entrypoint to start execution at")
	// scale := flag.Int("scale", 2, "Scale factor for window")
	flag.Parse()

	// emu.Run(rom, *scale)

	cpu := core.Cpu{}
	rom, err := core.NewRom(romPath)
	if err != nil {
		panic(err)
	}

	cpu.Memory.LoadROM(rom)
	cpu.Init()

	entrypointAddress, err := strconv.ParseUint(*entrypoint, 16, 16)
	if err != nil {
		panic(err)
	}

	cpu.SetEntrypoint(uint16(entrypointAddress))

	cpu.Start()
}
