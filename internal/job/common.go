package job

import (
	"github.com/elpsyr/saltfish/pkg/win"
	lxnWin "github.com/lxn/win"
)

func (m *Manager) BackToHome(hwnd uintptr) {
	win.PerformBackgroundClick(lxnWin.HWND(hwnd), 233, 817, 1)
}
