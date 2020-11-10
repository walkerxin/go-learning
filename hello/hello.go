// 包是一种将相关的 Go 代码组合到一起的方式。
package main

import "fmt"

const englishHelloPrefix = "Hello, "

func aHello(name string) string {
	if name == "" {
		name = "world"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(aHello("world"))
}