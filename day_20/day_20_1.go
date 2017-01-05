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
	ip_map := make(map[int]int)
	data, err := ioutil.ReadFile("input.txt")
	check(err)
	blocked_ranges := strings.Split(strings.Trim(string(data), "\n"), "\n")
	for _, blocked_range := range blocked_ranges {
	    ips := strings.Split(blocked_range, "-")

	    start, err := strconv.Atoi(ips[0])
	    check(err)
	    finish, err := strconv.Atoi(ips[1])
	    check(err)

	    ip_map[start] = finish
	}
	ips := make([]int, len(ip_map))
	i := 0
	for key := range ip_map {
		ips[i] = key
		i++
	}
	sort.Ints(ips)

	max_end := -1
	smallest := -1
	i = 0
	for smallest < 0 {
		end := ip_map[ips[i]]
		if end > max_end {
			max_end = end
		}

		if max_end + 1 < ips[i+1] {
			smallest = end + 1
		}
		i++
	}
	fmt.Println(smallest)
}

