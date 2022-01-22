package luabind


/*
#cgo CFLAGS: -std=gnu99
//#cgo CFLAGS: -DHAVE_USLEEP=1
//#cgo CFLAGS: -DLUA_USE_DLOPEN=1
#cgo CFLAGS: -DLUA_USE_READLINE=1
#cgo LDFLAGS: -L/usr/lib -L/usr/local/opt/lua/lib -ledit -llua
#cgo CFLAGS: -Wno-deprecated-declarations
//#cgo linux,!android CFLAGS: -DHAVE_PREAD64=1 -DHAVE_PWRITE64=1
#include "lua/lua.h"
#include "lua/lapi.h"
#include "lua/lauxlib.h"
#cgo CFLAGS: -DUSE_ONELUA=1
//#include "lua/onelua.c"
#include "lua/lua.c"
*/
import "C"
import "unsafe"

var (
	LUA_VERSION = C.LUA_VERSION
	LUA_VERSION_NUM = C.LUA_VERSION_NUM
	LUA_VERSION_RELEASE_NUM = C.LUA_VERSION_RELEASE_NUM
)


// Version returns SQLite library version information.
func Version() (libVersion string, libVersionNumber, releaseNum int) {
	libVersion = string(C.LUA_VERSION)
	libVersionNumber = int(C.LUA_VERSION_NUM)
	releaseNum = int(C.LUA_VERSION_RELEASE_NUM)
	return libVersion, libVersionNumber, releaseNum
}

func NewLuaState() (L *C.lua_State, err error) {
	L = C.luaL_newstate()
	err = nil
	return
}

func LuaMain(){
	args := [...]string{"lua"} //, "-e", "print('hello go-lua')"
	carg := make([]*C.char, 0)
	for _, i := range args {
		carg = append(carg, (*C.char)(unsafe.Pointer(C.CString(i))))
	}
	C.lua_main(C.int(len(args)), (**C.char)(unsafe.Pointer(&carg[0])))
}