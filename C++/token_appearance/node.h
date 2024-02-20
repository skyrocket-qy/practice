#ifndef NODE_H
#define NODE_H

#include <string>
#include <unordered_map>

class Node {
public:
    std::string Token;
    int Start;
    int Count;
    int End;
    std::unordered_map<std::string, Node*> Children;

    Node(std::string token);
    ~Node();
};

#endif // NODE_H
