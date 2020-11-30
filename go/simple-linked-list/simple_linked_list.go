package linkedlist

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
