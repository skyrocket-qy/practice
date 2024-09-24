#include <iostream>
#include <unordered_map>
#include <vector>
#include <string>

using namespace std;

// Trie node structure
struct TrieNode {
    unordered_map<char, TrieNode*> children; // map of children nodes (digits or chars)
    bool isEndOfNumber; // true if the node represents the end of a number

    TrieNode() {
        isEndOfNumber = false;
    }
};

// Trie (prefix tree) class
class Trie {
public:
    TrieNode* root;

    Trie() {
        root = new TrieNode();
    }

    // Function to insert a number into the Trie
    void insert(int num) {
        string s = to_string(num); // Convert the integer to a string
        TrieNode* node = root;
        for (char digit : s) {
            // If the digit is not already a child, create a new node
            if (node->children.find(digit) == node->children.end()) {
                node->children[digit] = new TrieNode();
            }
            node = node->children[digit]; // Move to the next node
        }
        node->isEndOfNumber = true; // Mark the end of the number
    }

    // Function to search for a prefix in the Trie
    bool searchPrefix(string prefix) {
        TrieNode* node = root;
        for (char digit : prefix) {
            // If the digit is not found, the prefix doesn't exist
            if (node->children.find(digit) == node->children.end()) {
                return false;
            }
            node = node->children[digit]; // Move to the next node
        }
        return true; // All digits in the prefix were found
    }

    // Function to check if a full number exists in the Trie
    bool searchNumber(int num) {
        string s = to_string(num); // Convert the integer to a string
        TrieNode* node = root;
        for (char digit : s) {
            if (node->children.find(digit) == node->children.end()) {
                return false; // If any digit is not found, the number doesn't exist
            }
            node = node->children[digit];
        }
        return node->isEndOfNumber; // Check if the end of the number is reached
    }
};

int main() {
    Trie trie;

    // Insert numbers into the Trie
    vector<int> numbers = {12345, 123, 456, 789, 4567, 1234};
    for (int num : numbers) {
        trie.insert(num);
    }

    // Search for a prefix
    string prefix = "123";
    if (trie.searchPrefix(prefix)) {
        cout << "Prefix " << prefix << " found in the Trie." << endl;
    } else {
        cout << "Prefix " << prefix << " not found in the Trie." << endl;
    }

    // Search for a full number
    int number = 1234;
    if (trie.searchNumber(number)) {
        cout << "Number " << number << " found in the Trie." << endl;
    } else {
        cout << "Number " << number << " not found in the Trie." << endl;
    }

    return 0;
}
