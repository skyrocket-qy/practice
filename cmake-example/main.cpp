#include <iostream>
#include <stdexcept>
#include <string>

// Define a simple logging macro
#define LOG_ERROR(message) std::cerr << "ERROR: " << message << std::endl

int main() {
    try {
        // Simulate an exception
        throw std::runtime_error("This is a demo exception");
    }
    catch (const std::exception& e) {
        // Use our LOG_ERROR macro to log the exception's message
        LOG_ERROR(std::string(e.what()));
    }
    return 0;
}
