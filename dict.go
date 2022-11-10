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
	return SplitDict(d.items)
}

func (d *dict[K, T]) Keys() (keys []K) {
	return DictKeys(d.items)
}

func (d *dict[K, T]) Values() (values []T) {
	return DictValues(d.items)
}

func (d *dict[K, T]) Has(key K) bool {
	_, ok := d.items[key]
	return ok
}

func (d *dict[K, T]) Get(key K) T {
	return d.items[key]
}

func (d *dict[K, T]) Set(key K, value T) {
	d.items[key] = value
}

func (d *dict[K, T]) Remove(key K) bool {
	if d.Has(key) {
		delete(d.items, key)
		return true
	}
	return false
}

func (d *dict[K, T]) Map(f func(K, T) (K, T)) *dict[K, T] {
	d.items = MapDict(d.items, f)
	return d
}

func (d *dict[K, T]) Filter(f func(K, T) bool) *dict[K, T] {
	d.items = FilterDict(d.items, f)
	return d
}

func (d *dict[K, T]) Value() map[K]T {
	return d.items
}

func SplitDict[K comparable, T interface{}](t map[K]T) (keys []K, values []T) {
	for k, v := range t {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

func DictKeys[K comparable, T interface{}](t map[K]T) (keys []K) {
	for k := range t {
		keys = append(keys, k)
	}
	return keys
}

func DictValues[K comparable, T interface{}](t map[K]T) (values []T) {
	for _, v := range t {
		values = append(values, v)
	}
	return values
}

func FilterDict[K comparable, T interface{}](dt map[K]T, f func(K, T) bool) map[K]T {
	m := make(map[K]T)
	for k, v := range dt {
		if f(k, v) {
			m[k] = v
		}
	}
	return m
}

func MapDict[K comparable, T interface{}](dt map[K]T, f func(K, T) (K, T)) map[K]T {
	m := make(map[K]T)
	for k, v := range dt {
		key, value := f(k, v)
		m[key] = value
	}
	return m
}
