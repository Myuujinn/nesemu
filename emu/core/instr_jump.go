package core

import "fmt"

func (c *Cpu) Jmp(address []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.pc = uint16(address[0]) | uint16(address[1])<<8
	return false, 3, fmt.Sprintf("JMP $%02X%02X", address[1], address[0])
}

func (c *Cpu) Jsr(address []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.pc += 2
	c.PushStack(uint8(c.registers.pc >> 8))
	c.PushStack(uint8(c.registers.pc))
	c.registers.pc = uint16(address[0]) | uint16(address[1])<<8
	return false, 6, fmt.Sprintf("JSR $%02X%02X", address[1], address[0])
}

func (c *Cpu) Rts(_ []byte) (incrPC bool, cycles byte, debug string) {
	c.registers.pc = uint16(c.PopStack()) + uint16(c.PopStack())<<8
	return true, 6, "RTS"
}
