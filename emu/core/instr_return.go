package core

func (c *Cpu) Rts(_ []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.pc = uint16(c.PopStack()) + uint16(c.PopStack())<<8
	return true, 6, "RTS"
}
