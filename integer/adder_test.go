package integers

import "fmt"
import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

/* TDD 是一门需要通过开发去实践的技能，
	通过将问题分解成更小的可测试的组件，你编写软件将会更加轻松。
*/

/*整数
编写一个加法函数测试用例

0. 包的问题，规范：每个目录下的文件同属于一个包
	[解释](https://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project)
	integer包，表示将聚集那些处理整数的函数
1. 占位符
	%d 十进制整数
	%s 字符串
	%q
2. 函数声明中参数类型相同时可缩写
	(a int, b int) -> (a, b int)
*/

/*
？一种“基于属性测试”的技巧，能帮助你查找程序漏洞。

3. “示例” - 示例函数 - 作为测试的一部分
	- 如 func ExampleAdd() {}
	- 典型测试 func TestXxx(t *testing.T) {}
	- 示例函数中必须包含注释[// Output: 6]，否则该示例函数会被编译，但不会被执行
		- 包含到doc中
go test -v(能查看到执行了哪些测试用例)
*/