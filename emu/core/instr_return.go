package core

func (c *Cpu) Rts(_ []byte) (incrPC bool, result byte, cycles byte, debug string) {
	c.registers.pc = uint16(c.PopStack()) + uint16(c.PopStack())<<8
	return true, 0, 6, "RTS"
}
