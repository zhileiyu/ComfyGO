package main

import "github.com/zhileiyu/comfyGO"

func main() {
	ci := comfyGO.NewComfyClient()
	println(ci.ISServerOK())
}
