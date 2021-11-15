package request

// The request module implements the Request structure and it's methods

// Example of a valid request:

// set /home/lucas/data.txt 30000
//
// status::ok
// host::com.toma
//
// (every data that fits 30000 bytes)

type Request struct {
	Method   string
	Key      string
	DataSize int
	Headers  []byte
}
