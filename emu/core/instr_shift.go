package core

import "fmt"

func (c *Cpu) AccLsr(_ []byte) (incrPC bool, cycles byte, debug string) {
	c.SetFlag(FlagCarry, c.registers.acc&0b00000001 == 0b00000001)
	c.registers.acc >>= 1

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 2, "LSR A"
}

func (c *Cpu) ZeropageLsr(address []byte) (incrPC bool, cycles byte, debug string) {
	value := c.Memory.Map(uint16(address[0]))
	c.SetFlag(FlagCarry, *value&0b00000001 == 0b00000001)
	debug = fmt.Sprintf("LSR $%02X = %02X", address[0], *value)
	*value >>= 1

	c.SetFlag(FlagZero, *value == 0)
	c.SetFlag(FlagNegative, *value&0b10000000 == 0b10000000)

	return true, 5, debug
}

func (c *Cpu) AbsoluteLsr(address []byte) (incrPC bool, cycles byte, debug string) {
	value := c.Memory.Map(uint16(address[0]) | uint16(address[1])<<8)
	c.SetFlag(FlagCarry, *value&0b00000001 == 0b00000001)
	debug = fmt.Sprintf("LSR $%02X%02X = %02X", address[1], address[0], *value)
	*value >>= 1

	c.SetFlag(FlagZero, *value == 0)
	c.SetFlag(FlagNegative, *value&0b10000000 == 0b10000000)

	return true, 6, debug
}

func (c *Cpu) AccAsl(_ []byte) (incrPC bool, cycles byte, debug string) {
	c.SetFlag(FlagCarry, c.registers.acc&0b10000000 == 0b10000000)
	c.registers.acc <<= 1

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 2, "ASL A"
}

func (c *Cpu) ZeropageAsl(address []byte) (incrPC bool, cycles byte, debug string) {
	value := c.Memory.Map(uint16(address[0]))
	c.SetFlag(FlagCarry, *value&0b10000000 == 0b10000000)
	debug = fmt.Sprintf("ASL $%02X = %02X", address[0], *value)
	*value <<= 1

	c.SetFlag(FlagZero, *value == 0)
	c.SetFlag(FlagNegative, *value&0b10000000 == 0b10000000)

	return true, 5, debug
}

func (c *Cpu) AbsoluteAsl(address []byte) (incrPC bool, cycles byte, debug string) {
	value := c.Memory.Map(uint16(address[0]) | uint16(address[1])<<8)
	c.SetFlag(FlagCarry, *value&0b10000000 == 0b10000000)
	debug = fmt.Sprintf("ASL $%02X%02X = %02X", address[1], address[0], *value)
	*value <<= 1

	c.SetFlag(FlagZero, *value == 0)
	c.SetFlag(FlagNegative, *value&0b10000000 == 0b10000000)

	return true, 6, debug
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

func (c *Cpu) ZeropageRor(address []byte) (incrPC bool, cycles byte, debug string) {
	value := c.Memory.Map(uint16(address[0]))
	oldCarry := c.IsFlagSet(FlagCarry)
	c.SetFlag(FlagCarry, *value&0b00000001 == 0b00000001)
	debug = fmt.Sprintf("ROR $%02X = %02X", address[0], *value)
	*value >>= 1
	if oldCarry {
		*value |= 0b10000000
	}

	c.SetFlag(FlagZero, *value == 0)
	c.SetFlag(FlagNegative, *value&0b10000000 == 0b10000000)

	return true, 5, debug
}

func (c *Cpu) AbsoluteRor(address []byte) (incrPC bool, cycles byte, debug string) {
	value := c.Memory.Map(uint16(address[0]) | uint16(address[1])<<8)
	oldCarry := c.IsFlagSet(FlagCarry)
	c.SetFlag(FlagCarry, *value&0b00000001 == 0b00000001)
	debug = fmt.Sprintf("ROR $%02X%02X = %02X", address[1], address[0], *value)
	*value >>= 1
	if oldCarry {
		*value |= 0b10000000
	}

	c.SetFlag(FlagZero, *value == 0)
	c.SetFlag(FlagNegative, *value&0b10000000 == 0b10000000)

	return true, 6, debug
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

func (c *Cpu) ZeropageRol(address []byte) (incrPC bool, cycles byte, debug string) {
	value := c.Memory.Map(uint16(address[0]))
	oldCarry := c.IsFlagSet(FlagCarry)
	c.SetFlag(FlagCarry, *value&0b10000000 == 0b10000000)
	debug = fmt.Sprintf("ROL $%02X = %02X", address[0], *value)
	*value <<= 1
	if oldCarry {
		*value |= 0b00000001
	}

	c.SetFlag(FlagZero, *value == 0)
	c.SetFlag(FlagNegative, *value&0b10000000 == 0b10000000)

	return true, 5, debug
}

func (c *Cpu) AbsoluteRol(address []byte) (incrPC bool, cycles byte, debug string) {
	value := c.Memory.Map(uint16(address[0]) | uint16(address[1])<<8)
	oldCarry := c.IsFlagSet(FlagCarry)
	c.SetFlag(FlagCarry, *value&0b10000000 == 0b10000000)
	debug = fmt.Sprintf("ROL $%02X%02X = %02X", address[1], address[0], *value)
	*value <<= 1
	if oldCarry {
		*value |= 0b00000001
	}

	c.SetFlag(FlagZero, *value == 0)
	c.SetFlag(FlagNegative, *value&0b10000000 == 0b10000000)

	return true, 6, debug
}
