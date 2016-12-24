package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)


func check(e error) {
    if e != nil {
      panic(e)
    }
}

// Fuck it. Giant if-then-else clause
func new_direction(facing string, turning string) string {
	if turning == "R" {
	  if facing == "N" {
		  return "E"
	  } else if facing == "E" {
		  return "S"
	  } else if facing == "S" {
		  return "W"
	  } else if facing == "W" {
		  return "N"
	  }
	} else if turning == "L" {
	  if facing == "N" {
		  return "W"
	  } else if facing == "E" {
		  return "N"
	  } else if facing == "S" {
		  return "E"
	  } else if facing == "W" {
		  return "S"
	  }
	}
	return "Not Found"
}


// I want to use global variables.
func take_step(steps int, dir string, x int, y int) (int, int) {
	if dir == "N" {
		y = y + steps
	} else if dir == "S" {
		y = y - steps
	} else if dir == "E" {
		x = x + steps
	} else {
		x = x - steps
	}

	return x, y
}

func abs(n int) int {
	if (n < 0) {
		return -n
	} else {
		return n
	}
}

func main() {
    data, err := ioutil.ReadFile("input.txt")
    check(err)
    directions := strings.Split(strings.Trim(string(data), "\n"), ", ")
    dir := "N"
    pos_x := 0
    pos_y := 0

    for _, steps := range directions {
	    turn := steps[0:1]
	    count, err := strconv.Atoi(steps[1:])
	    check(err)
	    dir = new_direction(dir, turn)
	    pos_x, pos_y = take_step(count, dir, pos_x, pos_y)
    }
    fmt.Println(abs(pos_x) + abs(pos_y))
}

