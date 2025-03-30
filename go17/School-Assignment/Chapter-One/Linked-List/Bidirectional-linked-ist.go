// // 用go语言实现个双向链表
// package main

// import (
// 	"fmt"
// )

// // Node 表示双向链表中的节点
// type Node struct {
// 	Value    interface{} // 存储任意类型的值
// 	Next     *Node       // 指向下一个节点
// 	Previous *Node       // 指向前一个节点
// }

// // DoublyLinkedList 表示双向链表
// type DoublyLinkedList struct {
// 	Head   *Node // 链表头部
// 	Tail   *Node // 链表尾部
// 	Length int   // 链表长度
// }

// // NewDoublyLinkedList 创建一个新的双向链表
// func NewDoublyLinkedList() *DoublyLinkedList { // 创建一个新的双向链表
// 	return &DoublyLinkedList{ // 返回一个新的双向链表
// 		Head:   nil,
// 		Tail:   nil,
// 		Length: 0,
// 	}
// }

// // IsEmpty 检查链表是否为空
// func (dll *DoublyLinkedList) IsEmpty() bool { // 检查链表是否为空
// 	return dll.Length == 0 // 返回链表是否为空
// }

// // Size 返回链表的大小
// func (dll *DoublyLinkedList) Size() int { // 返回链表的大小
// 	return dll.Length // 返回链表的大小
// }

// // AddFront 在链表头部添加节点
// func (dll *DoublyLinkedList) AddFront(value interface{}) { // 在链表头部添加节点
// 	newNode := &Node{ // 创建一个新的节点
// 		Value:    value,    // 设置节点的值
// 		Next:     dll.Head, // 设置节点的下一个节点
// 		Previous: nil,      // 设置节点的上一个节点
// 	}

// 	// 如果链表为空，则新节点同时是头部和尾部
// 	if dll.IsEmpty() { // 如果链表为空
// 		dll.Head = newNode // 设置链表的头部为新节点
// 		dll.Tail = newNode // 设置链表的尾部为新节点
// 	} else {
// 		// 将当前头部节点的Previous指向新节点
// 		dll.Head.Previous = newNode // 设置当前头部节点的上一个节点为新节点
// 		// 更新头部为新节点
// 		dll.Head = newNode
// 	}

// 	dll.Length++
// }

// // AddBack 在链表尾部添加节点
// func (dll *DoublyLinkedList) AddBack(value interface{}) { // 在链表尾部添加节点
// 	newNode := &Node{ // 创建一个新的节点
// 		Value:    value,    // 设置节点的值
// 		Next:     nil,      // 设置节点的下一个节点
// 		Previous: dll.Tail, // 设置节点的上一个节点
// 	}

// 	// 如果链表为空，则新节点同时是头部和尾部
// 	if dll.IsEmpty() { // 如果链表为空
// 		dll.Head = newNode
// 		dll.Tail = newNode
// 	} else {
// 		// 将当前尾部节点的Next指向新节点
// 		dll.Tail.Next = newNode
// 		// 更新尾部为新节点
// 		dll.Tail = newNode
// 	}

// 	dll.Length++
// }

// // RemoveFront 移除并返回链表头部节点的值
// func (dll *DoublyLinkedList) RemoveFront() (interface{}, error) { // 移除并返回链表头部节点的值
// 	if dll.IsEmpty() { // 如果链表为空
// 		return nil, fmt.Errorf("list is empty") // 返回错误
// 	}

// 	value := dll.Head.Value // 获取链表头部的值

// 	// 如果链表只有一个节点
// 	if dll.Head == dll.Tail { // 如果链表只有一个节点
// 		dll.Head = nil // 设置链表的头部为nil
// 		dll.Tail = nil // 设置链表的尾部为nil
// 	} else {
// 		// 更新头部为下一个节点
// 		dll.Head = dll.Head.Next // 设置链表的头部为下一个节点
// 		dll.Head.Previous = nil  // 设置链表的头部节点的上一个节点为nil
// 	}

// 	dll.Length--      // 减少链表的长度
// 	return value, nil // 返回链表头部的值
// }

// // RemoveBack 移除并返回链表尾部节点的值
// func (dll *DoublyLinkedList) RemoveBack() (interface{}, error) { // 移除并返回链表尾部节点的值
// 	if dll.IsEmpty() {
// 		return nil, fmt.Errorf("list is empty")
// 	}

// 	value := dll.Tail.Value

// 	// 如果链表只有一个节点
// 	if dll.Head == dll.Tail { // 如果链表只有一个节点
// 		dll.Head = nil
// 		dll.Tail = nil
// 	} else {
// 		// 更新尾部为前一个节点
// 		dll.Tail = dll.Tail.Previous // 设置链表的尾部为前一个节点
// 		dll.Tail.Next = nil
// 	}

// 	dll.Length-- // 减少链表的长度
// 	return value, nil
// }

// // Insert 在指定索引处插入节点
// func (dll *DoublyLinkedList) Insert(index int, value interface{}) error { // 在指定索引处插入节点
// 	// 检查索引是否有效
// 	if index < 0 || index > dll.Length { // 如果索引无效
// 		return fmt.Errorf("index out of range") // 返回错误
// 	}

// 	// 如果在头部插入
// 	if index == 0 { // 如果索引为0
// 		dll.AddFront(value) // 在头部插入节点
// 		return nil          // 返回nil
// 	}

// 	// 如果在尾部插入
// 	if index == dll.Length { // 如果索引为链表长度
// 		dll.AddBack(value) // 在尾部插入节点
// 		return nil         // 返回nil
// 	}

// 	// 创建新节点
// 	newNode := &Node{ // 创建一个新的节点
// 		Value: value, // 设置节点的值
// 	}

// 	// 找到要插入位置的当前节点
// 	current := dll.Head          // 获取链表头部的节点
// 	for i := 0; i < index; i++ { // 遍历链表
// 		current = current.Next // 设置当前节点为下一个节点
// 	}

// 	// 更新指针
// 	newNode.Previous = current.Previous // 设置新节点的上一个节点为当前节点的上一个节点
// 	newNode.Next = current              // 设置新节点的下一个节点为当前节点
// 	current.Previous.Next = newNode     // 设置当前节点的上一个节点的下一个节点为新节点
// 	current.Previous = newNode          // 设置当前节点的上一个节点为新节点

// 	dll.Length++ // 增加链表的长度
// 	return nil   // 返回nil
// }

// // Remove 移除指定索引处的节点并返回其值
// func (dll *DoublyLinkedList) Remove(index int) (interface{}, error) { // 移除指定索引处的节点并返回其值
// 	// 检查索引是否有效
// 	if index < 0 || index >= dll.Length { // 如果索引无效
// 		return nil, fmt.Errorf("index out of range") // 返回错误
// 	}

// 	// 如果移除头部节点
// 	if index == 0 { // 如果索引为0
// 		return dll.RemoveFront() // 移除头部节点
// 	}

// 	// 如果移除尾部节点
// 	if index == dll.Length-1 { // 如果索引为链表长度减1
// 		return dll.RemoveBack() // 移除尾部节点
// 	}

// 	// 找到要移除的节点
// 	current := dll.Head
// 	for i := 0; i < index; i++ { // 遍历链表
// 		current = current.Next // 设置当前节点为下一个节点
// 	}

// 	// 保存节点值
// 	value := current.Value // 获取节点的值

// 	// 更新指针，跳过当前节点
// 	current.Previous.Next = current.Next     // 设置当前节点的上一个节点的下一个节点为当前节点的下一个节点
// 	current.Next.Previous = current.Previous // 设置当前节点的下一个节点的上一个节点为当前节点的上一个节点

// 	dll.Length--      // 减少链表的长度
// 	return value, nil // 返回节点的值
// }

// // Get 获取指定索引处节点的值
// func (dll *DoublyLinkedList) Get(index int) (interface{}, error) { // 获取指定索引处节点的值
// 	// 检查索引是否有效
// 	if index < 0 || index >= dll.Length { // 如果索引无效
// 		return nil, fmt.Errorf("index out of range") // 返回错误
// 	}

// 	current := dll.Head          // 获取链表头部的节点
// 	for i := 0; i < index; i++ { // 遍历链表
// 		current = current.Next // 设置当前节点为下一个节点
// 	}

// 	return current.Value, nil // 返回节点的值
// }

// // PrintForward 从头到尾打印链表
// func (dll *DoublyLinkedList) PrintForward() { // 从头到尾打印链表
// 	if dll.IsEmpty() { // 如果链表为空
// 		fmt.Println("List is empty") // 打印链表为空
// 		return                       // 返回
// 	}

// 	current := dll.Head    // 获取链表头部的节点
// 	fmt.Print("Forward: ") // 打印链表头部
// 	for current != nil {   // 遍历链表
// 		fmt.Printf("%v ", current.Value) // 打印节点的值
// 		current = current.Next           // 设置当前节点为下一个节点
// 	}
// 	fmt.Println()
// }

// // PrintBackward 从尾到头打印链表
// func (dll *DoublyLinkedList) PrintBackward() { // 从尾到头打印链表
// 	if dll.IsEmpty() { // 如果链表为空
// 		fmt.Println("List is empty") // 打印链表为空
// 		return                       // 返回
// 	}

// 	current := dll.Tail     // 获取链表尾部的节点
// 	fmt.Print("Backward: ") // 打印链表尾部
// 	for current != nil {    // 遍历链表
// 		fmt.Printf("%v ", current.Value) // 打印节点的值
// 		current = current.Previous       // 设置当前节点为上一个节点
// 	}
// 	fmt.Println() // 打印换行符
// }

// func main() {
// 	// 创建一个新的双向链表
// 	list := NewDoublyLinkedList() // 创建一个新的双向链表

// 	// 添加元素
// 	list.AddBack(1)   // 在链表尾部添加元素
// 	list.AddBack(2)   // 在链表尾部添加元素
// 	list.AddBack(3)   // 在链表尾部添加元素
// 	list.AddFront(0)  // 在链表头部添加元素
// 	list.AddFront(-1) // 在链表头部添加元素

// 	// 打印链表
// 	fmt.Printf("List size: %d\n", list.Size()) // 打印链表的大小
// 	list.PrintForward()                        // 从头到尾打印链表
// 	list.PrintBackward()                       // 从尾到头打印链表

// 	// 在中间插入元素
// 	list.Insert(2, 100)                              // 在指定索引处插入元素
// 	fmt.Println("\nAfter inserting 100 at index 2:") // 打印插入后的链表
// 	list.PrintForward()                              // 从头到尾打印链表

// 	// 移除元素
// 	removed, _ := list.Remove(2)                              // 移除指定索引处的元素
// 	fmt.Printf("\nRemoved element at index 2: %v\n", removed) // 打印移除的元素
// 	list.PrintForward()                                       // 从头到尾打印链表

// 	// 获取元素
// 	val, _ := list.Get(1)                         // 获取指定索引处的元素
// 	fmt.Printf("\nElement at index 1: %v\n", val) // 打印获取的元素

// 	// 移除首尾元素
// 	front, _ := list.RemoveFront()                                     // 移除头部元素
// 	back, _ := list.RemoveBack()                                       // 移除尾部元素
// 	fmt.Printf("\nRemoved front: %v, Removed back: %v\n", front, back) // 打印移除的元素
// 	list.PrintForward()                                                // 从头到尾打印链表
// }

// 用go语言实现一个单链表
package main

import (
	"fmt"
)

// Node 表示链表中的节点
type Node struct {
	Value interface{} // 节点存储的值
	Next  *Node       // 指向下一个节点的指针
}

// LinkedList 表示单向链表
type LinkedList struct {
	Head   *Node // 链表头节点
	Length int   // 链表长度
}

// NewLinkedList 创建一个新的单向链表
func NewLinkedList() *LinkedList { // 创建一个新的单向链表
	return &LinkedList{ // 返回一个新的单向链表
		Head:   nil, // 设置链表头节点为nil
		Length: 0,   // 设置链表长度为0
	}
}

// IsEmpty 检查链表是否为空
func (ll *LinkedList) IsEmpty() bool { // 检查链表是否为空
	return ll.Length == 0 // 返回链表是否为空
}

// AddFront 在链表头部添加节点
func (ll *LinkedList) AddFront(value interface{}) { // 在链表头部添加节点
	newNode := &Node{ // 创建一个新的节点
		Value: value,   // 设置节点的值
		Next:  ll.Head, // 设置节点的下一个节点为链表头节点
	}
	ll.Head = newNode // 设置链表头节点为新节点
	ll.Length++       // 增加链表长度
}

// AddBack 在链表尾部添加节点
func (ll *LinkedList) AddBack(value interface{}) { // 在链表尾部添加节点
	newNode := &Node{ // 创建一个新的节点
		Value: value, // 设置节点的值
		Next:  nil,   // 设置节点的下一个节点为nil
	}

	if ll.IsEmpty() { // 如果链表为空
		ll.Head = newNode // 设置链表头节点为新节点
	} else {
		current := ll.Head        // 获取链表头节点
		for current.Next != nil { // 遍历链表
			current = current.Next // 设置当前节点为下一个节点
		}
		current.Next = newNode // 设置当前节点的下一个节点为新节点
	}
	ll.Length++ // 增加链表长度
}

// RemoveFront 移除并返回链表头部节点的值
func (ll *LinkedList) RemoveFront() (interface{}, error) { // 移除并返回链表头部节点的值
	if ll.IsEmpty() { // 如果链表为空
		return nil, fmt.Errorf("链表为空") // 返回错误
	}

	value := ll.Head.Value // 获取链表头节点的值
	ll.Head = ll.Head.Next // 设置链表头节点为下一个节点
	ll.Length--
	return value, nil
}

// Get 获取指定位置的节点值
func (ll *LinkedList) Get(index int) (interface{}, error) { // 获取指定位置的节点值
	if index < 0 || index >= ll.Length { // 如果索引越界
		return nil, fmt.Errorf("索引越界") // 返回错误
	}

	current := ll.Head           // 获取链表头节点
	for i := 0; i < index; i++ { // 遍历链表
		current = current.Next // 设置当前节点为下一个节点
	}
	return current.Value, nil // 返回节点的值
}

// Print 打印链表所有节点的值
func (ll *LinkedList) Print() { // 打印链表所有节点的值
	if ll.IsEmpty() { // 如果链表为空
		fmt.Println("链表为空") // 打印链表为空
		return              // 返回
	}

	current := ll.Head   // 获取链表头节点
	fmt.Print("链表内容: ")  // 打印链表内容
	for current != nil { // 遍历链表
		fmt.Printf("%v -> ", current.Value) // 打印节点的值
		current = current.Next              // 设置当前节点为下一个节点
	}
	fmt.Println("nil")                  // 打印换行符
	fmt.Printf("链表长度: %d\n", ll.Length) // 打印链表长度
}

func main() {
	// 创建新链表
	list := NewLinkedList() // 创建一个新的单向链表

	// 添加一些测试数据
	list.AddBack(1)   // 在链表尾部添加元素
	list.AddBack(2)   // 在链表尾部添加元素
	list.AddBack(3)   // 在链表尾部添加元素
	list.AddFront(0)  // 在链表头部添加元素
	list.AddFront(-1) // 在链表头部添加元素

	// 打印链表内容
	fmt.Println("初始链表:") // 打印初始链表
	list.Print()         // 打印链表内容

	// 测试获取元素
	if val, err := list.Get(2); err == nil { // 获取指定位置的元素值
		fmt.Printf("\n索引2的元素值: %v\n", val) // 打印指定位置的元素值
	}

	// 测试删除头部元素
	if val, err := list.RemoveFront(); err == nil { // 删除头部元素
		fmt.Printf("\n删除的头部元素: %v\n", val) // 打印删除的头部元素
		fmt.Println("删除头部元素后的链表:")         // 打印删除头部元素后的链表
		list.Print()                       // 打印链表内容
	}

	// 继续添加元素
	list.AddBack(4)             // 在链表尾部添加元素
	fmt.Println("\n添加新元素后的链表:") // 打印添加新元素后的链表
	list.Print()                // 打印链表内容
}
