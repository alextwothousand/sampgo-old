package main

/*
#cgo CFLAGS: -Wno-attributes

#define SAMPGDK_AMALGAMATION

#if defined __linux__
	#define LINUX
#elif defined __win32
	#define WIN32
#endif

#include "sampgdk.h"

extern int goOnGameModeInit();

PLUGIN_EXPORT bool PLUGIN_CALL OnGameModeInit() {
	return goOnGameModeInit();
}

PLUGIN_EXPORT unsigned int PLUGIN_CALL Supports() {
	return sampgdk_Supports() | SUPPORTS_PROCESS_TICK;
}

PLUGIN_EXPORT bool PLUGIN_CALL Load(void **ppData) {
	return sampgdk_Load(ppData);
}

PLUGIN_EXPORT void PLUGIN_CALL Unload() {
	sampgdk_Unload();
}

PLUGIN_EXPORT void PLUGIN_CALL ProcessTick() {
	sampgdk_ProcessTick();
}
*/
import "C"
import "fmt"

// export goOnGameModeInit
func goOnGameModeInit() C.int {
	fmt.Printf("Hello World!\n")
	return true
}

func main() {}
