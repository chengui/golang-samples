#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include "engine.h"
#include "wrapper.h"

using namespace std;

VoidPointer EngineResource_New(char *cfg) {
    ifstream cfg_is(cfg, ios_base::in);
    EngineResource *ptr = new EngineResource(cfg_is);
    cfg_is.close();
    return ptr;
}

void EngineResource_Delete(VoidPointer p) {
    delete static_cast<EngineResource *>(p);
}

VoidPointer EngineOptions_New(int argc, char **argv) {
    EngineOptions *ptr = new EngineOptions();
    ptr->Load(argc, argv);
    return ptr;
}

void EngineOptions_Delete(VoidPointer p) {
    delete static_cast<EngineOptions *>(p);
}


VoidPointer Engine_New(VoidPointer res, VoidPointer opts) {
    EngineResource *resource = static_cast<EngineResource *>(res);
    EngineOptions *options = static_cast<EngineOptions *>(opts);
    return new Engine(*resource, *options);
}

void Engine_Delete(VoidPointer p) {
    delete static_cast<Engine *>(p);
}

void Engine_Predict(VoidPointer p, char *data, int data_len, char *output) {
    vector<char> str(data, data+data_len);
    vector<char> ans = static_cast<Engine *>(p)->Predict(str);
    copy(ans.begin(), ans.end(), output);
}

