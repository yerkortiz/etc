#include <stdio.h>
#include <string.h>

void substrings(char* s, int s_size) {
    for (int i = 1; i <= s_size; ++i) {
        int len = i;
        for(int j = 0; j + len <= s_size; ++j) {
            for(int k = j; k < j + len; ++k) {
                printf("%c ", s[k]);
            }
            printf("\n");
        }
    }
}

void tupleK (int* nums, int nums_size) {

}

int main(int argv, char **argc) {
    char* s = "hello, world";
    int s_size = strlen(s);
    substrings(s, s_size);
}
