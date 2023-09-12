package job

import (
	"fmt"
	"github.com/elpsyr/saltfish/pkg/win"
	lxnWin "github.com/lxn/win"
	"time"
)

// GetFish 钓鱼
func (m *Manager) GetFish(hwnd uintptr) {
	instance := m.GetInstance()
	// instance.mu.Lock() 是阻塞操作，这里使用 TryLock进行判断
	tryLock := instance.mu.TryLock()
	if !tryLock {
		fmt.Println("Failed to acquire lock. Exiting...")
		// TODO:
		// 执行适当的退出操作，比如弹窗通知
		// 或者直接 return，结束当前函数或方法的执行
		return
	}
	defer instance.mu.Unlock()
	fmt.Println("Acquire lock. fishing...")
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
