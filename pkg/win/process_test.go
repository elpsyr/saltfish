package win

import (
	"fmt"
	"testing"
)

func TestGetProcessIDByHwnd(t *testing.T) {
	hwndByTitle := GetHwndByTitle("咸鱼之王")
	ID := GetProcessIDByHwnd(hwndByTitle)
	fmt.Println(ID)
}
