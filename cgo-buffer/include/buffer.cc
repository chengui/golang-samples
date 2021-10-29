#include "buffer.h"


Buffer::Buffer(int size) {
    this->s_ = new std::string(size, char('\0'));
}

Buffer::~Buffer() {
    delete this->s_;
}

int Buffer::Size() const {
    return this->s_->size();
}

char* Buffer::Data() {
    return (char*)this->s_->data();
}

void Buffer::Append(char *s) {
    this->s_->append(s);
}
