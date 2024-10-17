# Compiler and flags
CXX = g++
CXXFLAGS = -I src/core -I src/utils -Wall -O2
AR = ar rcs

# Directories
BUILD_DIR = build
LIB_DIR = lib
SRC_DIR = src

# Source files
CORE_SRC = $(SRC_DIR)/core/core.cpp
UTILS_SRC = $(SRC_DIR)/utils/utils.cpp

# Object files
CORE_OBJ = $(BUILD_DIR)/core.o
UTILS_OBJ = $(BUILD_DIR)/utils.o

# Libraries
CORE_LIB = $(LIB_DIR)/libcore.a
UTILS_LIB = $(LIB_DIR)/libutils.a

# Application source and target
APP_SRC = $(SRC_DIR)/app/main.cpp
APP_TARGET = app

# Build the application
all: $(APP_TARGET)

# Link the final executable using the static libraries
$(APP_TARGET): $(APP_SRC) $(CORE_LIB) $(UTILS_LIB)
	$(CXX) $(CXXFLAGS) -o $@ $^
	
# Build the static library for core
$(CORE_LIB): $(CORE_OBJ)
	$(AR) $(CORE_LIB) $(CORE_OBJ)

# Build the static library for utils
$(UTILS_LIB): $(UTILS_OBJ)
	$(AR) $(UTILS_LIB) $(UTILS_OBJ)

# Compile core object
$(CORE_OBJ): $(CORE_SRC) $(SRC_DIR)/core/core.hpp
	$(CXX) $(CXXFLAGS) -c $(CORE_SRC) -o $(CORE_OBJ)

# Compile utils object
$(UTILS_OBJ): $(UTILS_SRC) $(SRC_DIR)/utils/utils.hpp
	$(CXX) $(CXXFLAGS) -c $(UTILS_SRC) -o $(UTILS_OBJ)

# Clean the build
clean:
	rm -f $(BUILD_DIR)/*.o $(LIB_DIR)/*.a $(APP_TARGET)
