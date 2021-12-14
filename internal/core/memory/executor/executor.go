package executor

import (
	"github.com/moreira0102/zootoma/internal/core/action"
)

//Executor interface is part of Strategy Desing pattern
//to abstract how each method acts in memory related
//workloads
type Executor interface {
	ExecuteAction(action *action.Action) action.ActionResponse
	ParseActionResponse(actionresp *action.ActionResponse) (resp []byte)
}

//ExecutorMap is a global map relating methods to its executor constructors
var ExecutorMap map[string]func() Executor = map[string]func() Executor{

	"set": newSetExecutor,
	"get": newGetExecutor,
}

//Execute a given action based on the operation that method represents
func Execute(actn *action.Action) (ar action.ActionResponse, resp []byte) {

	executor := ExecutorMap[actn.Method]()
	ar = executor.ExecuteAction(actn)
	resp = executor.ParseActionResponse(&ar)

	return ar, resp
}
