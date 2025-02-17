package knowledge

type Tagged[T any] struct {
	Value T
	Tags  []string
}

func (t Tagged[T]) TaggedWith(tag string) bool {
	for _, t := range t.Tags {
		if t == tag {
			return true
		}
	}
	return false
}
