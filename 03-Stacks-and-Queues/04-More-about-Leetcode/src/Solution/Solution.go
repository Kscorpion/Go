package main

import (
	"Play-with-Data-Structures/03-Stacks-and-Queues/04-More-about-Leetcode/src/ArrayStack"
	"fmt"
)

// go 的字符串实际是 byte 类型组成的切片
func isValid(s string) bool {
	brackets := map[rune]rune{')': '(', ']': '[', '}': '{'}
	stack := ArrayStack.Constructor(20)

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			// 入栈
			stack.Push(char)
		} else if stack.GetSize() > 0 && brackets[char] == stack.Peek() {
			// 栈中有数据，且此元素与栈尾元素相同
			stack.Pop()
		} else {
			return false
		}
	}

	// 循环结束，栈中还有数据则 false
	return stack.GetSize() == 0
}

func main() {
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("([)]"))
}
