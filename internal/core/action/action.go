package action

// The request module implements the Request structure and it's methods

// Example of a valid request:

// set or get
//
// status::ok
// host::com.toma
//
// (every data that fits 30000 bytes)

type Action struct {
	Method   string
	Key      string
	DataSize int
	Data     *[]byte
	Headers  map[string][]byte
}

type ActionResponse struct {
	Method  string
	Status  int
	Data    *[]byte
	Message string
}
