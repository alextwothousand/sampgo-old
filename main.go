package main

/*
#include "samp-plugin-sdk/plugincommon.h"
#include "samp-plugin-sdk/amx/amx.h"

#define SAMPGDK_AMALGAMATION
#include "samp-plugin-sdk/sampgdk.h"

PLUGIN_EXPORT bool PLUGIN_CALL OnGameModeInit()
{
	extern bool goOnGameModeInit();
    return true;
}

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

/*export ProcessTick
func ProcessTick() {
	/ /C.sampgdk_ProcessTick
}*/

//export Load
func Load(ppData *unsafe.Pointer) bool {
	fmt.Println("Called main.go#Load")
	pAMXFunctions := unsafe.Pointer(uintptr(*ppData) + C.PLUGIN_DATA_AMX_EXPORTS)
	//C.sampgdk_Load(ppData, C.int(0))
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
	return 1
}

//export AmxUnload
func AmxUnload() uint {
	fmt.Println("Called main.go#AmxUnload")
	return C.AMX_ERR_NONE
}

//export goOnGameModeInit
func goOnGameModeInit() C.bool {
	fmt.Println("Called main.go#OnGameModeInit")
	return true
}

//export OnPlayerConnect
func OnPlayerConnect(playerid C.int) C.bool {
	fmt.Println("OnPlayerConnect called for %d", int(playerid))
	return true
}

func main() {}
