package core

type AddressingMode uint8

const (
	Implied AddressingMode = iota
	Accumulator
	Immediate
	ZeroPage
	ZeroPageX
	ZeroPageY
	Absolute
	AbsoluteX
	AbsoluteY
	Indirect
	IndexedIndirect
	IndirectIndexed
	Relative
)

type Instruction struct {
	Opcode  uint8
	Name    string
	Bytes   uint8
	Execute func(*Cpu, []byte) (incrPC bool, result byte, cycles byte)
	Mode    AddressingMode
}

var instructionMap = map[uint8]Instruction{
	0x18: {0x18, "CLC", 1, (*Cpu).Clc, Implied},
	0x20: {0x20, "JSR", 3, (*Cpu).Jsr, Absolute},
	0x38: {0x38, "SEC", 1, (*Cpu).Sec, Implied},
	0x4C: {0x4C, "JMP", 3, (*Cpu).Jmp, Absolute},
	0x86: {0x86, "STX", 2, (*Cpu).Stx, ZeroPage},
	0xA2: {0xA2, "LDX", 2, Negative(Zero((*Cpu).Ldx)), Immediate},
	0xB0: {0xB0, "BCS", 2, (*Cpu).Bcs, Relative},
	0xEA: {0xEA, "NOP", 1, (*Cpu).Nop, Implied},
}

func (c *Cpu) Ldx(value []byte) (incrPC bool, result byte, cycles byte) {
	c.registers.x = value[0]
	return true, c.registers.x, 2
}

func (c *Cpu) Jmp(address []byte) (incrPC bool, result byte, cycles byte) {
	c.registers.pc = uint16(address[0]) | uint16(address[1])<<8
	return false, 0, 3
}

func (c *Cpu) Stx(address []byte) (incrPC bool, result byte, cycles byte) {
	*c.Memory.Map(uint16(address[0])) = c.registers.x
	return true, c.registers.x, 3
}

func (c *Cpu) Jsr(address []byte) (incrPC bool, result byte, cycles byte) {
	c.PushStack(uint8(c.registers.pc >> 8))
	c.PushStack(uint8(c.registers.pc))
	c.registers.pc = uint16(address[0]) | uint16(address[1])<<8
	return false, 0, 6
}

func (c *Cpu) Nop(_ []byte) (incrPC bool, result byte, cycles byte) {
	return true, 0, 2
}

func (c *Cpu) Sec(_ []byte) (incrPC bool, result byte, cycles byte) {
	c.SetFlag(FlagCarry)
	return true, 0, 2
}

func (c *Cpu) Bcs(address []byte) (incrPC bool, result byte, cycles byte) {
	cycles = 2

	if c.IsFlagSet(FlagCarry) {
		targetAddress := c.registers.pc + uint16(address[0])
		c.registers.pc = targetAddress

		// If the branch jumps to a new page, add an extra cycle
		if (c.registers.pc & 0xFF00) != (targetAddress) {
			cycles += 2
		} else {
			cycles++
		}
	}

	return true, 0, cycles
}

func (c *Cpu) Clc(_ []byte) (incrPC bool, result byte, cycles byte) {
	c.ClearFlag(FlagCarry)
	return true, 0, 2
}
