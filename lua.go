package main


import (
	"golua/luabind"
)

func main(){
	println(luabind.Version())
	L , _:= luabind.NewLuaState()
	println(L)
	luabind.LuaMain()
}