package job

import "github.com/elpsyr/saltfish/pkg/win"

func (m *Manager) ShowMode() {

	hwnd := m.GetHwnd()
	if hwnd == 0 {
		return
	}
	win.ShowWindow(hwnd)
	win.SetTopWindow(hwnd)
}

func (m *Manager) HideMode() {
	hwnd := m.GetHwnd()
	if hwnd == 0 {
		return
	}
	win.SetTopWindow(hwnd)
	win.HideWindow(hwnd)
}
