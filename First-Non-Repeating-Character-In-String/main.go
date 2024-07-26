package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	// a := FirstNonRepeatingCharacter("sss1")
	str := "gggggggggggggggggggghfffffffffffffffffffff"
	FirstNonRepeatingCharacter(str)
	
	
	
}

func FirstNonRepeatingCharacter(str string) {

	var find = []string{}

	spltstring := strings.Split(str, "")
	// fmt.Println(q)
	count := 0
	for key, ch := range spltstring {
		if key == 0 {
			find = append(find, ch)
		} else {
			if slices.Contains(find, ch) {
				continue
			} else {
				count ++
			}
		}
		if count == 1 {
			fmt.Println(ch)
		}
	}
}