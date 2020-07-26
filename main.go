package main

/*
#include "samp-plugin-sdk/plugincommon.h"
#include "samp-plugin-sdk/amx/amx.h"
#cgo CFLAGS: -Wno-attributes

#include "natives.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

//export Test
func Test(amx *C.AMX, params *C.cell) {
	fmt.Println("Test func was called successfully")
}

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
	fmt.Printf("%v\n", amx)

	C.LoadNatives(amx)
	return 1
}

//export AmxUnload
func AmxUnload() uint {
	fmt.Println("Called main.go#AmxUnload")
	return C.AMX_ERR_NONE
}

func main() {}
