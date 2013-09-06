package main

import (
	"strconv"
	"syscall"
	"winapi"
)

func _TEXT(_srt string) *uint16 {
	return syscall.StringToUTF16Ptr(_srt)
}

func _toString(_n int32) string {
	return strconv.Itoa(int(_n))
}

func main() {
	var hwnd winapi.HWND
	//cxScreen := winapi.GetSystemMetrics(winapi.SM_CXSCREEN)
	//cyScreen := winapi.GetSystemMetrics(winapi.SM_CYSCREEN)
	winapi.MessageBox(hwnd, _TEXT("我是消息窗口"), _TEXT("Go语言消息"), winapi.MB_OK)
}
