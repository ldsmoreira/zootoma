package protocol

var (
	MAIN_HEADER_ITEMS          int    = 3 // MAIN_HEADER_ITEMS stands for "Main Header Components Quantity"
	METHODS_QTT                int    = 2
	MAX_DATA_SIZE              int    = 10e6
	METHOD_SET                 string = "set"
	METHOD_GET                 string = "get"
	MAIN_HEADER_SEPARATOR      []byte = []byte(" ")
	META_HEADER_SEPARATOR      []byte = []byte("::")
	STATEMENTS_DELIMITER       byte   = byte('\n')
	MAIN_HEADER_POSITION       int    = 0
	META_HEADER_BLOCK_POSITION int    = 1
	DATA_BLOCK_POSITION        int    = 2

	STD_PARSER_ORDER [3]int = [3]int{
		MAIN_HEADER_POSITION,
		META_HEADER_BLOCK_POSITION,
		DATA_BLOCK_POSITION}
	// Private mapping of possible methods
	METHODS = map[string]bool{
		"set": true,
		"get": true,
	}
)
