package protocol

import (
	"bytes"
	"fmt"
)

// IsValidMethod validades if the method passed
// is a valid one
func IsValidMethod(method []byte) bool {
	ms := string(method)
	if _, present := METHODS[ms]; present {
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
func IsValidSize(ds int) bool {
	if ds <= MAX_DATA_SIZE {
		return true
	} else {
		return false
	}
}

// IsValidMinorHeader validades if the Meta Header was constructed
// following the protocol needs
func IsValidMetaHeader(mh []byte) (key []byte, value []byte, valid bool) {
	fmt.Println(string(mh))
	smh := bytes.SplitN(mh, META_HEADER_SEPARATOR, 2)
	fmt.Println(smh)
	if len(smh[0]) == 0 {
		return nil, nil, false
	} else if len(smh[0]) == 1 && smh[0][0] == STATEMENTS_DELIMITER {

		return nil, nil, true

	} else {
		key = smh[0]
		value = smh[1]
		return key, value, true
	}
}
