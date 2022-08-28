package main

/*
// 开始配置头文件路径
#cgo CFLAGS : -Wall -Wextra -g -I./helloword/include
// 配置动态链接库路径，其中-L为动态链接库所在目录 -l为动态链接库名，注意动态链接库名称取lib
// 后面的值比如我们编译出来的是libhelloword.so,则动态链接库名为helloword
#cgo LDFLAGS: -L./helloword/output -lhelloword
// 导入头文件，头文件只包含接口定义而接口的实现则已经打包在动态链接库中
// cgo会根据头文件定义的接口取动态链接库中寻找对应的实现
#include "helloword.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

func CallHello(word string) {
	w := (*C.schar)(C.CString(word))
	// 调用c 接口
	C.Hello(w)

}
func main() {
	CallHello("World!")
}
