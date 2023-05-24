package sets

import "github.com/Janyd/pangu-common/container"

type Set[T container.Object[T]] interface {
	Add(elements ...T)
	Remove(elements ...T)
	Contains(elements ...T) bool
	container.Container[T]
}
