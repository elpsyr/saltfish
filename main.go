package main

import (
	"fmt"
	"github.com/elpsyr/saltfish/pkg/win"
)

func main() {
	windowTitle := "咸鱼之王"

	hwnd := win.GetHwndByTitle(windowTitle)

	processID := win.GetProcessIDByHwnd(hwnd)

	win.GetProgressMemInfo(processID)

	// 通过窗口句柄获取窗口大小
	width, height, err := win.GetWindowSize(hwnd)
	if err != nil {
		fmt.Printf("获取窗口大小失败：%v\n", err)
	} else {
		fmt.Printf("窗口大小：%d x %d\n", width, height)
	}

	// 40 350
	win.MockClick(hwnd, 40, 350)

	//win.HideWindow(hwnd)
	win.ShowWindow(hwnd)
	win.SetTopWindow(hwnd)

}
