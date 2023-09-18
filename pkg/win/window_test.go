package win

import "testing"

func TestSetWindowSize(t *testing.T) {
	hwndByTitle := GetHwndByTitle("咸鱼之王")
	//SetWindowSize(hwndByTitle, 225, 422)
	SetWindowSize(hwndByTitle, 450, 844)
}

func TestSetTrans(t *testing.T) {
	//hwndByTitle := GetHwndByTitle("咸鱼之王")
	hwndByTitle := GetHwndByTitle("salt fish @elpsyr")
	SetWindowAlpha(hwndByTitle, 255)
}
