package iteration

import "fmt"
import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 3)
	expected := "aaa"

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i:=0; i<b.N; i++ {
		Repeat("a", b.N)
	}
}

func ExampleRepeat() {
	result := Repeat("b", 8)
	fmt.Println(result)
	// Output: bbbbbbbb
}


// Go 里字符串只能是""包裹，只有占位符是''，这个太容易错了

/*三、迭代
1. Go 中循环和迭代的关键字只有 `for`
声明变量
	:= 包含两步（声明并初始化变量，推断类型)
	var 关键字显式声明，指定类型
	= 仍是赋值
2. %s 与%q的区别？
3. 编写**`基准测试（benchmarks）`**
	- 是Go语言的另一个一级特性，它与编写测试非常相似
	- 【声明】func BenchmarkRepeat(b *testing.B)
	- testing.B 可使你访问隐性命名（cryptically named）b.N
	- 基准测试运行时，代码会运行 b.N 次，并测量需要多长时间。
	- 代码运行的次数不会对你产生影响，测试框架会选择一个它所认为的最佳值，以便让你获得更合理的结果。
	- 用 go test -bench=. 来运行基准测试 （Powershell中 go test -bench="."）
	- 【输出说明】运行一次函数用了多久，运行次数

4. [strings包](https://golang.org/pkg/strings/)
*/