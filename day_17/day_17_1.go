package main

import (
	"fmt"
	"crypto/md5"
)


func check(e error) {
    if e != nil {
      panic(e)
    }
}

type state struct {
	path string
	x, y int
}

func is_open(x string) bool {
	result := ((x == "b") || (x == "c") || (x == "d") || (x == "e") || (x == "f"))
	return result
}

func main() {
	passcode := "qzthpkfp"
	initial_state := state{path: "", x: 0, y:0}
	states := []state{initial_state}
	final_path := ""
	for len(states) > 0 {
		state_to_check := states[0]
		states = states[1:]
		path := state_to_check.path
		x := state_to_check.x
		y := state_to_check.y
		if (x == 3 && y == 3) {
			final_path = path
			break
		}
		passcode_path := fmt.Sprintf("%s%s", passcode, state_to_check.path)
		encoded_path := fmt.Sprintf("%x", md5.Sum([]byte(passcode_path)))
		going_up := (y > 0) && is_open(string(encoded_path[0]))
		going_down := (y < 3) && is_open(string(encoded_path[1]))
		going_left := (x > 0) && is_open(string(encoded_path[2]))
		going_right := (x < 3) && is_open(string(encoded_path[3]))

		if going_up {
			up_state := state{
				path: fmt.Sprintf("%sU", path),
				x: x,
				y: y - 1,
			}
			states = append(states, up_state)
		}

		if going_down {
			down_state := state{
				path: fmt.Sprintf("%sD", path),
				x: x,
				y: y + 1,
			}
			states = append(states, down_state)
		}

		if going_left {
			left_state := state{
				path: fmt.Sprintf("%sL", path),
				x: x - 1,
				y: y,
			}
			states = append(states, left_state)
		}

		if going_right {
			right_state := state{
				path: fmt.Sprintf("%sR", path),
				x: x + 1,
				y: y,
			}
			states = append(states, right_state)
		}
	}
	fmt.Println(final_path)
}

