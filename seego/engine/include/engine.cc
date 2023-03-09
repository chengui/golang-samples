#include <iostream>
#include <vector>
#include <string>
#include "engine.h"

using namespace std;

Engine::Engine(EngineResource &res, EngineOptions &opts)
: _res(res), _opts(opts)
{
    _size = 0;
}

Engine::~Engine() {
}

vector<char> Engine::Predict(vector<char> &data) {
    vector<char> ans;
    for (vector<char>::reverse_iterator it = data.rbegin(); it != data.rend(); it++) {
        ans.push_back(*it);
    }
    return ans;
}
