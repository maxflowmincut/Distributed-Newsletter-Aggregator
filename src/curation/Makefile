CC := g++
CFLAGS := -std=c++11 -fPIC

# Detect OS
ifneq (,$(findstring Windows_NT,$(OS)))
    LDFLAGS := -shared
    LIB_EXT := .dll
    MKDIR := if not exist "cpp\lib" mkdir "cpp\lib"
    RM := del /f /q
    RM_PATH := $(subst /,\,$(TARGETS))
else ifneq (,$(findstring Linux,$(shell uname -s)))
    LDFLAGS := -shared
    LIB_EXT := .so
    MKDIR := mkdir -p $(LIB_DIR)
    RM := rm -f
else
    $(error Unsupported OS)
endif

CPP_DIR := cpp
LIB_DIR := $(CPP_DIR)/lib

TARGETS := $(LIB_DIR)/libcategorizer$(LIB_EXT) $(LIB_DIR)/libfilter$(LIB_EXT)

all: $(TARGETS)

$(LIB_DIR)/libcategorizer$(LIB_EXT): $(CPP_DIR)/categorizer.cpp
	$(MKDIR)
	$(CC) $(CFLAGS) $(LDFLAGS) $< -o $@

$(LIB_DIR)/libfilter$(LIB_EXT): $(CPP_DIR)/filter.cpp
	$(MKDIR)
	$(CC) $(CFLAGS) $(LDFLAGS) $< -o $@

clean:
	$(RM) $(subst /,\,$(TARGETS))