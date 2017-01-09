package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)


func check(e error) {
    if e != nil {
      panic(e)
    }
}

func copy_instruction(parsed_instrc []string, registers map[string]int) {
//	fmt.Println("copy")
	x := parsed_instrc[1]
	y := parsed_instrc[2]

	x_value, ok := strconv.Atoi(x)
	if ok == nil {
		registers[y] = x_value
	} else {
		x_reg_value := registers[x]
		registers[y] = x_reg_value
	}
}

func jump_instruction(parsed_instrc []string, registers map[string]int, toggled bool) int {
//	fmt.Println("jump")
	jump_decide := parsed_instrc[1]
	jump_amount := parsed_instrc[2]

	decide_value, ok := strconv.Atoi(jump_decide)
	if ok != nil {
		decide_value = registers[jump_decide]
	}
	amount_value, ok := strconv.Atoi(jump_amount)
	if ok != nil {
		amount_value = registers[jump_amount]
	}

	if decide_value == 0 {
		return 1
	} else {
		return amount_value
	}
}

func toggle_instruction(parsed_instrc []string, program_index int, registers map[string]int, program []*program_line) {
//	fmt.Println("toggle")
	tgl_decide := parsed_instrc[1]
	tgl_value := program_index + registers[tgl_decide]
	if tgl_value < len(program) {
		program[tgl_value].toggled = !program[tgl_value].toggled
	}
}

func run_instruction(instruction *program_line, registers map[string]int, program_index int, program []*program_line) int {
	parsed_instrc := strings.Split(instruction.line, " ")
	base_instrc := parsed_instrc[0]
	if base_instrc == "cpy" {
		if instruction.toggled {
		  jump_result := jump_instruction(parsed_instrc, registers, instruction.toggled)
		  program_index = program_index + jump_result
	        } else {
		  copy_instruction(parsed_instrc, registers)
		  program_index++
	        }
	} else if base_instrc == "inc" {
		if instruction.toggled {
	//	  fmt.Println("dec")
		  registers[parsed_instrc[1]]--
		  program_index++
		} else {
	//	  fmt.Println("inc")
		  registers[parsed_instrc[1]]++
		  program_index++
		}
	} else if base_instrc == "dec" {
		if instruction.toggled {
	//	  fmt.Println("inc")
		  registers[parsed_instrc[1]]++
		  program_index++
		} else {
	//	  fmt.Println("dec")
		  registers[parsed_instrc[1]]--
		  program_index++
		}
	} else if base_instrc == "jnz" {
		if instruction.toggled {
		  copy_instruction(parsed_instrc, registers)
		  program_index++
	        } else {
		  jump_result := jump_instruction(parsed_instrc, registers, instruction.toggled)
		  program_index = program_index + jump_result
	        }
	} else if base_instrc == "tgl" {
		if instruction.toggled {
	//	  fmt.Println("inc")
		  registers[parsed_instrc[1]]++
		  program_index++
		} else {
		  toggle_instruction(parsed_instrc, program_index, registers, program)
		  program_index++
		}
	} else {
		fmt.Println("uh-oh!")
		fmt.Println(instruction)
		program_index++
	}
	return program_index
}

type program_line struct {
	line string
	toggled bool
}

func main() {
    data, err := ioutil.ReadFile("input.txt")
    check(err)
    raw_program := strings.Split(strings.Trim(string(data), "\n"), "\n")
    /**raw_program = []string{
	    "cpy 2 a",
	    "tgl a",
	    "tgl a",
	    "tgl a",
	    "cpy 1 a",
	    "dec a",
	    "dec a",
    }**/

    program := []*program_line{}
    for _, raw_instruction := range raw_program {
	    program = append(program, &program_line{line: raw_instruction, toggled: false})
    }

    registers := make(map[string]int)
    registers["a"] = 7
    registers["b"] = 0
    registers["c"] = 0
    registers["d"] = 0

    program_index := 0
    for program_index < len(program) {
	//    fmt.Println(registers, program_index)
	    instruction := program[program_index]
	    program_index = run_instruction(instruction, registers, program_index, program)
    }
    fmt.Println(registers["a"])
}

