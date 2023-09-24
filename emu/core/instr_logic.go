package core

import (
	"fmt"
)

func (c *Cpu) And(value []byte) (incrPC bool, result byte, cycles byte, debug string) {
	c.registers.acc &= value[0]

	return true, c.registers.acc, 2, fmt.Sprintf("AND #$%02X", value[0])
}

func (c *Cpu) Ora(value []byte) (incrPC bool, result byte, cycles byte, debug string) {
	c.registers.acc |= value[0]

	return true, c.registers.acc, 2, fmt.Sprintf("ORA #$%02X", value[0])
}

func (c *Cpu) Eor(value []byte) (incrPC bool, result byte, cycles byte, debug string) {
	c.registers.acc ^= value[0]

	return true, c.registers.acc, 2, fmt.Sprintf("EOR #$%02X", value[0])
}
