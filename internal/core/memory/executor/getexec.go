package executor

import (
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
		Status:  0,
		Data:    nil,
		Message: "Key not found",
		Key:     actn.Key,
		Size:    actn.DataSize}
}

func newGetExecutor() Executor {
	return GetExecutor{}
}
