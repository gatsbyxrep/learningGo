package main

import (
	"fmt"
	"learning/algoritms"
	"learning/mystack"
)

func main() {
	s := myStack.Stack{}
	s.Push(5)
	fmt.Println(s.Pop().(int))

	slice := []int{1, -1, 0, 150, 13}
	algoritms.QuickSort(slice, 0, len(slice)-1, func(a, b int) bool { return a < b })
	fmt.Println(fmt.Sprint(slice))
}
