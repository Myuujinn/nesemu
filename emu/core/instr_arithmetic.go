package core

import "fmt"

func (c *Cpu) Adc(value []byte) (incrPC bool, result byte, cycles byte, debug string) {
	toAdd := value[0]
	if c.IsFlagSet(FlagCarry) {
		toAdd++
	}

	// If the result overflows, set the overflow flag
	c.SetFlag(FlagOverflow, int8(c.registers.acc) > 0 && int8(c.registers.acc)+int8(toAdd) < 0 || int8(c.registers.acc) < 0 && int8(toAdd) < 0 && int8(c.registers.acc)+int8(toAdd) > 0)

	c.registers.acc += toAdd

	return true, c.registers.acc, 2, fmt.Sprintf("ADC #$%02X", value[0])
}
