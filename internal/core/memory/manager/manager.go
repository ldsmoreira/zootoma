package manager

import (
	"sync"
	"zootoma/internal/core/memory/memorydata"
)

type NodeManager struct {
	MemoryStorageMap map[string]*memorydata.MemoryData
}

var onlyOnceCreation sync.Once

//newManager function Creates a NodeManager pointer
//and initialize the object it points to.
//Since there is just ONE nodeManager per zootoma
//proccess, this function just have practical effects
//in it's first call
func newManager() *NodeManager {
	var nm *NodeManager
	onlyOnceCreation.Do(func() {
		nm = new(NodeManager)
		nm.MemoryStorageMap = make(map[string]*memorydata.MemoryData)
	})
	if nm != nil {
		return nm
	} else {
		return nil
	}
}

var GblNodeManager *NodeManager = newManager()
