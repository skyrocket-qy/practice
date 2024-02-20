#ifndef TOKENCOUNTERIMPL_H
#define TOKENCOUNTERIMPL_H

#include <string>
#include <unordered_map>
#include "node.h"

class TokenCounterImpl {
public:
    std::unordered_map<std::string, Node*> TokenNodeMap;
    int Total;

    TokenCounterImpl();

    void Ingest(std::string input);

    float Appearance(std::string input);

    void Init();

    ~TokenCounterImpl();
};

#endif // TOKENCOUNTERIMPL_H
