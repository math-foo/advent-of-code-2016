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

func find_abas(ip_portion string, dict map[string]struct{}) {
	max_start := len(ip_portion) - 2
	for i := 0; i < max_start; i++ {
		a1 := ip_portion[i]
		a2 := ip_portion[i+2]
		if a1 == a2 {
			b := ip_portion[i+1]
			if b != a1 {
				aba := string(ip_portion[i:i+3])
				dict[aba] = struct{}{}
			}
		}
	}
}

func main() {
    data, err := ioutil.ReadFile("input.txt")
    check(err)
    ips := strings.Split(strings.Trim(string(data), "\n"), "\n")
    ssl_enabled := 0
    for _, ip_address := range ips {
	    hypertext_dict := make(map[string]struct{})
	    dict := make(map[string]struct{})
	    start_rest := strings.Split(ip_address, "[")
	    start := start_rest[0]
	    rest := start_rest[1:]
	    find_abas(start, dict)
	    for _, ip_portion := range rest {
		    hypertext_section := strings.Split(ip_portion, "]")
		    section := hypertext_section[1]
		    find_abas(section, dict)
		    hypertext := hypertext_section[0]
		    find_abas(hypertext, hypertext_dict)
	    }
	    for aba := range dict {
		    bab := fmt.Sprintf("%s%s",aba[1:2],aba[0:2])
		    _, pres := hypertext_dict[bab]
		    if pres {
			    ssl_enabled++
			    break
		    }
	    }
    }
    fmt.Println(ssl_enabled)
}

