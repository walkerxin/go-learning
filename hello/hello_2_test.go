package main

import "testing"

/* 子测试，进行分组测试

测试用例中重复的代码
*/

func TestHello2(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper() // 告诉测试套件该方法只是个辅助函数（helper）。使报错信息位置更精准
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Michael", "")
		want := "Hello, Michael"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Jane", "Spanish")
		want := "Hola, Jane"
		assertCorrectMessage(t, got, want)

	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Polo", "French")
		want := "Bonjour, Polo"
		assertCorrectMessage(t, got, want)
	})
}