package executor

import (
	"encoding/json"

	"github.com/moreira0102/zootoma/internal/core/action"
	"github.com/moreira0102/zootoma/internal/core/memory/manager"
)

//GetExecutor is an abstract struct that implements
//the Executor interface
type GetExecutor struct{}

//Implementation of the Executor Interface
func (ge GetExecutor) ExecuteAction(actn *action.Action) action.ActionResponse {

	if val, ok := manager.GblNodeManager.MemoryStorageMap[actn.Key]; ok {
		return action.ActionResponse{
			Method:  actn.Method,
			Status:  0,
			Data:    val.Data,
			Message: "Data retrived with success",
			Key:     val.Key,
			Size:    val.Size}
	}

	return action.ActionResponse{
		Method:  actn.Method,
		Status:  -1,
		Data:    nil,
		Message: "Key not found",
		Key:     actn.Key,
		Size:    -1}
}

func (ge GetExecutor) ParseActionResponse(ar *action.ActionResponse) (resp []byte) {
	if ar.Data != nil {
		return *ar.Data
	}
	resp, _ = json.Marshal(ar)
	return resp
}

func newGetExecutor() Executor {
	return GetExecutor{}
}
