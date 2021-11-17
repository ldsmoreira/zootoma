package protocol

var (
	// MAIN_HEADER_ITEMS stands for "Main Header Components Quantity"
	MAIN_HEADER_ITEMS int = 3

	// METHODS_QTT stands for Methods Quantity. It means that at the
	// present time, there are just METHODS_QTT methods available
	METHODS_QTT int = 2

	MAX_DATA_SIZE int = 10e6

	METHOD_SET string = "set"

	METHOD_GET string = "get"

	MAIN_HEADER_SEPARATOR []byte = []byte(" ")

	META_HEADER_SEPARATOR []byte = []byte("::")

	STATEMENTS_DELIMITER byte = byte('\n')

	// Private mapping of possible methods
	METHODS = map[string]bool{
		"set": true,
		"get": true,
	}
)
