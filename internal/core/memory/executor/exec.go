package executor

import (
	"zootoma/internal/core/action"
)

type Executor interface {
	ExecuteAction(action *action.Action) action.ActionResponse
}
