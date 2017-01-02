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

func jump_instruction(parsed_instrc []string, registers map[string]int) int {
	jump_decide := parsed_instrc[1]
	jump_amount := parsed_instrc[2]

	decide_value, ok := strconv.Atoi(jump_decide)
	if ok != nil {
		decide_value = registers[jump_decide]
	}

	if decide_value == 0 {
		return 1
	} else {
		amount_value, err := strconv.Atoi(jump_amount)
		check(err)
		return amount_value
	}
}

func run_instruction(instruction string, registers map[string]int, program_index int) int {
	parsed_instrc := strings.Split(instruction, " ")
	base_instrc := parsed_instrc[0]
	if base_instrc == "cpy" {
		copy_instruction(parsed_instrc, registers)
		program_index++
	} else if base_instrc == "inc"{
		registers[parsed_instrc[1]]++
		program_index++
	} else if base_instrc == "dec"{
		registers[parsed_instrc[1]]--
		program_index++
	} else if base_instrc == "jnz"{
		jump_result := jump_instruction(parsed_instrc, registers)
		program_index = program_index + jump_result
	} else {
		fmt.Println("uh-oh!")
		fmt.Println(instruction)
		program_index++
	}
	return program_index
}

func main() {
    data, err := ioutil.ReadFile("input.txt")
    check(err)
    program := strings.Split(strings.Trim(string(data), "\n"), "\n")

    registers := make(map[string]int)
    registers["a"] = 0
    registers["b"] = 0
    registers["c"] = 0
    registers["d"] = 0

    program_index := 0
    for program_index < len(program) {
	    instruction := program[program_index]
	    program_index = run_instruction(instruction, registers, program_index)
    }
    fmt.Println(registers["a"])
}

