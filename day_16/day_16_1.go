package main

import (
	"fmt"
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
	b1 := reverse(a)
	b2 := ""

	for _, i := range b1 {
		if string(i) == "0" {
			b2 = fmt.Sprintf("%s%s", b2 , "1")
		} else {
			b2 = fmt.Sprintf("%s%s", b2 , "0")
		}
	}

	result := fmt.Sprintf("%s0%s", a, b2)
	return result
}

func check_sum(a string) string {
	check_str := ""

	for i := 0; i < len(a); i = i + 2 {
		if a[i] == a[i+1] {
			check_str = fmt.Sprintf("%s1", check_str)
		} else {
			check_str = fmt.Sprintf("%s0", check_str)
		}
	}

	if (len(check_str) % 2) == 0 {
		return check_sum(check_str)
	} else {
		return check_str
	}
}


func main() {
	initial_state := "10001001100000001"
	target_length := 272
	dragon_string := initial_state
	for len(dragon_string) < target_length {
		dragon_string = dragon_string_double(dragon_string)
	}
	dragon_string = dragon_string[:target_length]
	check_sum_str := check_sum(dragon_string)
	fmt.Println(check_sum_str)
}

