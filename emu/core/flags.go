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

func (c *Cpu) SetFlag(flag uint8, set bool) {
	if set {
		c.registers.status |= flag
	} else {
		c.registers.status &= ^flag
	}
}
