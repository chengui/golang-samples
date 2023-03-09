#ifndef _WRAPPER_H_
#define _WRAPPER_H_

#include <stdio.h>
#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef void* handle;

handle NewBuffer(int size);
void DeleteBuffer(handle);
int Buffer_Size(handle);
char* Buffer_Data(handle);
void Buffer_Append(handle, char *str, char *out);

#ifdef __cplusplus
}
#endif

#endif
