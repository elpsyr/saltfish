package job

import (
	"github.com/elpsyr/saltfish/pkg/win"
	lxnWin "github.com/lxn/win"
	"time"
)

// GetFish 钓鱼
func (m Manager) GetFish(hwnd uintptr) {
	//if !m.mu.TryLock() {
	//	fmt.Println("Unable to acquire lock")
	//	return
	//}
	m.mu.Lock()
	defer m.mu.Unlock()
	// reset
	m.BackToHome(hwnd)
	// do job
	win.PerformBackgroundClick(lxnWin.HWND(hwnd), 400, 667, 1)
	time.Sleep(time.Second * 2)
	win.PerformBackgroundClick(lxnWin.HWND(hwnd), 30, 380, 1)
	time.Sleep(time.Second * 2)
	win.PerformBackgroundClick(lxnWin.HWND(hwnd), 130, 680, 1)
	time.Sleep(time.Second * 5)

	// reset
	win.PerformBackgroundClick(lxnWin.HWND(hwnd), 45, 750, 1)
	time.Sleep(time.Second * 3)
	win.PerformBackgroundClick(lxnWin.HWND(hwnd), 45, 750, 1)
	time.Sleep(time.Second * 3)
	win.PerformBackgroundClick(lxnWin.HWND(hwnd), 45, 750, 1)
}
