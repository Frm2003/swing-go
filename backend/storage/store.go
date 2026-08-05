package storage

type allocator[K comparable] interface {
	NewId() K
}

type Item[K comparable] interface {
	GetId() K
}

type Store[K comparable, V Item[K]] struct {
	allocator allocator[K]
	content   map[K]V
}

func NewStore[K comparable, V Item[K]](a allocator[K]) *Store[K, V] {
	return &Store[K, V]{
		allocator: a,
		content:   make(map[K]V),
	}
}

func (s *Store[K, V]) NewId() K {
	return s.allocator.NewId()
}

func (s *Store[K, V]) Get(k K) (V, bool) {
	v, ok := s.content[k]
	return v, ok
}

func (s *Store[K, V]) Register(v V) {
	s.content[v.GetId()] = v
}
