package core

/*
+---------------+-------+------------------------------------------------+
| Address range | Size  | Device                                         |
+---------------+-------+------------------------------------------------+
| $0000–$07FF   | $0800 | 2 KB internal RAM                              |
| $0800–$0FFF   | $0800 | Mirrors of $0000–$07FF                         |
| $1000–$17FF   | $0800 | Mirrors of $0000–$07FF                         |
| $1800–$1FFF   | $0800 | Mirrors of $0000–$07FF                         |
| $2000–$2007   | $0008 | NES PPU registers                              |
| $2008–$3FFF   | $1FF8 | Mirrors of $2000–$2007 (repeats every 8 bytes) |
| $4000–$401F   | $0020 | NES APU and I/O registers                      |
| $4020–$5FFF   | $1980 | Expansion ROM                                  |
| $6000-$7FFF   | $2000 | SRAM                                           |
| $8000-$9FFF   | $2000 | Program ROM                                    |
+---------------+-------+------------------------------------------------+
*/

const (
	ramStart    uint16 = 0x0000
	ramEnd      uint16 = 0x0800
	ramSize     uint16 = ramEnd - ramStart
	StackStart  uint16 = 0x0100
	StackEnd    uint16 = 0x0200
	ppuStart    uint16 = 0x2000
	ppuEnd      uint16 = 0x2008
	ppuSize     uint16 = ppuEnd - ppuStart
	ioStart     uint16 = 0x4000
	ioEnd       uint16 = 0x4020
	ioSize      uint16 = ioEnd - ioStart
	eromStart   uint16 = 0x4020
	eromEnd     uint16 = 0x6000
	eromSize    uint16 = eromEnd - eromStart
	sramStart   uint16 = 0x6000
	sramEnd     uint16 = 0x8000
	sramSize    uint16 = sramEnd - sramStart
	prgromStart uint16 = 0x8000
	prgromEnd   uint16 = 0xFFFF
	prgromSize  uint16 = prgromEnd - prgromStart + 1
)

type Memory struct {
	ram    [ramSize]byte
	ppu    [ppuSize]byte
	io     [ioSize]byte
	erom   [eromSize]byte
	sram   [sramSize]byte
	prgrom [prgromSize]byte
}

func (m *Memory) Map(addr uint16) *byte {
	switch {
	case addr < ramEnd:
		return &m.ram[addr]
	case addr < ppuStart:
		return &m.ram[addr%ramSize]
	case addr < ioStart:
		return &m.ppu[(addr-ppuStart)%ppuSize]
	case addr < eromStart:
		return &m.io[addr-ioStart]
	case addr < sramStart:
		return &m.erom[addr-eromStart]
	case addr < prgromStart:
		return &m.sram[addr-sramStart]
	default:
		return &m.prgrom[addr-prgromStart]
	}
}

func (m *Memory) LoadROM(rom *Rom) {
	copy(m.prgrom[:], rom.Data)

	if rom.Header.PrgromBanks == 1 {
		copy(m.prgrom[0x4000:], rom.Data)
	}
}
