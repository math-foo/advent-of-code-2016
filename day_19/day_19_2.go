package main

import (
	"fmt"
)


func check(e error) {
    if e != nil {
      panic(e)
    }
}


func main() {
	number_of_elves := 3014387
	elves := []int{}
	for i := 1; i < number_of_elves + 1; i++ {
		elves = append(elves, i)
	}

	for len(elves) > 1 {
		current_elf := elves[0]
		opposite_index := len(elves)/2
		elves = append(elves[1:opposite_index], elves[opposite_index+1:]...)
		elves = append(elves, current_elf)
	}

	fmt.Println(elves[0])
}

