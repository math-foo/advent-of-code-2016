package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)


func check(e error) {
    if e != nil {
      panic(e)
    }
}

func find_abba(ip_portion string) bool {
	max_start := len(ip_portion) - 3
	abba_found := false

	for i := 0; i < max_start; i++ {
		a1 := ip_portion[i]
		a2 := ip_portion[i+3]
		if a1 == a2 {
			b1 := ip_portion[i+1]
			b2 := ip_portion[i+2]
			if b1 == b2 && b1 != a1 {
				abba_found = true
				break
			}
		}
	}

	return abba_found
}

func main() {
    data, err := ioutil.ReadFile("input.txt")
    check(err)
    ips := strings.Split(strings.Trim(string(data), "\n"), "\n")
    transport_enabled := 0
    for _, ip_address := range ips {
	    start_rest := strings.Split(ip_address, "[")
	    start := start_rest[0]
	    rest := start_rest[1:]
	    abba_found := find_abba(start)
	    hypertext_abba_found := false
	    for _, ip_portion := range rest {
		    hypertext_section := strings.Split(ip_portion, "]")
		    section := hypertext_section[1]
		    abba_found = abba_found || find_abba(section)
		    hypertext := hypertext_section[0]
		    hypertext_abba_found = hypertext_abba_found || find_abba(hypertext)
	    }
	    if abba_found && !hypertext_abba_found {
		    transport_enabled++
	    }
    }
    fmt.Println(transport_enabled)
}

