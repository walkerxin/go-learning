package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris", "")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}

// Hello 方法测试用例
/*
- 命名
	1. 测试文件名 xxx_test.go
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
	【需求2】Hello("") 未指定时，返回 "Hello, world"
	【需求3】添加第二个参数，指定hello的语言
*/

/*
**常量**
常量定义 const englishHelloPrefix = "Hello, "
优点：易于理解+提高性能（避免相同的字符串其实例多次创建）

**swith 语句**
**提取出功能函数**
命名返回值 <=> 声明了返回值变量 prefix，
	【使用场景】通常应该在结果的含义在上下文中不明确时使用
		[more](https://github.com/golang/go/wiki/CodeReviewComments#named-result-parameters)
	1. 它将被分配「零」值。这取决于类型，例如 int 是 0，对于字符串它是 ""。
		- 你只需调用 return 而不是 return prefix 即可返回所设置的值。
	2. 这将显示在 Go Doc 中，使你的代码更加清晰。

函数命名
	greetingPrefix 以小写字母开头。
	在 Go 中，
		公共函数以大写字母开始，
		私有函数以小写字母开头。
	我们不希望我们算法的内部结构暴露给外部，所以我们将这个功能私有化。

godoc -http :8000 运行本地文档
*/