package once

// Do ...
func Do(name string, fn func()) {
	global.Do(name, fn)
}
