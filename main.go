package main

/*
#include "samp-plugin-sdk/plugincommon.h"
#include "samp-plugin-sdk/amx/amx.h"
#include "samp-plugin-sdk/sampgdk.h"

#cgo CFLAGS: -Wno-attributes
*/
import "C"

import (
	"fmt"
	"unsafe"
)

//export Supports
func Supports() uint {
	fmt.Println("Called main.go#Supports")
	return C.SUPPORTS_VERSION | C.SUPPORTS_AMX_NATIVES
}

//export Load
func Load(ppData *unsafe.Pointer) bool {
	fmt.Println("Called main.go#Load")
	return true
}

//export Unload
func Unload() bool {
	fmt.Println("Called main.go#Unload")
	return true
}

//export AmxLoad
func AmxLoad(amx *C.AMX) C.int {
	fmt.Println("Called main.go#AmxLoad")
	C.OnGameModeInit()
	C.OnPlayerCommandText(C.int(0), C.CString("/hello world"))
	return 1
}

//export AmxUnload
func AmxUnload() uint {
	fmt.Println("Called main.go#AmxUnload")
	return C.AMX_ERR_NONE
}

//export OnGameModeInit
func OnGameModeInit() C.bool {
	fmt.Println("Called main.go#OnGameModeInit")
	return true
}

//export OnPlayerCommandText
func OnPlayerCommandText(playerid C.int, cmdtext *C.char) C.bool {
	fmt.Println("OnPlayerCommandText called, ", C.GoString(cmdtext))
	return true
}
