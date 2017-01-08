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

type contents struct {
	total, used, avail int
}

type state struct {
	steps int
	space, target string
}

func find_adjacent(node_address string) []string {
	max_x := 32
	max_y := 29
	coordinates := strings.Split(node_address, "-")
	x, err := strconv.Atoi(coordinates[0][1:])
	check(err)
	y, err := strconv.Atoi(coordinates[1][1:])
	check(err)

	adjacent_nodes := []string{}
	if x > 0 {
		adjacent_nodes = append(adjacent_nodes, fmt.Sprintf("x%d-y%d", x - 1, y))
	}
	if x < max_x {
		adjacent_nodes = append(adjacent_nodes, fmt.Sprintf("x%d-y%d", x + 1, y))
	}

	if y > 0 {
		adjacent_nodes = append(adjacent_nodes, fmt.Sprintf("x%d-y%d", x, y - 1))
	}
	if y < max_y {
		adjacent_nodes = append(adjacent_nodes, fmt.Sprintf("x%d-y%d", x, y + 1))
	}


	return adjacent_nodes
}

func find_terabytes(entry string) int {
	l := len(entry)
	t := string(entry[l-1])
	if t != "T" {
		fmt.Println("Hey! look at this")
		fmt.Println(t)
	}
	total, err := strconv.Atoi(entry[:l-1])
	check(err)
	return total
}

func find_contents(parsed_node []string, node_code string) contents {
	total := -1
	used := -1
	avail := -1
	for _, entry := range parsed_node[1:] {
		if entry == "" {
			continue
		}

		if total < 0 {
			total = find_terabytes(entry)
		} else if used < 0 {
			used = find_terabytes(entry)
		} else if avail < 0 {
			avail = find_terabytes(entry)
		}
	}

	result := contents{
		total: total,
		used: used,
		avail: avail,
	}
	return result
}

func find_state_code(state_to_check state) string {
	return fmt.Sprintf("%s:%s", state_to_check.space, state_to_check.target)
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	check(err)
	node_usage := strings.Split(strings.Trim(string(data), "\n"), "\n")[2:]
	node_map := make(map[string]contents)
	for _, node := range node_usage {
		parsed_node := strings.Split(node, " ")
		node_code := parsed_node[0][15:]
		node_contents := find_contents(parsed_node, node_code)
		node_map[node_code] = node_contents
	}
	spaces := make(map[string]struct{})
	for node_a, usage_a := range node_map {
		for node_b, usage_b := range node_map {
			if (usage_a.used > 0) && (usage_a.used <= usage_b.avail) && node_a != node_b {
				spaces[node_b] = struct{}{}
			}
		}
	}
	rocks := make(map[string]struct{})
	for node, usage := range node_map {
		is_rock := true
		for space := range spaces {
			if usage.used <= node_map[space].avail {
				is_rock = false
			}
		}

		if is_rock {
			rocks[node] = struct{}{}
		}
	}
	if len(spaces) != 1 {
		fmt.Printf("This solution is only for grids with 1 space, this one has %d\n", len(spaces))
		return
	}
	space := ""
	for entry := range spaces {
		space = entry
	}
	target_file := "x32-y0"
	target_location := "x0-y0"
	start_state := state{steps: 0, space: space, target: target_file}
	states_seen := make(map[string]struct{})
	states_seen[fmt.Sprintf("%s:%s", space, target_file)] = struct{}{}
	states := []state{start_state}

	steps_to_move_target := -1
	//max_step := -1
	for len(states) > 0 {
		//fmt.Println(states)
		state_to_check := states[0]
		states = states[1:]
		if state_to_check.target == target_location {
			steps_to_move_target = state_to_check.steps
			break
		}

		adjacent_nodes := find_adjacent(state_to_check.space)
		for _, adj_node := range adjacent_nodes {
			_, pres := rocks[adj_node]
			// skip the rocks
			if pres {
				continue
			}

			new_state := state{
				steps: state_to_check.steps + 1,
				space: adj_node,
				target: state_to_check.target,
			}

			// the target file was moved
			if adj_node == state_to_check.target {
				new_state.target = state_to_check.space
			}

			// have we seen this state before
			state_code := find_state_code(new_state)
			_, pres = states_seen[state_code]
			if pres {
				continue
			} else {
				states_seen[state_code] = struct{}{}
			}

			states = append(states, new_state)
		}
	}
	fmt.Println(steps_to_move_target)
}

