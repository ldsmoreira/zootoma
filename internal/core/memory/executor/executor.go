package executor

import (
	"zootoma/internal/core/action"
)

//Executor interface is part of Strategy Desing pattern
//to abstract how each method acts in memory related
//workloads
type Executor interface {
	ExecuteAction(action *action.Action) action.ActionResponse
}

//ExecutorMap is a global map relating methods to its executor constructors
var ExecutorMap map[string]func() Executor = map[string]func() Executor{

	"set": newSetExecutor,
	"get": newGetExecutor,
}

//Execute a given action based on the operation that method represents
func Execute(actn *action.Action) (ar action.ActionResponse) {

	executor := ExecutorMap[actn.Method]()
	ar = executor.ExecuteAction(actn)

	return action.ActionResponse{
		Status:  ar.Status,
		Method:  ar.Method,
		Data:    ar.Data,
		Message: ar.Message}
}
