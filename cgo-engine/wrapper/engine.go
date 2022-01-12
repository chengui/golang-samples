package wrapper

// #cgo LDFLAGS: -L../lib -lengine -lstdc++
// #cgo CXXFLAGS: -I../include -std=c++11
// #cgo CFLAGS: -I../include
// #include <stdlib.h>
// #include "wrapper.h"
import "C"
import (
	"strings"
	"unsafe"
)

type Engine struct {
	cptr C.VoidPointer
}

func NewEngine(res *EngineResource, opts *EngineOptions) *Engine {
	return &Engine{
		cptr: C.Engine_New(res.cptr, opts.cptr),
	}
}

func (p *Engine) Delete() {
	C.Engine_Delete(p.cptr)
}

func (p *Engine) Predict(data []byte) (string, error) {
	clen := C.int(len(data))
	cdata := (*C.char)(unsafe.Pointer(&data[0]))
	cbuff := (*C.char)(C.malloc(C.size_t(1024)))
	defer C.free(unsafe.Pointer(cbuff))
	C.Engine_Predict(p.cptr, cdata, clen, cbuff)
	return C.GoString(cbuff), nil
}

type EngineOptions struct {
	cptr C.VoidPointer
}

func NewEngineOptions(options string) *EngineOptions {
	cmds := strings.Split(options, ",")
	argc := len(cmds)
	args := make([](*C.char), argc)
	for i := range cmds {
		ptr := C.CString(cmds[i])
		defer C.free(unsafe.Pointer(ptr))
		args[i] = (*C.char)(unsafe.Pointer(ptr))
	}
	if options == "" {
		argc = 0
	}
	return &EngineOptions{
		cptr: C.EngineOptions_New(C.int(argc), &args[0]),
	}
}

func (p *EngineOptions) Delete() {
	C.EngineOptions_Delete(p.cptr)
}

type EngineResource struct {
	cptr C.VoidPointer
}

func NewEngineResource(resource string) *EngineResource {
	cresource := C.CString(resource)
	defer C.free(unsafe.Pointer(cresource))
	return &EngineResource{
		cptr: C.EngineResource_New(cresource),
	}
}

func (p *EngineResource) Delete() {
	C.EngineResource_Delete(p.cptr)
}
