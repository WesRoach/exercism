package linkedlist

import "errors"

// Element defines
type Element struct {
	data int
	next *Element
}

// List defines
type List struct {
	head *Element
	size int
}

// New creates List given []int
func New(intArray []int) *List {
	intArrayLength := len(intArray)
	if intArrayLength == 0 {
		return &List{size: intArrayLength}
	}

	if intArrayLength == 1 {
		return &List{size: intArrayLength, head: &Element{data: intArray[0]}}
	}

	// Otherwise: Create List; Init Head; Iterate through
	ls := List{size: intArrayLength, head: &Element{data: intArray[0]}}
	currentElement := ls.head
	for i := 1; i < intArrayLength; i++ {
		currentElement.next = &Element{data: intArray[i]}
		currentElement = currentElement.next
	}

	return &ls
}

// Size returns size of List
func (ls *List) Size() int {
	return ls.size
}

// Array converts List into []int
func (ls *List) Array() []int {
	if ls.Size() == 0 {
		return make([]int, ls.Size())
	}
	arr := make([]int, ls.Size())
	currentElement := ls.head
	for i := 0; i < ls.Size(); i++ {
		arr[i] = currentElement.data
		currentElement = currentElement.next
	}
	return arr
}

// Pop returns the last value from List
func (ls *List) Pop() (int, error) {
	if ls.Size() == 0 {
		return 0, errors.New("attempted to Pop from empty List")
	}

	currentElement := ls.head
	previousElement := currentElement
	for currentElement.next != nil {
		previousElement = currentElement
		currentElement = currentElement.next
	}
	previousElement.next = nil
	ls.size = ls.size - 1
	return currentElement.data, nil
}

// Push places an item at the end of List
func (ls *List) Push(val int) {

	if ls.Size() == 0 {
		ls.head = &Element{data: val}
	}

	ls.size = ls.size + 1

	currentElement := ls.head
	for currentElement.next != nil {
		currentElement = currentElement.next
	}
	currentElement.next = &Element{data: val}

}

// func (*List) Reverse() *List
