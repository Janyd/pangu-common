package arraylist

import (
	"fmt"
	"github.com/Janyd/pangu-common/container/comparator"
	"github.com/Janyd/pangu-common/container/util"
	"reflect"
	"strings"
)

type List[T any] struct {
	elements   []T
	size       int
	comparator comparator.Comparator[T]
}

const (
	growthFactor = float32(2.0)
	shrinkFactor = float32(0.25)
)

func New[T any](values ...T) *List[T] {
	list := &List[T]{
		comparator: func(a, b T) int {
			result := reflect.DeepEqual(a, b)
			if result == true {
				return 0
			} else {
				return -1
			}
		},
	}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

func NewWith[T any](comparator comparator.Comparator[T]) *List[T] {
	list := &List[T]{
		comparator: comparator,
	}
	return list
}
func (list *List[T]) Add(values ...T) {
	list.growBy(len(values))
	for _, value := range values {
		list.elements[list.size] = value
		list.size++
	}
}

func (list *List[T]) AddAll(values *List[T]) {
	list.Add(values.Values()...)
}

func (list *List[T]) Get(index int) T {
	return list.elements[index]
}

func (list *List[T]) Remove(index int) {
	if !list.WithinRange(index) {
		return
	}
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size--
	list.shrink()
}

func (list *List[T]) RemoveAll(deleteList *List[T]) {
	for _, value := range deleteList.Values() {
		list.Remove(list.IndexOf(value))
	}
}

// Contains checks if elements (one or more) are present in the set.
// All elements have to be present in the set for the method to return true.
// Performance time complexity of n^2.
// Returns true if no arguments are passed at all, i.e. set is always super-set of empty set.
func (list *List[T]) Contains(values ...T) bool {

	for _, searchValue := range values {
		found := false
		for index := 0; index < list.size; index++ {
			if list.comparator(list.elements[index], searchValue) == 0 {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (list *List[T]) Values() []T {
	newElements := make([]T, list.size, list.size)
	copy(newElements, list.elements[:list.size])
	return newElements
}

// IndexOf returns index of provided element
func (list *List[T]) IndexOf(value T) int {
	if list.size == 0 {
		return -1
	}
	for index, element := range list.elements {
		if list.comparator(value, element) == 0 {
			return index
		}
	}
	return -1
}

// Empty returns true if list does not contain any elements.
func (list *List[T]) Empty() bool {
	return list.size == 0
}

// Size returns number of elements within the list.
func (list *List[T]) Size() int {
	return list.size
}

func (list *List[T]) Clear() {
	list.size = 0
	list.elements = []T{}
}

func (list *List[T]) Sort(comparator comparator.Comparator[T]) {
	if len(list.elements) < 2 {
		return
	}

	util.Sort(list.elements[:list.size], comparator)
}

func (list *List[T]) Swap(i, j int) {
	if list.WithinRange(i) || list.WithinRange(j) {
		list.elements[i], list.elements[j] = list.elements[j], list.elements[i]
	}
}

func (list *List[T]) Insert(index int, values ...T) {
	if !list.WithinRange(index) {
		//Append
		if index == list.size {
			list.Add(values...)
		}
		return
	}

	l := len(values)
	list.growBy(l)
	list.size += l
	copy(list.elements[index+l:], list.elements[index:list.size-l])
	copy(list.elements[index:], values)

}

func (list *List[T]) Set(index int, value T) {
	if !list.WithinRange(index) {
		if index == list.size {
			list.Add(value)
		}
		return
	}

	list.elements[index] = value
}

func (list *List[T]) String() string {
	str := "ArrayList\n"
	var values []string
	for _, value := range list.elements[:list.size] {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// WithinRange Check that the index is within bounds of the list
func (list *List[T]) WithinRange(index int) bool {
	return index >= 0 && index < list.size
}

func (list *List[T]) growBy(n int) {
	// When capacity is reached, grow by a factor of growthFactor and add number of elements
	currentCapacity := cap(list.elements)
	if list.size+n >= currentCapacity {
		newCapacity := int(growthFactor * float32(currentCapacity+n))
		list.resize(newCapacity)
	}
}

func (list *List[T]) resize(cap int) {
	newElements := make([]T, cap, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}

// Shrink the array if necessary, i.e. when size is shrinkFactor percent of current capacity
func (list *List[T]) shrink() {
	if shrinkFactor == 0.0 {
		return
	}
	// Shrink when size is at shrinkFactor * capacity
	currentCapacity := cap(list.elements)
	if list.size <= int(float32(currentCapacity)*shrinkFactor) {
		list.resize(list.size)
	}
}
