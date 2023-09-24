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
	Execute func(*Cpu, []byte) (incrPC bool, result byte, cycles byte, debug string)
	Mode    AddressingMode
}

var instructionMap = map[uint8]Instruction{
	0x08: {0x08, "PHP", 1, (*Cpu).Php, Implied},
	0x09: {0x09, "ORA", 2, Negative(Zero((*Cpu).Ora)), Immediate},
	0x10: {0x10, "BPL", 2, (*Cpu).Bpl, Relative},
	0x18: {0x18, "CLC", 1, (*Cpu).Clc, Implied},
	0x20: {0x20, "JSR", 3, (*Cpu).Jsr, Absolute},
	0x24: {0x24, "BIT", 2, (*Cpu).Bit, ZeroPage},
	0x28: {0x28, "PLP", 1, (*Cpu).Plp, Implied},
	0x29: {0x29, "AND", 2, Negative(Zero((*Cpu).And)), Immediate},
	0x30: {0x30, "BMI", 2, (*Cpu).Bmi, Relative},
	0x38: {0x38, "SEC", 1, (*Cpu).Sec, Implied},
	0x48: {0x48, "PHA", 1, (*Cpu).Pha, Implied},
	0x4C: {0x4C, "JMP", 3, (*Cpu).Jmp, Absolute},
	0x50: {0x50, "BVC", 2, (*Cpu).Bvc, Relative},
	0x60: {0x60, "RTS", 1, (*Cpu).Rts, Implied},
	0x68: {0x68, "PLA", 1, (*Cpu).Pla, Implied},
	0x70: {0x70, "BVS", 2, (*Cpu).Bvs, Relative},
	0x78: {0x78, "SEI", 1, (*Cpu).Sei, Implied},
	0x85: {0x85, "STA", 2, (*Cpu).Sta, ZeroPage},
	0x86: {0x86, "STX", 2, (*Cpu).Stx, ZeroPage},
	0x90: {0x90, "BCC", 2, (*Cpu).Bcc, Relative},
	0xA2: {0xA2, "LDX", 2, Negative(Zero((*Cpu).Ldx)), Immediate},
	0xA9: {0xA9, "LDA", 2, Negative(Zero((*Cpu).Lda)), Immediate},
	0xB0: {0xB0, "BCS", 2, (*Cpu).Bcs, Relative},
	0xC9: {0xC9, "CMP", 2, (*Cpu).Cmp, Immediate},
	0xEA: {0xEA, "NOP", 1, (*Cpu).Nop, Implied},
	0xD0: {0xD0, "BNE", 2, (*Cpu).Bne, Relative},
	0xD8: {0xD8, "CLD", 1, (*Cpu).Cld, Implied},
	0xF0: {0xF0, "BEQ", 2, (*Cpu).Beq, Relative},
	0xF8: {0xF8, "SED", 1, (*Cpu).Sed, Implied},
}
