package core

import "fmt"

func (c *Cpu) Cmp(value []byte) (incrPC bool, cycles byte, debug string) {
	c.SetFlag(FlagCarry, c.registers.acc >= value[0])
	c.SetFlag(FlagZero, c.registers.acc == value[0])
	c.SetFlag(FlagNegative, (c.registers.acc-value[0])&0x80 != 0)

	return true, 2, fmt.Sprintf("CMP #$%02X", value[0])
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
