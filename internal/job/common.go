package job

import (
	"github.com/elpsyr/saltfish/pkg/win"
	lxnWin "github.com/lxn/win"
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
