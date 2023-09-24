package core

import "fmt"

func (c *Cpu) Bcs(address []byte) (incrPC bool, result byte, cycles byte, debug string) {
	cycles = 2
	incrPC = true
	targetAddress := c.registers.pc + uint16(address[0]) + 2

	if c.IsFlagSet(FlagCarry) {
		// If the branch jumps to a new page, add an extra cycle
		if c.registers.pc&0xFF00 != targetAddress&0xFF00 {
			cycles += 2
		} else {
			cycles++
		}

		c.registers.pc = targetAddress
		incrPC = false
	}

	return incrPC, 0, cycles, fmt.Sprintf("BCS $%04X", targetAddress)
}

func (c *Cpu) Bcc(address []byte) (incrPC bool, result byte, cycles byte, debug string) {
	cycles = 2
	incrPC = true
	targetAddress := c.registers.pc + uint16(address[0]) + 2

	if !c.IsFlagSet(FlagCarry) {
		// If the branch jumps to a new page, add an extra cycle
		if c.registers.pc&0xFF00 != targetAddress&0xFF00 {
			cycles += 2
		} else {
			cycles++
		}

		c.registers.pc = targetAddress
		incrPC = false
	}

	return incrPC, 0, cycles, fmt.Sprintf("BCC $%04X", targetAddress)
}

func (c *Cpu) Beq(address []byte) (incrPC bool, result byte, cycles byte, debug string) {
	cycles = 2
	incrPC = true
	targetAddress := c.registers.pc + uint16(address[0]) + 2

	if c.IsFlagSet(FlagZero) {
		// If the branch jumps to a new page, add an extra cycle
		if c.registers.pc&0xFF00 != targetAddress&0xFF00 {
			cycles += 2
		} else {
			cycles++
		}

		c.registers.pc = targetAddress
		incrPC = false
	}

	return incrPC, 0, cycles, fmt.Sprintf("BEQ $%04X", targetAddress)
}

func (c *Cpu) Bne(address []byte) (incrPC bool, result byte, cycles byte, debug string) {
	cycles = 2
	incrPC = true
	targetAddress := c.registers.pc + uint16(address[0]) + 2

	if !c.IsFlagSet(FlagZero) {
		// If the branch jumps to a new page, add an extra cycle
		if c.registers.pc&0xFF00 != targetAddress&0xFF00 {
			cycles += 2
		} else {
			cycles++
		}

		c.registers.pc = targetAddress
		incrPC = false
	}

	return incrPC, 0, cycles, fmt.Sprintf("BNE $%04X", targetAddress)
}

func (c *Cpu) Bvs(address []byte) (incrPC bool, result byte, cycles byte, debug string) {
	cycles = 2
	incrPC = true
	targetAddress := c.registers.pc + uint16(address[0]) + 2

	if c.IsFlagSet(FlagOverflow) {
		// If the branch jumps to a new page, add an extra cycle
		if c.registers.pc&0xFF00 != targetAddress&0xFF00 {
			cycles += 2
		} else {
			cycles++
		}

		c.registers.pc = targetAddress
		incrPC = false
	}

	return incrPC, 0, cycles, fmt.Sprintf("BVS $%04X", targetAddress)
}

func (c *Cpu) Bvc(address []byte) (incrPC bool, result byte, cycles byte, debug string) {
	cycles = 2
	incrPC = true
	targetAddress := c.registers.pc + uint16(address[0]) + 2

	if !c.IsFlagSet(FlagOverflow) {
		// If the branch jumps to a new page, add an extra cycle
		if c.registers.pc&0xFF00 != targetAddress&0xFF00 {
			cycles += 2
		} else {
			cycles++
		}

		c.registers.pc = targetAddress
		incrPC = false
	}

	return incrPC, 0, cycles, fmt.Sprintf("BVC $%04X", targetAddress)
}

func (c *Cpu) Bpl(address []byte) (incrPC bool, result byte, cycles byte, debug string) {
	cycles = 2
	incrPC = true
	targetAddress := c.registers.pc + uint16(address[0]) + 2

	if !c.IsFlagSet(FlagNegative) {
		// If the branch jumps to a new page, add an extra cycle
		if c.registers.pc&0xFF00 != targetAddress&0xFF00 {
			cycles += 2
		} else {
			cycles++
		}

		c.registers.pc = targetAddress
		incrPC = false
	}

	return incrPC, 0, cycles, fmt.Sprintf("BPL $%04X", targetAddress)
}

func (c *Cpu) Bmi(address []byte) (incrPC bool, result byte, cycles byte, debug string) {
	cycles = 2
	incrPC = true
	targetAddress := c.registers.pc + uint16(address[0]) + 2

	if c.IsFlagSet(FlagNegative) {
		// If the branch jumps to a new page, add an extra cycle
		if c.registers.pc&0xFF00 != targetAddress&0xFF00 {
			cycles += 2
		} else {
			cycles++
		}

		c.registers.pc = targetAddress
		incrPC = false
	}

	return incrPC, 0, cycles, fmt.Sprintf("BMI $%04X", targetAddress)
}
