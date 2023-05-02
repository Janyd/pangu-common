package arraylist

func (list *List[T]) Each(f func(index int, value T)) {
	iterator := list.Iterator()
	for iterator.Next() {
		f(iterator.Index(), iterator.Value())
	}
}

func (list *List[T]) Filter(f func(index int, value T) bool) *List[T] {
	newList := &List[T]{comparator: list.comparator}
	it := list.Iterator()
	for it.Next() {
		if f(it.Index(), it.Value()) {
			newList.Add(it.Value())
		}
	}
	return newList
}

func (list *List[T]) Any(f func(index int, value T) bool) bool {
	it := list.Iterator()
	for it.Next() {
		if f(it.Index(), it.Value()) {
			return true
		}
	}
	return false
}

func (list *List[T]) All(f func(index int, value T) bool) bool {
	it := list.Iterator()
	for it.Next() {
		if !f(it.Index(), it.Value()) {
			return false
		}
	}
	return true
}
