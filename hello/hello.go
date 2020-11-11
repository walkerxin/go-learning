// 包是一种将相关的 Go 代码组合到一起的方式。
package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func aHello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := englishHelloPrefix

	switch language {
		case french:
			prefix = frenchHelloPrefix
		case spanish:
			prefix = spanishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(aHello("world", ""))
}