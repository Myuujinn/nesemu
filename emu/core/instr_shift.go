package core

func (c *Cpu) AccLsr(_ []byte) (incrPC bool, cycles byte, debug string) {
	c.SetFlag(FlagCarry, c.registers.acc&0b00000001 == 0b00000001)
	c.registers.acc >>= 1

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 2, "LSR A"
}

func (c *Cpu) AccAsl(_ []byte) (incrPC bool, cycles byte, debug string) {
	c.SetFlag(FlagCarry, c.registers.acc&0b10000000 == 0b10000000)
	c.registers.acc <<= 1

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 2, "ASL A"
}

func (c *Cpu) AccRor(_ []byte) (incrPC bool, cycles byte, debug string) {
	oldCarry := c.IsFlagSet(FlagCarry)
	c.SetFlag(FlagCarry, c.registers.acc&0b00000001 == 0b00000001)
	c.registers.acc >>= 1
	if oldCarry {
		c.registers.acc |= 0b10000000
	}

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 2, "ROR A"
}

func (c *Cpu) AccRol(_ []byte) (incrPC bool, cycles byte, debug string) {
	oldCarry := c.IsFlagSet(FlagCarry)
	c.SetFlag(FlagCarry, c.registers.acc&0b10000000 == 0b10000000)
	c.registers.acc <<= 1
	if oldCarry {
		c.registers.acc |= 0b00000001
	}

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 2, "ROL A"
}
