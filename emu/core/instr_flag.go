package core

func (c *Cpu) Sec(_ []byte) (incrPC bool, cycles byte, debug string) {
	c.SetFlag(FlagCarry, true)
	return true, 2, "SEC"
}

func (c *Cpu) Clc(_ []byte) (incrPC bool, cycles byte, debug string) {
	c.SetFlag(FlagCarry, false)
	return true, 2, "CLC"
}

func (c *Cpu) Sei(_ []byte) (incrPC bool, cycles byte, debug string) {
	c.SetFlag(FlagInterrupt, true)
	return true, 2, "SEI"
}

func (c *Cpu) Sed(_ []byte) (incrPC bool, cycles byte, debug string) {
	c.SetFlag(FlagDecimal, true)
	return true, 2, "SED"
}

func (c *Cpu) Cld(_ []byte) (incrPC bool, cycles byte, debug string) {
	c.SetFlag(FlagDecimal, false)
	return true, 2, "CLD"
}

func (c *Cpu) Clv(_ []byte) (incrPC bool, cycles byte, debug string) {
	c.SetFlag(FlagOverflow, false)
	return true, 2, "CLV"
}

func (c *Cpu) Cli(_ []byte) (incrPC bool, cycles byte, debug string) {
	c.SetFlag(FlagInterrupt, false)
	return true, 2, "CLI"
}
