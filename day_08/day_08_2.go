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

func int_min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func display_screen(lights [][]bool) {
	for _, row := range lights {
		row_sl := []byte{}
		for _, light := range row {
			if light {
				row_sl = append(row_sl, '#')
			} else {
				row_sl = append(row_sl, ' ')
			}
		}
		fmt.Println(string(row_sl))
	}
}

func place_rect(rect_command []string, lights [][]bool) {
	rect_dimensions := strings.Split(rect_command[1], "x")

	width, err := strconv.Atoi(rect_dimensions[0])
	check(err)
	width = int_min(width, len(lights[0]))

	height, err := strconv.Atoi(rect_dimensions[1])
	check(err)
	height = int_min(height, len(lights))

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			lights[i][j] = true
		}
	}
}

func rotate(rotate_command []string, lights [][]bool) {
	value, err := strconv.Atoi(rotate_command[4])
	check(err)

	place, err := strconv.Atoi(strings.Split(rotate_command[2], "=")[1])
	check(err)

	direction := rotate_command[1]
	if direction == "row" {
		rotate_row(place, value, lights)
	} else if direction == "column" {
		rotate_column(place, value, lights)
	} else {
		fmt.Println("Ivalid direction")
		fmt.Println(rotate_command)
	}

}

func rotate_row(row int, value int, lights [][]bool) {
	base := len(lights[0])
	original_row := []bool{}
	for i := 0; i < base; i++ {
		original_row = append(original_row, lights[row][i])
	}

	for i := 0; i < base; i++ {
		j := (i + value) % base
		lights[row][j] = original_row[i]
	}
}

func rotate_column(column int, value int, lights [][]bool) {
	base := len(lights)
	original_col := []bool{}
	for i := 0; i < base; i++ {
		original_col = append(original_col, lights[i][column])
	}

	for i := 0; i < base; i++ {
		j := (i + value) % base
		lights[j][column] = original_col[i]
	}
}

func main() {
    data, err := ioutil.ReadFile("input.txt")
    check(err)
    instructions := strings.Split(strings.Trim(string(data), "\n"), "\n")
    height := 6
    width := 50
    lights := make([][]bool, height)
    for i := 0; i < height; i++ {
	    lights[i] = make([]bool, width)
	    for j := 0; j < width; j++ {
		    lights[i][j] = false
	    }
    }

    for _, instruction := range instructions {
	    parts := strings.Split(instruction, " ")
	    command := parts[0]
	    if command == "rect" {
		    place_rect(parts, lights)
	    } else if command == "rotate" {
		    rotate(parts, lights)
	    } else {
		    fmt.Println("Bad command!")
		    fmt.Println(instruction)
		    break
	    }
    }
    display_screen(lights)
}

