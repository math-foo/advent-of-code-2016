package main

import (
	"fmt"
)


func check(e error) {
    if e != nil {
      panic(e)
    }
}

func check_state(state []int, spaces map[string]bool, visited map[string]struct{}) bool {
	x := state[1]
	y := state[2]

	if (x < 0) || (y < 0) {
		return false
	}

	if is_open(x, y, spaces) {
		key := fmt.Sprintf("%d_%d", x, y)
		_, pres := visited[key]
		if pres {
			return false
		} else {
			visited[key] = struct{}{}
			return true
		}

	} else {
		return false
	}
}

func is_open(x int, y int, spaces map[string]bool) bool {
	key := fmt.Sprintf("%d_%d", x, y)
	value, pres := spaces[key]
	if pres {
		return value
	}

	input_value := 1364
	int_value := x*x + 3*x + 2*x*y + y + y*y + input_value
	ones_count := 0
	for int_value > 0 {
		if (int_value % 2) == 1 {
			ones_count++
		}
		int_value = int_value / 2
	}
	if (ones_count % 2) == 1 {
		spaces[key] = false
		return false
	} else {
		spaces[key] = true
		return true
	}
}

func main() {
	spaces := make(map[string]bool)
	visited := make(map[string]struct{})
	// steps, x, y
	initial_state := []int{0, 1, 1}

	states := [][]int{initial_state}
	visited["1_1"] = struct{}{}

	for true {
		state_to_check := states[0]
		states = states[1:]
		steps := state_to_check[0]
		x := state_to_check[1]
		y := state_to_check[2]
		if steps == 50 {
			break
		}

		up_state := []int{steps + 1, x, y - 1}
		if check_state(up_state, spaces, visited) {
			states = append(states, up_state)
		}
		down_state := []int{steps + 1, x, y + 1}
		if check_state(down_state, spaces, visited) {
			states = append(states, down_state)
		}
		right_state := []int{steps + 1, x + 1, y}
		if check_state(right_state, spaces, visited) {
			states = append(states, right_state)
		}
		left_state := []int{steps + 1, x - 1, y}
		if check_state(left_state, spaces, visited) {
			states = append(states, left_state)
		}
	}
	fmt.Println(len(visited))
}

