package job

import "sync"

type Manager struct {
	mu sync.Mutex
}
