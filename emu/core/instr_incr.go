package core

func (c *Cpu) Dec(address []byte) (incrPC bool, cycles byte, debug string) {
	value := c.Memory.Map(uint16(address[0]))
	*value--

	c.SetFlag(FlagZero, *value == 0)
	c.SetFlag(FlagNegative, *value&0b10000000 == 0b10000000)

	return true, 2, "DEC"
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
