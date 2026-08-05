package main

import (
	"swing-go/backend/wayland"
)

func main() {
	runtime := wayland.NewRuntime()

	runtime.Bootstrap()

	select {}
}
