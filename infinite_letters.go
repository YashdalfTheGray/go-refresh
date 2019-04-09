package main

// InfiniteLetters implements io.Reader and hands out
// an infinite stream of the same letter
type InfiniteLetters struct {
	letter rune
}

// Read will fill the given slice with the letter that
// InfiniteLetters was initialized with and returns
// either an error or how many bytes were filled
func (il InfiniteLetters) Read(b []byte) (n int, err error) {
	numBytes := len(b)

	for i := 0; i < numBytes; i++ {
		b[i] = byte(il.letter)
	}

	return numBytes, nil
}
