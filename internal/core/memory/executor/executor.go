package executor

import (
	"fmt"
	"zootoma/internal/core/action"
)

type Executor interface {
	ExecuteAction(action *action.Action) action.ActionResponse
}

var ExecutorMap map[string]Executor = map[string]Executor{
	"set": newSetExecutor(),
	"get": newGetExecutor(),
}

func Execute(actn *action.Action) action.ActionResponse{
	exec:=ExecutorMap[actn.Method].ExecuteAction(actn)
	fmt.Print(exec)
	return action.ActionResponse{Status: exec.Status, Method: exec.Method, Data: exec.Data, Message: exec.Message}
}
