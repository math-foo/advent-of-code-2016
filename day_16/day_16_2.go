package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)


func check(e error) {
    if e != nil {
      panic(e)
    }
}

// You have got to be fucking kidding me golang
func reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func dragon_string_double(a string) string {
	b_slice := []string{}

	for _, i := range a {
		if string(i) == "0" {
			b_slice = append(b_slice, "1")
		} else {
			b_slice = append(b_slice, "0")
		}
	}

	b := reverse(strings.Join(b_slice, ""))
	result := fmt.Sprintf("%s0%s", a, b)
	return result
}

func check_sum(a string) string {
	check_str_slice := []string{}

	for i := 0; i < len(a); i = i + 2 {
		if a[i] == a[i+1] {
			check_str_slice = append(check_str_slice, "1")
		} else {
			check_str_slice = append(check_str_slice, "0")
		}
	}

	check_str := strings.Join(check_str_slice, "")
	if (len(check_str) % 2) == 0 {
		return check_sum(check_str)
	} else {
		return check_str
	}
}


func main() {
	initial_state := "10001001100000001"
	target_length := 35651584
	dragon_string := initial_state
	for len(dragon_string) < target_length {
		dragon_string = dragon_string_double(dragon_string)
	}
	dragon_string = dragon_string[:target_length]
	check_sum_str := check_sum(dragon_string)
	fmt.Println(check_sum_str)
}

