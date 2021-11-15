package protocol

const (
	// MainHeaderCompQtt stands for "Main Header Components Quantity"
	MainHeaderCompQtt int = 3

	// MethodsQtt stands for Methods Quantity. It means that at the
	// present time, there are just MethodsQtt methods available
	MethodsQtt int = 2

	MaxDataSize int = 10e6

	MethodSet string = "set"

	MethodGet string = "get"
)

var (
	MainHeaderSeparator []byte = []byte(" ")

	MetaHeaderSeparator []byte = []byte("::")

	// Private mapping of possible methods
	methods = map[string]bool{
		"set": true,
		"get": true,
	}
)
