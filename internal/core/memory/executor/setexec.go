package executor

import (
	"zootoma/internal/core/action"
	"zootoma/internal/core/memory/manager"
	"zootoma/internal/core/memory/memorydata"
)

type SetExecutor struct{
}

func (se SetExecutor) ExecuteAction(actn *action.Action) action.ActionResponse {
	memoryData := memorydata.MemoryData{Key: actn.Key, Data: actn.Data, Size: actn.DataSize}
	manager.GblNodeManager.MemoryStorageMap[actn.Key] = &memoryData
	return action.ActionResponse{Status: 1000, Method: actn.Method, Data: actn.Data, Message: actn.Key}
}

func newSetExecutor() SetExecutor {
	return SetExecutor{}
}
