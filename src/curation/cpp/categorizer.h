#pragma once

#ifdef __cplusplus
extern "C" {
#endif

int* categorize(const char* description, int* size);
void free_categories(int* categories);
void my_free(void* ptr);

#ifdef __cplusplus
}
#endif