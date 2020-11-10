package main

import "testing"

func TestHello(t *testing.T) {
	got := aHello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}

// aHello 方法测试用例
/*
- 命名
	1. 文件名 xxx_test.go
	2. 测试函数名要以 Test+Xxx 开头
	3. t *testing.T 作为参数
- if := t.Errorf
*/

// 探讨方法和函数之间的区别？

// 【总结】 测试编写和声明函数
// 还有本地查看 Go 文档

/* 后面都用测试驱动开发（TDD）
	当使用像 Go 这样的静态类型语言时，聆听编译器 是很重要的
	【需求1】指定hello的人

*/