package once

import "fmt"

// Get ...
func Get[T any](name string, fn func() T) T {
	v := global.Get(name, func() any {
		return fn()
	})

	vt, ok := v.(T)
	if !ok {
		panic(fmt.Sprintf("once.Get: register the same name(%s) with different type", name))
	}

	return vt
}
