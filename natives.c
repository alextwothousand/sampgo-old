#include "natives.h"

// extern cell AMX_NATIVE_CALL Test(AMX *amx, cell *params);

AMX_NATIVE_INFO PluginNatives[] =
{
    {"HelloWorld", Test},
    {0, 0}
};

void LoadNatives(AMX *amx) {
    amx_Register(amx, PluginNatives, -1);
}