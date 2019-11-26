#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int n = 100;

char* NewBuf()  {
    char* buf = (char*)malloc(n);
    memset(buf, 0, n);
    memcpy(buf, "hello cgo", 9);
    return buf;
}

void PrintBuf(char* s) {
    sleep(1);
    puts(s);
    sleep(1);
}


void FreeBuf(char* s) {
    free(s);
}

void SayHello(const char *s) {
    puts(s);
}

// int main(){
//     SayHello("Dabin");
// }