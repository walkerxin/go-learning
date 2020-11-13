package structs

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	expected := 40.0

	if got != expected {
		t.Errorf("expected %.2f but got %.2f", expected, got)
	}
}

func TestArea(t *testing.T) {
	got := Area(10.0, 10.0)
	expected := 100.0

	if got != expected {
		t.Errorf("expected %.2f but got %.2f", expected, got)
	}

}

/*
新的格式化字符串了吗？%.2f
	float64，.2 表示输出 2 位小数
*/
/*
1. 利用保留字 struct 来定义自己的类型

上面的周长 2*(a+b) 和面积 a*b
我们可以仅仅给这些函数命名成像 RectangleArea 一样更具体的名字。
但是更简洁的方案是定义我们自己的类型 Rectangle，它可以封装长方形的信息。
一个通过 struct 定义出来的类型是一些已命名的域的集合，这些域用来保存数据。

*/