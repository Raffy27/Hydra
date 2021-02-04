package util

import "log"

//Handle catches and handles errors.
func Handle(err error) {
	if err == nil {
		return
	}
	log.Panicln(err)
}
