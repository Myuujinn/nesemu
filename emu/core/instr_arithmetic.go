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

func (c *Cpu) Sbc(value []byte) (incrPC bool, cycles byte, debug string) {
	c.Adc([]byte{^value[0]})

	return true, 2, fmt.Sprintf("SBC #$%02X", value[0])
}
