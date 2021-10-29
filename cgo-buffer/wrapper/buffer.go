package buffer

// #cgo LDFLAGS: -L../lib -lbuffer -lstdc++
// #cgo CXXFLAGS: -I../include -std=c++11
// #cgo CFLAGS: -I../include
// #include "wrapper.h"
import "C"
import "unsafe"

type Buffer struct {
    cptr C.handle
}

func NewBuffer(size int) *Buffer {
    return &Buffer{
        cptr: C.NewBuffer(C.int(size)),
    }
}

func (p *Buffer) Delete() {
    C.DeleteBuffer(p.cptr)
}

func (p *Buffer) Data() []byte {
    data := C.Buffer_Data(p.cptr)
    size := C.Buffer_Size(p.cptr)

    return ((*[1 << 31]byte)(unsafe.Pointer(data)))[0:int(size):int(size)]
}

func (p *Buffer) Size() int {
    size := C.Buffer_Size(p.cptr)
    return int(size)
}

func (p *Buffer) Append(str string) string {
    cstr := C.CString(str)
    defer C.free(unsafe.Pointer(cstr))
    cdata := (*C.char)(C.malloc(20))
    C.Buffer_Append(p.cptr, cstr, cdata)
    return C.GoString(cdata)

}

func (p *Buffer) Print() {
    data := p.Data()
    C.puts((*C.char)(unsafe.Pointer(&data[0])))
}
