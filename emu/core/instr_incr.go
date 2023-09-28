package core

import "fmt"

func (c *Cpu) Dec(address []byte) (incrPC bool, cycles byte, debug string) {
	value := c.Memory.Map(uint16(address[0]))
	*value--

	c.SetFlag(FlagZero, *value == 0)
	c.SetFlag(FlagNegative, *value&0b10000000 == 0b10000000)

	return true, 2, "DEC"
}

func (c *Cpu) ZeropageDec(address []byte) (incrPC bool, cycles byte, debug string) {
	value := c.Memory.Map(uint16(address[0]))
	debug = fmt.Sprintf("DEC $%02X = %02X", address[0], *value)
	*value--

	c.SetFlag(FlagZero, *value == 0)
	c.SetFlag(FlagNegative, *value&0b10000000 == 0b10000000)

	return true, 5, debug
}

func (c *Cpu) AbsoluteDec(address []byte) (incrPC bool, cycles byte, debug string) {
	value := c.Memory.Map(uint16(address[0]) | uint16(address[1])<<8)
	debug = fmt.Sprintf("DEC $%02X%02X = %02X", address[1], address[0], *value)
	*value--

	c.SetFlag(FlagZero, *value == 0)
	c.SetFlag(FlagNegative, *value&0b10000000 == 0b10000000)

	return true, 6, debug
}

func (c *Cpu) Dex(_ []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.x--

	c.SetFlag(FlagZero, c.registers.x == 0)
	c.SetFlag(FlagNegative, c.registers.x&0b10000000 == 0b10000000)

	return true, 2, "DEX"
}

func (c *Cpu) Dey(_ []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.y--

	c.SetFlag(FlagZero, c.registers.y == 0)
	c.SetFlag(FlagNegative, c.registers.y&0b10000000 == 0b10000000)

	return true, 2, "DEY"
}

func (c *Cpu) Inc(address []byte) (incrPC bool, cycles byte, debug string) {
	value := c.Memory.Map(uint16(address[0]))
	*value++

	c.SetFlag(FlagZero, *value == 0)
	c.SetFlag(FlagNegative, *value&0b10000000 == 0b10000000)

	return true, 2, "INC"
}

func (c *Cpu) ZeropageInc(address []byte) (incrPC bool, cycles byte, debug string) {
	value := c.Memory.Map(uint16(address[0]))
	debug = fmt.Sprintf("INC $%02X = %02X", address[0], *value)
	*value++

	c.SetFlag(FlagZero, *value == 0)
	c.SetFlag(FlagNegative, *value&0b10000000 == 0b10000000)

	return true, 5, debug
}

func (c *Cpu) AbsoluteInc(address []byte) (incrPC bool, cycles byte, debug string) {
	value := c.Memory.Map(uint16(address[0]) | uint16(address[1])<<8)
	debug = fmt.Sprintf("INC $%02X%02X = %02X", address[1], address[0], *value)
	*value++

	c.SetFlag(FlagZero, *value == 0)
	c.SetFlag(FlagNegative, *value&0b10000000 == 0b10000000)

	return true, 6, debug
}

func (c *Cpu) Inx(_ []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.x++

	c.SetFlag(FlagZero, c.registers.x == 0)
	c.SetFlag(FlagNegative, c.registers.x&0b10000000 == 0b10000000)

	return true, 2, "INX"
}

func (c *Cpu) Iny(_ []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.y++

	c.SetFlag(FlagZero, c.registers.y == 0)
	c.SetFlag(FlagNegative, c.registers.y&0b10000000 == 0b10000000)

	return true, 2, "INY"
}
