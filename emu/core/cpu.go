package core

import (
	"fmt"
)

type Registers struct {
	pc     uint16
	sp     uint8
	acc    uint8
	x      uint8
	y      uint8
	status uint8
}

type Cpu struct {
	cycle     uint64
	registers Registers
	Memory    Memory
}

func (c *Cpu) String() string {
	return fmt.Sprintf("A:%02X X:%02X Y:%02X P:%02X SP:%02X PPU:%3d,%3d CYC:%d", c.registers.acc, c.registers.x, c.registers.y, c.registers.status, c.registers.sp, 0, 0, c.cycle)
}

func (c *Cpu) Init() {
	c.registers.status = FlagInterrupt | FlagUnused
	c.registers.sp = 0xFD
	c.cycle = 7

	c.SetEntrypoint(uint16(*c.Memory.Map(0xFFFC)) | uint16(*c.Memory.Map(0xFFFD))<<8)
}

func (c *Cpu) SetEntrypoint(entrypoint uint16) {
	c.registers.pc = entrypoint
}

func (c *Cpu) PushStack(value uint8) {
	*c.Memory.Map(StackStart + uint16(c.registers.sp)) = value
	c.registers.sp--
}

func (c *Cpu) PopStack() uint8 {
	c.registers.sp++
	return *c.Memory.Map(StackStart + uint16(c.registers.sp))
}

func (c *Cpu) Start() {
	fmt.Printf("PC: $%X\n", c.registers.pc)

	for {
		c.Cycle()
	}
}

func (c *Cpu) PrintState(i *Instruction, operands []byte) {
	bytes := fmt.Sprintf("%02X ", i.Opcode)
	for _, op := range operands {
		bytes += fmt.Sprintf("%02X ", op)
	}

	absolute := ""
	if i.Mode == Immediate {
		absolute = "#"
	}

	hex := "$"
	if i.Bytes < 2 {
		hex = ""
	}

	operand := ""
	if len(operands) > 1 {
		operand = fmt.Sprintf("%02X%02X", operands[1], operands[0])
	} else if len(operands) > 0 {
		operand = fmt.Sprintf("%02X", operands[0])
	}
	instr := fmt.Sprintf("%s %s%s%s", i.Name, absolute, hex, operand)

	fmt.Printf("%04X  %-9s %-30s  %s\n", c.registers.pc, bytes, instr, c)
}

func (c *Cpu) Cycle() {
	opcode := *c.Memory.Map(c.registers.pc)

	instruction, ok := instructionMap[opcode]
	if !ok {
		panic(fmt.Sprintf("Unknown opcode: %X", opcode))
	}

	operands := make([]byte, instruction.Bytes-1)
	for i := 0; i < len(operands); i++ {
		operands[i] = *c.Memory.Map(c.registers.pc + uint16(i+1))
	}

	c.PrintState(&instruction, operands)

	incrPC, _, cycles := instruction.Execute(c, operands)
	if incrPC {
		c.registers.pc += uint16(instruction.Bytes)
	}

	c.cycle += uint64(cycles)
}
