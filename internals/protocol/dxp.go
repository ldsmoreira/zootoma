package dxp

import (
	"bytes"
	"fmt"
	"strconv"
)

type DxpProtocolMapping struct {
	Method    []byte
	Path      []byte
	Data_size int
	Data      []byte
}

func byteSliceToInt(slice []byte) int {
	data, err := strconv.Atoi(string(slice))
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func NewDxpProtocolMapping(method []byte, path []byte, data_size []byte) *DxpProtocolMapping {
	dxp := new(DxpProtocolMapping)
	dxp.Method = method
	dxp.Path = path

	parsed_data_size := bytes.Split(data_size, []byte("::"))[1]
	parsed_data_size = parsed_data_size[:bytes.Index(parsed_data_size, []byte{0})]
	dxp.Data_size = byteSliceToInt(parsed_data_size)
	return dxp
}
