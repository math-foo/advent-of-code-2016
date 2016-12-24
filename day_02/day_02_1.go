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

func max(a int, b int) int {
	if (a > b) {
		return a
	}

	return b
}


func min(a int, b int) int {
	if (a < b) {
		return a
	}

	return b
}

func step(x int, y int, d string) (int, int){
	if d == "D" {
		y = y + 1
		y = min(2, y)
	} else if d == "U" {
		y = y - 1
		y = max(0, y)
	} else if d == "R" {
		x = x + 1
		x = min(2, x)
	} else {
		x = x - 1
		x = max(0, x)
	}
	return x, y
}

func main() {
    data, err := ioutil.ReadFile("input.txt")
    check(err)
    directions := strings.Split(strings.Trim(string(data), "\n"), "\n")
    x := 1
    y := 1
    code := ""
    for _, entry_string := range directions {
	    for _, d := range strings.Trim(entry_string, "\n") {
		    x, y = step(x, y, string(d))
	    }
	    v := (y*3) + x + 1
	    code += fmt.Sprintf("%d", v)
    }
    fmt.Println(code)
}

