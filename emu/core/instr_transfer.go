package core

import "fmt"

func (c *Cpu) Ldx(value []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.x = value[0]

	c.SetFlag(FlagZero, c.registers.x == 0)
	c.SetFlag(FlagNegative, c.registers.x&0b10000000 == 0b10000000)

	return true, 2, fmt.Sprintf("LDX #$%02X", value[0])
}

func (c *Cpu) Ldy(value []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.y = value[0]

	c.SetFlag(FlagZero, c.registers.y == 0)
	c.SetFlag(FlagNegative, c.registers.y&0b10000000 == 0b10000000)

	return true, 2, fmt.Sprintf("LDY #$%02X", value[0])
}

func (c *Cpu) Lda(value []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.acc = value[0]

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 2, fmt.Sprintf("LDA #$%02X", value[0])
}

func (c *Cpu) Stx(address []byte) (incrPC bool, cycles byte, debug string) {
	*c.Memory.Map(uint16(address[0])) = c.registers.x
	return true, 3, fmt.Sprintf("STX $%02X = %02X", address[0], c.registers.x)
}

func (c *Cpu) Sty(address []byte) (incrPC bool, cycles byte, debug string) {
	*c.Memory.Map(uint16(address[0])) = c.registers.y
	return true, 3, fmt.Sprintf("STY $%02X = %02X", address[0], c.registers.y)
}

func (c *Cpu) Sta(address []byte) (incrPC bool, cycles byte, debug string) {
	*c.Memory.Map(uint16(address[0])) = c.registers.acc
	return true, 3, fmt.Sprintf("STA $%02X = %02X", address[0], c.registers.acc)
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
