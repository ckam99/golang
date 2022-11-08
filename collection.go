package collection

import (
	"fmt"
	"strings"
)

type collection[T interface{}] struct {
	arr []T
}

func Collect[T interface{}](t []T) *collection[T] {
	return &collection[T]{
		arr: t,
	}
}

func (c *collection[T]) Map(f func(T, int64) T) *collection[T] {
	var tmp []T
	for k, v := range c.arr {
		tmp = append(tmp, f(v, int64(k)))
	}
	c.arr = tmp
	return c
}

func (c *collection[T]) Filter(f func(T, int64) bool) *collection[T] {
	var tmp []T
	for k, v := range c.arr {
		if f(v, int64(k)) {
			tmp = append(tmp, v)
		}
	}
	c.arr = tmp
	return c
}

func (c *collection[T]) Shift() *collection[T] {
	c.arr = c.arr[1:]
	return c
}

func (c *collection[T]) Pop() *collection[T] {
	c.arr = c.arr[:len(c.arr)-1]
	return c
}

func (c *collection[T]) Reverse() *collection[T] {
	var result []T
	for i := len(c.arr) - 1; i >= 0; i-- {
		result = append(result, c.arr[i])
	}
	c.arr = result
	return c
}

func (c *collection[T]) Join(sep string) string {
	sb := strings.Builder{}
	for k, v := range c.arr {
		sb.WriteString(fmt.Sprintf("%v", v))
		if len(c.arr) > k+1 {
			sb.WriteString(sep)
		}
	}
	return sb.String()
}

func (c *collection[T]) Remove(index int64) *collection[T] {
	c.arr = append(c.arr[:index], c.arr[index+1:]...)
	return c
}

func (c *collection[T]) ToList() []T {
	return c.arr
}

func (c *collection[T]) First() T {
	return c.arr[0]
}

func (c *collection[T]) Last() T {
	return c.arr[len(c.arr)-1]
}

func (c *collection[T]) Get(index int64) T {
	return c.arr[index]
}

func Map[T interface{}](t []T, f func(T, int) T) []T {
	var tp []T
	for k, v := range t {
		tp = append(tp, f(v, k))
	}
	return tp
}

func Filter[T interface{}](t []T, f func(T, int) bool) []T {
	var o []T
	for k, v := range t {
		if f(v, k) {
			o = append(o, v)
		}
	}
	return o
}

func Concat[T interface{}](t1 []T, t2 []T) []T {
	return append(t1[:], t2[:]...)
}

func Shift[T interface{}](arr []T) []T {
	if len(arr) > 0 {
		return arr[1:]
	}
	return arr
}

func Remove[T interface{}](arr []T, index int) []T {
	if index < len(arr) {
		arr = append(arr[:index], arr[index+1:]...)
	}
	return arr
}

func Pop[T interface{}](arr []T) []T {
	if len(arr) > 0 {
		return arr[:len(arr)-1]
	}
	return arr
}

func Set[T any](arr []T) []T {
	var sets []T
	for _, v := range arr {
		exist := false
		for _, v2 := range sets {
			if fmt.Sprintf("%v", v) == fmt.Sprintf("%v", v2) {
				exist = true
				break
			}
		}
		if !exist {
			sets = append(sets, v)
		}
	}
	return sets
}

func Join[T interface{}](tab []T, sep string) string {
	sb := strings.Builder{}
	for k, v := range tab {
		sb.WriteString(fmt.Sprintf("%v", v))
		if len(tab) > k+1 {
			sb.WriteString(sep)
		}
	}
	return sb.String()
}

func Reverse[T interface{}](arr []T) []T {
	var result []T
	for i := len(arr) - 1; i >= 0; i-- {
		result = append(result, arr[i])
	}
	return result
}
