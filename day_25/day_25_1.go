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

func out_instruction(parsed_instrc []string, registers map[string]int, output string) string {
	out_entry := parsed_instrc[1]

	out_value, ok := strconv.Atoi(out_entry)
	if ok != nil {
		out_value = registers[out_entry]
	}

	len_output := len(output)

	// return "" when output is invalid
	if (out_value != 0) && (out_value != 1) {
		return ""
	} else if len_output > 0 {
		last_value, err := strconv.Atoi(string(output[len_output - 1]))
		check(err)
		if last_value == out_value {
			return ""
		} else {
			return fmt.Sprintf("%s%d", output, out_value)
		}
	} else {
		return fmt.Sprintf("%s%d", output, out_value)
	}
}

type program_result struct {
	index int
	output string
}

func run_instruction(instruction string, registers map[string]int, program_index int, output string) program_result {
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
	} else if base_instrc == "out"{
		output = out_instruction(parsed_instrc, registers, output)
		// output was invalid
		if (len(output) == 0) {
			program_index = 99999
		}
		program_index++
	} else {
		fmt.Println("uh-oh!")
		fmt.Println(instruction)
		program_index++
	}

	return program_result{index: program_index, output: output}
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

    init_a := 1
    for true {
	    registers["a"] = init_a
            program_index := 0
	    output := ""
	    states_seen := make(map[string]struct{})
	    for program_index < len(program) {
		    state_code := fmt.Sprintf("i:%d-a:%d-b:%d-c:%d-d:%d", program_index, registers["a"], registers["b"], registers["c"], registers["d"])
		    _, pres := states_seen[state_code]
		    if pres {
			    break
		    } else {
			    states_seen[state_code] = struct{}{}
		    }
		    instruction := program[program_index]
		    program_output := run_instruction(instruction, registers, program_index, output)
		    program_index = program_output.index
		    output = program_output.output
            }

	    if program_index < len(program) {
		    break
	    }

	    init_a++
    }

    fmt.Println(init_a)
}

