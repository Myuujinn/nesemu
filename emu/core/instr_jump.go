package core

import "fmt"

func (c *Cpu) Jmp(address []byte) (incrPC bool, result byte, cycles byte, debug string) {
	c.registers.pc = uint16(address[0]) | uint16(address[1])<<8
	return false, 0, 3, fmt.Sprintf("JMP $%02X%02X", address[1], address[0])
}

func (c *Cpu) Jsr(address []byte) (incrPC bool, result byte, cycles byte, debug string) {
	c.PushStack(uint8(c.registers.pc >> 8))
	c.PushStack(uint8(c.registers.pc + 2))
	c.registers.pc = uint16(address[0]) | uint16(address[1])<<8
	return false, 0, 6, fmt.Sprintf("JSR $%02X%02X", address[1], address[0])
}
