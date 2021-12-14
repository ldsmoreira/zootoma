package executor

import (
	"encoding/json"

	"github.com/moreira0102/zootoma/internal/core/action"
	"github.com/moreira0102/zootoma/internal/core/memory/manager"
	"github.com/moreira0102/zootoma/internal/core/memory/memorydata"
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
		Status:  0,
		Method:  actn.Method,
		Data:    nil,
		Message: "Data set with success"}
}

func (se SetExecutor) ParseActionResponse(ar *action.ActionResponse) (resp []byte) {
	resp, _ = json.Marshal(ar)
	return resp
}

func newSetExecutor() Executor {
	return SetExecutor{}
}
