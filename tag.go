package knowledge

type Tagged[T any] struct {
	value T
	tags  []string
}

func (t Tagged[T]) TaggedWith(tag string) bool {
	for _, t := range t.tags {
		if t == tag {
			return true
		}
	}
	return false
}
