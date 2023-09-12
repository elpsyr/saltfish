package job

import (
	"fmt"
	"github.com/elpsyr/saltfish/pkg/win"
	lxnWin "github.com/lxn/win"
	"time"
)

func (m *Manager) GetReward(hwnd uintptr) {
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
	fmt.Println("Acquire lock. get reward...")
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
