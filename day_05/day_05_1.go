package main

import (
	"fmt"
	"crypto/md5"
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
	password := []byte{}
	for letters_found < 8 {
		value_to_hash := fmt.Sprintf("%s%d", door_id, i)
		hashed_value := fmt.Sprintf("%x", md5.Sum([]byte(value_to_hash)))
		match := true
		for j := 0; j < 5; j++ {
			match = match && (string(hashed_value[j]) == "0")
		}
		if match {
			password = append(password, hashed_value[5])
			letters_found++
		}
		i++
	}
	fmt.Println(string(password))
}

