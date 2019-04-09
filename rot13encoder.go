package main

import "io"

func rot13(b byte) byte {
	var a, z byte
	switch {
	case 'a' <= b && b <= 'z':
		a, z = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}
	return (b-a+13)%(z-a+1) + a
}

// Message holds a message and various encoders that can
// apply to the message.
type Message string

// NewRot13Encoder returns a new Rot13Encoder that implements
// io.Reader
func (m Message) NewRot13Encoder() Rot13Encoder {
	return Rot13Encoder{message: m, offset: 0}
}

// Rot13Encoder takes in a string and implements io.Reader
// as a way to read the encoded rot13 string out
type Rot13Encoder struct {
	message Message
	offset  int64
}

// Read fills the passed in byte slice and returns the number
// of bytes filled or any errors that happened
func (r *Rot13Encoder) Read(b []byte) (n int, err error) {
	if r.offset >= int64(len(r.message)) {
		return 0, io.EOF
	}

	n = copy(b, r.message[r.offset:])
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	r.offset += int64(n)
	return
}
