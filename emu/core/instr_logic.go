package core

import (
	"fmt"
)

func (c *Cpu) ImmediateAnd(value []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.acc &= value[0]

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 2, fmt.Sprintf("AND #$%02X", value[0])
}

func (c *Cpu) ZeropageXAnd(address []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.acc &= *c.Memory.Map((uint16(address[0]) + uint16(c.registers.x)) % 0x100)

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 4, fmt.Sprintf("AND $%02X,X", address[0])
}

func (c *Cpu) IndexedIndirectAnd(address []byte) (incrPC bool, cycles byte, debug string) {
	target := uint16(address[0]) + uint16(c.registers.x)
	targetAddress := uint16(*c.Memory.Map(target % 0x100)) + uint16(*c.Memory.Map((target + 1) % 0x100))*0x100
	value := c.Memory.Map(targetAddress)
	debug = fmt.Sprintf("AND ($%02X,X) @ %02X = %04X = %02X", address[0], uint8(target), targetAddress, *value)
	c.registers.acc &= *value

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 6, debug
}

func (c *Cpu) Ora(value []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.acc |= value[0]

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 2, fmt.Sprintf("ORA #$%02X", value[0])
}

func (c *Cpu) IndexedIndirectOra(address []byte) (incrPC bool, cycles byte, debug string) {
	target := uint16(address[0]) + uint16(c.registers.x)
	targetAddress := uint16(*c.Memory.Map(target % 0x100)) + uint16(*c.Memory.Map((target + 1) % 0x100))*0x100
	c.registers.acc |= *c.Memory.Map(targetAddress)

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 6, fmt.Sprintf("ORA ($%02X,X) @ %02X = %04X = %02X", address[0], uint8(target), targetAddress, *c.Memory.Map(targetAddress))
}

func (c *Cpu) Eor(value []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.acc ^= value[0]

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 2, fmt.Sprintf("EOR #$%02X", value[0])
}

func (c *Cpu) IndexedIndirectEor(address []byte) (incrPC bool, cycles byte, debug string) {
	target := uint16(address[0]) + uint16(c.registers.x)
	targetAddress := uint16(*c.Memory.Map(target % 0x100)) + uint16(*c.Memory.Map((target + 1) % 0x100))*0x100
	value := c.Memory.Map(targetAddress)
	debug = fmt.Sprintf("EOR ($%02X,X) @ %02X = %04X = %02X", address[0], uint8(target), targetAddress, *value)
	c.registers.acc ^= *value

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 6, debug
}
