package main

/*
#include "../samp-plugin-sdk/plugincommon.h"
#include "../samp-plugin-sdk/amx/amx.h"
#cgo CFLAGS: -Wno-attributes

extern int amx_Register(AMX *amx, const AMX_NATIVE_INFO *NativeList, int number);

AMX_NATIVE_INFO PluginNatives[] =
{
    {"Test", Test},
    {0, 0}
};
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
func Unload() {
	fmt.Println("Called main.go#Unload")
}

//export AmxLoad
func AmxLoad(amx *C.AMX) {
	fmt.Println("Called main.go#AmxLoad")
	fmt.Printf("%v\n", amx)

	C.amx_Register(amx, C.PluginNatives, -1)
}

//export AmxUnload
func AmxUnload() C.int {
	fmt.Println("Called main.go#AmxUnload")
	return 0 // equal to AMX_ERR_NONE
}

//export Test
func Test() C.cell {
	fmt.Println("Hello world!")
	return 1
}

func main() {}
