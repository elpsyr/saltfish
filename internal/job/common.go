package job

import (
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
