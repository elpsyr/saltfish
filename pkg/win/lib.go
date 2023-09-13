package win

import "syscall"

var (
	user32DLL                = syscall.NewLazyDLL("user32.dll")
	findWindow               = user32DLL.NewProc("FindWindowW")
	getWindowText            = user32DLL.NewProc("GetWindowTextW")
	getWindowThreadProcessID = user32DLL.NewProc("GetWindowThreadProcessId")
	getWindowRect            = user32DLL.NewProc("GetWindowRect")
	setWindowPos             = user32DLL.NewProc("SetWindowPos")

	// 点击
	setCursorPos = user32DLL.NewProc("SetCursorPos")
	mouseEvent   = user32DLL.NewProc("mouse_event")
)

var (
	setForegroundWin    = user32DLL.NewProc("SetForegroundWindow")
	showWindow          = user32DLL.NewProc("ShowWindow")
	setForegroundWindow = user32DLL.NewProc("SetForegroundWindow")
)
