package panic_helper

import "fmt"

func RecoverOnError(err *error) {
	if r := recover(); r != nil {
		var ok bool
		*err, ok = r.(error)
		if !ok {
			*err = fmt.Errorf("pkg: %v", r)
		}
	}
}
