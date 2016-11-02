package main

import (
	"fmt"
)

type Element struct {
	data string

}
type Stack struct {
	elements []*Element
}

func NewStack() *Stack {
	return &Stack {elements: make([]*Element, 0)}
}

func (s *Stack) Push(data string) {
	s.elements = append(s.elements, &Element{data: data})
}

func (s *Stack) Pop() *Element {
	if s.IsEmpty() {
		return nil
	}
	returnValue := s.elements[len(s.elements) - 1]
	s.elements = s.elements[0:len(s.elements) - 1]
	return returnValue
}

func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack) Len() int {
	return len(s.elements)
}

func main() {
	stack := NewStack()
	stack.Push("bottom")
	stack.Push("middle")
	stack.Push("top")

	fmt.Println(stack.Pop().data)
	fmt.Println(stack.Pop().data)
	fmt.Println(stack.Pop().data)
}
