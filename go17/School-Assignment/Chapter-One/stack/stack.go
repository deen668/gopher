// 用go语言实现一个栈
package main

import (
	"errors"
	"fmt"
)

// Stack 表示栈数据结构
type Stack struct {
	items []interface{} // 使用切片存储栈中的元素
}

// NewStack 创建并返回一个新的空栈
func NewStack() *Stack {
	return &Stack{
		items: make([]interface{}, 0),
	}
}

// Push 将元素压入栈顶
func (s *Stack) Push(item interface{}) { // 将元素压入栈顶
	s.items = append(s.items, item) // 将元素压入栈顶
}

// Pop 弹出并返回栈顶元素，如果栈为空则返回错误
func (s *Stack) Pop() (interface{}, error) { // 弹出并返回栈顶元素，如果栈为空则返回错误
	if s.IsEmpty() { // 如果栈为空，则返回错误
		return nil, errors.New("stack is empty") // 返回错误
	}

	// 获取栈顶元素（最后一个元素）
	index := len(s.items) - 1 // 获取栈顶元素（最后一个元素）
	item := s.items[index]    // 获取栈顶元素（最后一个元素）

	// 从切片中移除栈顶元素
	s.items = s.items[:index] // 从切片中移除栈顶元素

	return item, nil
}

// Peek 返回栈顶元素但不弹出，如果栈为空则返回错误
func (s *Stack) Peek() (interface{}, error) { // 返回栈顶元素但不弹出，如果栈为空则返回错误
	if s.IsEmpty() { // 如果栈为空，则返回错误
		return nil, errors.New("stack is empty") // 返回错误
	}

	return s.items[len(s.items)-1], nil
}

// IsEmpty 判断栈是否为空
func (s *Stack) IsEmpty() bool { // 判断栈是否为空
	return len(s.items) == 0 // 返回栈是否为空
}

// Size 返回栈中元素的数量
func (s *Stack) Size() int { // 返回栈中元素的数量
	return len(s.items) // 返回栈中元素的数量
}

// Clear 清空栈
func (s *Stack) Clear() { // 清空栈
	s.items = make([]interface{}, 0) // 清空栈
}

// Print 打印栈的内容
func (s *Stack) Print() { // 打印栈的内容
	if s.IsEmpty() {
		fmt.Println("Stack: []")
		return
	}

	fmt.Print("Stack: [")
	for i, item := range s.items { // 打印栈的内容
		if i > 0 {
			fmt.Print(", ") // 打印栈的内容
		}
		fmt.Printf("%v", item) // 打印栈的内容
	}
	fmt.Println("]")                                        // 打印栈的内容
	fmt.Println("Top of stack: →", s.items[len(s.items)-1]) // 打印栈顶元素
}

// PrintReversed 以栈顶到栈底的顺序打印栈的内容
func (s *Stack) PrintReversed() { // 以栈顶到栈底的顺序打印栈的内容
	if s.IsEmpty() {
		fmt.Println("Stack (top to bottom): []") // 如果栈为空，则打印空栈
		return
	}

	fmt.Print("Stack (top to bottom): [")    // 打印栈的内容
	for i := len(s.items) - 1; i >= 0; i-- { // 从栈顶到栈底打印栈的内容
		if i < len(s.items)-1 { // 如果不是最后一个元素，则打印逗号
			fmt.Print(", ")
		}
		fmt.Printf("%v", s.items[i]) // 打印栈的内容
	}
	fmt.Println("]")                                        // 打印栈的内容
	fmt.Println("Top of stack: →", s.items[len(s.items)-1]) // 打印栈顶元素
}

// 创建一个泛型版本的栈（Go 1.18+）
type GenericStack[T any] struct { // 创建一个泛型版本的栈
	items []T // 使用切片存储栈中的元素
}

func NewGenericStack[T any]() *GenericStack[T] { // 创建一个泛型版本的栈
	return &GenericStack[T]{
		items: make([]T, 0),
	}
}

func (s *GenericStack[T]) Push(item T) { // 将元素压入栈顶
	s.items = append(s.items, item)
}

func (s *GenericStack[T]) Pop() (T, error) { // 弹出并返回栈顶元素，如果栈为空则返回错误
	var zero T
	if s.IsEmpty() {
		return zero, errors.New("stack is empty")
	}

	index := len(s.items) - 1 // 获取栈顶元素（最后一个元素）
	item := s.items[index]
	s.items = s.items[:index]

	return item, nil // 返回栈顶元素
}

func (s *GenericStack[T]) Peek() (T, error) { // 返回栈顶元素但不弹出，如果栈为空则返回错误
	var zero T
	if s.IsEmpty() { // 如果栈为空，则返回错误
		return zero, errors.New("stack is empty") // 返回错误
	}

	return s.items[len(s.items)-1], nil
}

func (s *GenericStack[T]) IsEmpty() bool { // 判断栈是否为空
	return len(s.items) == 0
}

func (s *GenericStack[T]) Size() int { // 返回栈中元素的数量
	return len(s.items)
}

func (s *GenericStack[T]) Print() { // 打印栈的内容
	if s.IsEmpty() {
		fmt.Println("Generic Stack: []")
		return
	}

	fmt.Print("Generic Stack: [") // 打印栈的内容
	for i, item := range s.items {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%v", item) // 打印栈的内容
	}
	fmt.Println("]")                                        // 打印栈的内容
	fmt.Println("Top of stack: →", s.items[len(s.items)-1]) // 打印栈顶元素
}

func main() {
	// 演示普通栈的用法
	fmt.Println("标准栈演示:") // 打印栈的内容
	stack := NewStack()

	// 压入元素
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	stack.Push(50)

	// 打印栈内容
	fmt.Println("栈内容（从栈底到栈顶）:")
	stack.Print()

	fmt.Println("\n栈内容（从栈顶到栈底）:")
	stack.PrintReversed()

	// 查看栈顶元素
	top, _ := stack.Peek()
	fmt.Printf("\n栈顶元素: %v\n", top)

	// 弹出元素
	popped, _ := stack.Pop()
	fmt.Printf("弹出的元素: %v\n", popped)

	// 再次打印栈内容
	fmt.Println("\n弹出一个元素后的栈:")
	stack.Print()

	// 清空栈
	fmt.Println("\n清空栈...")
	stack.Clear()
	fmt.Printf("栈是否为空: %v\n", stack.IsEmpty())
	stack.Print()

	// 泛型栈演示（需要 Go 1.18+）
	fmt.Println("\n\n泛型栈演示 (字符串类型):")     // 打印栈的内容
	strStack := NewGenericStack[string]() // 创建一个泛型版本的栈

	strStack.Push("Hello")       // 将元素压入栈顶
	strStack.Push("World")       // 将元素压入栈顶
	strStack.Push("Go")          // 将元素压入栈顶
	strStack.Push("Programming") // 将元素压入栈顶

	strStack.Print() // 打印栈的内容

	strPopped, _ := strStack.Pop()        // 弹出并返回栈顶元素，如果栈为空则返回错误
	fmt.Printf("弹出的字符串: %s\n", strPopped) // 打印栈的内容
	strStack.Print()                      // 打印栈的内容
}
