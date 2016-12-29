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



func main() {
    data, err := ioutil.ReadFile("input.txt")
    check(err)
    raw_lines := strings.Split(strings.Trim(string(data), "\n"), "\n")
    compressed_string := ""
    for _, raw_line := range raw_lines {
	    compressed_string = fmt.Sprintf("%s%s", compressed_string, raw_line)
    }
    decompressed_string := ""
    i := 0
    max_i := len(compressed_string)
    for i < max_i {
	    if string(compressed_string[i:i+1]) == "(" {
		    start := i+1
		    end := i
		    for string(compressed_string[end:end+1]) != ")" {
			    end++
		    }
		    repeat_command := strings.Split(compressed_string[start:end], "x")
		    length, err := strconv.Atoi(repeat_command[0])
		    check(err)
		    repeats, err := strconv.Atoi(repeat_command[1])
		    check(err)
		    string_to_repeat := compressed_string[end+1:end+length+1]
		    for j := 0; j < repeats; j++ {
			    decompressed_string = fmt.Sprintf("%s%s", decompressed_string, string_to_repeat)
		    }
		    i = end + length + 1
	    } else {
		    decompressed_string = fmt.Sprintf("%s%s", decompressed_string, string(compressed_string[i:i+1]))
		    i++
	    }
    }
    fmt.Println(len(decompressed_string))
}

