package win

import (
	"fmt"
	"github.com/lxn/win"
	"syscall"
	"time"
)

const (
	MOUSEEVENTF_MOVE       = 0x0001
	MOUSEEVENTF_ABSOLUTE   = 0x8000
	MOUSEEVENTF_LEFTDOWN   = 0x0002
	MOUSEEVENTF_LEFTUP     = 0x0004
	MOUSEEVENTF_MIDDLEDOWN = 0x0020
	MOUSEEVENTF_MIDDLEUP   = 0x0040
	MOUSEEVENTF_RIGHTDOWN  = 0x0008
	MOUSEEVENTF_RIGHTUP    = 0x0010
)
const (
	WindowWidth  = 450
	WindowHeight = 844
)

func MockClick(hwndAddress uintptr, x, y int) {
	// 替换为你要点击的窗口句柄
	hwnd := syscall.Handle(hwndAddress)

	// 获取窗口的位置和大小
	rect, err := GetWindowRectByHandle(hwnd)
	if err != nil {
		fmt.Printf("获取窗口位置和大小失败：%v\n", err)
		return
	}

	width, height, err := GetWindowSize(hwndAddress)
	if err != nil {
		fmt.Printf("获取窗口大小失败：%v\n", err)
	} else {
		fmt.Printf("窗口大小：%d x %d\n", width, height)
	}

	// 计算要点击的位置
	clickX := rect.Left + int32(float64(width)*float64(x)/WindowWidth)
	clickY := rect.Top + int32(float64(height)*float64(y)/WindowHeight)

	// 移动鼠标到指定位置
	_, _, _ = setCursorPos.Call(uintptr(clickX), uintptr(clickY))

	// 模拟鼠标点击
	_, _, _ = mouseEvent.Call(
		uintptr(MOUSEEVENTF_LEFTDOWN),
		0,
		0,
		0,
		0,
	)
	_, _, _ = mouseEvent.Call(
		uintptr(MOUSEEVENTF_LEFTUP),
		0,
		0,
		0,
		0,
	)

	fmt.Printf("模拟点击窗口 %v 的位置：%d, %d\n", hwnd, clickX, clickY)
}

func PerformBackgroundClick(hwnd win.HWND, x, y int, cnt int) {

	lParam := win.MAKELONG(uint16(x), uint16(y))
	count := 0
	for count < cnt {
		win.SendMessage(hwnd, win.WM_LBUTTONDOWN, win.MK_LBUTTON, uintptr(lParam))
		win.SendMessage(hwnd, win.WM_LBUTTONUP, 0, uintptr(lParam))
		count++
		time.Sleep(200 * time.Millisecond)
	}
}
