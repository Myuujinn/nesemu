package core

import "fmt"

func (c *Cpu) Cmp(value []byte) (incrPC bool, cycles byte, debug string) {
	c.SetFlag(FlagCarry, c.registers.acc >= value[0])
	c.SetFlag(FlagZero, c.registers.acc == value[0])
	c.SetFlag(FlagNegative, (c.registers.acc-value[0])&0x80 != 0)

	return true, 2, fmt.Sprintf("CMP #$%02X", value[0])
}

func (c *Cpu) IndexedIndirectCmp(address []byte) (incrPC bool, cycles byte, debug string) {
	target := uint16(address[0]) + uint16(c.registers.x)
	targetAddress := uint16(*c.Memory.Map(target % 0x100)) + uint16(*c.Memory.Map((target + 1) % 0x100))*0x100
	value := *c.Memory.Map(targetAddress)

	c.SetFlag(FlagCarry, c.registers.acc >= value)
	c.SetFlag(FlagZero, c.registers.acc == value)
	c.SetFlag(FlagNegative, (c.registers.acc-value)&0x80 != 0)

	return true, 6, fmt.Sprintf("CMP ($%02X,X) @ %02X = %04X = %02X", address[0], uint8(target), targetAddress, value)
}

func (c *Cpu) Cpy(value []byte) (incrPC bool, cycles byte, debug string) {
	c.SetFlag(FlagCarry, c.registers.y >= value[0])
	c.SetFlag(FlagZero, c.registers.y == value[0])
	c.SetFlag(FlagNegative, (c.registers.y-value[0])&0x80 != 0)

	return true, 2, fmt.Sprintf("CPY #$%02X", value[0])
}

func (c *Cpu) Cpx(value []byte) (incrPC bool, cycles byte, debug string) {
	c.SetFlag(FlagCarry, c.registers.x >= value[0])
	c.SetFlag(FlagZero, c.registers.x == value[0])
	c.SetFlag(FlagNegative, (c.registers.x-value[0])&0x80 != 0)

	return true, 2, fmt.Sprintf("CPX #$%02X", value[0])
}
