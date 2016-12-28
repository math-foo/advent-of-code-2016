package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)


func check(e error) {
    if e != nil {
      panic(e)
    }
}


func main() {
    data, err := ioutil.ReadFile("input.txt")
    check(err)
    possible := 0
    triangles := strings.Split(strings.Trim(string(data), "\n"), "\n")
    for _, triangle := range triangles {
	    triangle_values := []int{0,0,0}
	    tri_index := 0
	    for _, value := range strings.Split(strings.Trim(triangle, " \n"), " ") {
		    int_string := strings.Trim(value, " ")
		    if int_string == "" {
			    continue
		    }
		    value, err := strconv.Atoi(int_string)
		    check(err)
		    triangle_values[tri_index] = value
		    tri_index = tri_index + 1
	    }
	    sort.Ints(triangle_values)
	    if (triangle_values[0] + triangle_values[1] > triangle_values[2]) {
		    possible = possible + 1
	    }
    }
    fmt.Println(possible)
}

