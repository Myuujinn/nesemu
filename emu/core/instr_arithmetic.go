package core

import (
	"fmt"
	"math"
)

func (c *Cpu) Adc(value []byte) (incrPC bool, cycles byte, debug string) {
	toAdd := value[0]
	overflow := int16(int8(c.registers.acc)) + int16(int8(toAdd))
	if c.IsFlagSet(FlagCarry) {
		overflow++
		toAdd++
	}

	c.SetFlag(FlagOverflow, overflow > math.MaxInt8 || overflow < math.MinInt8)
	c.SetFlag(FlagCarry, uint16(c.registers.acc)+uint16(toAdd) > 0xFF)

	c.registers.acc += toAdd

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 2, fmt.Sprintf("ADC #$%02X", value[0])
}

func (c *Cpu) IndexedIndirectAdc(address []byte) (incrPC bool, cycles byte, debug string) {
	target := uint16(address[0]) + uint16(c.registers.x)
	targetAddress := uint16(*c.Memory.Map(target % 0x100)) + uint16(*c.Memory.Map((target + 1) % 0x100))*0x100
	toAdd := *c.Memory.Map(targetAddress)
	overflow := int16(int8(c.registers.acc)) + int16(int8(toAdd))
	if c.IsFlagSet(FlagCarry) {
		overflow++
		toAdd++
	}

	c.SetFlag(FlagOverflow, overflow > math.MaxInt8 || overflow < math.MinInt8)
	c.SetFlag(FlagCarry, uint16(c.registers.acc)+uint16(toAdd) > 0xFF)

	c.registers.acc += toAdd

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 6, fmt.Sprintf("ADC ($%02X,X) @ %02X = %04X = %02X", address[0], uint8(target), targetAddress, *c.Memory.Map(targetAddress))
}

func (c *Cpu) ZeropageAdc(address []byte) (incrPC bool, cycles byte, debug string) {
	toAdd := *c.Memory.Map(uint16(address[0]))
	overflow := int16(int8(c.registers.acc)) + int16(int8(toAdd))
	if c.IsFlagSet(FlagCarry) {
		overflow++
		toAdd++
	}

	c.SetFlag(FlagOverflow, overflow > math.MaxInt8 || overflow < math.MinInt8)
	c.SetFlag(FlagCarry, uint16(c.registers.acc)+uint16(toAdd) > 0xFF)

	c.registers.acc += toAdd

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 3, fmt.Sprintf("ADC $%02X = %02X", address[0], *c.Memory.Map(uint16(address[0])))
}

func (c *Cpu) AbsoluteAdc(address []byte) (incrPC bool, cycles byte, debug string) {
	toAdd := *c.Memory.Map(uint16(address[0]) | uint16(address[1])<<8)
	overflow := int16(int8(c.registers.acc)) + int16(int8(toAdd))
	if c.IsFlagSet(FlagCarry) {
		overflow++
		toAdd++
	}

	c.SetFlag(FlagOverflow, overflow > math.MaxInt8 || overflow < math.MinInt8)
	c.SetFlag(FlagCarry, uint16(c.registers.acc)+uint16(toAdd) > 0xFF)

	c.registers.acc += toAdd

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 4, fmt.Sprintf("ADC $%02X%02X = %02X", address[1], address[0], *c.Memory.Map(uint16(address[0]) | uint16(address[1])<<8))
}

func (c *Cpu) Sbc(value []byte) (incrPC bool, cycles byte, debug string) {
	c.Adc([]byte{^value[0]})

	return true, 2, fmt.Sprintf("SBC #$%02X", value[0])
}

func (c *Cpu) ZeropageSbc(address []byte) (incrPC bool, cycles byte, debug string) {
	target := uint16(address[0])
	toAdd := ^(*c.Memory.Map(target))
	overflow := int16(int8(c.registers.acc)) + int16(int8(toAdd))
	if c.IsFlagSet(FlagCarry) {
		overflow++
		toAdd++
	}

	c.SetFlag(FlagOverflow, overflow > math.MaxInt8 || overflow < math.MinInt8)
	c.SetFlag(FlagCarry, uint16(c.registers.acc)+uint16(toAdd) > 0xFF)

	c.registers.acc += toAdd

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 3, fmt.Sprintf("SBC $%02X = %02X", address[0], *c.Memory.Map(target))
}

func (c *Cpu) IndexedIndirectSbc(address []byte) (incrPC bool, cycles byte, debug string) {
	target := uint16(address[0]) + uint16(c.registers.x)
	targetAddress := uint16(*c.Memory.Map(target % 0x100)) + uint16(*c.Memory.Map((target + 1) % 0x100))*0x100
	toAdd := ^(*c.Memory.Map(targetAddress))
	overflow := int16(int8(c.registers.acc)) + int16(int8(toAdd))
	if c.IsFlagSet(FlagCarry) {
		overflow++
		toAdd++
	}

	c.SetFlag(FlagOverflow, overflow > math.MaxInt8 || overflow < math.MinInt8)
	c.SetFlag(FlagCarry, uint16(c.registers.acc)+uint16(toAdd) > 0xFF)

	c.registers.acc += toAdd

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 6, fmt.Sprintf("SBC ($%02X,X) @ %02X = %04X = %02X", address[0], uint8(target), targetAddress, *c.Memory.Map(targetAddress))
}

func (c *Cpu) AbsoluteSbc(address []byte) (incrPC bool, cycles byte, debug string) {
	target := uint16(address[0]) | uint16(address[1])<<8
	toAdd := ^(*c.Memory.Map(target))
	overflow := int16(int8(c.registers.acc)) + int16(int8(toAdd))
	if c.IsFlagSet(FlagCarry) {
		overflow++
		toAdd++
	}

	c.SetFlag(FlagOverflow, overflow > math.MaxInt8 || overflow < math.MinInt8)
	c.SetFlag(FlagCarry, uint16(c.registers.acc)+uint16(toAdd) > 0xFF)

	c.registers.acc += toAdd

	c.SetFlag(FlagZero, c.registers.acc == 0)
	c.SetFlag(FlagNegative, c.registers.acc&0b10000000 == 0b10000000)

	return true, 4, fmt.Sprintf("SBC $%02X%02X = %02X", address[1], address[0], *c.Memory.Map(target))
}
