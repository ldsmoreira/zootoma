package executor

import (
	"zootoma/internal/core/action"
	"zootoma/internal/core/memory/manager"
)

type GetExecutor struct{}

func (ge GetExecutor) ExecuteAction(actn *action.Action) action.ActionResponse {
	if val, ok := manager.GblNodeManager.MemoryStorageMap[actn.Key]; ok {
		return action.ActionResponse{Method: actn.Method, Status: 0, Data: val.Data, Message: "receba", Key: val.Key, Size: val.Size}
	}
	return action.ActionResponse{Status: -1}
}

func newGetExecutor() GetExecutor {
	return GetExecutor{}
}
