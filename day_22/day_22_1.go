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

type usage struct {
	total, used, avail int
}

func find_terabytes(entry string) int {
	l := len(entry)
	t := string(entry[l-1])
	if t != "T" {
		fmt.Println("Hey! look at this")
		fmt.Println(t)
	}
	total, err := strconv.Atoi(entry[:l-1])
	check(err)
	return total
}

func find_usage(parsed_node []string) usage {
	total := -1
	used := -1
	avail := -1
	for _, entry := range parsed_node[1:] {
		if entry == "" {
			continue
		}

		if total < 0 {
			total = find_terabytes(entry)
		} else if used < 0 {
			used = find_terabytes(entry)
		} else if avail < 0 {
			avail = find_terabytes(entry)
		}
	}

	result := usage{
		total: total,
		used: used,
		avail: avail,
	}
	return result
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	check(err)
	node_usage := strings.Split(strings.Trim(string(data), "\n"), "\n")[2:]
	node_map := make(map[string]usage)
	for _, node := range node_usage {
		parsed_node := strings.Split(node, " ")
		node_code := parsed_node[0][15:]
		node_contents := find_usage(parsed_node)
		node_map[node_code] = node_contents
	}
	viable_pairs := 0
	for node_a, usage_a := range node_map {
		for node_b, usage_b := range node_map {
			if (usage_a.used > 0) && (usage_a.used <= usage_b.avail) && node_a != node_b {
				viable_pairs++
			}
		}
	}
	fmt.Println(viable_pairs)
}

