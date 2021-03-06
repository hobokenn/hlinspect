package cvar

import (
	"hlinspect/internal/engine"
	"strconv"
	"unsafe"
)

import "C"

var Wallhack = makeCVar("hli_wallhack", "0")

func makeCVar(name, value string) engine.CVar {
	floatVal, _ := strconv.ParseFloat(value, 32)
	cvar := engine.RawCVar{
		// Probably don't need to free these strings?
		Name:   uintptr(unsafe.Pointer(C.CString(name))),
		String: uintptr(unsafe.Pointer(C.CString(value))),
		Value:  float32(floatVal),
	}
	return engine.MakeCVar(unsafe.Pointer(&cvar))
}
