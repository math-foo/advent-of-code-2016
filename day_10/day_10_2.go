package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)


type bot struct {
	id, chip1, chip2, high_id, low_id int
	high_type, low_type string
}

func reset_bot(b *bot) {
	b.chip1 = -1
	b.chip2 = -1
}

func activate(b *bot, bots map[int]*bot, outputs map[int]int) []int {
	low := -1
	high := -1
	if b.chip1 < b.chip2 {
		low = b.chip1
		high = b.chip2
	} else {
		low = b.chip2
		high = b.chip1
	}

	next := false
	next_bots := []int{}

	if b.high_type == "bot" {
		high_bot := bots[b.high_id]
		next = give(high_bot, high)
		if next {
			next_bots = append(next_bots, b.high_id)
		}
	} else {
		outputs[b.high_id] = high
	}

	if b.low_type == "bot" {
		low_bot := bots[b.low_id]
		next = give(low_bot, low)
		if next {
			next_bots = append(next_bots, b.low_id)
		}
	} else {
		outputs[b.low_id] = low
	}
	reset_bot(b)

	return next_bots
}

func give(b *bot, chip int) bool {
  if b.chip1 < 0 {
	  b.chip1 = chip
	  return false
  } else {
	  b.chip2 = chip
	  return true
  }
}

func create_bot(instructions []string) bot {
	bot_id, err := strconv.Atoi(instructions[1])
	check(err)
	low_type_str := instructions[5]
	low_id_value, err := strconv.Atoi(instructions[6])
	check(err)

	high_type_str := instructions[10]
	high_id_value, err := strconv.Atoi(instructions[11])
	check(err)

	b := bot{
		id: bot_id,
		low_type: low_type_str,
		low_id: low_id_value,
		high_type: high_type_str,
		high_id: high_id_value,
		chip1: -1,
		chip2: -1}
	return b
}

func check(e error) {
    if e != nil {
      panic(e)
    }
}


func main() {
    data, err := ioutil.ReadFile("input.txt")
    check(err)
    instructions := strings.Split(strings.Trim(string(data), "\n"), "\n")
    bot_instructions := [][]string{}
    value_instructions := [][]string{}
    bots := make(map[int]*bot)
    outputs := make(map[int]int)
    for _, raw_instruction := range instructions {
	    instruction := strings.Split(strings.Trim(raw_instruction, "\n"), " ")
	    instruction_type := instruction[0]
	    if instruction_type == "bot" {
		    bot_instructions = append(bot_instructions, instruction)
	    } else if instruction_type == "value" {
		    value_instructions = append(value_instructions, instruction)
	    } else {
		    fmt.Println("Bad instruction")
		    fmt.Println(instruction)
		    break
	    }
    }

    for _, bot_instruction := range bot_instructions {
	    new_bot := create_bot(bot_instruction)
	    bots[new_bot.id] = &new_bot
    }
    active_bots := []*bot{}
    for _, value_instruction := range value_instructions {
	    chip_id, err := strconv.Atoi(value_instruction[1])
	    check(err)

	    bot_id, err := strconv.Atoi(value_instruction[5])
	    check(err)

	    bot := bots[bot_id]
	    activated := give(bot, chip_id)
	    if activated {
		    active_bots = append(active_bots, bot)
	    }
    }
    for len(active_bots) > 0 {
	    next_bots := []int{}
	    for _, active_bot := range active_bots {
		    these_bots := activate(active_bot, bots, outputs)
		    next_bots = append(next_bots, these_bots...)
	    }
	    active_bots = []*bot{}
	    for _, bot_id := range next_bots {
		    active_bots = append(active_bots, bots[bot_id])
	    }
    }
    a, _ := outputs[0]
    b, _ := outputs[1]
    c, _ := outputs[2]
    fmt.Println(a * b * c)
}

