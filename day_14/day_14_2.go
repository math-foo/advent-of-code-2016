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

func md5_2017_times(value string) string{
	for i :=0; i < 2017; i++ {
		value = fmt.Sprintf("%x", md5.Sum([]byte(value)))
	}

	return value
}

func triple_index(hash string) int {
	index := -1
	for i := 2; i < len(hash); i++ {
		if (hash[i] == hash[i-1]) && (hash[i] == hash[i-2]) {
			index = i
			break
		}
	}

	return index
}

func quint_index(hash string, value byte) int {
	index := -1
	for i := 4; i < len(hash); i++ {
		if (hash[i] == value && hash[i] == hash[i-1]) && (hash[i] == hash[i-2] &&
	            hash[i] == hash[i-3] && hash[i] == hash[i-4]) {
			index = i
			break
		}
	}

	return index
}

func get_md5_hash(x int, hash_dict map[string]string) string {
	salt := "ngcjuoqr"
	str_to_hash := fmt.Sprintf("%s%d", salt, x)

	value, pres := hash_dict[str_to_hash]
	if pres {
		return value
	} else {
		//hashed_value := fmt.Sprintf("%x", md5.Sum([]byte(str_to_hash)))
		hashed_value := md5_2017_times(str_to_hash)
		hash_dict[str_to_hash] = hashed_value
		return hashed_value
	}
}

func main() {
	hash_dict := make(map[string]string)
	keys_found := 0
	found_64th_key_index := -1
	current_index := 1
	for found_64th_key_index < 0 {
		hash := get_md5_hash(current_index, hash_dict)
		new_triple_index := triple_index(hash)
		if new_triple_index > 0 {
			value := hash[new_triple_index]
			found_quint_index := false
			for j := current_index + 1; j < current_index + 1000; j++ {
				quint_hash := get_md5_hash(j, hash_dict)
				new_quint_index := quint_index(quint_hash, value)
				if new_quint_index > 0 {
					found_quint_index = true
					break
				}
			}
			if found_quint_index {
				keys_found++
				if keys_found == 64 {
					found_64th_key_index = current_index
				}
			}
		}
		current_index++
	}

	fmt.Println(found_64th_key_index)
}

