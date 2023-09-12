package job

import (
	"github.com/elpsyr/saltfish/pkg/win"
	lxnWin "github.com/lxn/win"
	"time"
)

func (m Manager) GetReward(hwnd uintptr) {
	m.mu.Lock()
	//if !m.mu.TryLock() {
	//	fmt.Println("Unable to acquire lock")
	//	return
	//}
	defer m.mu.Unlock()
	// reset
	m.BackToHome(hwnd)
	// do job
	win.PerformBackgroundClick(lxnWin.HWND(hwnd), 40, 350, 1)
	win.PerformBackgroundClick(lxnWin.HWND(hwnd), 310, 675, 5)
	win.PerformBackgroundClick(lxnWin.HWND(hwnd), 150, 675, 5)
	// reset
	m.BackToHome(hwnd)
	time.Sleep(time.Second * 3)

}
