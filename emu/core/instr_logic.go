package core

import (
	"fmt"
)

func (c *Cpu) And(value []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.acc &= value[0]

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 2, fmt.Sprintf("AND #$%02X", value[0])
}

func (c *Cpu) Ora(value []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.acc |= value[0]

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 2, fmt.Sprintf("ORA #$%02X", value[0])
}

func (c *Cpu) Eor(value []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.acc ^= value[0]

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 2, fmt.Sprintf("EOR #$%02X", value[0])
}
