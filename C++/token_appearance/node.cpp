#include "node.h"

Node::Node(std::string token) : Token(token), Start(0), Count(0), End(0) {}

Node::~Node() {
    for (auto& pair : Children) {
        delete pair.second;
    }
    Children.clear();
}