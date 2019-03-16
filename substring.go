package main

import (
	"strings"
)

func buildMapOfWords(str string) (result map[string]int) {
	result = make(map[string]int)
	words := strings.Fields(str)

	for i := 0; i < len(words); i++ {
		if _, ok := result[words[i]]; !ok {
			result[words[i]] = 1
		} else {
			result[words[i]]++
		}
	}

	return result
}

// SubstringWithMaps splits a given string into words
// and checks if the given substring can be made using
// the word pieces
func SubstringWithMaps(s, sub string) (result bool, err error) {
	sMap := buildMapOfWords(s)
	subMap := buildMapOfWords(sub)

	if len(sMap) == 0 {
		return false, ArgError{"s", "s needs to have at least 1 word separated by spaces"}
	}

	if len(subMap) == 0 {
		return false, ArgError{"sub", "sub needs to have at least 1 word separated by spaces"}
	}

	if len(subMap) > len(s) {
		return false, ArgError{"s", "s needs to be larger than sub"}
	}

	result = true

	for k, v := range subMap {
		if _, ok := sMap[k]; !ok {
			return false, nil
		}

		result = result && sMap[k] >= v
	}

	return result, nil
}
