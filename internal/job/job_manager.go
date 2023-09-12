package job

import (
	"sync"
)

type Manager struct {
}

var lock = &sync.Mutex{}

type Single struct {
	mu sync.Mutex
}

var singleInstance *Single

// GetInstance 双重校验获取单例
func (*Manager) GetInstance() *Single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &Single{}
		} else {
		}
	} else {
	}
	return singleInstance
}
