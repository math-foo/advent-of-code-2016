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


func main() {
    data, err := ioutil.ReadFile("input.txt")
    check(err)
    disc_values := []int{0}
    disc_bases := []int{1}
    disc_states := strings.Split(strings.Trim(string(data), ".\n"), ".\n")
    for _, disc_state := range disc_states {
	    split_disc_state := strings.Split(disc_state, " ")

	    base, err := strconv.Atoi(split_disc_state[3])
	    check(err)
	    disc_bases = append(disc_bases, base)

	    value, err := strconv.Atoi(split_disc_state[11])
	    check(err)
	    disc_values = append(disc_values, value)
    }
    time_found := -1
    time := 0
    for time_found < 0 {
	    correct_time := true
	    for index, value := range disc_values {
		    value := (time + value + index) % disc_bases[index]
		    if value != 0 {
			    correct_time = false
			    time++
			    break
		    }
	    }
	    if correct_time {
		    time_found = time
	    }
    }
    fmt.Println(time_found)
}

