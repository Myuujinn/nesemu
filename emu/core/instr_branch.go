package core

import "fmt"

func (c *Cpu) Bcs(address []byte) (incrPC bool, cycles byte, debug string) {
	targetAddress := c.registers.pc + uint16(address[0]) + 2

	if c.IsFlagSet(FlagCarry) {
		// If the branch jumps to a new page, add an extra cycle
		if c.registers.pc&0xFF00 != targetAddress&0xFF00 {
			cycles += 2
		} else {
			cycles++
		}

		c.registers.pc = targetAddress
	} else {
		incrPC = true
	}

	return incrPC, cycles + 2, fmt.Sprintf("BCS $%04X", targetAddress)
}

func (c *Cpu) Bcc(address []byte) (incrPC bool, cycles byte, debug string) {
	targetAddress := c.registers.pc + uint16(address[0]) + 2

	if !c.IsFlagSet(FlagCarry) {
		// If the branch jumps to a new page, add an extra cycle
		if c.registers.pc&0xFF00 != targetAddress&0xFF00 {
			cycles += 2
		} else {
			cycles++
		}

		c.registers.pc = targetAddress
	} else {
		incrPC = true
	}

	return incrPC, cycles + 2, fmt.Sprintf("BCC $%04X", targetAddress)
}

func (c *Cpu) Beq(address []byte) (incrPC bool, cycles byte, debug string) {
	targetAddress := c.registers.pc + uint16(address[0])

	if c.IsFlagSet(FlagZero) {
		// If the branch jumps to a new page, add an extra cycle
		// We calculate after the branch, so we need to add 2 to the target address
		if (c.registers.pc+2)&0xFF00 != targetAddress&0xFF00 {
			cycles += 2
		} else {
			cycles++
		}

		c.registers.pc = targetAddress
	}

	// This is just for the debug string
	targetAddress += 2

	return true, cycles + 2, fmt.Sprintf("BEQ $%04X", targetAddress)
}

func (c *Cpu) Bne(address []byte) (incrPC bool, cycles byte, debug string) {
	targetAddress := c.registers.pc + uint16(address[0]) + 2

	if !c.IsFlagSet(FlagZero) {
		// If the branch jumps to a new page, add an extra cycle
		if c.registers.pc&0xFF00 != targetAddress&0xFF00 {
			cycles += 2
		} else {
			cycles++
		}

		c.registers.pc = targetAddress
	} else {
		incrPC = true
	}

	return incrPC, cycles + 2, fmt.Sprintf("BNE $%04X", targetAddress)
}

func (c *Cpu) Bvs(address []byte) (incrPC bool, cycles byte, debug string) {
	targetAddress := c.registers.pc + uint16(address[0]) + 2

	if c.IsFlagSet(FlagOverflow) {
		// If the branch jumps to a new page, add an extra cycle
		if c.registers.pc&0xFF00 != targetAddress&0xFF00 {
			cycles += 2
		} else {
			cycles++
		}

		c.registers.pc = targetAddress
	} else {
		incrPC = true
	}

	return incrPC, cycles + 2, fmt.Sprintf("BVS $%04X", targetAddress)
}

func (c *Cpu) Bvc(address []byte) (incrPC bool, cycles byte, debug string) {
	targetAddress := c.registers.pc + uint16(address[0]) + 2

	if !c.IsFlagSet(FlagOverflow) {
		// If the branch jumps to a new page, add an extra cycle
		if c.registers.pc&0xFF00 != targetAddress&0xFF00 {
			cycles += 2
		} else {
			cycles++
		}

		c.registers.pc = targetAddress
	} else {
		incrPC = true
	}

	return incrPC, cycles + 2, fmt.Sprintf("BVC $%04X", targetAddress)
}

func (c *Cpu) Bpl(address []byte) (incrPC bool, cycles byte, debug string) {
	targetAddress := c.registers.pc + uint16(address[0]) + 2

	if !c.IsFlagSet(FlagNegative) {
		// If the branch jumps to a new page, add an extra cycle
		if c.registers.pc&0xFF00 != targetAddress&0xFF00 {
			cycles += 2
		} else {
			cycles++
		}

		c.registers.pc = targetAddress
	} else {
		incrPC = true
	}

	return incrPC, cycles + 2, fmt.Sprintf("BPL $%04X", targetAddress)
}

func (c *Cpu) Bmi(address []byte) (incrPC bool, cycles byte, debug string) {
	targetAddress := c.registers.pc + uint16(address[0]) + 2

	if c.IsFlagSet(FlagNegative) {
		// If the branch jumps to a new page, add an extra cycle
		if c.registers.pc&0xFF00 != targetAddress&0xFF00 {
			cycles += 2
		} else {
			cycles++
		}

		c.registers.pc = targetAddress
	} else {
		incrPC = true
	}

	return incrPC, cycles + 2, fmt.Sprintf("BMI $%04X", targetAddress)
}
