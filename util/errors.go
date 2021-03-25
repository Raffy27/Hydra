package util

import "log"

//Calm consumes top-level errors.
//Defer this if you want to be absolutely certain that errors should be ignored.
func Calm() {
	if r := recover(); r != nil {
		log.Println("Recovered from", r)
	}
}

//Handle will panic when given a valid error and print some debug info.
func Handle(err error, str ...interface{}) {
	if err == nil {
		return
	}
	if len(str) > 0 {
		log.Panicln(str[0].(string)+",", err)
	} else {
		log.Panicln(err)
	}
}
