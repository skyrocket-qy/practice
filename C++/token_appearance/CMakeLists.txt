cmake_minimum_required(VERSION 3.20)
project(main)

set(CMAKE_CXX_STANDARD 20)  # Specify C++20
set(CMAKE_CXX_STANDARD_REQUIRED ON)  # Require C++20
set(CMAKE_CXX_EXTENSIONS OFF)  # Do not enable compiler extensions

# Add your source files here
add_executable(main main.cpp node.cpp TokenCounterImpl.cpp)
