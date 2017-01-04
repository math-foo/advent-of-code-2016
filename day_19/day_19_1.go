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
		next_elves := []int{}
		for i := 0; i < len(elves); i = i + 2 {
			next_elves = append(next_elves, elves[i])
		}
		if len(elves)%2 == 0 {
			elves = next_elves
		} else {
			elves = next_elves[1:]
		}
	}

	fmt.Println(elves[0])
}

