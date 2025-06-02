package learnlab

import (
	"fmt"
)

// Creating Generic Stack Data Structure
type Stack[T any] struct {
	elements []T
}

// Pushing to Stack
func (s *Stack[T]) push(ele T) {
	s.elements = append(s.elements, ele)
} 

// Popping from Stack
func (s *Stack[T]) pop() (T, bool) {
	
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	
	ele := s.elements[len(s.elements) - 1]
	s.elements = s.elements[:len(s.elements) - 1]
	return ele, true
}

// Print All elements
func (s *Stack[T]) printAll() {
	fmt.Print("Stack Elements : ")
	for _, v := range s.elements {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

// Checks whether the stack is empty or not.
func (s *Stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}

func Generics() {
	s1 := Stack[int32]{}
	fmt.Println(s1.isEmpty())
	s1.push(10)
	s1.push(20)
	s1.push(30)
	s1.push(40)
	s1.push(50)
	s1.printAll()
	s1.pop()
	s1.printAll()

	s2 := Stack[string]{}
	s2.push("I")
	s2.push("am")
	s2.push("Ashish")
	s2.push("Verma")
	s2.printAll()
	s2.pop()
	s2.pop()
	s2.pop()
	s2.printAll()
}
