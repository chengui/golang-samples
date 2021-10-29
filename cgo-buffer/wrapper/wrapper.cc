#include <cstring>
#include <vector>
#include "buffer.h"
#include "wrapper.h"

handle NewBuffer(int size) {
    return new Buffer(size);
}

void DeleteBuffer(handle p) {
    delete static_cast<Buffer*>(p);
}

char* Buffer_Data(handle p) {
    return static_cast<Buffer*>(p)->Data();
}

int Buffer_Size(handle p) {
    return static_cast<Buffer*>(p)->Size();
}

void Buffer_Append(handle p, char * s, char *o) {
    static_cast<Buffer*>(p)->Append(s);
}
