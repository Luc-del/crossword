package v4

type sorted[E any] struct {
	items []E
	less  func(a, b E) bool
}

func newSortedSlice[E any](s []E, less func(a, b E) bool) sorted[E] {
	return sorted[E]{s, less} // TODO
}

func (s sorted[E]) Insert() {}

func (s sorted[E]) Patch(id string, item E) {}

func (s sorted[E]) GetMax() E {
	return *new(E) // TODO
}
