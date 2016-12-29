package main

import (
	"fmt"
	"io/ioutil"
//	"sort"
	"strconv"
	"strings"
)


func check(e error) {
    if e != nil {
      panic(e)
    }
}

func decompress_length(text string) int {
    i := 0
    max_i := len(text)
    decompressed_length := 0
    for i < max_i {
	    if string(text[i:i+1]) == "(" {
		    start := i+1
		    end := i
		    for string(text[end:end+1]) != ")" {
			    end++
		    }
		    repeat_command := strings.Split(text[start:end], "x")
		    length, err := strconv.Atoi(repeat_command[0])
		    check(err)
		    repeats, err := strconv.Atoi(repeat_command[1])
		    check(err)

		    string_to_repeat := text[end+1:end+length+1]
		    string_to_repeat_real_length := decompress_length(string_to_repeat)
		    decompressed_length = decompressed_length + repeats * string_to_repeat_real_length
		    i = end + length + 1
	    } else {
		    decompressed_length++
		    i++
	    }
    }
    return decompressed_length
}

func main() {
    data, err := ioutil.ReadFile("input.txt")
    check(err)
    raw_lines := strings.Split(strings.Trim(string(data), "\n"), "\n")
    compressed_string := ""
    for _, raw_line := range raw_lines {
	    compressed_string = fmt.Sprintf("%s%s", compressed_string, raw_line)
    }
    decompressed_length := decompress_length(compressed_string)
    fmt.Println(decompressed_length)
}

