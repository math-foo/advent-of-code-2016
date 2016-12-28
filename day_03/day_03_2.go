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
    all_values := [][]int{}
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
	    all_values = append(all_values, triangle_values)
    }
    count := len(all_values)/3
    all_triangles := [][]int{}
    for i := 0; i < count; i++ {
	    first_triangle := []int{
		    all_values[i * 3][0],
		    all_values[i * 3 + 1][0],
		    all_values[i * 3 + 2][0],
	    }
	    second_triangle := []int{
		    all_values[i * 3][1],
		    all_values[i * 3 + 1][1],
		    all_values[i * 3 + 2][1],
	    }
	    third_triangle := []int{
		    all_values[i * 3][2],
		    all_values[i * 3 + 1][2],
		    all_values[i * 3 + 2][2],
	    }
	    all_triangles = append(all_triangles, first_triangle)
	    all_triangles = append(all_triangles, second_triangle)
	    all_triangles = append(all_triangles, third_triangle)
    }
    for _, triangle := range all_triangles {
	    sort.Ints(triangle)
	    if (triangle[0] + triangle[1] > triangle[2]) {
		    possible++
	    }
    }
    fmt.Println(possible)
}

