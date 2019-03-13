package main

// FindOccurences takes an alphabet, a string length and a
// character to to detect. It returns how many occurances
// of that character are in the string
func FindOccurences(alphabet string, n int, char string) (result int, err error) {
	if len(char) != 1 {
		return 0, &ArgError{"char", "Only a single character can be provided"}
	}

	if n < 0 {
		return 0, &ArgError{"n", "A positive value of n is required"}
	}

	if len(alphabet) == 0 {
		return 0, &ArgError{"alphabet", "Need at least a string of length 1 to repeat"}
	}

	if len(alphabet) == 1 {
		if alphabet != char {
			return 0, nil
		} else if alphabet == char {
			return n, nil
		}
	}

	remainingLength := n % len(alphabet)
	charAsRune := []rune(char)[0]

	if n > len(alphabet) {
		charsInAlphabet := 0
		for _, rune := range alphabet {
			if rune == charAsRune {
				charsInAlphabet++
			}
		}

		result = charsInAlphabet * (n / len(alphabet))
	}

	for _, rune := range alphabet[0:remainingLength] {
		if rune == charAsRune {
			result++
		}
	}

	return result, nil
}
