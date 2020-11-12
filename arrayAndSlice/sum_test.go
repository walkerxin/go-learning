package arrayAndSlice

import "testing"

//【需求1】创建一个 Sum 函数，它使用 for 来循环获取数组中的元素并返回所有元素的总和。
func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int {1, 2, 3}
		
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got '%d' want '%d' given, %v", got, want, numbers)
		}
	})
}


/*
1. 如何声明并初始化一个数组
	[length]type{value1, value2, ..., valueN}
	[...]type{value1, value2, ..., valueN}
1-2. %v 默认输出格式
	[格式字符串more](https://golang.org/pkg/fmt/)

2. 遍历数组，支持 array[index] 索引取值
3. `range` 语法重构遍历数组
	- range 会迭代数组，每次迭代都会`返回数组元素的索引和值`。
		我们选择使用 _ `空白标志符` 来忽略索引。
4. 数组的长度也属于它的类型
如 [5]int 和 [4]int
./sum_test.go:8:12: cannot use numbers (type [5]int) as type [4]int in argument to Sum
FAIL    github.com/amanda/arrayAndSlice [build failed]

因为这个原因，所以数组比较笨重，大多数情况下我们都不会使用它。
5. 使用**切片类型**改写，以支持动态大小
	> Go 的切片（slice）类型不会将集合的长度保存在类型中，因此它的长度是不固定的。
	- 语法上和数组非常相似，只是在声明的时候不指定长度
		mySlice := []int{1, 2, 3}
		mySlice := [3]int{1, 2, 3}
	- 执行`覆盖率测试工具` go test -cover


*/