package main

import (
	"github.com/electricbubble/go-toast"
)

func main() {
	// _ = toast.Push("test message")
	_ = toast.Push("我是傻逼", toast.WithTitle("我宣布个事"))
}
