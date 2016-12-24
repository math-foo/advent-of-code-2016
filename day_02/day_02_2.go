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

func key(x int, y int) string {
	if (y == 0) {
		return "1"
	} else if (y == 1) {
		return fmt.Sprintf("%d", 1 + x)
	} else if (y == 2) {
		return fmt.Sprintf("%d", 5 + x)
	} else if (y == 3) {
		if (x == 1) {
			return "A"
		} else if (x == 2) {
			return "B"
		} else {
			return "C"
		}
	} else {
		return "D"
	}
}

func step(x int, y int, d string) (int, int){
	if d == "D" {
		if (x == 0 || x == 4) {
			return x, y
		} else if (x == 1 || x == 3) {
			y = y + 1
			y = min(y, 3)
		} else {
			y = y + 1
			y = min(y, 4)
		}
	} else if d == "U" {
		if (x == 1 || x == 3) {
			y = y - 1
			y = max(y, 1)
		} else if x == 2 {
			y = y - 1
			y = max(y, 0)
		}
	} else if d == "R" {
		if (y == 1 || y == 3) {
			x = x + 1
			x = min(x, 3)
		} else if y == 2 {
			x = x + 1
			x = min(x, 4)
		}
	} else {
		if (y == 1 || y == 3) {
			x = x - 1
			x = max(x, 1)
		} else if y == 2 {
			x = x - 1
			x = max(x, 0)
		}
	}
	return x, y
}

func main() {
    data, err := ioutil.ReadFile("input.txt")
    check(err)
    directions := strings.Split(strings.Trim(string(data), "\n"), "\n")
    x := 0
    y := 2
    code := ""
    for _, entry_string := range directions {
	    for _, d := range strings.Trim(entry_string, "\n") {
		    x, y = step(x, y, string(d))
	    }
	    v := key(x, y)
	    code += v
    }
    fmt.Println(code)
}

