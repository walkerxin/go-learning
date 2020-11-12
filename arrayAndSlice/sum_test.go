package arrayAndSlice

import "testing"
import "reflect"

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

//【需求2】SumAll 函数，它接受多个切片，并返回由每个切片元素的总和组成的新切片。
func TestSumAll(t *testing.T) {
	numbers1 := []int{1, 2, 3}
	numbers2 := []int{4, 5}

	got := SumAll(numbers1, numbers2)
	want := []int{6, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

/*【需求3】进化为 SumAllTails，
	把每个切片的尾部元素相加（尾部的意思就是除去第一个元素以外的其他元素）*/
func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		numbers1 := []int{1, 2, 3}
		numbers2 := []int{4, 5}
		got := SumAllTails(numbers1, numbers2)
		want := []int{5, 5}
		checkSums(t, got, want)
	})

	// 测试空切片传入时可能发生的越界 slice bounds out of range
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{4, 5})
		want := []int{0, 5}
		checkSums(t, got, want)
	})
}

/*数组：`有序`、`多个相同类型的元素`
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

/*【需求2】SumAll
6. 使用`可变参数` ...[例子](https://gobyexample.com/variadic-functions)
7. 切片比较是否相等的方法
	invalid operation: got != want (slice can only be compared to nil)
	在 Go 中不能对切片使用等号运算符。
	- 你可以写一个函数迭代每个元素来检查它们的值。
	-【比较简单】是使用 reflect.DeepEqual，它在判断两个变量是否相等时十分有用。
	-【注意 reflect.DeepEqual 不是「类型安全」的】即使比较的是不同类型，也能通过编译

8. make([]int, length) 
	- 创建指定类型和长度(容量)的切片

9. 使用 `append` 函数，从原来的切片中创建一个新切片
	- 数组有越界问题，属于编译时错误
	- 切片相当于长度不是类型的数组，有容量的概念，所以也存在越界
	- 切片发生越界，属于运行时错误
	可以通过sums = append(sums, xxx)来实现追加

10. 如何获取部分切片
	- 使用语法 slice[low:high] 获取部分切片。
		如果在冒号的一侧没有数字就会一直取到最边缘的元素。
		我们使用 numbers[1:] 取到从索引 1 到最后一个元素。
		你可能需要花费一些时间才能熟悉切片的操作。

11. 获取数组和切片的长度，如 len([]int{1, 2, 3})

12. 通过重构，提取出checkSums公用函数，一定程度上增加了代码的类型安全
*/