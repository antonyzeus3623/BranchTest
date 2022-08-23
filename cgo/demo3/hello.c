#include "hello.h" // 保证函数的实现满足模块对外公开的接口
#include <stdio.h>
void Hello(char* s) {
    puts(s);
}