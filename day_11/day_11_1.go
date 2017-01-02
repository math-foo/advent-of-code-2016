package main

import (
	"fmt"
	"reflect"
)


func check(e error) {
    if e != nil {
      panic(e)
    }
}

func find_items_on_floor(state []int) []int {
	floor := state[1]
	item_indexes := []int{}

	for index, item_floor := range state {
		if index < 2 {
			continue
		}
		if item_floor == floor {
			item_indexes = append(item_indexes, index)
		}
	}
	return item_indexes
}

func check_state(state []int, states_seen map[string]struct{}) bool {
	floors_with_chips := make(map[int]struct{})
	floors_with_gens := make(map[int]struct{})

	for i := 2; i < len(state); i = i + 2 {
		floors_with_gens[state[i]] = struct{}{}
		if state[i] != state[i+1] {
			floors_with_chips[state[i+1]] = struct{}{}
		}
	}

	fried_chip := false
	for floor, _ := range floors_with_chips {
		_, pres := floors_with_gens[floor]
		fried_chip = fried_chip || pres
	}
	// State causes chip to fried
	if fried_chip {
		return false
	}

	state_code := ""
	for _, entry := range state[1:] {
		state_code = fmt.Sprintf("%s_%d", state_code, entry)
	}
	_, already_seen := states_seen[state_code]
	if already_seen {
		return false
	}
	states_seen[state_code] = struct{}{}
	return true
}

func main() {
	// Steps, Elevator, SrG, SrC, PuG, PuC, RuG, RuC, TmG, TmC, CmG, Cmc
	initial_state := []int{0, 1, 1, 1, 1, 1, 2, 2, 2, 3, 2, 2}

	target_state := []int{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4}
	states_seen := make(map[string]struct{})
	state_queue := [][]int{initial_state}
	not_found := true
	steps := -1

	for not_found {
		state_to_check := state_queue[0]
		state_queue = state_queue[1:]
		next_states := [][]int{}
		if reflect.DeepEqual(state_to_check[1:], target_state) {
			not_found = false
			steps = state_to_check[0]
		} else {
			current_floor := state_to_check[1]
			items_on_floor := find_items_on_floor(state_to_check)
			for _, item := range items_on_floor {
				if current_floor < 4 {
					new_state := make([]int, len(state_to_check))
					copy(new_state, state_to_check)
					new_state[0] = state_to_check[0] + 1
					new_state[1] = current_floor + 1
					new_state[item] = current_floor + 1
					good_state := check_state(new_state, states_seen)
					if good_state {
						next_states = append(next_states, new_state)
					}
				}

				if current_floor > 1 {
					new_state := make([]int, len(state_to_check))
					copy(new_state, state_to_check)
					new_state[0] = state_to_check[0] + 1
					new_state[1] = current_floor - 1
					new_state[item] = current_floor - 1
					good_state := check_state(new_state, states_seen)
					if good_state {
						next_states = append(next_states, new_state)
					}
				}
			}

			for i := 0; i < len(items_on_floor); i++ {
				for j:= 0; j < i; j++ {
					item := items_on_floor[i]
					other_item := items_on_floor[j]
					if current_floor < 4 {
						new_state := make([]int, len(state_to_check))
						copy(new_state, state_to_check)
						new_state[0] = state_to_check[0] + 1
						new_state[1] = current_floor + 1
						new_state[item] = current_floor + 1
						new_state[other_item] = current_floor + 1
						good_state := check_state(new_state, states_seen)
						if good_state {
							next_states = append(next_states, new_state)
						}
					}

					if current_floor > 1 {
						new_state := make([]int, len(state_to_check))
						copy(new_state, state_to_check)
						new_state[0] = state_to_check[0] + 1
						new_state[1] = current_floor - 1
						new_state[item] = current_floor - 1
						new_state[other_item] = current_floor - 1
						good_state := check_state(new_state, states_seen)
						if good_state {
							next_states = append(next_states, new_state)
						}
					}
				}
			}
			state_queue = append(state_queue, next_states...)
		}
	}
	fmt.Println(steps)
}

