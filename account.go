package knowledge

type Account[T any] struct {
	primary     T
	subsidaries []T
}
