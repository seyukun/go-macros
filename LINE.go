package gomacros

import "runtime"

func M__LINE__() int {
	_, _, line, ok := runtime.Caller(1)
	if !ok {
		return -1
	}
	return line
}
