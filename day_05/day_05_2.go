package main

import (
	"fmt"
	"crypto/md5"
	"strconv"
)


func check(e error) {
    if e != nil {
      panic(e)
    }
}



func main() {
	door_id := "ojvtpuvg"
	letters_found := 0
	i := 0
	password := []byte{0,0,0,0,0,0,0,0}
	for letters_found < 8 {
		value_to_hash := fmt.Sprintf("%s%d", door_id, i)
		hashed_value := fmt.Sprintf("%x", md5.Sum([]byte(value_to_hash)))
		match := true
		for j := 0; j < 5; j++ {
			match = match && (string(hashed_value[j]) == "0")
		}
		if match {
			index_str := string(hashed_value[5])
			index, err := strconv.Atoi(index_str)
			if err == nil {
				if (index > -1 && index < 8) && (password[index] == 0) {
					password[index] = hashed_value[6]
					letters_found++
				}
			}
		}
		i++
	}
	fmt.Println(string(password))
}

