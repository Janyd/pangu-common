package lists

import (
	"gitee.com/ChuckChan/pangu-common/container"
	"gitee.com/ChuckChan/pangu-common/container/comparator"
)

type List[T container.Object[T]] interface {
	Get(index int) (T, bool)
	Remove(index int)
	Add(values ...T)
	Contains(values ...T) bool
	Sort(comparator comparator.Comparator[T])
	Swap(index1, index2 int)
	Insert(index int, values ...T)
	Set(index int, value T)
	AddAll(values List[T])

	container.Container[T]
}
