package collection

type Set[T comparable] map[T]struct{}

func MakeSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Remove(v T) {
	delete(s, v)
}

func (s Set[T]) Clear() {
	for k := range s {
		delete(s, k)
	}
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s Set[T]) GetAnyElement() T {
	for k := range s {
		return k
	}
	panic("Tried to get element from empty set.")
}

func (s Set[T]) GetOnlyElement() T {
	if s.Len() > 1 {
		panic("Tried to get only element from a set with multiple elements.")
	}
	for k := range s {
		return k
	}
	panic("Tried to get only element from empty set.")
}
