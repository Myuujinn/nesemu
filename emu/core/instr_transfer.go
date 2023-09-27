package core

import "fmt"

func (c *Cpu) ImmediateLdx(value []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.x = value[0]

	c.SetFlag(FlagZero, c.registers.x == 0)
	c.SetFlag(FlagNegative, c.registers.x&0b10000000 == 0b10000000)

	return true, 2, fmt.Sprintf("LDX #$%02X", value[0])
}

func (c *Cpu) AbsoluteLdx(address []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.x = *c.Memory.Map(uint16(address[0]) | uint16(address[1])<<8)

	c.SetFlag(FlagZero, c.registers.x == 0)
	c.SetFlag(FlagNegative, c.registers.x&0b10000000 == 0b10000000)

	return true, 4, fmt.Sprintf("LDX $%02X%02X = %02X", address[1], address[0], c.registers.x)
}

func (c *Cpu) Ldy(value []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.y = value[0]

	c.SetFlag(FlagZero, c.registers.y == 0)
	c.SetFlag(FlagNegative, c.registers.y&0b10000000 == 0b10000000)

	return true, 2, fmt.Sprintf("LDY #$%02X", value[0])
}

func (c *Cpu) ImmediateLda(value []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.acc = value[0]

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 2, fmt.Sprintf("LDA #$%02X", value[0])
}

func (c *Cpu) AbsoluteLda(address []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.acc = *c.Memory.Map(uint16(address[0]) | uint16(address[1])<<8)

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 4, fmt.Sprintf("LDA $%02X%02X = %02X", address[1], address[0], c.registers.acc)
}

func (c *Cpu) ZeropageLda(address []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.acc = *c.Memory.Map(uint16(address[0]))

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 3, fmt.Sprintf("LDA $%02X = %02X", address[0], c.registers.acc)
}

func (c *Cpu) IndexedIndirectLda(address []byte) (incrPC bool, cycles byte, debug string) {
	target := uint16(address[0]) + uint16(c.registers.x)
	targetAddress := uint16(*c.Memory.Map(target % 0x100)) + uint16(*c.Memory.Map((target + 1) % 0x100))*0x100
	c.registers.acc = *c.Memory.Map(targetAddress)

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 6, fmt.Sprintf("LDA ($%02X,X) @ %02X = %04X = %02X", address[0], uint8(target), targetAddress, c.registers.acc)
}

func (c *Cpu) ZeropageStx(address []byte) (incrPC bool, cycles byte, debug string) {
	*c.Memory.Map(uint16(address[0])) = c.registers.x
	return true, 3, fmt.Sprintf("STX $%02X = %02X", address[0], c.registers.x)
}

func (c *Cpu) AbsoluteStx(address []byte) (incrPC bool, cycles byte, debug string) {
	value := c.Memory.Map(uint16(address[0]) | uint16(address[1])<<8)
	debug = fmt.Sprintf("STX $%02X%02X = %02X", address[1], address[0], *value)
	*value = c.registers.x
	return true, 4, debug
}

func (c *Cpu) Sty(address []byte) (incrPC bool, cycles byte, debug string) {
	*c.Memory.Map(uint16(address[0])) = c.registers.y
	return true, 3, fmt.Sprintf("STY $%02X = %02X", address[0], c.registers.y)
}

func (c *Cpu) ZeropageSta(address []byte) (incrPC bool, cycles byte, debug string) {
	value := c.Memory.Map(uint16(address[0]))
	debug = fmt.Sprintf("STA $%02X = %02X", address[0], *value)
	*value = c.registers.acc
	return true, 3, debug
}

func (c *Cpu) AbsoluteSta(address []byte) (incrPC bool, cycles byte, debug string) {
	value := c.Memory.Map(uint16(address[0]) | uint16(address[1])<<8)
	debug = fmt.Sprintf("STA $%02X%02X = %02X", address[1], address[0], *value)
	*value = c.registers.acc
	return true, 4, debug
}

func (c *Cpu) IndexedIndirectSta(address []byte) (incrPC bool, cycles byte, debug string) {
	target := uint16(address[0]) + uint16(c.registers.x)
	targetAddress := uint16(*c.Memory.Map(target % 0x100)) + uint16(*c.Memory.Map((target + 1) % 0x100))*0x100
	value := c.Memory.Map(targetAddress)
	debug = fmt.Sprintf("STA ($%02X,X) @ %02X = %04X = %02X", address[0], uint8(target), targetAddress, *value)
	*value = c.registers.acc
	return true, 6, debug
}

func (c *Cpu) Tax(_ []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.x = c.registers.acc

	c.SetFlag(FlagZero, c.registers.x == 0)
	c.SetFlag(FlagNegative, c.registers.x&0b10000000 == 0b10000000)

	return true, 2, "TAX"
}

func (c *Cpu) Tay(_ []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.y = c.registers.acc

	c.SetFlag(FlagZero, c.registers.y == 0)
	c.SetFlag(FlagNegative, c.registers.y&0b10000000 == 0b10000000)

	return true, 2, "TAY"
}

func (c *Cpu) Tya(_ []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.acc = c.registers.y

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 2, "TYA"
}

func (c *Cpu) Txa(_ []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.acc = c.registers.x

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 2, "TXA"
}

func (c *Cpu) Tsx(_ []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.x = c.registers.sp

	c.SetFlag(FlagZero, c.registers.x == 0)
	c.SetFlag(FlagNegative, c.registers.x&0b10000000 == 0b10000000)

	return true, 2, "TSX"
}

func (c *Cpu) Txs(_ []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.sp = c.registers.x

	return true, 2, "TXS"
}
