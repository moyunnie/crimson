package main

import (
	"crimson/global"
	"crimson/initalize"
)

func main() {
	global.Docker = initalize.InitDocker()
	initalize.InitRouter()
}
