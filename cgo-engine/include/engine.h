#ifndef _ENGINE_H_
#define _ENGINE_H_

#include <iostream>
#include <vector>
#include <string>
using namespace std;

class EngineResource {
public:
    EngineResource(ifstream &stream) {
        _size = 0;
    }

    ~EngineResource() {
    }

private:
    int _size;
};

class EngineOptions {
public:
    EngineOptions() {
        _size = 0;
    }

    ~EngineOptions() {
    }

    void Load(int argc, char **argv) {
        cout << "Load Engine Options: ";
        for (int i = 0; i < argc; i++) {
            cout << argv[i] << ",";
        }
        cout << "." << endl;
    }

private:
    int _size;
};

class Engine {
public:
    Engine(EngineResource &res, EngineOptions &opts);
    ~Engine();

    vector<char> Predict(vector<char> &data);

private:
    int _size;
    EngineResource &_res;
    EngineOptions &_opts;
};

#endif /* ifndef _ENGINE_H_ */
