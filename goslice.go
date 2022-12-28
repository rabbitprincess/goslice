package goslice

type Slice[T any] []T

func New[T any](size int) *Slice[T] {
	s := make(Slice[T], size)
	return &s
}

func From[T any](v ...T) *Slice[T] {
	var s Slice[T] = v
	return &s
}
