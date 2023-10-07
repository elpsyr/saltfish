package job

import (
	"errors"
	"fmt"
	"github.com/elpsyr/saltfish/pkg/win"
	lxnWin "github.com/lxn/win"
	"os/exec"
	"time"
)

func (m *Manager) BackToHome(hwnd uintptr) {
	win.PerformBackgroundClick(lxnWin.HWND(hwnd), 233, 817, 1)
}

func (m *Manager) ResizeWindow() {
	hwnd := m.GetHwnd()
	if hwnd == 0 {
		return
	}
	win.SetWindowSize(hwnd, 450, 844)
}

func (m *Manager) AlphaWindow(alpha int) {
	hwnd := m.GetHwnd()
	if hwnd == 0 {
		return
	}
	win.SetWindowAlpha(hwnd, alpha)
}

func (m *Manager) Restart() {

	hwnd := m.GetHwnd()

	win.CloseWindow(hwnd)
	// 定义要执行的命令和参数
	cmd := exec.Command("C:\\Program Files (x86)\\Tencent\\WeChat\\WechatAppLauncher.exe", "-launch_appid=wx0840558555a454ed")

	// 执行命令
	_, err := cmd.Output()

	if err != nil {
		fmt.Println("命令执行失败:", err)
		return
	}
	time.Sleep(3 * time.Second)
	hwndRestart := m.GetHwnd()
	if hwndRestart == 0 {
		return
	}
	win.HideWindow(hwndRestart)
	for i := 0; i < 10; i++ {
		win.PerformBackgroundClick(lxnWin.HWND(hwndRestart), 230, 800, 1)
		time.Sleep(time.Second)
	}
}

// CheckSaltfishAppAlived 检查咸鱼之王小程序是否打开
//
//	会先执行点击屏幕中央的操作,检查是否被顶号,此时若已经顶号则会自动关闭小程序,并返回 false
//	点击后若句柄仍可被获取,则表示小程序存活,且没被顶号,此时返回 true
//	若返回的 error 为非空,则表示函数执行错误
func (m *Manager) CheckSaltfishAppAlived() (bool, error) {
	instance := m.GetInstance()
	tryLock := instance.mu.TryLock()
	if !tryLock {
		fmt.Println("Failed to acquire lock. Exiting...")
		return false, errors.New("Failed to acquire lock. Exiting...")
	}
	defer instance.mu.Unlock()
	fmt.Println("Acquire lock. get reward...")

	// 获取句柄
	hwnd := m.GetHwnd()
	if hwnd == 0 { // 获取失败，小程序已被关闭
		return false, nil
	}

	// 尝试点击屏幕中间部位,试探是否被顶号
	win.PerformBackgroundClick(lxnWin.HWND(hwnd), 230, 540, 3)
	time.Sleep(time.Second * 3)

	// 休眠后再次尝试获取句柄
	hwnd = m.GetHwnd()
	if hwnd == 0 { // 获取失败，小程序已被关闭
		return false, nil
	}

	// 程序存活
	return true, nil
}
