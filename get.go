package once

// Get ...
func Get[T any](name string, fn func() T) T {
	v := global.Get(name, func() any {
		return fn()
	})

	return v.(T)
}
