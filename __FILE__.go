package gomacros

import "runtime"

func M__FILE__() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}
	return file
}
