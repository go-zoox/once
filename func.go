package once

// Func ...
func Func(name string, fn func()) func() {
	return func() {
		Do(name, fn)
	}
}
