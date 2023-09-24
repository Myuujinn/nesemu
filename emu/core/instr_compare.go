package core

import "fmt"

func (c *Cpu) Cmp(value []byte) (incrPC bool, result byte, cycles byte, debug string) {

	c.SetFlag(FlagCarry, c.registers.acc >= value[0])
	c.SetFlag(FlagZero, c.registers.acc == value[0])
	c.SetFlag(FlagNegative, (c.registers.acc-value[0])&0x80 != 0)

	return true, 0, 2, fmt.Sprintf("CMP #$%02X", value[0])
}
