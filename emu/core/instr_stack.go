package core

func (c *Cpu) Php(_ []byte) (incrPC bool, result byte, cycles byte, debug string) {
	c.PushStack(c.registers.status | FlagBreak)
	return true, 0, 3, "PHP"
}

func (c *Cpu) Pla(_ []byte) (incrPC bool, result byte, cycles byte, debug string) {
	c.registers.acc = c.PopStack()
	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0x80 != 0)
	return true, 0, 4, "PLA"
}

func (c *Cpu) Pha(_ []byte) (incrPC bool, result byte, cycles byte, debug string) {
	c.PushStack(c.registers.acc)
	return true, 0, 3, "PHA"
}

func (c *Cpu) Plp(_ []byte) (incrPC bool, result byte, cycles byte, debug string) {
	// Set unused flag to 1 and break flag to old value
	oldStatus := c.registers.status
	c.registers.status = c.PopStack() | FlagUnused
	c.SetFlag(FlagBreak, oldStatus&FlagBreak == FlagBreak)

	return true, 0, 4, "PLP"
}
