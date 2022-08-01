package main

import (
	"fmt"
	"strings"
)

func checkDuplicate(a string) bool {
	a = strings.ToLower(a)
	save := make(map[string]bool)

	unique := true
	for i := range a {
		if _, ok := save[string(a[i])]; ok {
			unique = false
			return unique
		} else {
			unique = true
			save[string(a[i])] = true
		}
	}

	return unique
}

func main() {
	a := []string{"safget", "ssttFFddggff", "DfsaweDffF", "fasfasrttetqg", "abCdefAaf", "abcd", "aabcd"}

	for i := range a {
		fmt.Printf("%s is %v\n", a[i], checkDuplicate(a[i]))
	}
}
