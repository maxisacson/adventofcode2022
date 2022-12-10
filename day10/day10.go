package day10

import (
	"aoc22/utils"
	"fmt"
	"strconv"
	"strings"
)

type Result struct {
	part1 int
	part2 int
}

type Instruction struct {
	op   string
	data int
}

type Cpu struct {
	X       int
	program []string
	pc      int
	counter int
	cache   Instruction
}

func (c *Cpu) LoadProgram(program []string) {
	c.program = program
	c.pc = 0
	c.X = 1
}

func (c *Cpu) BeginCycle() {
	if c.counter == 0 {
		c.LoadInstruction()
	}
}

func (c *Cpu) EndCycle() {
	c.counter--
	c.ExecuteInstruction()
}

func (c *Cpu) ExecuteInstruction() {
	if c.counter > 0 {
		return
	}

	switch c.cache.op {
	case "noop":
	case "addx":
		c.X += c.cache.data
	default:
		panic(fmt.Sprintf("Unexpected opcode: %s", c.cache.op))
	}

}

func (c *Cpu) LoadInstruction() {
	line := c.program[c.pc]

	if line == "noop" {
		c.counter = 1
		c.cache.op = "noop"
	} else {
		fields := strings.Fields(line)
		c.cache.op = fields[0]
		data, err := strconv.Atoi(fields[1])

		if err != nil {
			panic(err)
		}

		c.cache.data = data

		switch c.cache.op {
		case "addx":
			c.counter = 2
		default:
			panic(fmt.Sprintf("Unexpected opcode: %s", c.cache.op))
		}
	}

	c.pc++
}

func (c *Cpu) HasCycles() bool {
	return c.pc < len(c.program) || c.counter > 0
}

type Crt struct {
	cols   int
	rows   int
	cursor int
	buffer []rune
}

func NewCrt(cols, rows int) Crt {
	return Crt{
		cols:   cols,
		rows:   rows,
		cursor: 0,
		buffer: make([]rune, cols*rows),
	}
}

func (c *Crt) Draw(spritePos int) {
	col := c.cursor % c.cols
	if spritePos-1 <= col && col <= spritePos+1 {
		c.buffer[c.cursor] = '#'
	} else {
		c.buffer[c.cursor] = '.'
	}

	c.cursor += 1
	c.cursor = c.cursor % (c.rows * c.cols)
}

func (c *Crt) GetScreen() string {
	screen := ""
	for row := 0; row < c.rows; row++ {
		screen = screen + string(c.buffer[row*c.cols:row*c.cols+c.cols]) + "\n"
	}
	return screen
}

func Run(fileName string) Result {
	lines := utils.ReadFileToLines(fileName)

	cpu := Cpu{}
	cpu.LoadProgram(lines)

	crt := NewCrt(40, 6)

	inspect := map[int]bool{
		20:  true,
		60:  true,
		100: true,
		140: true,
		180: true,
		220: true,
	}

	sum := 0
	for cycle := 1; cpu.HasCycles(); cycle++ {
		cpu.BeginCycle()
		if inspect[cycle] {
			sum += cycle * cpu.X
		}
		crt.Draw(cpu.X)
		cpu.EndCycle()
	}

	image := crt.GetScreen()
	fmt.Println(image) // ZCBAJFJZ

	return Result{sum, 0}
}
