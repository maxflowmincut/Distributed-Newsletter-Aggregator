#include "filter.h"
#include <string>

extern "C" {

bool isRelevant(const char* description) {
    std::string desc(description);
    if (desc.size() < 50) {
        return false;
    }
    return true;
}

}