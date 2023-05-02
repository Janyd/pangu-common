package container

type Object[T any] interface {
	comparable
	Equals(other T) bool
}
