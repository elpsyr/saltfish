package win

import (
	"fmt"
	"github.com/mitchellh/go-ps"
	"github.com/shirou/gopsutil/process"
	"unsafe"
)

// GetProcessIDByHwnd 根据句柄获取进程id
func GetProcessIDByHwnd(hwnd uintptr) uint32 {
	// hwnd := syscall.Handle(0x12345678) // 用窗口句柄替换这里

	var processID uint32
	_, _, err := getWindowThreadProcessID.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&processID)))
	if err != nil && err.Error() != "The operation completed successfully." {
		fmt.Println("获取进程ID失败:", err)
	} else {
		fmt.Printf("窗口句柄 %v 对应的进程ID是 %d\n", hwnd, processID)
	}
	return processID
}

func GetProgressMemInfo(pid uint32) {

	// 获取进程信息
	p, err := ps.FindProcess(int(pid))
	if err != nil {
		fmt.Println("Error finding process:", err)
		return
	}

	// 使用gopsutil打开进程
	proc, err := process.NewProcess(int32(p.Pid()))
	if err != nil {
		fmt.Println("Error opening process:", err)
		return
	}

	// 获取进程的内存信息
	memInfo, err := proc.MemoryInfo()
	if err != nil {
		fmt.Println("Error getting memory info:", err)
		return
	}

	fmt.Printf("Process Name: %s\n", p.Executable())
	fmt.Printf("Process ID: %d\n", p.Pid())
	// 驻留内存（RSS）和虚拟内存（VMS）信息
	fmt.Printf("Resident Memory: %d bytes\n", memInfo.RSS)
	fmt.Printf("Virtual Memory: %d bytes\n", memInfo.VMS)
}
