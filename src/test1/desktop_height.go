// +build USE_WINSYSTEMMETRICS

package main

import "github.com/lxn/win"

func getDesktopHeight() int {
	return int(win.GetSystemMetrics(win.SM_CYSCREEN))
}
