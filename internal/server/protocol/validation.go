package protocol

import (
	"bytes"
	"fmt"
	"strconv"
)

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
func IsValidKey(key []byte) bool {
	return true
}

// IsValidSize validades if the size bytes passed
// are integer convertible and if it is less than
// the maximum allowed size for data
func IsValidSize(size []byte) bool {
	ds, err := strconv.Atoi(string(size))
	fmt.Println(ds <= MaxDataSize)
	fmt.Println(err)
	if err == nil && ds <= MaxDataSize {
		return true
	} else {
		return false
	}
}

// IsValidMinorHeader validades if the Meta Header was constructed
// following the protocol needs
func IsValidMetaHeader(mh []byte) (key []byte, value []byte, valid bool) {
	smh := bytes.SplitN(mh, MetaHeaderSeparator, 2)
	if len(smh) != 2 {
		return nil, nil, false
	} else {
		return key, value, true
	}
}
