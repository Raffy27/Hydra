package util

import "log"

//Handle catches and handles errors.
func Handle(err error) {
	if err == nil {
		return
	}
	log.Panicln(err)
}

func Panicln(err error, str string) {
	if err == nil {
		return
	}
	log.Panicln(str+",", err)
}
