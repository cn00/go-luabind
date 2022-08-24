package main

import (
	"golua/luabind"
)

func main() {
	print("go-")
	println(luabind.Version())
	// L , _:= luabind.NewLuaState()
	// println(L)
	luabind.LuaMain()
}
