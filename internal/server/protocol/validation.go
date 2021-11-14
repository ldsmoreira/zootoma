package protocol

import "strconv"

// IsValidMethod validades if the method passed
// is a valid one
func IsValidMethod(method []byte) bool {
	ms := string(method)
	if _, present := methods[ms]; present {
		return true
	} else {
		return false
	}
}

// IsValidKey validades if the key passed
// is a valid one
// In the moment of it has been written every key was allowed
func IsValidKey(method []byte) bool {
	return true
}

// IsValidSize validades if the size bytes passed
// are integer convertible and if it is less than
// the maximum allowed size for data
func IsValidSize(size []byte) bool {
	ds, err := strconv.Atoi(string(size))
	if err != nil && ds <= MaxDataSize {
		return true
	} else {
		return false
	}
}
