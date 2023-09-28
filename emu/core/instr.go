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
	0x01: {0x01, "ORA", 2, (*Cpu).IndexedIndirectOra, IndexedIndirect},
	0x05: {0x05, "ORA", 2, (*Cpu).ZeropageOra, ZeroPage},
	0x06: {0x06, "ASL", 2, (*Cpu).ZeropageAsl, ZeroPage},
	0x08: {0x08, "PHP", 1, (*Cpu).Php, Implied},
	0x09: {0x09, "ORA", 2, (*Cpu).Ora, Immediate},
	0x0A: {0x0A, "ASL", 1, (*Cpu).AccAsl, Accumulator},
	0x0D: {0x0D, "ORA", 3, (*Cpu).AbsoluteOra, Absolute},
	0x0E: {0x0E, "ASL", 3, (*Cpu).AbsoluteAsl, Absolute},
	0x10: {0x10, "BPL", 2, (*Cpu).Bpl, Relative},
	0x18: {0x18, "CLC", 1, (*Cpu).Clc, Implied},
	0x20: {0x20, "JSR", 3, (*Cpu).Jsr, Absolute},
	0x21: {0x21, "AND", 2, (*Cpu).IndexedIndirectAnd, IndexedIndirect},
	0x24: {0x24, "BIT", 2, (*Cpu).Bit, ZeroPage},
	0x25: {0x25, "AND", 2, (*Cpu).ZeropageAnd, ZeroPageX},
	0x26: {0x26, "ROL", 2, (*Cpu).ZeropageRol, ZeroPage},
	0x28: {0x28, "PLP", 1, (*Cpu).Plp, Implied},
	0x29: {0x29, "AND", 2, (*Cpu).ImmediateAnd, Immediate},
	0x2A: {0x2A, "ROL", 1, (*Cpu).AccRol, Accumulator},
	0x2C: {0x2C, "BIT", 3, (*Cpu).AbsoluteBit, Absolute},
	0x2D: {0x2D, "AND", 3, (*Cpu).AbsoluteAnd, Absolute},
	0x2E: {0x2E, "ROL", 3, (*Cpu).AbsoluteRol, Absolute},
	0x30: {0x30, "BMI", 2, (*Cpu).Bmi, Relative},
	0x35: {0x35, "AND", 2, (*Cpu).ZeropageXAnd, ZeroPageX},
	0x38: {0x38, "SEC", 1, (*Cpu).Sec, Implied},
	0x40: {0x40, "RTI", 1, (*Cpu).Rti, Implied},
	0x45: {0x45, "EOR", 2, (*Cpu).ZeropageEor, ZeroPage},
	0x41: {0x41, "EOR", 2, (*Cpu).IndexedIndirectEor, IndexedIndirect},
	0x46: {0x46, "LSR", 2, (*Cpu).ZeropageLsr, ZeroPage},
	0x48: {0x48, "PHA", 1, (*Cpu).Pha, Implied},
	0x49: {0x49, "EOR", 2, (*Cpu).Eor, Immediate},
	0x4A: {0x4A, "LSR", 1, (*Cpu).AccLsr, Accumulator},
	0x4C: {0x4C, "JMP", 3, (*Cpu).Jmp, Absolute},
	0x4D: {0x4D, "EOR", 3, (*Cpu).AbsoluteEor, Absolute},
	0x4E: {0x4E, "LSR", 3, (*Cpu).AbsoluteLsr, Absolute},
	0x50: {0x50, "BVC", 2, (*Cpu).Bvc, Relative},
	0x58: {0x58, "CLI", 1, (*Cpu).Cli, Implied},
	0x60: {0x60, "RTS", 1, (*Cpu).Rts, Implied},
	0x61: {0x61, "ADC", 2, (*Cpu).IndexedIndirectAdc, IndexedIndirect},
	0x65: {0x65, "ADC", 2, (*Cpu).ZeropageAdc, ZeroPage},
	0x66: {0x66, "ROR", 2, (*Cpu).ZeropageRor, ZeroPage},
	0x68: {0x68, "PLA", 1, (*Cpu).Pla, Implied},
	0x69: {0x69, "ADC", 2, (*Cpu).Adc, Immediate},
	0x6A: {0x6A, "ROR", 1, (*Cpu).AccRor, Accumulator},
	0x6D: {0x6D, "ADC", 3, (*Cpu).AbsoluteAdc, Absolute},
	0x6E: {0x6E, "ROR", 3, (*Cpu).AbsoluteRor, Absolute},
	0x70: {0x70, "BVS", 2, (*Cpu).Bvs, Relative},
	0x78: {0x78, "SEI", 1, (*Cpu).Sei, Implied},
	0x81: {0x81, "STA", 2, (*Cpu).IndexedIndirectSta, IndexedIndirect},
	0x84: {0x84, "STY", 2, (*Cpu).Sty, ZeroPage},
	0x85: {0x85, "STA", 2, (*Cpu).ZeropageSta, ZeroPage},
	0x86: {0x86, "STX", 2, (*Cpu).ZeropageStx, ZeroPage},
	0x88: {0x88, "DEY", 1, (*Cpu).Dey, Implied},
	0x8A: {0x8A, "TXA", 1, (*Cpu).Txa, Implied},
	0x8C: {0x8C, "STY", 3, (*Cpu).AbsoluteSty, Absolute},
	0x8D: {0x8D, "STA", 3, (*Cpu).AbsoluteSta, Absolute},
	0x8E: {0x8E, "STX", 3, (*Cpu).AbsoluteStx, Absolute},
	0x90: {0x90, "BCC", 2, (*Cpu).Bcc, Relative},
	0x98: {0x98, "TYA", 1, (*Cpu).Tya, Implied},
	0x9A: {0x9A, "TXS", 1, (*Cpu).Txs, Implied},
	0xA0: {0xA0, "LDY", 2, (*Cpu).ImmediateLdy, Immediate},
	0xA1: {0xA1, "LDA", 2, (*Cpu).IndexedIndirectLda, IndexedIndirect},
	0xA2: {0xA2, "LDX", 2, (*Cpu).ImmediateLdx, Immediate},
	0xA4: {0xA4, "LDY", 2, (*Cpu).ZeropageLdy, ZeroPage},
	0xA5: {0xA5, "LDA", 2, (*Cpu).ZeropageLda, ZeroPage},
	0xA6: {0xA6, "LDX", 2, (*Cpu).ZeropageLdx, ZeroPage},
	0xA8: {0xA8, "TAY", 1, (*Cpu).Tay, Implied},
	0xA9: {0xA9, "LDA", 2, (*Cpu).ImmediateLda, Immediate},
	0xAA: {0xAA, "TAX", 1, (*Cpu).Tax, Implied},
	0xAC: {0xAC, "LDY", 3, (*Cpu).AbsoluteLdy, Absolute},
	0xAE: {0xAE, "LDX", 3, (*Cpu).AbsoluteLdx, Absolute},
	0xAD: {0xAD, "LDA", 3, (*Cpu).AbsoluteLda, Absolute},
	0xB0: {0xB0, "BCS", 2, (*Cpu).Bcs, Relative},
	0xB1: {0xB1, "LDA", 2, (*Cpu).IndirectIndexedLda, IndirectIndexed},
	0xB8: {0xB8, "CLV", 1, (*Cpu).Clv, Implied},
	0xBA: {0xBA, "TSX", 1, (*Cpu).Tsx, Implied},
	0xC0: {0xC0, "CPY", 2, (*Cpu).Cpy, Immediate},
	0xC1: {0xC1, "CMP", 2, (*Cpu).IndexedIndirectCmp, IndexedIndirect},
	0xC4: {0xC4, "CPY", 2, (*Cpu).ZeropageCpy, ZeroPage},
	0xC5: {0xC5, "CMP", 2, (*Cpu).ZeropageCmp, ZeroPage},
	0xC6: {0xC6, "DEC", 2, (*Cpu).ZeropageDec, ZeroPage},
	0xC8: {0xC8, "INY", 1, (*Cpu).Iny, Implied},
	0xC9: {0xC9, "CMP", 2, (*Cpu).Cmp, Immediate},
	0xCA: {0xCA, "DEX", 1, (*Cpu).Dex, Implied},
	0xCC: {0xCC, "CPY", 3, (*Cpu).AbsoluteCpy, Absolute},
	0xCD: {0xCD, "CMP", 3, (*Cpu).AbsoluteCmp, Absolute},
	0xCE: {0xCE, "DEC", 3, (*Cpu).AbsoluteDec, Absolute},
	0xE0: {0xE0, "CPX", 2, (*Cpu).Cpx, Immediate},
	0xE1: {0xE1, "SBC", 2, (*Cpu).IndexedIndirectSbc, IndexedIndirect},
	0xE4: {0xE4, "CPX", 2, (*Cpu).ZeropageCpx, ZeroPage},
	0xE5: {0xE5, "SBC", 2, (*Cpu).ZeropageSbc, ZeroPage},
	0xE6: {0xE6, "INC", 2, (*Cpu).ZeropageInc, ZeroPage},
	0xE8: {0xE8, "INX", 1, (*Cpu).Inx, Implied},
	0xE9: {0xE9, "SBC", 2, (*Cpu).Sbc, Immediate},
	0xEA: {0xEA, "NOP", 1, (*Cpu).Nop, Implied},
	0xEC: {0xEC, "CPX", 3, (*Cpu).AbsoluteCpx, Absolute},
	0xED: {0xED, "SBC", 3, (*Cpu).AbsoluteSbc, Absolute},
	0xEE: {0xEE, "INC", 3, (*Cpu).AbsoluteInc, Absolute},
	0xD0: {0xD0, "BNE", 2, (*Cpu).Bne, Relative},
	0xD8: {0xD8, "CLD", 1, (*Cpu).Cld, Implied},
	0xF0: {0xF0, "BEQ", 2, (*Cpu).Beq, Relative},
	0xF8: {0xF8, "SED", 1, (*Cpu).Sed, Implied},
}
