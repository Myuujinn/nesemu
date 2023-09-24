package core

import (
	"encoding/binary"
	"fmt"
	"os"
)

type Header struct {
	Nes         [3]byte // NES
	Magic       byte    // 0x1A
	PrgromBanks byte
	ChrromBanks byte
	Control1    byte
	Control2    byte
	PrgramBanks byte
	Zero        [7]byte // should be zero
}

type Rom struct {
	Header Header
	Mapper byte
	Data   []byte
}

func (h *Header) String() string {
	return fmt.Sprintf("%s, Magic: %x, PRGROM Banks: %d, CHRROM Banks: %d, Control1: %d, Control2: %d, PRGRAM Banks: %d, Zero: %v",
		h.Nes, h.Magic, h.PrgromBanks, h.ChrromBanks, h.Control1, h.Control2, h.PrgramBanks, h.Zero)
}

func (r *Rom) String() string {
	return fmt.Sprintf("Header: %s, Mapper: %s", &r.Header, MapperName(r.Mapper))
}

func NewRom(romPath *string) (*Rom, error) {
	file, err := os.Open(*romPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	rom := &Rom{}
	err = binary.Read(file, binary.LittleEndian, &rom.Header)
	if err != nil {
		return nil, err
	}

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	rom.Data = make([]byte, stat.Size()-16) // header is 16 bytes
	err = binary.Read(file, binary.LittleEndian, &rom.Data)
	if err != nil {
		return nil, err
	}

	rom.Mapper = (rom.Header.Control1 & 0xF0) | (rom.Header.Control2 >> 4)

	fmt.Println(rom)
	return rom, nil
}
