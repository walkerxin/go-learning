package main

import "testing"

/* 子测试，进行分组测试

*/
func TestHello2(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := aHello("Michael")
		want := "Hello, Michael"

		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := aHello("")
		want := "Hello, world"

		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})
}