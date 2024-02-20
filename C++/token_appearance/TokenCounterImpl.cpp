#include "TokenCounterImpl.h"
#include <unordered_map>
#include <vector>
#include "node.h"

TokenCounterImpl::TokenCounterImpl() : Total(0), TokenNodeMap() {}

void TokenCounterImpl::Ingest(std::string input) {
    Total++;
    std::vector<std::string> tokens;
    size_t pos = 0;
    std::string token;
    while ((pos = input.find(":")) != std::string::npos) {
        token = input.substr(0, pos);
        tokens.push_back(token);
        input.erase(0, pos + 1);
    }
    tokens.push_back(input);

    Node* curNode = nullptr;
    for (size_t i = 0; i < tokens.size(); ++i) {
        std::string token = tokens[i];
        Node* node = nullptr;

        // get or create node
        auto it = TokenNodeMap.find(token);
        if (it == TokenNodeMap.end()) {
            node = new Node(token);
            TokenNodeMap[token] = node;
        } else {
            node = it->second;
        }

        // add children
        if (curNode != nullptr) {
            curNode->Children[token] = node;
        }
        curNode = node;
        curNode->Count++;

        if (i == 0) {
            curNode->Start++;
        }
        if (i == tokens.size() - 1) {
            curNode->End++;
        }
    }
}

float TokenCounterImpl::Appearance(std::string input) {
    std::vector<std::string> tokens;
    size_t pos = 0;
    std::string token;
    while ((pos = input.find(":")) != std::string::npos) {
        token = input.substr(0, pos);
        tokens.push_back(token);
        input.erase(0, pos + 1);
    }
    tokens.push_back(input);

    int process = 0;
    Node* curNode = nullptr;
    for (size_t i = 0; i < tokens.size(); ++i) {
        std::string token = tokens[i];
        if (curNode == nullptr) {
            auto it = TokenNodeMap.find(token);
            if (it == TokenNodeMap.end()) {
                return 0;
            }
            curNode = it->second;
        } else {
            bool isFind = false;
            for (auto& [key, child] : curNode->Children) {
                if (child->Token == token) {
                    curNode = child;
                    isFind = true;
                    break;
                }
            }
            if (!isFind) {
                return 0;
            }
        }

        if (i == 0) {
            process = curNode->Count;
        } else if (process > curNode->Count - curNode->Start) {
            process = curNode->Count - curNode->Start;
        }
        if (i == tokens.size() - 1) {
            return static_cast<float>(process) / static_cast<float>(Total);
        }

        process -= curNode->End;
        if (process <= 0) {
            return 0;
        }
    }

    return 0;
}

void TokenCounterImpl::Init() {
    for (auto& [key, node] : TokenNodeMap) {
        delete node;
    }
    TokenNodeMap.clear();
    Total = 0;
}

TokenCounterImpl::~TokenCounterImpl() {
    for (auto& [key, node] : TokenNodeMap) {
        delete node;
    }
    TokenNodeMap.clear();
}
