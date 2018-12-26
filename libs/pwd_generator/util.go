package pwd_generator

import "time"

// slice initialize
// = make([]T, len, cap)
var allSets = make([]byte, 0, 100)

// generate byte array from string
// = []byte("string")
// combine bytes into string
// = string([]byte)
var lowerAlphabetSets = []byte("abcdefghijklmnopqrstuvwxyz")
var upperAlphabetSets = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numbericSets = []byte("0123456789")
var superSets = []byte("!@#$*&")

//
type RandomResult struct {
	Len       int
	data      string
	timestamp time.Time
}

func (result *RandomResult) Data() string{
	return result.data
}
func (result *RandomResult) Timestamp() string{
	return result.timestamp.Format("2006-01-02 15:04:05")
}
func (result *RandomResult) Error() string{
	return ""
}

func getDestSet(size int, data []byte) *[]byte {
	dest := make([]byte, 0, size)
	if size > len(data) {
		dest = data
	} else {
		dest = data[:size]
	}
	return &dest
}
