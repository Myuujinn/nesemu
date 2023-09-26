package core

import "fmt"

func (c *Cpu) Bit(address []byte) (incrPC bool, cycles byte, debug string) {
	result := *c.Memory.Map(uint16(address[0]))

	c.SetFlag(FlagZero, result&c.registers.acc == 0)
	c.SetFlag(FlagNegative, result&0x80 != 0)
	c.SetFlag(FlagOverflow, result&0x40 != 0)

	return true, 3, fmt.Sprintf("BIT $%02X = %02X", address[0], result)
}

func (c *Cpu) Nop(_ []byte) (incrPC bool, cycles byte, debug string) {
	return true, 2, "NOP"
}
