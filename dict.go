package collection

type dict[K comparable, T any] struct {
	items map[K]T
}

func Dict[K comparable, T any](m map[K]T) *dict[K, T] {
	return &dict[K, T]{
		items: m,
	}
}

func (d *dict[K, T]) Split() (keys []K, values []T) {
	for k, v := range d.items {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

func (d *dict[K, T]) Keys(t map[K]T) (keys []K) {
	for k := range t {
		keys = append(keys, k)
	}
	return keys
}

func (d *dict[K, T]) Values() (values []T) {
	for _, v := range d.items {
		values = append(values, v)
	}
	return values
}

func SplitMap[K comparable, T interface{}](t map[K]T) (keys []K, values []T) {
	for k, v := range t {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

func MapKeys[K comparable, T interface{}](t map[K]T) (keys []K) {
	for k := range t {
		keys = append(keys, k)
	}
	return keys
}

func MapValues[K comparable, T interface{}](t map[K]T) (values []T) {
	for _, v := range t {
		values = append(values, v)
	}
	return values
}
