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

func reverse(a string) string {
	in := []rune(a)
	out := []rune{}
	i := len(a) - 1

	for i >= 0 {
		out = append(out, in[i])
		i--
	}

	return string(out)
}

func swap_letters(parsed_direction []string, password string) string {
	x := strings.Index(password, parsed_direction[2])
	y := strings.Index(password, parsed_direction[5])
	return swap(x, y, password)
}

func swap_positions(parsed_direction []string, password string) string {
	x, err := strconv.Atoi(parsed_direction[2])
	check(err)
	y, err := strconv.Atoi(parsed_direction[5])
	check(err)
	return swap(x, y, password)
}


func reverse_range(parsed_direction []string, password string) string {
	x, err := strconv.Atoi(parsed_direction[2])
	check(err)
	y, err := strconv.Atoi(parsed_direction[4])
	check(err)
	new_password := ""
	if x > 0 {
		new_password = password[:x]
	}

	new_password += reverse(password[x:y+1])

	if y < len(password) - 1 {
		new_password += password[y+1:]
	}

	return new_password
}

func swap(x int, y int, password string) string {
	new_runes := []rune(password)
	tmp := new_runes[x]
	new_runes[x] = new_runes[y]
	new_runes[y] = tmp
	password = string(new_runes)

	return password
}

func rotate_by_letter(parsed_direction []string, password string) string {
	x := strings.Index(password, parsed_direction[6])

	prev := 7
	if (x % 2) == 1 {
		prev = (x - 1) /2
	} else if x > 0 {
		prev = (x/2) + 3
	}
	diff := x - prev
	if diff < 0 {
		diff = diff + len(password)
	}

	if diff == 0 {
		return password
	}

	return rotate(diff, password)
}

func rotate_left(parsed_direction []string, password string) string {
	x, err := strconv.Atoi(parsed_direction[2])
	check(err)
	x = len(password) - x
	if x < 0 {
		x = x + len(password)
	}
	return rotate(x, password)
}

func rotate_right(parsed_direction []string, password string) string {
	x, err := strconv.Atoi(parsed_direction[2])
	check(err)
	return rotate(x, password)
}

func rotate(x int, password string) string {
	x = x % len(password)
	return fmt.Sprintf("%s%s", password[x:], password[:x])
}

func move(parsed_direction []string, password string) string {
	x, err := strconv.Atoi(parsed_direction[2])
	check(err)
	y, err := strconv.Atoi(parsed_direction[5])
	check(err)
	runes := []rune(password)
	tmp := runes[y]
	runes = append(runes[:y], runes[y+1:]...)
	tmp_str := string(runes[x:])
	runes = append(runes[:x], tmp)
	runes = append(runes, []rune(tmp_str)...)
	password = string(runes)
	return password
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	check(err)
	directions := strings.Split(strings.Trim(string(data), "\n"), "\n")
	password := "fbgdceah"
	for i := len(directions) - 1; i >= 0; i-- {
		direction := directions[i]
		parsed_direction := strings.Split(direction, " ")
		if parsed_direction[0] == "swap" {
			if parsed_direction[1] == "position" {
				password = swap_positions(parsed_direction, password)
			} else if parsed_direction[1] == "letter" {
				password = swap_letters(parsed_direction, password)
			}
		} else if parsed_direction[0] == "reverse" {
			password = reverse_range(parsed_direction, password)
		} else if parsed_direction[0] == "rotate" {
			if parsed_direction[1] == "left" {
				password = rotate_left(parsed_direction, password)
			} else if parsed_direction[1] == "right" {
				password = rotate_right(parsed_direction, password)
			} else if parsed_direction[1] == "based" {
				password = rotate_by_letter(parsed_direction, password)
			}
		} else if parsed_direction[0] == "move" {
			password = move(parsed_direction, password)
		}
	}
	fmt.Println(password)
}

