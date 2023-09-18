package win

import (
	"syscall"
	"unsafe"
)

const (
	SwHide = 0
	SwShow = 1
)

type RECT struct {
	Left, Top, Right, Bottom int32
}

// GetHwndByTitle 根据title获取窗口句柄
func GetHwndByTitle(windowTitle string) uintptr {
	// 查找窗口句柄
	hwnd, _, _ := findWindow.Call(0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(windowTitle))))
	if hwnd != 0 {
		// 获取窗口标题
		var buffer [512]uint16
		getWindowText.Call(hwnd, uintptr(unsafe.Pointer(&buffer[0])), uintptr(len(buffer)))
		//title := syscall.UTF16ToString(buffer[:])
		_ = syscall.UTF16ToString(buffer[:])
		return hwnd
	} else {
		// 未找到窗口
		return hwnd
	}
}

// GetWindowRectByHandle 根据句柄获取窗口四点位坐标
func GetWindowRectByHandle(hwnd syscall.Handle) (RECT, error) {
	var rect RECT
	_, _, err := getWindowRect.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&rect)))
	if err != nil && err.Error() != "The operation completed successfully." {
		return rect, err
	}
	return rect, nil
}

// GetWindowSize 获取窗口大小
func GetWindowSize(hwnd uintptr) (int, int, error) {
	var rect RECT
	_, _, err := getWindowRect.Call(hwnd, uintptr(unsafe.Pointer(&rect)))
	if err != nil && err.Error() != "The operation completed successfully." {
		return 0, 0, err
	}
	width := int(rect.Right - rect.Left)
	height := int(rect.Bottom - rect.Top)
	return width, height, nil
}

func HideWindow(hwnd uintptr) bool {
	if hwnd == 0 {
		return false
	}

	//setForegroundWin.Call(hwnd)
	showWindow.Call(hwnd, SwHide)
	return true
}

func ShowWindow(hwnd uintptr) bool {
	if hwnd == 0 {
		return false
	}

	//setForegroundWin.Call(hwnd)
	showWindow.Call(hwnd, SwShow)
	return true
}

// SetTopWindow 窗口置顶
func SetTopWindow(hwnd uintptr) {
	setForegroundWin.Call(hwnd)
}

func SetWindowSize(hwnd uintptr, width, height int) {
	setWindowPos.Call(
		uintptr(hwnd), 0, 0, 0, uintptr(width), uintptr(height), 0,
	)
}

const (
	GWL_EXSTYLE      = 0xFFFFFFEC
	WS_EX_TOOLWINDOW = 0x00000080
	WS_EX_LAYERED    = 0x80000

	LWA_COLORKEY = 0x1
	LWA_ALPHA    = 0x2
)

func SetWindowAlpha(hwnd uintptr, alpha int) error {

	style, _, _ := getWindowLong.Call(uintptr(hwnd), GWL_EXSTYLE)
	setWindowLong.Call(uintptr(hwnd), GWL_EXSTYLE, style|WS_EX_LAYERED)
	setLayeredWindowAttributes.Call(uintptr(hwnd), 0, uintptr(alpha), LWA_ALPHA)
	_, _, err := setLayeredWindowAttributes.Call(uintptr(hwnd), 0, uintptr(alpha), LWA_ALPHA)
	if err != nil && err.Error() != "The operation completed successfully." {
		return err
	}
	return nil
}
