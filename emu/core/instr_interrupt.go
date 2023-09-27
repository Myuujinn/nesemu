package core

func (c *Cpu) Brk(_ []byte) (incrPC bool, cycles byte, debug string) {
	c.PushStack(byte(c.registers.pc >> 8))
	c.PushStack(byte(c.registers.pc & 0xFF))
	c.PushStack(c.registers.status)
	c.SetFlag(FlagInterrupt, true)
	c.registers.pc = uint16(*c.Memory.Map(0xFFFE)) + uint16(*c.Memory.Map(0xFFFF))<<8
	return false, 7, "BRK"
}

func (c *Cpu) Rti(_ []byte) (incrPC bool, cycles byte, debug string) {
	setUnusedFlag := c.IsFlagSet(FlagUnused)
	c.registers.status = c.PopStack()
	c.SetFlag(FlagUnused, setUnusedFlag)
	c.registers.pc = uint16(c.PopStack()) + uint16(c.PopStack())<<8
	return false, 6, "RTI"
}
