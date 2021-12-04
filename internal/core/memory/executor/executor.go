package executor

import (
	"zootoma/internal/core/action"
)

type Executor interface {
	ExecuteAction(action *action.Action) action.ActionResponse
}

var ExecutorMap map[string]interface{} = map[string]interface{}{
	"set": newSetExecutor,
	"get": newGetExecutor,
}
