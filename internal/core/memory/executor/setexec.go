package executor

import (
	"zootoma/internal/core/action"
	"zootoma/internal/core/memory/manager"
	"zootoma/internal/core/memory/memorydata"
)

//GetExecutor is an abstract struct that implements
//the Executor interface
type SetExecutor struct{}

//Implementation of the Executor Interface
func (se SetExecutor) ExecuteAction(actn *action.Action) action.ActionResponse {

	memoryData := memorydata.MemoryData{
		Key:  actn.Key,
		Data: actn.Data,
		Size: actn.DataSize}

	manager.GblNodeManager.MemoryStorageMap[actn.Key] = &memoryData

	return action.ActionResponse{
		Status:  1000,
		Method:  actn.Method,
		Data:    nil,
		Message: "Data set with success"}
}

func newSetExecutor() Executor {
	return SetExecutor{}
}
