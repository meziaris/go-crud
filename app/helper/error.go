package helper

import "log"

func FatalIfNeeded(err interface{}) {
	if err != nil {
		log.Fatal(err)
	}
}
