package core

const (
	FlagCarry     uint8 = 0b00000001
	FlagZero      uint8 = 0b00000010
	FlagInterrupt uint8 = 0b00000100
	FlagDecimal   uint8 = 0b00001000
	FlagBreak     uint8 = 0b00010000
	FlagUnused    uint8 = 0b00100000
	FlagOverflow  uint8 = 0b01000000
	FlagNegative  uint8 = 0b10000000
)

func (c *Cpu) IsFlagSet(flag uint8) bool {
	return c.registers.status&flag == flag
}

func (c *Cpu) SetFlag(flag uint8) {
	c.registers.status |= flag
}

func (c *Cpu) ClearFlag(flag uint8) {
	c.registers.status &= ^flag
}

func Zero(instr func(*Cpu, []byte) (incrPC bool, result byte, cycles byte)) func(*Cpu, []byte) (incrPC bool, result byte, cycles byte) {
	return func(cpu *Cpu, operands []byte) (incrPC bool, result byte, cycles byte) {
		incrPC, result, cycles = instr(cpu, operands)

		if result == 0 {
			cpu.SetFlag(FlagZero)
		} else {
			cpu.ClearFlag(FlagZero)
		}

		return incrPC, result, cycles
	}
}

func Negative(instr func(*Cpu, []byte) (incrPC bool, result byte, cycles byte)) func(*Cpu, []byte) (incrPC bool, result byte, cycles byte) {
	return func(cpu *Cpu, operands []byte) (incrPC bool, result byte, cycles byte) {
		incrPC, result, cycles = instr(cpu, operands)

		if result&0b10000000 == 0b10000000 {
			cpu.SetFlag(FlagNegative)
		} else {
			cpu.ClearFlag(FlagNegative)
		}

		return incrPC, result, cycles
	}
}
