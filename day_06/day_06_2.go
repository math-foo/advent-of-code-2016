package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)


func check(e error) {
    if e != nil {
      panic(e)
    }
}


func top_letter(letter_count map[string]int) string {
	count_letter := map[int][]string{}
	var counts[]int
	for k, v := range letter_count {
		count_letter[v] = append(count_letter[v], k)
	}
	for k := range count_letter {
		counts = append(counts, k)
	}
	sort.Sort(sort.IntSlice(counts))
	for _, v := range count_letter {
		sort.Strings(v)
	}

	top_count := counts[0]
	return count_letter[top_count][0]
}

func main() {
    data, err := ioutil.ReadFile("input.txt")
    check(err)
    transmissions := strings.Split(strings.Trim(string(data), "\n"), "\n")
    frequencies := []map[string]int{
	    make(map[string]int),
	    make(map[string]int),
	    make(map[string]int),
	    make(map[string]int),
	    make(map[string]int),
	    make(map[string]int),
	    make(map[string]int),
	    make(map[string]int),
    }
    for _, transmission := range transmissions {
	    for i, entry := range transmission {
		    entry_str := string(entry)
		    value, pres := frequencies[i][entry_str]
		    if pres {
			    frequencies[i][entry_str] = value + 1
		    } else {
			    frequencies[i][entry_str] = 1
		    }
	    }
    }
    signal := []string{}
    for _, frequency := range frequencies {
	    real_letter := top_letter(frequency)
	    signal = append(signal, real_letter)
    }
    fmt.Println(signal)
}

