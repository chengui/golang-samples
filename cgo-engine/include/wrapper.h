#ifndef _WRAPPER_H_
#define _WRAPPER_H_

#ifdef __cplusplus
extern "C" {
#endif


typedef void* VoidPointer;

VoidPointer EngineResource_New(char *cfg);
void EngineResource_Delete(VoidPointer p);

VoidPointer EngineOptions_New(int argc, char** argv);
void EngineOptions_Delete(VoidPointer p);

VoidPointer Engine_New(VoidPointer res, VoidPointer opts);
void Engine_Delete(VoidPointer p);
void Engine_Predict(VoidPointer p, char *data, int datalen, char *output);


#ifdef __cplusplus
}
#endif

#endif
