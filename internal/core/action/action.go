package action

// The request module implements the Request structure and it's methods

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
	Key     string
	Size    int
}
