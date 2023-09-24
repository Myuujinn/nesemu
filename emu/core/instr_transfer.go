package core

import "fmt"

func (c *Cpu) Ldx(value []byte) (incrPC bool, result byte, cycles byte, debug string) {
	c.registers.x = value[0]
	return true, c.registers.x, 2, fmt.Sprintf("LDX #$%02X", value[0])
}

func (c *Cpu) Lda(value []byte) (incrPC bool, result byte, cycles byte, debug string) {
	c.registers.acc = value[0]
	return true, c.registers.acc, 2, fmt.Sprintf("LDA #$%02X", value[0])
}

func (c *Cpu) Stx(address []byte) (incrPC bool, result byte, cycles byte, debug string) {
	*c.Memory.Map(uint16(address[0])) = c.registers.x
	return true, c.registers.x, 3, fmt.Sprintf("STX $%02X = %02X", address[0], c.registers.x)
}

func (c *Cpu) Sta(address []byte) (incrPC bool, result byte, cycles byte, debug string) {
	*c.Memory.Map(uint16(address[0])) = c.registers.acc
	return true, c.registers.acc, 3, fmt.Sprintf("STA $%02X = %02X", address[0], c.registers.acc)
}
