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

func top_5_letters(letter_count map[string]int) []string {
	count_letter := map[int][]string{}
	var counts[]int
	for k, v := range letter_count {
		count_letter[v] = append(count_letter[v], k)
	}
	for k := range count_letter {
		counts = append(counts, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))
	for _, v := range count_letter {
		sort.Strings(v)
	}

	letters := []string{}
	letter_index := 0

	for k := range counts {
		count := counts[k]
		for _, v := range count_letter[count] {
			if letter_index >= 5 {
				break
			} else {
				letters = append(letters, v)
				letter_index++
			}
		}
	}
	return letters
}

func shift_decode(encrypted string, shift int) string {
	shift = shift % 26
	z_byte := 122
	base_shift := 26
	decrypted := []byte{}
	for _, i := range encrypted {
		j := int(i) + shift
		if j > z_byte {
			j = j - base_shift
		}
		decrypted = append(decrypted, byte(j))
	}
	return string(decrypted)
}

func main() {
    data, err := ioutil.ReadFile("input.txt")
    check(err)
    room_codes := strings.Split(strings.Trim(string(data), "\n"), "\n")
    for _, room_code := range room_codes {
	    room_code_sections := strings.Split(strings.Trim(room_code, "\n"), "-")
	    sector_checksum := strings.Split(room_code_sections[len(room_code_sections) - 1], "[")
            sector, err := strconv.Atoi(sector_checksum[0])
	    check(err)
	    checksum := strings.Trim(sector_checksum[1], "]")
	    names := room_code_sections[:len(room_code_sections)-1]

            letter_count := make(map[string]int)
	    for _, name := range names {
		    for _, letter := range name {
			    letter_str := string(letter)
			    value, prs := letter_count[letter_str]
			    if prs {
				    letter_count[letter_str] = value + 1
			    } else {
				    letter_count[letter_str] = 1
			    }
		    }
	    }
	    expected_checksum := top_5_letters(letter_count)
	    match := true
	    for i, letters := range checksum {
		    match = match && (string(letters) == expected_checksum[i])
	    }
	    if match {
		    decrypted_names := make(map[string]struct{})
		    for _, name := range names {
			    decrypt_name := shift_decode(name, sector)
			    decrypted_names[decrypt_name] = struct{}{}
		    }

		    _, found := decrypted_names["northpole"]
		    if found {
			    fmt.Println(sector)
		    }
	    }
    }
}

