#ifndef _BUFFER_H_
#define _BUFFER_H_

#include <string>

class Buffer {
private:
    std::string* s_;
public:
    Buffer(int size);

    ~Buffer();

    int Size() const;

    char* Data();

    void Append(char *s);
};

#endif /* ifndef _BUFFER_H_ */
