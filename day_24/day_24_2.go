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

type posn struct {
	x, y, id int
}

type state struct {
	steps int
	p posn
}

func find_adj_posns(a posn) []posn {
	max_y := 43
	max_x := 181
	adj_posns := make([]posn, 4)

	if a.x < max_x {
		adj_posns = append(adj_posns, posn{x: a.x + 1, y: a.y})
	}
	if a.x > 0 {
		adj_posns = append(adj_posns, posn{x: a.x - 1, y: a.y})
	}
	if a.y < max_y {
		adj_posns = append(adj_posns, posn{x: a.x, y: a.y + 1})
	}
	if a.y > 0 {
		adj_posns = append(adj_posns, posn{x: a.x, y: a.y - 1})
	}

	return adj_posns
}

func find_pair_path(a posn, b posn, duct_map [][]bool) int {
	seen_posns := make(map[string]struct{})
	seen_posns[fmt.Sprintf("%d:%d", a.x, a.y)] = struct{}{}
	start_state := state{steps: 0, p: a}
	states := []state{start_state}
	min_dist := -1
	for len(states) > 0 {
		state_to_check := states[0]
		states = states[1:]

		if (state_to_check.p.x == b.x) && (state_to_check.p.y == b.y) {
			min_dist = state_to_check.steps
		}

		adj_posns := find_adj_posns(state_to_check.p)
		for _, adj_posn := range adj_posns {
			posn_code := fmt.Sprintf("%d-%d", adj_posn.x, adj_posn.y)
			_, pres := seen_posns[posn_code]

			// already reached this spot
			if pres {
				continue
			} else {
				seen_posns[posn_code] = struct{}{}
			}

			if duct_map[adj_posn.y][adj_posn.x] {
				new_state := state{steps: state_to_check.steps + 1, p: adj_posn}
				states = append(states, new_state)
			}
		}
	}
	return min_dist
}

func find_shortest_path(start_posn posn, remaining []posn, pair_map map[string]int) int {
	// base case
	if len(remaining) == 2 {
		start_a, ok := pair_map[fmt.Sprintf("%d-%d", start_posn.id, remaining[0].id)]
		end_a, ok2 := pair_map[fmt.Sprintf("%d-%d", 0, remaining[0].id)]
		start_b, ok3 := pair_map[fmt.Sprintf("%d-%d", start_posn.id, remaining[1].id)]
		end_b, ok4 := pair_map[fmt.Sprintf("%d-%d", 0, remaining[1].id)]
		a_b, ok5 := pair_map[fmt.Sprintf("%d-%d", remaining[0].id, remaining[1].id)]

		if !(ok && ok2 && ok3 && ok4 && ok5) {
			fmt.Println("uh-oh quintuple")
		}

		min_3_path := 0
		if (start_a + end_b) < (start_b + end_a) {
			min_3_path = start_a + a_b + end_b
		} else {
			min_3_path = start_b + a_b + end_a
		}
		return min_3_path
	}

	// more general case
	min_path := -1
	for i, next_posn := range remaining {
		value, ok := pair_map[fmt.Sprintf("%d-%d", start_posn.id, next_posn.id)]
		if !ok {
			fmt.Println("uh-oh")
		}

		// No point in checking this one
		if (min_path > 0) && (value >= min_path) {
			continue
		}

		new_remaining := []posn{}
		for j, entry := range remaining {
			if i != j {
				new_remaining = append(new_remaining, entry)
			}
		}

		value = value + find_shortest_path(next_posn, new_remaining, pair_map)

		if (min_path < 0) || (value < min_path) {
			min_path = value
		}

	}

	return min_path
}

func main() {
    data, err := ioutil.ReadFile("input.txt")
    check(err)
    duct_map_strings := strings.Split(strings.Trim(string(data), "\n"), "\n")

    duct_map := [][]bool{}
    start := posn{x:-1,y:-1}
    places := make([]posn, 8)

    for y, duct_map_string := range duct_map_strings {
	    duct_map_row := []bool{}
	    for x, entry := range duct_map_string {
		    entry_str := string(entry)
		    if entry_str == "#" {
			    duct_map_row = append(duct_map_row, false)
		    } else if entry_str == "." {
			    duct_map_row = append(duct_map_row, true)
		    } else {
			    duct_map_row = append(duct_map_row, true)
			    a, err := strconv.Atoi(entry_str)
			    check(err)

			    places[a] = posn{x: x, y: y, id: a}
			    if a == 0 {
				    start = places[a]
			    }
		    }
	    }
	    duct_map = append(duct_map, duct_map_row)
    }

    path_map := make(map[string]int)

    for i, place_i := range places {
	    if i == len(places) - 1 {
		    continue
	    }

	    for j, place_j := range places[i+1:] {
		    real_j := j + i + 1
		    dist := find_pair_path(place_i, place_j, duct_map)
		    path_code_ij := fmt.Sprintf("%d-%d",i,real_j)
		    path_code_ji := fmt.Sprintf("%d-%d",real_j,i)
		    path_map[path_code_ij] = dist
		    path_map[path_code_ji] = dist
	    }
    }

    result := find_shortest_path(start, places[1:], path_map)
    fmt.Println(result)
}


