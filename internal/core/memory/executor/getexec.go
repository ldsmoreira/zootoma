package executor

import "zootoma/internal/core/action"

type GetExecutor struct{}

func (_ GetExecutor) ExecuteAction(action *action.Action) action.ActionResponse {

}

func newGetExecutor() GetExecutor {
	return GetExecutor{}
}
