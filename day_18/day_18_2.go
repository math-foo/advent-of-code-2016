package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)


func check(e error) {
    if e != nil {
      panic(e)
    }
}

func is_safe(x string) bool {
	return x == "."
}

func safe_count(row string) int {
	count := 0
	for _, entry := range row {
		if is_safe(string(entry)) {
			count++

		}
	}

	return count
}

func next_row(row string) string {
	next_row_slice := []string{}
	row_length := len(row)
	for index, _ := range row {
		left_safe := false
		right_safe := false

		if index < 1 {
			left_safe = true
		} else {
			left_safe = is_safe(string(row[index-1]))
		}

		if index > row_length - 2 {
			right_safe = true
		} else {
			right_safe = is_safe(string(row[index+1]))
		}

		if left_safe == right_safe {
			next_row_slice = append(next_row_slice, ".")
		} else {
			next_row_slice = append(next_row_slice, "^")
		}
	}

	next_row := strings.Join(next_row_slice, "")
	return next_row
}

func main() {
    data, err := ioutil.ReadFile("input.txt")
    check(err)
    first_row_tiles := strings.Trim(string(data), "\n")
    current_row := first_row_tiles
    row_count := 0
    safe_tile_count := 0
    for row_count < 400000 {
	    safe_tile_count = safe_tile_count + safe_count(current_row)
	    row_count++
	    current_row = next_row(current_row)
    }

    fmt.Println(safe_tile_count)
}

