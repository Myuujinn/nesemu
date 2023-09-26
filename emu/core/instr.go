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
	Execute func(*Cpu, []byte) (incrPC bool, cycles byte, debug string)
	Mode    AddressingMode
}

var instructionMap = map[uint8]Instruction{
	0x08: {0x08, "PHP", 1, (*Cpu).Php, Implied},
	0x09: {0x09, "ORA", 2, (*Cpu).Ora, Immediate},
	0x10: {0x10, "BPL", 2, (*Cpu).Bpl, Relative},
	0x18: {0x18, "CLC", 1, (*Cpu).Clc, Implied},
	0x20: {0x20, "JSR", 3, (*Cpu).Jsr, Absolute},
	0x24: {0x24, "BIT", 2, (*Cpu).Bit, ZeroPage},
	0x28: {0x28, "PLP", 1, (*Cpu).Plp, Implied},
	0x29: {0x29, "AND", 2, (*Cpu).And, Immediate},
	0x30: {0x30, "BMI", 2, (*Cpu).Bmi, Relative},
	0x38: {0x38, "SEC", 1, (*Cpu).Sec, Implied},
	0x48: {0x48, "PHA", 1, (*Cpu).Pha, Implied},
	0x49: {0x49, "EOR", 2, (*Cpu).Eor, Immediate},
	0x4C: {0x4C, "JMP", 3, (*Cpu).Jmp, Absolute},
	0x50: {0x50, "BVC", 2, (*Cpu).Bvc, Relative},
	0x60: {0x60, "RTS", 1, (*Cpu).Rts, Implied},
	0x68: {0x68, "PLA", 1, (*Cpu).Pla, Implied},
	0x69: {0x69, "ADC", 2, (*Cpu).Adc, Immediate},
	0x70: {0x70, "BVS", 2, (*Cpu).Bvs, Relative},
	0x78: {0x78, "SEI", 1, (*Cpu).Sei, Implied},
	0x84: {0x84, "STY", 2, (*Cpu).Sty, ZeroPage},
	0x85: {0x85, "STA", 2, (*Cpu).Sta, ZeroPage},
	0x86: {0x86, "STX", 2, (*Cpu).Stx, ZeroPage},
	0x88: {0x88, "DEY", 1, (*Cpu).Dey, Implied},
	0x8A: {0x8A, "TXA", 1, (*Cpu).Txa, Implied},
	0x90: {0x90, "BCC", 2, (*Cpu).Bcc, Relative},
	0x98: {0x98, "TYA", 1, (*Cpu).Tya, Implied},
	0xA0: {0xA0, "LDY", 2, (*Cpu).Ldy, Immediate},
	0xA2: {0xA2, "LDX", 2, (*Cpu).Ldx, Immediate},
	0xA8: {0xA8, "TAY", 1, (*Cpu).Tay, Implied},
	0xA9: {0xA9, "LDA", 2, (*Cpu).Lda, Immediate},
	0xAA: {0xAA, "TAX", 1, (*Cpu).Tax, Implied},
	0xB0: {0xB0, "BCS", 2, (*Cpu).Bcs, Relative},
	0xB8: {0xB8, "CLV", 1, (*Cpu).Clv, Implied},
	0xBA: {0xBA, "TSX", 1, (*Cpu).Tsx, Implied},
	0xC0: {0xC0, "CPY", 2, (*Cpu).Cpy, Immediate},
	0xC8: {0xC8, "INY", 1, (*Cpu).Iny, Implied},
	0xC9: {0xC9, "CMP", 2, (*Cpu).Cmp, Immediate},
	0xCA: {0xCA, "DEX", 1, (*Cpu).Dex, Implied},
	0xE0: {0xE0, "CPX", 2, (*Cpu).Cpx, Immediate},
	0xE8: {0xE8, "INX", 1, (*Cpu).Inx, Implied},
	0xE9: {0xE9, "SBC", 2, (*Cpu).Sbc, Immediate},
	0xEA: {0xEA, "NOP", 1, (*Cpu).Nop, Implied},
	0xD0: {0xD0, "BNE", 2, (*Cpu).Bne, Relative},
	0xD8: {0xD8, "CLD", 1, (*Cpu).Cld, Implied},
	0xF0: {0xF0, "BEQ", 2, (*Cpu).Beq, Relative},
	0xF8: {0xF8, "SED", 1, (*Cpu).Sed, Implied},
}
