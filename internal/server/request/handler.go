package request

import (
	action "github.com/moreira0102/zootoma/internal/core/action"
)

type Handler struct {
	Parser Parser
}

func NewHandler() (handler Handler) {

	parser := Parser{
		Request: new(Request),
		Action:  new(action.Action),
	}

	handler = Handler{
		Parser: parser}

	handler.Parser.Action.Headers = make(map[string][]byte)

	return handler
}
